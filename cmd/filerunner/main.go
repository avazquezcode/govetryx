package main

import (
	"log"
	"os"

	"github.com/avazquezcode/govetryx/internal/adapter"
)

func main() {
	err := adapter.RunFile(os.Args[1], os.Stdout)
	if err != nil {
		log.Fatalf("failed interpreting the script: %s", err.Error())
	}
}
