/*
Hill cipher is a polygraphic substitution algorithm based on Linear Algebra.
The algorithm is...

C = (P * K) mod n

where:
C = cipher text block (vector)
P = plain text block (vector)
K = square matrix (Key)
n = number of symbols to be used

The symbol space for this implemention of Hill Cipher contains all the uppercase
alphabets (A - Z) and 3 extra symbols:

- SPACE (' ')
- Period ('.')
- Question Mark ('?')

Making n = 29 (Prime) for this implementation.
Decryption is just the reverse of encryption:

P = (IK * C) mod n

where,
IK = inverse of Key matrix
C = cipher text block (vector)
P = plain text block (vector)
*/

package hill

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/yade/utils"
)

// some constants
const (
	SPACE_BYTE    = 26
	PERIOD_BYTE   = 27
	QUESTION_BYTE = 28
	MOD           = 29
	A_BYTE        = 65
	Aa_DIFF       = 32
)

// Return the alphabet list
func _alpha() []byte {
	return []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O',
		'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', ' ', '.', '?'}
}

// Helper function to create a byte square matrix of size N.
// Input provided by the user.
func matCreateSquare(keyText string, n int) ([][]byte, error) {
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
			// for 3 extra symbols
			if bKeyText[pos] == ' ' {
				mat[i][j] = SPACE_BYTE
			} else if bKeyText[pos] == '.' {
				mat[i][j] = PERIOD_BYTE
			} else if bKeyText[pos] == '?' {
				mat[i][j] = QUESTION_BYTE
			} else {
				mat[i][j] -= A_BYTE
			}
			pos++
		}
	}

	return mat, nil
}

// Let the user enter the Key
func userInputKey(keysize int) ([][]byte, error) {
	var keyText string
	fmt.Printf("Provide the Key matrix text of size %d (no spaces): ", keysize)
	fmt.Scanf("%s", &keyText)

	return matCreateSquare(keyText, keysize)
}

// Generate random Key matrix from dimensions
func generateRandomKey(matsize int) ([][]byte, error) {
	fmt.Printf("Generating random Key matrix of size %d:\n", matsize)
	return utils.MatGenerateSquare(MOD, matsize)
}

// Gets the Key matrix by generating it randomly or taking user input
// mode 0: Generate randomly
// mode 1: User input
func GetKeyMatrix(mode int) ([][]byte, error) {
	var n int
	fmt.Print("Provide Key matrix size (n x n): ")
	fmt.Scanf("%d", &n)

	switch mode {
	case 0:
		return generateRandomKey(n)
	case 1:
		return userInputKey(n)
	default:
		log.Panicf("Invalid mode %d provided as input...", mode)
	}

	return nil, errors.New("Invalid input provided...")
}

// ---------------- Encryption -------------------
func _encrypt(ptext []byte, key [][]byte) ([]byte, error) {
	// use the square matrix row count as key size
	keySize := len(key)

	// --- PREPROCESSING
	// Modify plaintext to uppercase and be divisible by key size
	var pTextMod []byte = ptext

	// ------ Upper case
	for i := range pTextMod {
		if pTextMod[i] >= 97 && pTextMod[i] <= 122 {
			pTextMod[i] -= Aa_DIFF
		}
		if pTextMod[i] == ' ' {
			pTextMod[i] = SPACE_BYTE
		} else if pTextMod[i] == '.' {
			pTextMod[i] = PERIOD_BYTE
		} else if pTextMod[i] == '?' {
			pTextMod[i] = QUESTION_BYTE
		} else {
			pTextMod[i] -= A_BYTE
		}
	}

	// ------ When ptext length is less than key size or it is not divisible,
	// ------ append extra 'X'
	lptm := len(pTextMod)
	diff := 0
	if lptm < keySize {
		diff = keySize - lptm
	} else if lptm%keySize != 0 {
		diff = keySize*(int(lptm/keySize)+1) - lptm
	}
	for i := 0; i < diff; i++ {
		pTextMod = append(pTextMod, byte(24))
	}

	// fmt.Println("Plain text in byte is: ", pTextMod)

	// --- ENCRYPTION
	var cipher []byte

	for i := 0; i < len(pTextMod); i += keySize {
		tempVec := pTextMod[i : i+keySize]

		// perform matrix vector multiplication
		tempEnc, err := utils.MatProdVec(key, tempVec)
		if err != nil {
			return nil, err
		}
		// modulus 29 (coz of 3 extra symbols)
		tempEnc1, err := utils.VecModConst(tempEnc, MOD)
		if err != nil {
			return nil, err
		}

		// append to cipher
		cipher = append(cipher, tempEnc1...)
	}

	// fmt.Println("Cipher text in byte: ", cipher)

	// --- POST PROCESSING (Make cipher a byte array)
	finalcip := make([]byte, len(cipher))

	for i := range finalcip {
		finalcip[i] = byte(_alpha()[cipher[i]])
	}

	return finalcip, nil
}

func Encrypt(filename, plaintext string) ([]byte, error) {
	// TODO: Key matrix creation or generation here
	var m int
	fmt.Println("\nThere are two options...")
	fmt.Print("0: Generate a random Key matrix\n")
	fmt.Print("1: Input a Key matrix text yourself\n")
	fmt.Print("\nSelect one mode (0 / 1): ")

	fmt.Scanf("%d", &m)

	// got the Key matrix
	keyMat, err := GetKeyMatrix(m)
	if err != nil {
		return nil, err
	}

	utils.MatPrint("Key Matrix", keyMat)

	// When file is provided
	if filename != "" && plaintext == "" {
		data, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		// TODO: use data to encrypt
		cip, err := _encrypt(data, keyMat)
		if err != nil {
			return nil, err
		}
		return cip, nil
	}

	// When plain text is provided
	if filename == "" && plaintext != "" {
		data := []byte(plaintext)

		cip, err := _encrypt(data, keyMat)
		if err != nil {
			return nil, err
		}
		return cip, nil
	}

	return nil, errors.New("Invalid input provided...")
}
