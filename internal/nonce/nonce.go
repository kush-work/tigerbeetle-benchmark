package nonce

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
)

// NonceManager keeps track of used nonces to prevent replay attacks.
type NonceManager struct {
	usedNonces map[string]bool
	mu         sync.Mutex
}

var NonceInstance *NonceManager

func NewNonceManager() {
	NonceInstance = &NonceManager{
		usedNonces: make(map[string]bool),
	}
}

func (m *NonceManager) GenerateNonce() (string, error) {
	nonceBytes := make([]byte, 16) // 128-bit nonce
	if _, err := rand.Read(nonceBytes); err != nil {
		return "", err
	}
	nonce := hex.EncodeToString(nonceBytes)
	m.mu.Lock()
	m.usedNonces[nonce] = false
	m.mu.Unlock()
	return nonce, nil
}

func (m *NonceManager) VerifyNonce(nonce string) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if nonce is already used
	if used, exists := m.usedNonces[nonce]; !exists || used {
		return false, errors.New("nonce is used")
	}

	// Mark the nonce as used
	m.usedNonces[nonce] = true
	return true, nil
}
