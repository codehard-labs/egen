package cli

import (
	"encoding/hex"
	"log"
	"os"

	"github.com/codehard-labs/egen/core"
)

var (
	BasePath    = ".keys/"
	AESKeyPath  = BasePath + "aes/"
	EncPKeyPath = BasePath + "enc-pkey/"
)

func init() {
	makeDirIfNotExists(AESKeyPath)
}

func GenerateNewAESKey() {
	name, err := GetInput("Please name the new AES key [leave empty for random name]: \n", true)
	if err != nil {
		log.Fatal(err)
	}
	if name == "" {
		name = randomString(16)
	}
	if fileExisted(AESKeyPath + name + ".key") {
		log.Fatal("key file with a same name already exists")
	}
	aesKey := core.NewAESEncryptionKey()
	f, err := os.Create(AESKeyPath + name + ".key")
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(hex.EncodeToString(aesKey))
}

func GenerateNewPkeyWithLocalAESKey() (string, string) {
	aesKey := ReadLocalHexAESKey()
	addr, pkey, err := core.GenerateNewPkey()
	if err != nil {
		log.Fatal(err)
	}
	encPkey, err := core.AESEncrypt(pkey, aesKey)
	if err != nil {
		log.Fatal(err)
	}
	return addr.Hex(), hex.EncodeToString(encPkey)
}

func ReadLocalHexAESKey() []byte {
	path, err := GetInput("Please insert path to the aes key: \n", true)
	if err != nil {
		log.Fatal(err)
	}
	key, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	aesKey, err := hex.DecodeString(string(key))
	if err != nil {
		log.Fatal(err)
	}
	return aesKey
}
