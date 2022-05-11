package main

import (
	"context"
	"os"
	"time"

	"github.com/athoune/chevillette/input/loki"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	u := os.Args[1]
	l, err := loki.NewLoki(u)
	if err != nil {
		panic(err)
	}
	err = l.Tail(context.TODO(), `{job="nginx"}`,
		5*time.Second, 42, time.Now().Add(-15*time.Minute), func(data *loki.Tail) error {
			spew.Dump(data)
			return nil
		})
	if err != nil {
		panic(err)
	}
}
