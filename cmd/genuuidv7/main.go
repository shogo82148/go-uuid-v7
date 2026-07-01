package main

import (
	"flag"
	"fmt"

	gouuidv7 "github.com/shogo82148/go-uuid-v7"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 1, "number of UUIDs to generate")
	flag.Parse()

	ids := make([]gouuidv7.UUID, 0, n)
	for i := 0; i < n; i++ {
		ids = append(ids, gouuidv7.NewV7())
	}
	for _, id := range ids {
		fmt.Println(id.String())
	}
}
