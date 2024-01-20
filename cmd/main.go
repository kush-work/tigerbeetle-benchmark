package main

import (
	_ "crypto/elliptic"
	"github.com/kushthedude/tigerbeetle-benchmark/internal/router"
)

func main() {
	serverObject := router.NewRouter(3001)
	err := serverObject.RunServer()
	if err != nil {
		panic("failed to run server")
	}
}
