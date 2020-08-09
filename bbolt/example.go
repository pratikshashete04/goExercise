package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var example = []byte("Example")

func main() {

	// opening newdata.db
	// it will create if not exist
	db, err := bolt.Open("C:/Users/ABC/Desktop/bbolt/newdata.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// assigning key and value
	key := []byte("Practice")
	value := []byte("Practice Example!")

	// storing data in database
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(example)
		if err != nil {
			return err
		}

		// bucket.Put will insert the value in the database
		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	// checking if error is there
	if err != nil {
		log.Fatal(err)
	}

	// retrieving the data from database
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(example)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", example)
		}

		// bucket.Get will give/show the value from the database
		val := bucket.Get(key)

		// printing the string output
		fmt.Println(string(val))
		return nil
	})
}
