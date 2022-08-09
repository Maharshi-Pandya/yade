package main

import (
	"log"

	"github.com/yade/symclassic/hill"
	"github.com/yade/utils"
)

func main() {
	// cip, _ := caesar.Encrpyt("", "Hello bro!", -1)
	mat, err := hill.GetKeyMatrix(1)

	if err != nil {
		log.Panic(err)
	}

	utils.MatPrint(mat)
}
