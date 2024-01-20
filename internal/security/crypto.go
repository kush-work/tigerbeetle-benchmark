package security

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	. "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

// SignTransaction signs the transaction with the sender's private key.
func SignTransaction(tx *Transfer, privateKey *ecdsa.PrivateKey) (string, error) {
	// Serialize the transaction into a string or byte slice for signing
	signatureData := fmt.Sprintf("%s:%s:%d:%d", tx.DebitAccountID, tx.CreditAccountID, tx.Amount, tx.Ledger)
	hash := sha256.Sum256([]byte(signatureData))

	// Sign the hash with the private key
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	// Combine the two big.Ints into a slice of bytes
	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	return hex.EncodeToString(signature), nil
}

// VerifySignature verifies that the transaction was signed by the sender.
func VerifySignature(tx *Transfer, signatureHex string, publicKey *ecdsa.PublicKey) (bool, error) {
	signature, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false, err
	}

	// Recover r and s values from the signature
	r := new(big.Int).SetBytes(signature[:len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])

	// Serialize the transaction the same way as during signing
	signatureData := fmt.Sprintf("%s:%s:%d:%d", tx.DebitAccountID, tx.CreditAccountID, tx.Amount, tx.Ledger)
	hash := sha256.Sum256([]byte(signatureData))

	// Verify the signature
	return ecdsa.Verify(publicKey, hash[:], r, s), nil
}
