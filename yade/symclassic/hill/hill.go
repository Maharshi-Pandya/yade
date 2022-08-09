/*
Hill Encryption technique is a polygraphic substitution cipher
based on Linear Algebra.

The user will provide with a plaintext phrase, and the information
regarding the Key matrix (square matrix). Hill cipher is:

C = (P * K) mod 26

where,
C = cipher matrix
P = plaintext blocks (vector)
K = key matrix

Decryption works the same way, just the Key matrix is inverted and multiplied
with blocks of Cipher text.

P = (IK * C) mod 26

where,
IK = inverse of Key matrix
C = ciphertext blocks (vector)
*/

package hill

import (
	"errors"
	"fmt"

	"github.com/yade/utils"
)

// Let the user enter the Key
func userInputKey(keysize int) ([][]byte, error) {
	return utils.MatCreateSquare(keysize)
}

// // Generate random Key matrix from dimensions
// func generateRandomKey(matsize int) ([][]byte, error) {

// }

// Gets the Key matrix by generating it randomly or taking user input
// mode 0: Generate randomly
// mode 1: User input
func GetKeyMatrix(mode int) ([][]byte, error) {
	var n int
	fmt.Print("Provide Key matrix size (n x n): ")
	fmt.Scanf("%d", &n)

	switch mode {
	// case 0:
	// 	return generateRandomKey(n)
	case 1:
		return userInputKey(n)
	}

	return nil, errors.New("Invalid input provided...")
}
