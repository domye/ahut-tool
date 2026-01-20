package pay

import (
	"ahut-tool/backend/utils"
	"testing"
)

func TestEncryptPasswordWithRSA(t *testing.T) {
	password := "test123"
	encrypted, err := utils.EncryptPasswordWithRSA(password)

	if err != nil {
		t.Errorf("Error encrypting password: %v", err)
	}

	if encrypted == "" {
		t.Error("Encrypted password is empty")
	}

	t.Logf("Original: %s, Encrypted: %s", password, encrypted)
}

//func TestSendLoginRequest(t *testing.T) {
//	service := NewService()
//
//	// 这里使用测试账号进行测试（实际使用时请替换为真实凭据）
//	username := "249074506"
//	password := "3tmwjdaW"
//
//	//statusCode := service.GetToken(username, password)
//
//	t.Logf("Login response status code: %d", statusCode)
//	// 注意：由于我们没有有效的测试凭据，这里可能收到401或其他错误状态码是正常的
//}
