package utils

import (
	"regexp"
	"strconv"
	"strings"

	"ahut-tool/backend/models"

	"github.com/PuerkitoBio/goquery"
)

// ParseClassSchedule 从HTML中解析课程表信息
func ParseClassSchedule(html string) ([]models.Class, error) {
	var classes []models.Class

	// 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	// 遍历课表表格中的每一行（跳过表头）
	doc.Find("#kbtable tr").Each(func(rowIndex int, row *goquery.Selection) {
		// 跳过表头
		if rowIndex == 0 {
			return
		}

		// 获取当前行的节次信息
		periodCell := row.Find("th").First()
		periodText := strings.TrimSpace(periodCell.Text())
		// 提取节次信息，例如"第1,2节" -> "1-2"
		period := extractPeriod(periodText)

		// 遍历该行的所有单元格（课程信息）
		row.Find("td").Each(func(colIndex int, cell *goquery.Selection) {
			// 第一个单元格是节次信息，跳过
			if colIndex == 0 {
				return
			}

			// 首先查找详细信息的div元素（优先级更高）
			cell.Find("div.kbcontent").Each(func(i int, div *goquery.Selection) {
				// 获取div的内容来解析课程信息
				content, _ := div.Html()
				if content == "" || strings.TrimSpace(div.Text()) == "&nbsp;" || strings.TrimSpace(div.Text()) == "" {
					return
				}

				// 解析详细课程信息
				classInfo := parseDetailedClassInfo(content, colIndex, period)
				if classInfo.Name != "" {
					classes = append(classes, classInfo)
				}
			})

			// 如果没有找到详细信息，则查找简要信息
			cell.Find("div.kbcontent1:not(.sykb1)").Each(func(i int, div *goquery.Selection) {
				// 检查是否已经有对应详细信息的课程，避免重复
				hasDetailVersion := false
				for _, existingClass := range classes {
					if existingClass.DayOfWeek == colIndex && existingClass.Period == period {
						hasDetailVersion = true
						break
					}
				}

				if !hasDetailVersion {
					content := div.Text()
					if content == "" || strings.TrimSpace(content) == "&nbsp;" {
						return
					}

					// 解析简要课程信息
					classInfo := parseSimpleClassInfo(content, colIndex, period)
					if classInfo.Name != "" {
						classes = append(classes, classInfo)
					}
				}
			})
		})
	})

	return classes, nil
}

// extractPeriod 从节次文本中提取节次信息
func extractPeriod(periodText string) string {
	// 匹配"第X,Y节"或"第X-Y节"的模式
	re := regexp.MustCompile(`第([0-9,,-]+)节`)
	matches := re.FindStringSubmatch(periodText)
	if len(matches) > 1 {
		result := strings.ReplaceAll(matches[1], ",", "-")
		return result
	}
	return "未知"
}

// parseDetailedClassInfo 解析详细课程信息
func parseDetailedClassInfo(contentHtml string, dayOfWeek int, period string) models.Class {
	var classInfo models.Class

	// 创建文档用于解析HTML内容
	doc, err := goquery.NewDocumentFromReader(strings.NewReader("<div>" + contentHtml + "</div>"))
	if err != nil {
		return classInfo
	}

	// 直接处理HTML内容
	textNodes := []string{}
	doc.Find("*").Contents().Each(func(i int, s *goquery.Selection) {
		node := s.Get(0)
		if node != nil && node.Type == 3 { // 文本节点
			text := strings.TrimSpace(node.Data)
			if text != "" && text != "&nbsp;" {
				textNodes = append(textNodes, text)
			}
		}
	})

	// 处理文本节点，按<br>标签分割的内容
	content := strings.ReplaceAll(contentHtml, "<br/>", "<br>")
	content = strings.ReplaceAll(content, "<br />", "<br>")
	parts := strings.Split(content, "<br>")

	for i, part := range parts {
		part = cleanText(part)
		if part == "" || part == "&nbsp;" {
			continue
		}

		// 第一行通常是课程名称
		if i == 0 && classInfo.Name == "" {
			classInfo.Name = part
		} else if strings.Contains(part, "老师") {
			// 提取教师姓名
			colonIndex := strings.Index(part, ">")
			if colonIndex != -1 && len(part) > colonIndex+1 {
				classInfo.Teacher = cleanText(part[colonIndex+1:])
			} else {
				// 如果没有找到">"，尝试其他方式提取
				classInfo.Teacher = extractTeacherFromText(part)
			}
		} else if strings.Contains(part, "教室") {
			// 提取教室信息
			colonIndex := strings.Index(part, ">")
			if colonIndex != -1 && len(part) > colonIndex+1 {
				classInfo.Classroom = cleanText(part[colonIndex+1:])
			}
		} else if strings.Contains(part, "周次") {
			// 提取周次信息
			classInfo.WeekNumbers = extractWeekNumbers(part)
		}
	}

	// 如果课程名称仍未设置，尝试从文本节点中获取
	if classInfo.Name == "" && len(textNodes) > 0 {
		classInfo.Name = textNodes[0]
	}

	// 设置星期和节次
	classInfo.DayOfWeek = dayOfWeek
	classInfo.Period = period

	return classInfo
}

