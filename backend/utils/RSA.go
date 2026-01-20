package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// RSA公钥
const RSAPublicKey = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCCCUg7rT5UBlDcqoISt9PR/p1qaf2Tj+qZgzV/J764hBJAinMcOGWlcTkGlcL69P8waHti4HsOYYo4Tk5Fx9dqHzEtJha/BtcFUysD/BKiyeJfMyWNMNlgggghG5BuY2M3AYY8qII1Q7xCN6XuQb4pAYJ8qVmIqqAqRvyFA0y4vQIDAQAB`

// ParseHardcodedPublicKey 解析硬编码的公钥
func ParseHardcodedPublicKey(pubKeyStr string) (*rsa.PublicKey, error) {
	// 直接解码Base64公钥
	pubBytes, err := base64.StdEncoding.DecodeString(pubKeyStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key: %v", err)
	}

	// 解析公钥
	pubKey, err := x509.ParsePKIXPublicKey(pubBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return rsaPubKey, nil
}

// EncryptPasswordWithRSA 使用RSA加密密码
func EncryptPasswordWithRSA(password string) (string, error) {
	// 解析公钥
	publicKey, err := ParseHardcodedPublicKey(RSAPublicKey)
	if err != nil {
		return "", err
	}

	// 对密码进行Base64编码后使用RSA公钥加密
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(password))
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(base64Encoded))
	if err != nil {
		return "", err
	}

	// 将加密后的字节数组转为Base64字符串返回
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}
