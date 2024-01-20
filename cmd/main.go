package main

import (
	_ "crypto/elliptic"
	"fmt"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/constants"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/tigerbeetle"
)

func main() {
	constants.InitPrivateKey()

	// Assume these are constants you've defined for your application:

	success, err := tigerbeetle.CreateAccount(5, 3)
	if !success || err != nil {
		print(success)
		print(err)
		panic("failed in create accoint")
	}
	success, err = tigerbeetle.CreateAccount(6, 3)
	if !success || err != nil {
		print(success)
		print(err)
		panic("failed in create accoint")
	}

	success, err = tigerbeetle.PostCredits(1, 2, 100, 3)
	if !success || err != nil {
		panic("failed in create credit")
	}
	fmt.Println("PASS")
}
