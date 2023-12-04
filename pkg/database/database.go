package database

import (
	"errors"
	"log"

	"github.com/boltdb/bolt"
)


func CreateKey(db *bolt.DB, key, value string, b ...string) error {
    bucket := "MyBucket"
    if len(b) > 0 {
        bucket = b[0]
    } 
    err := db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucket))
        v := b.Get([]byte(key))
        if len(v) != 0 {
            return errors.New("Key Already Exists")
        }
        err := b.Put([]byte(key), []byte(value))
        return err
    })

    return err
}

func GetKey(db *bolt.DB, key string, b ...string) (string, error) {
    bucket := "MyBucket"
    if len(b) > 0 {
        bucket = b[0]
    } 
    var value string
    err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucket))
        value = string(b.Get([]byte(key)))
        return nil
    })
    if len(value) < 1 {
        return "", errors.New("Key not found")
    }

    return value, err
}

func UpdateKey(db *bolt.DB, key, value string, b ...string) error{
    bucket := "MyBucket"
    if len(b) > 0 {
        bucket = b[0]
    } 
    err := db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucket))
        v := b.Get([]byte(key))
        if len(v) == 0 {
            return errors.New("Key doesn't exist")
        }
        err := b.Put([]byte(key), []byte(value))
        return err
    })

    return err
}

func DeleteKey(db *bolt.DB, key string, b ...string) error{
    bucket := "MyBucket"
    if len(b) > 0 {
        bucket = b[0]
    } 
    err := db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucket))
        v := b.Get([]byte(key))
        if len(v) == 0 {
            return errors.New("Key doesn't exist")
        }
        err := b.Delete([]byte(key))
        if err != nil {
            log.Fatal(err)
        }
        return err
    })

    return err
}

func GetAllKey(db *bolt.DB, b ...string) (map[string]string, error) {
    bucket := "MyBucket"
    if len(b) > 0 {
        bucket = b[0]
    } 
    values := make(map[string]string)
    err := db.View(func(tx *bolt.Tx) error {
        // Assume bucket exists and has keys
        b := tx.Bucket([]byte(bucket))
        c := b.Cursor()
        for k, v := c.First(); k != nil; k, v = c.Next() {
            values[string(k)] = string(v)
        }
        return nil
    })

    return values, err
}

