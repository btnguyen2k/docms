package gvabe

import (
	"crypto/rand"
	"crypto/rsa"
)

// available since template-v0.2.0
func genRsaKey(numBits int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, numBits)
}
