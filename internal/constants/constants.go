package constants

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	tigerbeetle_go "github.com/tigerbeetle/tigerbeetle-go"
	tigerbeetle_type "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

const (
	TigerbeetleServerAddress = "127.0.0.1:3000"
	AccountTypeCode          = uint16(718)
	AccountFlags             = uint16(0)
	ClusterID                = 0
)

var PrivateKey *ecdsa.PrivateKey
var TigerBeetleClient tigerbeetle_go.Client

func InitPrivateKey() {
	PrivateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func GetPrivateKey() *ecdsa.PrivateKey {
	return PrivateKey
}

func GetPublicKey() ecdsa.PublicKey {
	return PrivateKey.PublicKey
}

func InitTigerBeetleClient() {
	client, err := tigerbeetle_go.NewClient(tigerbeetle_type.ToUint128(ClusterID), []string{TigerbeetleServerAddress}, uint(1000))
	if err != nil {
		panic("failed to init client")
	}
	TigerBeetleClient = client
}
