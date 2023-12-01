package main

import "fmt"

// Caesar Cipher Encryption Function
func caesarEncrypt(plainText string, shift int) string {
	cipherText := ""

	for _, char := range plainText {
		if char >= 'A' && char <= 'Z' {
			cipherText += string((char+rune(shift)-'A')%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			cipherText += string((char+rune(shift)-'a')%26 + 'a')
		} else {
			cipherText += string(char)
		}
	}

	return cipherText
}

// Caesar Cipher Decryption Function
func caesarDecrypt(cipherText string, shift int) string {
	plainText := ""

	for _, char := range cipherText {
		if char >= 'A' && char <= 'Z' {
			plainText += string((char-rune(shift)-'A')%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			plainText += string((char-rune(shift)-'a')%26 + 'a')
		} else {
			plainText += string(char)
		}
	}

	return plainText
}

func main() {
	plainText := "HELLO WORLD"
	shift := 3

	cipherText := caesarEncrypt(plainText, shift)

	fmt.Printf("Plaintext: %s\n", plainText)
	fmt.Printf("Shift: %d\n", shift)
	fmt.Printf("Ciphertext: %s\n", cipherText)

	plainText = caesarDecrypt(cipherText, shift)
	fmt.Printf("Plaintext: %s\n", plainText)
}
