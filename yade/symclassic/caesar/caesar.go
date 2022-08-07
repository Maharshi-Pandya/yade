package caesar

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// ---------- Encryption -------------
func _encrypt(ptext []byte, key int) []byte {
	var cipher []byte

	for _, ch := range ptext {
		var tor byte
		if ch >= 65 && ch <= 90 {
			tor = (ch+byte(key)-65)%26 + 65
		} else if ch >= 97 && ch <= 122 {
			tor = (ch+byte(key)-97)%26 + 97
		} else {
			tor = ch
		}

		cipher = append(cipher, tor)
	}
	return cipher
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
