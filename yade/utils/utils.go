package utils

import (
	"errors"
	"fmt"
)

// Helper function to print a byte matrix in character form
func MatPrint(mat [][]byte) {
	for i := range mat {
		fmt.Println()
		for j := range mat[0] {
			fmt.Printf("%c ", mat[i][j])
		}
	}

	fmt.Println()
}

// Helper function to create a byte square matrix of size N.
// Input provided by the user.
func MatCreateSquare(n int) ([][]byte, error) {
	var keyText string
	fmt.Printf("Provide the Key matrix text of size %d (no spaces): ", n)
	fmt.Scanf("%s", &keyText)

	if len(keyText) != n*n {
		return nil, errors.New(fmt.Sprintf("Invalid Key length for square matrix of size %d", n))
	}

	// init matrix
	bKeyText := []byte(keyText)
	mat := make([][]byte, n)
	for i := range mat {
		mat[i] = make([]byte, n)
	}

	// build it
	pos := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = bKeyText[pos]
			pos++
		}
	}

	return mat, nil
}
