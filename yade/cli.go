package main

import (
	"fmt"
	"log"

	"github.com/yade/symclassic/hill"
)

func main() {
	cip, err := hill.Encrypt("test.txt", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nCipher text: %s\n", cip)
}
