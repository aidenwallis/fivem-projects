package backend

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
)

const tokenDictionary = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var tokenDictionaryLength = big.NewInt(int64(len(tokenDictionary)))

// RandomToken generates a crytographically random token.
func RandomToken(length int) (string, error) {
	bs := make([]byte, length)
	for i := range bs {
		randIndex, err := rand.Int(rand.Reader, tokenDictionaryLength)
		if err != nil {
			return "", err
		}
		bs[i] = tokenDictionary[randIndex.Uint64()]
	}
	return string(bs), nil
}

func hashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return base64.RawURLEncoding.EncodeToString(h[:])
}
