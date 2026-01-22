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

// formatDate 将日期从"2026年01月05日"格式转换为"2026-01-05"格式
func formatDate(dateStr string) string {
	// 使用正则表达式匹配"年月日"格式的日期
	re := regexp.MustCompile(`(\d{4})年(\d{1,2})月(\d{1,2})日`)
	matches := re.FindStringSubmatch(dateStr)
	if len(matches) == 4 {
		year := matches[1]
		month := matches[2]
		day := matches[3]
		// 确保月份和日期是两位数
		if len(month) == 1 {
			month = "0" + month
		}
		if len(day) == 1 {
			day = "0" + day
		}
		return year + "-" + month + "-" + day
	}
	// 如果不是"年月日"格式，直接返回原字符串
	return dateStr
}
