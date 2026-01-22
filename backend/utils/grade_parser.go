package utils

import (
	"regexp"
	"strconv"
	"strings"

	"ahut-tool/backend/models"

	"github.com/PuerkitoBio/goquery"
)

// 预编译正则表达式以提高性能
var (
	summaryRegex = regexp.MustCompile(`所修门数:(\d+) 所修总学分:([\d.]+) 平均学分绩点:([\d.]+) 平均成绩:([\d.]+)`)
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
	summaryText := strings.TrimSpace(doc.Find("div").First().Text())
	summaryMatches := summaryRegex.FindStringSubmatch(summaryText)
	if len(summaryMatches) == 5 {
		summary.CourseCount, _ = strconv.Atoi(summaryMatches[1])
		summary.TotalCredit, _ = strconv.ParseFloat(summaryMatches[2], 64)
		summary.AvgGPA, _ = strconv.ParseFloat(summaryMatches[3], 64)
		summary.AvgScore, _ = strconv.ParseFloat(summaryMatches[4], 64)
	}

	// 预分配切片容量以提高性能（假设平均约100门课程）
	grades = make([]models.Grade, 0, 100)

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
		scoreText := strings.TrimSpace(scoreCell.Text())
		if link := scoreCell.Find("a"); link.Length() > 0 {
			scoreText = strings.TrimSpace(link.Text())
		}

		grade := models.Grade{
			Index:        parseInt(cells.Eq(0).Text()),
			Semester:     normalizeText(cells.Eq(1).Text()),
			CourseID:     normalizeText(cells.Eq(2).Text()),
			CourseName:   normalizeText(cells.Eq(3).Text()),
			GroupName:    normalizeText(cells.Eq(4).Text()),
			Score:        scoreText,
			ScoreFlag:    normalizeText(cells.Eq(6).Text()),
			Credit:       parseFloat(cells.Eq(7).Text()),
			TotalHours:   parseInt(cells.Eq(8).Text()),
			GPA:          parseFloat(cells.Eq(9).Text()),
			RetakeSem:    normalizeText(cells.Eq(10).Text()),
			ExamMode:     normalizeText(cells.Eq(11).Text()),
			ExamType:     normalizeText(cells.Eq(12).Text()),
			CourseAttr:   normalizeText(cells.Eq(13).Text()),
			CourseNature: normalizeText(cells.Eq(14).Text()),
			GEType:       normalizeText(cells.Eq(15).Text()),
		}

		grades = append(grades, grade)
	})

	return grades, &summary, nil
}

// normalizeText 已移至 common.go
// parseInt 已移至 common.go
// parseFloat 已移至 common.go
