package bbolt

import (
	"fmt"
	"github.com/urfave/cli/v2"
	bolt "go.etcd.io/bbolt"
)

func Get(ctx *cli.Context) error {
	path := ctx.String("path")
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		return err
	}

	bucket := ctx.String("bucket")
	k := ctx.String("key")

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(k))
		fmt.Printf("%s\n", v)
		return nil
	})

	defer db.Close()

	return nil
}
