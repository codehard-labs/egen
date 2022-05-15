package main

import (
	"fmt"

	"github.com/codehard-labs/egen/cli"
)

func main() {
	//cli.GenerateNewAESKey()
	fmt.Println(cli.GenerateNewPkeyWithLocalAESKey())
}
