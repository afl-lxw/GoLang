package utils

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// 密码加密
func PasswordHash(pwd string) (string, error) {
	// GenerateFromPassword 方法对密码进行加密操作
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func RandSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return salt, err
	}
	fmt.Println("---- salt ----%v", salt)

	return salt, err
}

// 密码验证
func PasswordVerify(pwd, hash string) bool {
	// CompareHashAndPassword 方法将加密的数据与原始数据进行对比
	err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(hash))

	return err == nil
}
