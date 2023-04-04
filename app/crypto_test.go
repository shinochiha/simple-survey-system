package app

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	plaintext := "76efa54cda364c1ca82f6995008d7915"
	encrypted, err := Crypto().Encrypt(plaintext)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	decrypted, err := Crypto().Decrypt(encrypted)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	if decrypted != plaintext {
		t.Errorf("Expected decrypted [%v], got [%v]", plaintext, decrypted)
	}
}
