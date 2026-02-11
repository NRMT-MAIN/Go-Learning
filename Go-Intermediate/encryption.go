package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func encrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func encryptionExample() {
	key := make([]byte, 32)
    if _, err := rand.Read(key); err != nil {
        panic(err)
    }
    
    plaintext := []byte("Secret message to encrypt")
    
    // Encrypt
    ciphertext, err := encrypt(plaintext, key)
    if err != nil {
        fmt.Println("Encryption error:", err)
        return
    }
    fmt.Println("Encrypted:", hex.EncodeToString(ciphertext))
    
    // Decrypt
    decrypted, err := decrypt(ciphertext, key)
    if err != nil {
        fmt.Println("Decryption error:", err)
        return
    }
    fmt.Println("Decrypted:", string(decrypted))
}