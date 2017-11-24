package main

import (
	"log"
	"time"
	"github.com/boltdb/bolt"
	"github.com/satori/go.uuid"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	log.Println("opened database")

	for i := 0; i < 20; i++ {
		var key []byte
		if i == 0 {
			key = []byte("hello")
		} else {
			key = []byte(uuid.NewV1().String())
		}
		value := []byte("Hello World!")
		err = db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte("borks"))
			if err != nil {
				return err
			}

			err = bucket.Put(key, value)
			if err != nil {
				return err
			}
			log.Printf("stored key\n")
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("borks"))
		v := bucket.Get([]byte("hello"))
		log.Printf("The answer is: %s\n", v)
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("borks"))

		b.ForEach(func(k, v []byte) error {
			log.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
}
