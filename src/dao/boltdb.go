package dao

import (
	"github.com/boltdb/bolt"
)

func Update(bucket string, key string, value string) {

	db := getDb()

	err := db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte(bucket))

		if b != nil {
			err := b.Put([]byte(key), []byte(value))
			if err != nil {
				panic(err)
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	db.Close()
}

func View(bucket string, key string) string {
	var value string
	db := getDb()

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b != nil {
			value = string(b.Get([]byte(key)))
		}
		return nil
	})

	db.Close()

	if err != nil {
		panic(err)
	}

	return value
}

func Delete(bucket string, key string) {
	db := getDb()

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		b.Writable()
		if b != nil {
			err := b.Delete([]byte(key))
			return err
		}
		return nil
	})

	db.Close()

	if err != nil {
		panic(err)
	}
}

func List(bucket string, pageNum int, pageSize int) ([]string, int) {
	var list []string
	db := getDb()

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b != nil {
			b.ForEach(func(k, v []byte) error {
				list = append(list, string(v))
				return nil
			})
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	db.Close()

	total := len(list)

	var end int

	if total < pageNum*pageSize {
		end = total
	} else {
		end = pageNum * pageSize
	}

	list = list[(pageNum-1)*pageSize : end]

	return list, total
}

func getDb() *bolt.DB {

	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		panic("数据库连接出错")
	}
	return db
}
