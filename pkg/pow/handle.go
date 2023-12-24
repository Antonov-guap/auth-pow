package pow

import (
	"crypto/rand"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/scrypt"
)

const (
	saltSize      = 16 // Размер соли для Scrypt
	keyLen        = 32 // Длина ключа для Scrypt
	scryptN       = 8  // Параметр CPU/Mem сложности для Scrypt
	scryptR       = 4  // Параметр размера блока для Scrypt
	scryptP       = 1  // Параметр параллелизма для Scrypt
	requiredZeros = 2  // Количество требуемых начальных нулей в хеше
)

func Handle(rw io.ReadWriter, onSuccess func(w io.Writer) error) error {
	salt, err := generateSalt()
	if err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}

	if _, err = rw.Write(salt); err != nil {
		return fmt.Errorf("failed to send salt to client: %w", err)
	}
	log.Println("salt sent to client:", salt)

	var clientNonce int
	if _, err := fmt.Fscan(rw, &clientNonce); err != nil {
		return fmt.Errorf("failed to read nonce from client: %w", err)
	}
	log.Println("nonce received from client:", clientNonce)

	hashIsValid, err := verifyPoW(salt, clientNonce)
	if err != nil {
		return fmt.Errorf("failed to verify pow: %w", err)
	}

	if !hashIsValid {
		if _, err = fmt.Fprintln(rw, "invalid hash"); err != nil {
			return fmt.Errorf("failed to respond to client that hash is invalid: %w", err)
		}
		log.Println("error sent to client:", salt)
	}

	if err = onSuccess(rw); err != nil {
		return fmt.Errorf("failed to call onSuccess: %w", err)
	}

	return nil
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("failed to read random bytes: %w", err)
	}
	return salt, nil
}

func verifyPoW(salt []byte, clientNonce int) (bool, error) {
	// Вычисляем хеш с использованием Scrypt
	nonceBytes := []byte(fmt.Sprintf("%d", clientNonce))
	hash, err := scrypt.Key(nonceBytes, salt, scryptN, scryptR, scryptP, keyLen)
	if err != nil {
		return false, fmt.Errorf("failed to derive scrypt hash: %w", err)
	}

	// Проверяем, начинается ли хеш с определенного количества нулей
	for i := 0; i < requiredZeros; i++ {
		if hash[i] != 0 {
			return false, nil
		}
	}

	return true, nil
}
