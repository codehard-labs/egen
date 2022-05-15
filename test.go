package main

import (
	"encoding/hex"
	"fmt"

	"github.com/codehard-labs/egen/core"
)

func main() {
	address, pkey, err := core.GenerateNewPkey()
	if err != nil {
		fmt.Println("Error generating", err)
	}

	fmt.Println("Generated pkey", address.Hex())

	aesKey := core.NewAESEncryptionKey()
	ciphertext, err := core.AESEncrypt(pkey, aesKey)
	if err != nil {
		fmt.Println("Error AESEncrypt", err)
	}
	encryptedPkey := hex.EncodeToString(ciphertext)

	fmt.Println("Encrypted", encryptedPkey)
	fmt.Println("AESKEY", hex.EncodeToString(aesKey))

	//

	pkey1, err := core.DecryptPkey(encryptedPkey, hex.EncodeToString(aesKey))
	if err != nil {
		fmt.Println("Error DecryptPkey", err)
	}
	fmt.Println(core.GetAddressFromPkey(pkey1))
}
