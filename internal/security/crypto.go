package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Clave estática para el cifrado AES-256 (debe ser de 32 bytes).
// En un entorno real esto iría en una variable de entorno.
var key = []byte("platinum-gym-secret-key-12345678")

// EncryptDeterministic cifra un texto usando AES-GCM con un nonce determinista.
// Esto permite que el resultado siempre sea el mismo para el mismo texto,
// preservando las restricciones UNIQUE de la base de datos (ej. para el DNI).
func EncryptDeterministic(plaintext string) (string, error) {
	if plaintext == "" {
		return "", nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Creamos un nonce determinista a partir del hash del texto
	hash := sha256.Sum256([]byte(plaintext))
	nonce := hash[:aesgcm.NonceSize()]

	// Añadimos el nonce al principio del texto cifrado para poder descifrarlo luego
	ciphertext := aesgcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(ciphertext), nil
}

// DecryptDeterministic descifra un texto cifrado con EncryptDeterministic.
func DecryptDeterministic(cryptoText string) (string, error) {
	if cryptoText == "" {
		return "", nil
	}
	data, err := hex.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesgcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("texto cifrado demasiado corto")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
