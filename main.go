package main

import (
	"github.com/deqing/micro-batching"
	"github.com/deqing/micro-batching-service/internal"

	"log"
	"os"
)

func main() {
	b, err := batching.NewBatching()
	if err != nil {
		path, _ := os.Getwd()
		log.Fatalf("creating batching service failed: %v in %s", err, path)
		return
	}

	internal.SetupHandler(&b)
}
