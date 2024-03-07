package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"math/big"
	mathRand "math/rand"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword for generate random password
func GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var (
		lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
		upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		specialCharSet = "!@#$%&*"
		numberSet      = "0123456789"
		allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
	)
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(specialCharSet))))
		password.WriteString(string(specialCharSet[random.BitLen()]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(numberSet))))
		password.WriteString(string(numberSet[random.BitLen()]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(upperCharSet))))
		password.WriteString(string(upperCharSet[random.BitLen()]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(allCharSet))))
		password.WriteString(string(allCharSet[random.BitLen()]))
	}
	inRune := []rune(password.String())
	mathRand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

// CipherEncrypt for encrypt data with AES algorithm
func CipherEncrypt(plaintext, key string) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err == nil {
		gcm, err := cipher.NewGCM(c)
		if err == nil {
			nonce := make([]byte, gcm.NonceSize())
			if _, err = io.ReadFull(rand.Reader, nonce); err == nil {
				return gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
			}
		}
	}

	return nil, err
}

// CipherDecrypt for decrypt data with AES algorithm
func CipherDecrypt(ciphertext []byte, key string) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err == nil {
		gcm, err := cipher.NewGCM(c)
		if err == nil {
			nonceSize := gcm.NonceSize()
			if len(ciphertext) < nonceSize {
				return nil, errors.New("ciphertext too short")
			}
			nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
			return gcm.Open(nil, nonce, ciphertext, nil)
		}
	}
	return nil, err
}

// RandomChars func
func RandomChars(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[mathRand.Intn(len(letters))]
	}
	return string(b)
}

// RandomNumber generates a unique number of the specified length
func RandomNumber(length int) string {
	characters := "0123456789"
	code := ""
	for i := 0; i < length; i++ {
		code += string(characters[mathRand.Intn(len(characters))])
	}
	return code
}

// HashPassword func
func HashPassword(password string) (string, error) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), nil
}

// CheckPasswordHash func
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
