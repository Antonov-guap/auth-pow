package pow

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"strconv"

	"golang.org/x/crypto/scrypt"
)

func Read(rw io.ReadWriter, onSuccess func(r io.Reader) error) error {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return fmt.Errorf("unable to read random bytes: %w", err)
	}

	if _, err := rw.Read(salt); err != nil {
		return fmt.Errorf("failed to read salt from server: %w", err)
	}
	log.Println("salt is read from server:", salt)

	nonce, err := findNonce(salt)

	if _, err := fmt.Fprintf(rw, "%d\n", nonce); err != nil {
		return fmt.Errorf("failed to send nonce to server: %w", err)
	}
	log.Println("nonce is sent to server:", nonce)

	err = onSuccess(rw)
	if err != nil {
		return fmt.Errorf("failed to read response from server: %w", err)
	}

	return nil
}

func findNonce(salt []byte) (int, error) {

	for nonce := 0; ; nonce++ {
		nonceBytes := []byte(strconv.Itoa(nonce))

		hash, err := scrypt.Key(nonceBytes, salt, scryptN, scryptR, scryptP, keyLen)
		if err != nil {
			return 0, fmt.Errorf("failed to generate scrypt hash: %w", err)
		}

		validHash := true
		for i := 0; i < requiredZeros; i++ {
			if hash[i] != 0 {
				validHash = false
				break
			}
		}

		if validHash {
			return nonce, nil
		}
	}
}
