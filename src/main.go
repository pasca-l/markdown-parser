package main

import (
	"log"

	"github.com/pasca-l/markdown-parser/server"
)

func main() {
	err := server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
