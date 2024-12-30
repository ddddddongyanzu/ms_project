package encrypts

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	plain := "100123213123"
	key := "abcdefgehjhijkmlkjjwwoew"
	cipherByte, err := Encrypt(plain, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ==> %s\n", plain, cipherByte)
	plainText, err := Decrypt(cipherByte, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s ==> %s\n", cipherByte, plainText)
}
