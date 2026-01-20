package utils

import (
	"regexp"
	"strconv"
	"strings"

	"ahut-tool/backend/models"
	"github.com/PuerkitoBio/goquery"
)

// ParseGrades 从HTML中解析成绩信息
func ParseGrades(html string) ([]models.Grade, *models.GradeSummary, error) {
	var grades []models.Grade
	var summary models.GradeSummary

	// 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, nil, err
	}

	// 提取成绩汇总信息
	summaryRegex := regexp.MustCompile(`所修门数:(\d+) 所修总学分:([\d.]+) 平均学分绩点:([\d.]+) 平均成绩:([\d.]+)`)
	summaryText := doc.Find("div").First().Text()
	summaryMatches := summaryRegex.FindStringSubmatch(summaryText)
	if len(summaryMatches) == 5 {
		summary.CourseCount, _ = strconv.Atoi(summaryMatches[1])
		summary.TotalCredit, _ = strconv.ParseFloat(summaryMatches[2], 64)
		summary.AvgGPA, _ = strconv.ParseFloat(summaryMatches[3], 64)
		summary.AvgScore, _ = strconv.ParseFloat(summaryMatches[4], 64)
	}

	// 查找成绩表格
	doc.Find("#dataList tr").Each(func(i int, s *goquery.Selection) {
		// 跳过表头
		if i == 0 {
			return
		}

		// 获取所有单元格
		cells := s.Find("td")
		if cells.Length() < 16 {
			return
		}

		// 提取成绩，可能在a标签中或直接在td中
		scoreCell := cells.Eq(5)
		scoreText := scoreCell.Text()
		if link := scoreCell.Find("a"); link.Length() > 0 {
			scoreText = link.Text()
		}

		grade := models.Grade{
			Index:        parseInt(cells.Eq(0).Text()),
			Semester:     strings.TrimSpace(cells.Eq(1).Text()),
			CourseID:     strings.TrimSpace(cells.Eq(2).Text()),
			CourseName:   strings.TrimSpace(cells.Eq(3).Text()),
			GroupName:    strings.TrimSpace(cells.Eq(4).Text()),
			Score:        strings.TrimSpace(scoreText),
			ScoreFlag:    strings.TrimSpace(cells.Eq(6).Text()),
			Credit:       parseFloat(cells.Eq(7).Text()),
			TotalHours:   parseInt(cells.Eq(8).Text()),
			GPA:          parseFloat(cells.Eq(9).Text()),
			RetakeSem:    strings.TrimSpace(cells.Eq(10).Text()),
			ExamMode:     strings.TrimSpace(cells.Eq(11).Text()),
			ExamType:     strings.TrimSpace(cells.Eq(12).Text()),
			CourseAttr:   strings.TrimSpace(cells.Eq(13).Text()),
			CourseNature: strings.TrimSpace(cells.Eq(14).Text()),
			GEType:       strings.TrimSpace(cells.Eq(15).Text()),
		}

		grades = append(grades, grade)
	})

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
