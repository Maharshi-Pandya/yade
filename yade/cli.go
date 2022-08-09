package main

import (
	"fmt"

	"github.com/yade/symclassic/caesar"
)

func main() {
	cip, _ := caesar.Rot13("", "Hello bro!")

	fmt.Printf("%s\n", cip)
}
