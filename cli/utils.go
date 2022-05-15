package cli

import (
	"crypto/rand"
	"math/big"
	"os"
)

func fileExisted(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func makeDirIfNotExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// path/to/whatever does not exist
		return os.MkdirAll(path, 0755)
	}
	return err
}

func randomString(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		ri, _ := rand.Int(rand.Reader, big.NewInt(36))
		s[i] = letters[ri.Int64()]
	}
	return string(s)
}
