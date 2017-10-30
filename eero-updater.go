package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"io"
	"log"
	"os"
	"os/exec"
	"net/http"
	"strings"
)

var debug bool

func main() {

	if len(os.Args) != 3 {
                log.Fatal("usage: eero-updater url other-root-device")
	}
        url := os.Args[1]
	ord := os.Args[2]

	d := os.Getenv("DEBUG")
	switch d[0] {
	case 'Y', 'y', '1':
		debug = true
		log.Println("debugging on")
	}

	// Split out the end filename component of the URL
        f := strings.Split(url, "?")
        s := f[0]
        f = strings.Split(s, "/")
        s = f[len(f) - 1 ]
        lastcomp := "/tmp/" + s

	//
	// We read from standard input
	//
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		log.Println("connected to", url)
	}

	//
	// Because we don't provide a byte-reader while reading
	// data, we have to read the data on the fly, calculating
	// a sha256, stash a copy in a cache file (in case we
	// have to restart), and then send the rest down an internal
	// pipeline that will uncompress and read the uncompressed
	// tarball. The tarball should have a kernel image (which
	// we'll write to /tmp) and a rootfs image (which we'll overwrite
	// the alternate partition with). It's a bit clunky but works.
	//
	gzip_pipe_in, gzip_pipe_out := io.Pipe()

	go func() {
		z := make([]byte, 4096)

		gzf, err := gzip.NewReader(gzip_pipe_in)
		if err != nil {
			log.Fatal(err)
		}
		gzf.Multistream(false)

		pin, pout := io.Pipe()
		go func() {
			tr := tar.NewReader(pin)
			for {
				hdr, err := tr.Next()
				if err != nil {
					if err == io.EOF {
						break
					}
					log.Fatal(err)
				}
				switch hdr.Typeflag {
				case tar.TypeDir:
					log.Fatal("we aren't expecting a directory type")
				case tar.TypeReg:
					switch hdr.Name {
					case "loki-kernel.itb":
						if debug {
							log.Println("saving loki-kernel.itb to /tmp")
						}
						fil, err := os.Create("/tmp/" + hdr.Name)
						if err != nil {
							log.Fatal(err)
						}
						_, err = io.Copy(fil, tr)
						if err != nil {
							log.Fatal(err)
						}
						err = fil.Close()
						if err != nil {
							log.Fatal(err)
						}
					case "loki-rootfs.img":
						if debug {
							log.Println("copying loki-rootfs.img to", ord)
						}
						dev, err := os.OpenFile(ord, os.O_WRONLY, 0)
						if err != nil {
							log.Fatal(err)
						}
						_, err = dev.Seek(0, 0)
						if err != nil {
							log.Fatal(err)
						}
						_, err = io.Copy(dev, tr)
						if err != nil {
							log.Fatal(err)
						}
						err = dev.Close()
						if err != nil {
							log.Fatal(err)
						}
						res := exec.Command("/sbin/e2fsck", "-fy", ord)
						output, err := res.CombinedOutput()
						if err != nil {
							log.Fatalf("%v: output: %v\n", err, string(output))
						}
						res = exec.Command("/sbin/resize2fs", ord)
						output, err = res.CombinedOutput()
						if err != nil {
							log.Fatalf("%v: output: %v\n", err, string(output))
						}
					default:
						log.Fatal("should not have any other name:", hdr.Name)
					}
				default:
					log.Fatal("unable to figure out filetype", hdr)
				}
			}
			pin.Close()
		}()

		for {
			n, err := gzf.Read(z)
			if err != nil {
				if err == io.EOF {
					pout.Close()
					break
				}
				log.Fatal(err)
			}
			pout.Write(z[:n])
		}
		gzip_pipe_in.Close()
	}()

	//
	// We create a sha256 instance for checking our incoming data
	//
	sha := sha256.New()

	//
	// Check for a lookaside file and read from that first?
	//
	l, err := os.Create(lastcomp)
	if err != nil {
		log.Fatal(err)
	}

	//
	// We have a small buffer we stage so that sha256 can look at
	// it before we unzip it and send it down to the tar function
	//
	b := make([]byte, 4096)

	tot := 0
	tmb := 1

	for {
		n, err := res.Body.Read(b)
		if err == nil || err == io.EOF {
			if n != 0 {
				tot = tot + n
				if (debug  && (tot >> 20) > tmb) {
					log.Println("read", tmb, "MiB")
					tmb++
				}
				// copy bytes to our lookaside
				l.Write(b[:n])
				// update sha256sum
				sha.Write(b[:n])
				// Write one byte at a time to simulate a byte reader
				for j := 0; j < n; j++ {
					gzip_pipe_out.Write(b[j : j+1])
				}
			}
		} else if err != nil {
			log.Fatal(err)
		}

		if err == io.EOF {
			break
		}
	}
	gzip_pipe_out.Close()
	res.Body.Close()
	log.Printf("%x /tmp/loki-kernel.itb\n", sha.Sum(nil))
}
