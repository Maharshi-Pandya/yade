package caesar

// ROT-13 is a special case of caesar cipher where
// the shift/rotation is by 13 letters (key = 13). Moreover,
// since there are 26 letters in the alphabet (13 x 2),
//
// Rot13(Rot13(p)) = p
//
// This property holds true.
func Rot13(filename, plaintext string) ([]byte, error) {
	return Encrpyt(filename, plaintext, 13)
}
