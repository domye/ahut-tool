package utils

import (
	"ahut-tool/backend/models"
	"regexp"
	"strconv"
	"strings"
)

// ParseGrades 从HTML中解析成绩信息
func ParseGrades(html string) ([]models.Grade, *models.GradeSummary, error) {
	var grades []models.Grade
	var summary models.GradeSummary

	// 提取成绩汇总信息
	summaryRegex := regexp.MustCompile(`所修门数:(\d+) 所修总学分:([\d.]+) 平均学分绩点:([\d.]+) 平均成绩:([\d.]+)`)
	summaryMatches := summaryRegex.FindStringSubmatch(html)
	if len(summaryMatches) == 5 {
		summary.CourseCount, _ = strconv.Atoi(summaryMatches[1])
		summary.TotalCredit, _ = strconv.ParseFloat(summaryMatches[2], 64)
		summary.AvgGPA, _ = strconv.ParseFloat(summaryMatches[3], 64)
		summary.AvgScore, _ = strconv.ParseFloat(summaryMatches[4], 64)
	}

	// 提取成绩表格行
	// 使用更简单的正则表达式，先匹配整个tr标签
	trRegex := regexp.MustCompile(`<tr>\s*<td>(\d+)</td>\s*<td>([^<]+)</td>\s*<td[^>]*>([^<]+)</td>\s*<td[^>]*>([^<]+)</td>\s*<td[^>]*>([^<]*)</td>\s*<td[^>]*>\s*(?:<a[^>]*>([^<]+)</a>|([^<]+))\s*</td>\s*(?:<!--.*?-->\s*)*</td>\s*<td>([^<]*)</td>\s*<td>([^<]+)</td>\s*<td>(\d+)</td>\s*<td>([\d.]+)</td>\s*<td>([^<]*)</td>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*<td>([^<]+)</td>\s*<td>([^<]*)</td>\s*</tr>`)
	trMatches := trRegex.FindAllStringSubmatch(html, -1)

	for _, match := range trMatches {
		if len(match) < 18 {
			continue
		}

		grade := models.Grade{
			Index:        parseInt(match[1]),
			Semester:     strings.TrimSpace(match[2]),
			CourseID:     strings.TrimSpace(match[3]),
			CourseName:   strings.TrimSpace(match[4]),
			GroupName:    strings.TrimSpace(match[5]),
			Score:        strings.TrimSpace(getNonEmpty(match[6], match[7])),
			ScoreFlag:    strings.TrimSpace(match[8]),
			Credit:       parseFloat(match[9]),
			TotalHours:   parseInt(match[10]),
			GPA:          parseFloat(match[11]),
			RetakeSem:    strings.TrimSpace(match[12]),
			ExamMode:     strings.TrimSpace(match[13]),
			ExamType:     strings.TrimSpace(match[14]),
			CourseAttr:   strings.TrimSpace(match[15]),
			CourseNature: strings.TrimSpace(match[16]),
			GEType:       strings.TrimSpace(match[17]),
		}

		grades = append(grades, grade)
	}

	return grades, &summary, nil
}

// parseInt 将字符串转换为整数
func parseInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return 0
	}
	return i
}

// parseFloat 将字符串转换为浮点数
func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0
	}
	return f
}

// getNonEmpty 获取非空字符串
func getNonEmpty(s1, s2 string) string {
	if strings.TrimSpace(s1) != "" {
		return s1
	}
	return s2
}
