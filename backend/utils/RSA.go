package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

// RSA公钥
const RSAPublicKey = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCCCUg7rT5UBlDcqoISt9PR/p1qaf2Tj+qZgzV/J764hBJAinMcOGWlcTkGlcL69P8waHti4HsOYYo4Tk5Fx9dqHzEtJha/BtcFUysD/BKiyeJfMyWNMNlgggghG5BuY2M3AYY8qII1Q7xCN6XuQb4pAYJ8qVmIqqAqRvyFA0y4vQIDAQAB`

// ParseHardcodedPublicKey 解析硬编码的公钥
func ParseHardcodedPublicKey(pubKeyStr string) (*rsa.PublicKey, error) {
	// 将字符串包装成PEM格式
	pemData := fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", FormatPEMString(pubKeyStr))

	// 解码PEM
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	// 解析公钥
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return rsaPubKey, nil
}

// FormatPEMString 格式化PEM字符串（每行64个字符）
func FormatPEMString(s string) string {
	const lineLen = 64
	if len(s) <= lineLen {
		return s
	}

	result := make([]byte, 0, len(s)+(len(s)/lineLen)+1)
	for i, b := range []byte(s) {
		if i > 0 && i%lineLen == 0 {
			result = append(result, '\n')
		}
		result = append(result, b)
	}
	return string(result)
}

// EncryptPasswordWithRSA 使用RSA加密密码
func EncryptPasswordWithRSA(password string) (string, error) {
	// 首先对密码进行Base64编码
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(password))

	// 解析公钥
	publicKey, err := ParseHardcodedPublicKey(RSAPublicKey)
	if err != nil {
		return "", err
	}

	// 使用RSA公钥加密Base64编码后的密码
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(base64Encoded))
	if err != nil {
		return "", err
	}

	// 将加密后的字节数组转为Base64字符串返回
	print(base64.StdEncoding.EncodeToString(encryptedBytes))
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}
