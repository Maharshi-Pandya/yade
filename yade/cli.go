package main

import (
	"fmt"

	"github.com/yade/symclassic/caesar"
)

func main() {
	cip, _ := caesar.Encrpyt("test.txt", "")

	fmt.Printf("%s\n", cip)
}
