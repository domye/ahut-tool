package utils

import "encoding/base64"

// Base64Encode 对输入字符串进行Base64编码
func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}
