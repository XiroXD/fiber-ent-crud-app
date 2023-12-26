package hash

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

func CreateHash(password string) []byte {
	salt, err := randomBytes(16)
	if err != nil {
		panic(err)
	}

	hash := argon2.Key([]byte(password), salt, 3, 64*1024, 4, 32)

	return hash
}

func randomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
