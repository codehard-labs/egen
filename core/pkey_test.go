package core

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestPkey(t *testing.T) {
	address, pkey, err := GenerateNewPkey()
	if err != nil {
		log.Fatal("Error generating", err)
	}

	fmt.Println("Generated pkey", address.Hex())

	aesKey := NewAESEncryptionKey()
	ciphertext, err := AESEncrypt(pkey, aesKey)
	if err != nil {
		log.Fatal("Error AESEncrypt", err)
	}
	encryptedPkey := hex.EncodeToString(ciphertext)

	fmt.Println("Encrypted", encryptedPkey)
	fmt.Println("AESKEY", hex.EncodeToString(aesKey))

	//

	pkey1, err := DecryptPkey(encryptedPkey, hex.EncodeToString(aesKey))
	if err != nil {
		log.Fatal("Error DecryptPkey", err)
	}

	address1, err := GetAddressFromPkey(pkey1)
	if err != nil {
		log.Fatal(err)
	}

	if address != address1 {
		log.Fatal("not equal")
	}
}