// parseSimpleClassInfo 解析简要课程信息
func parseSimpleClassInfo(content string, dayOfWeek int, period string) models.Class {
	var classInfo models.Class

	// 直接处理文本内容
	content = strings.ReplaceAll(content, "<br/>", "<br>")
	content = strings.ReplaceAll(content, "<br />", "<br>")
	parts := strings.Split(content, "<br>")

	for i, part := range parts {
		part = cleanText(part)
		if part == "" || part == "&nbsp;" {
			continue
		}

		// 第一部分通常是课程名称
		if i == 0 && classInfo.Name == "" {
			classInfo.Name = part
		} else if strings.Contains(part, "教室") {
			// 提取教室信息
			colonIndex := strings.Index(part, ">")
			if colonIndex != -1 && len(part) > colonIndex+1 {
				classInfo.Classroom = cleanText(part[colonIndex+1:])
			}
		} else if strings.Contains(part, "周次") {
			// 提取周次信息
			classInfo.WeekNumbers = extractWeekNumbers(part)
		}
	}

	// 设置星期和节次
	classInfo.DayOfWeek = dayOfWeek
	classInfo.Period = period

	return classInfo
}

// extractTeacherFromText 从文本中提取教师姓名
func extractTeacherFromText(text string) string {
	// 移除"老师"字样和周围的标点
	text = strings.ReplaceAll(text, "老师", "")
	text = strings.Trim(text, ":： ")
	return text
}

// extractWeekNumbers 从文本中提取周次信息
func extractWeekNumbers(text string) []int {
	var weeks []int

	// 匹配"X-Y(周)"模式，如"1-16(周)"
	re := regexp.MustCompile(`(\d+)-(\d+)\(周\)`)
	matches := re.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			start, _ := strconv.Atoi(match[1])
			end, _ := strconv.Atoi(match[2])
			for w := start; w <= end; w++ {
				weeks = append(weeks, w)
			}
		}
	}

	// 如果没有匹配到范围，尝试匹配单独的数字
	if len(weeks) == 0 {
		numRe := regexp.MustCompile(`\d+`)
		numbers := numRe.FindAllString(text, -1)
		for _, numStr := range numbers {
			if num, err := strconv.Atoi(numStr); err == nil {
				// 确保是合理的周次（1-30周）
				if num >= 1 && num <= 30 {
					weeks = append(weeks, num)
				}
			}
		}
		// 排序并去重
		weeks = removeDuplicatesAndSort(weeks)
	}

	return weeks
}

// removeDuplicatesAndSort 去除重复项并排序
func removeDuplicatesAndSort(slice []int) []int {
	seen := make(map[int]bool)
	var result []int

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	// 简单冒泡排序（由于数组通常较小，效率足够）
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}

// cleanText 清理文本
func cleanText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "&nbsp;", "")
	text = strings.Trim(text, ":： ")
	return text
}
