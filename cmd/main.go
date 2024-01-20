package main

import (
	_ "crypto/elliptic"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/constants"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/nonce"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/router"
)

func main() {
	nonce.NewNonceManager()
	constants.InitPrivateKey()
	constants.InitTigerBeetleClient()
	serverObject := router.NewRouter(3001)
	err := serverObject.RunServer()
	if err != nil {
		panic("failed to run server")
	}
}
