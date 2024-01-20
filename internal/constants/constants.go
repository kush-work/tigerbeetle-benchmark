package constants

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

const (
	TigerbeetleServerAddress = "127.0.0.1:3000"
	AccountTypeCode          = uint16(718)
	AccountFlags             = uint16(0)
	ClusterID                = 0
)

var PrivateKey *ecdsa.PrivateKey

func InitPrivateKey() {
	PrivateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func GetPrivateKey() *ecdsa.PrivateKey {
	return PrivateKey
}

func GetPublicKey() ecdsa.PublicKey {
	return PrivateKey.PublicKey
}
