package main

import (
	"flag"
	"log"

	"github.com/pathcl/go-password/password"
)

func main() {

	length := flag.Int("n", 20, "an int between 20 and 64")
	flag.Parse()

	res, err := password.Generate(*length, 10, 10, false, false)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf(res)
}
