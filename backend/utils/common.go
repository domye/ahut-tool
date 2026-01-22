package utils

import (
	"encoding/base64"
	"regexp"
	"strconv"
	"strings"
)

// 预编译正则表达式以提高性能
var (
	htmlTagRegexp = regexp.MustCompile(`<[^>]*>`)
	whitespaceRe  = regexp.MustCompile(`\s+`)
)

// Base64Encode 对输入字符串进行Base64编码
func Base64Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// cleanText 清理文本
func cleanText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "&nbsp;", "")
	text = strings.Trim(text, ":： ")
	return text
}

// removeHtmlTags 去除HTML标签
func removeHtmlTags(text string) string {
	// 使用预编译正则表达式去除HTML标签
	text = htmlTagRegexp.ReplaceAllString(text, "")
	// 去除多余的"P"标记（通常表示实践课）
	text = strings.ReplaceAll(text, " P", "")
	text = strings.ReplaceAll(text, "P", "")
	// 清理多余的空格
	text = strings.TrimSpace(text)
	return text
}

// normalizeText 规范化文本（去除多余空白字符）
func normalizeText(s string) string {
	s = strings.TrimSpace(s)
	return whitespaceRe.ReplaceAllString(s, " ")
}

// parseInt 将字符串转换为整数
func parseInt(s string) int {
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// parseFloat 将字符串转换为浮点数
func parseFloat(s string) float64 {
	s = strings.TrimSpace(s)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}
