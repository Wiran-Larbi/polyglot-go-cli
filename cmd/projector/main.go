package main

import (
	"fmt"
	"log"
	"polyglot-go-cli/pkg/projector"
)

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		log.Fatal("unable to get options %v", err)
	}
	fmt.Printf("opts: %+v", opts)
}
