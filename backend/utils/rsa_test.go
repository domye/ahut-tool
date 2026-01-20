package utils

import (
	"testing"
)

func TestEncryptPasswordWithRSA(t *testing.T) {
	// 测试正常情况
	encrypted, err := EncryptPasswordWithRSA("3tmwjdaW")
	if err != nil {
		t.Errorf("EncryptPasswordWithRSA failed: %v", err)
	}
	if encrypted == "" {
		t.Error("Expected non-empty encrypted string")
	}
}
