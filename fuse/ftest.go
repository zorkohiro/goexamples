package main

import (
	"io"
	"os"
	"syscall"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
	log "github.com/sirupsen/logrus"
)

var srcdir string

type testfs struct {
	pathfs.FileSystem
}


// linux syscall stat to fuse Attr
func ls_to_fs(a *syscall.Stat_t, b *fuse.Attr) {
	b.Size = uint64(a.Size)
	b.Blocks = uint64(a.Blocks)
	b.Atime = uint64(a.Atim.Sec)
	b.Atimensec  = uint32(a.Atim.Nsec)
	b.Mtime = uint64(a.Mtim.Sec)
	b.Mtimensec = uint32(a.Mtim.Nsec)
	b.Ctime = uint64(a.Ctim.Sec)
	b.Ctimensec = uint32(a.Ctim.Nsec)
	b.Mode = a.Mode
	b.Blksize = uint32(a.Blksize)
}

func (me *testfs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	var fu fuse.Attr
	var sf syscall.Stat_t
	fullname := srcdir + "/" + name
	log.Debug("Gettatr ", fullname)
	err := syscall.Stat(fullname, &sf)
	if err != nil {
		if err != syscall.ENOENT {
			log.Errorf("GetAttr: %s, %v", fullname, err)
		}
		return nil, fuse.EIO
	}
	ls_to_fs(&sf, &fu)
	return &fu, fuse.OK
}

func (me *testfs) OpenDir(name string, context *fuse.Context) ([]fuse.DirEntry, fuse.Status) {
	var sl []fuse.DirEntry
	fullname := srcdir + "/" + name
	log.Debug("OpenDir ", fullname)
	f, err := os.Open(fullname)
	if err != nil {
		log.Error(err)
		return nil, fuse.EIO
	}
	for {
		ls, err := f.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Error(err)
			return nil, fuse.EIO
		}
		ne := new(fuse.DirEntry)
		ne.Name = ls[0].Name()
		ne.Mode = uint32(ls[0].Mode())
		sl = append(sl, *ne)
	}
	return sl, fuse.OK
}

//func (me *testfs) Access(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
//	fullname := srcdir + "/" + name
//	log.Infof("Access %s", fullname)
//	return fuse.ENOSYS
//}


func (me *testfs) Open(name string, flags uint32, context *fuse.Context) (nodefs.File, fuse.Status) {
        if flags&fuse.O_ANYWRITE != 0 {
                return nil, fuse.EPERM
        }
	fullname := srcdir + "/" + name
	log.Infof("OpenDir %s", fullname)
	return nil, fuse.EPERM
	// return nodefs.NewDataFile([]byte(name)), fuse.OK
}


func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: testfs SRCDIR MOUNTPOINT")
	}
	srcdir = os.Args[1]
        fs := pathfs.NewPathNodeFs(&testfs{FileSystem: pathfs.NewReadonlyFileSystem(pathfs.NewDefaultFileSystem())}, nil)
	conn := nodefs.NewFileSystemConnector(fs.Root(), nil)
	mountopts := fuse.MountOptions{}
	mountopts.FsName = "ftest"
	mountopts.Name = "ftest"
	mountopts.DisableXAttrs = true
	mountopts.Options = append(mountopts.Options, "allow_other")
	server, err := fuse.NewServer(conn.RawFS(), os.Args[2], &mountopts)
	if err != nil {
		log.Fatal(err)
	}
	server.Serve()
}
