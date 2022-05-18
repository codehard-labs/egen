package main

import (
	"fmt"

	"github.com/codehard-labs/egen/cli"
)

func main() {
	//cli.GenerateNewAESKey()
	//cli.GenerateNewPkeyWithLocalAESKey("test")
	fmt.Println(cli.VerifyLocalPkeyWithLocalAESKey(
		"0x95f397cFE62Fbcc40C30D3D7870E772267E9bE7b",
		"test",
	))
}
