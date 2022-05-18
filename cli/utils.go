package cli

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func GetInput(s string, echo bool) (string, error) {
	fmt.Print(s)
	if echo {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(input), nil
	} else {
		input, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(input)), err
	}
}

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
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	s := make([]rune, n)
	for i := range s {
		ri, _ := rand.Int(rand.Reader, big.NewInt(36))
		s[i] = letters[ri.Int64()]
	}
	return string(s)
}
