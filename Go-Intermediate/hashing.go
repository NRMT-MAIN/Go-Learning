package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashingExample() {
	data := "Hello , Go"

	hash := sha256.Sum256([]byte(data))

	hashString := hex.EncodeToString(hash[:])

	fmt.Println(hashString)

	hash512 := sha512.Sum512([]byte(data))
    fmt.Println("SHA-512:", hex.EncodeToString(hash512[:]))
    
    // SHA-384 (truncated SHA-512)
    hash384 := sha512.Sum384([]byte(data))
    fmt.Println("SHA-384:", hex.EncodeToString(hash384[:]))

    password := "my_secure_password"
	hashedPassword, err := hashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	fmt.Println("Hashed Password:", hashedPassword)

	isValid := verifyPassword(password, hashedPassword)
	fmt.Println("Password valid:", isValid)
	
}

func hashPassword(password string) (string , error) {
	hash , err :=  bcrypt.GenerateFromPassword([]byte(password) , bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error in password hashing")
		return "" , err
	}
	return string(hash) , nil
}

func verifyPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}