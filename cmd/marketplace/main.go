package main

import (
	"barbz.dev/marketplace/cmd/marketplace/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
