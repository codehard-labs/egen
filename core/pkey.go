package core

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateNewPkey() (common.Address, []byte, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return common.Address{}, nil, err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, nil, errors.New("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, privateKeyBytes, nil
}

func DecryptPkey(encryptedPkey string, aesKey string) ([]byte, error) {
	a, err := hex.DecodeString(encryptedPkey)
	if err != nil {
		return nil, errors.New("cannot decode encryptedPkey")
	}
	b, err := hex.DecodeString(aesKey)
	if err != nil {
		return nil, errors.New("cannot decode aesKey")
	}
	return AESDecrypt(a, b)
}

// This should be deprecated.
// It decrypts the hex pkey string which is UTF-8 decoded
func DecryptPkeyUTF8(encryptedPkey string, aesKey string) ([]byte, error) {
	a, err := hex.DecodeString(encryptedPkey)
	if err != nil {
		return nil, errors.New("cannot decode encryptedPkey")
	}
	b, err := hex.DecodeString(aesKey)
	if err != nil {
		return nil, errors.New("cannot decode aesKey")
	}
	c, err := AESDecrypt(a, b)
	if err != nil {
		return nil, errors.New("AESDecrypt failed")
	}
	return hex.DecodeString(string(c))
}

func GetAddressFromPkey(pkey []byte) (common.Address, error) {
	privateKey, err := crypto.ToECDSA(pkey)
	if err != nil {
		return common.Address{}, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("get publicKeyECDSA failed")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, nil
}
