package caesar

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// ---------- Encryption -------------
func _encrypt(ptext []byte, key int) []byte {
	return []byte("testing")
}

func Encrpyt(filename, ptext string) ([]byte, error) {
	var k int
	fmt.Print("Enter the key value (1-25): ")
	fmt.Scanf("%d", &k)

	// when file name is provided
	if filename != "" && ptext == "" {
		data, err := os.ReadFile(filename)

		if err != nil {
			log.Panicf("Failed reading data from file %s\n", err)
			return []byte(""), errors.New("Failed to read file data")
		}

		cipher := _encrypt(data, k)
		return cipher, nil
	}

	// when plain text is directly provided
	if filename == "" && ptext != "" {
		plain := []byte(ptext)

		cipher := _encrypt(plain, k)
		return cipher, nil
	}

	return []byte(""), errors.New("Invalid input...")
}
