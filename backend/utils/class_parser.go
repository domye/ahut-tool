package utils

import (
	"fmt"
	"regexp"
	"sort"
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

	// 创建一个映射来跟踪已存在的课程，防止重复
	existingCourses := make(map[string]*models.Class)

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
			// 星期几（colIndex + 1，因为表格第一列是节次）
			dayOfWeek := colIndex + 1

			// 收集该单元格中所有显示的div元素（详细信息版本）
			cell.Find("div.kbcontent").Each(func(i int, div *goquery.Selection) {
				content, _ := div.Html()
				if content == "" || strings.TrimSpace(div.Text()) == "&nbsp;" || strings.TrimSpace(div.Text()) == "" {
					return
				}

				classInfo := parseDetailedClassInfo(content, dayOfWeek, period)
				if classInfo.Name != "" {
					// 生成唯一键值用于去重
					key := fmt.Sprintf("%s_%d_%s_%s", classInfo.Name, classInfo.DayOfWeek, classInfo.Period, strings.Join(intSliceToStringSlice(classInfo.WeekNumbers), ","))

					// 如果不存在相同课程，则添加
					if _, exists := existingCourses[key]; !exists {
						classes = append(classes, classInfo)
						existingCourses[key] = &classInfo
					}
				}
			})

			// 同时也要检查简要信息（非放大模式下的信息）
			cell.Find("div.kbcontent1:not(.sykb1)").Each(func(i int, div *goquery.Selection) {
				content, _ := div.Html()
				if content == "" || strings.TrimSpace(div.Text()) == "&nbsp;" || strings.TrimSpace(div.Text()) == "" {
					return
				}

				classInfo := parseSimpleClassInfo(content, dayOfWeek, period)
				if classInfo.Name != "" {
					// 生成唯一键值用于去重
					key := fmt.Sprintf("%s_%d_%s_%s", classInfo.Name, classInfo.DayOfWeek, classInfo.Period, strings.Join(intSliceToStringSlice(classInfo.WeekNumbers), ","))

					// 检查是否已存在相同的课程
					if _, exists := existingCourses[key]; !exists {
						// 检查是否能与已有的详细课程信息合并
						found := false
						for idx, existingClass := range classes {
							if existingClass.Name == classInfo.Name &&
								existingClass.DayOfWeek == classInfo.DayOfWeek &&
								existingClass.Period == classInfo.Period {
								// 更新现有课程的信息
								if classes[idx].Classroom == "" && classInfo.Classroom != "" {
									classes[idx].Classroom = classInfo.Classroom
								}
								if len(classes[idx].WeekNumbers) == 0 && len(classInfo.WeekNumbers) > 0 {
									classes[idx].WeekNumbers = classInfo.WeekNumbers
								}
								found = true
								break
							}
						}

						if !found {
							classes = append(classes, classInfo)
							existingCourses[key] = &classInfo
						}
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

	// 提取所有文本节点和带title属性的font标签内容
	var lines []string

	// 遍历div中的所有直接子元素和文本节点
	doc.Find("div").Children().Each(func(i int, s *goquery.Selection) {
		tagName := goquery.NodeName(s)
		if tagName == "font" {
			title, _ := s.Attr("title")
			text := s.Text()
			if title != "" {
				lines = append(lines, fmt.Sprintf("%s: %s", title, text))
			} else {
				if text != "" && text != "&nbsp;" {
					lines = append(lines, text)
				}
			}
		} else {
			text := s.Text()
			if text != "" && text != "&nbsp;" {
				lines = append(lines, text)
			}
		}
	})

	// 如果Children没捕获到，尝试直接获取所有内容
	if len(lines) == 0 {
		content := doc.Text()
		content = strings.ReplaceAll(contentHtml, "<br/>", "<br>")
		content = strings.ReplaceAll(content, "<br />", "<br>")
		content = strings.ReplaceAll(content, "<br >", "<br>")
		parts := strings.Split(content, "<br>")

		for _, part := range parts {
			cleanPart := cleanText(part)
			if cleanPart != "" && cleanPart != "&nbsp;" {
				lines = append(lines, cleanPart)
			}
		}
	}

	// 解析提取到的行
	className := ""
	for _, line := range lines {
		line = cleanText(line)
		if line == "" || line == "&nbsp;" {
			continue
		}

		// 检查是否包含教师信息
		if strings.Contains(line, "老师:") || strings.Contains(line, "老师：") || strings.HasPrefix(line, "老师") {
			classInfo.Teacher = extractValueFromLine(line)
		} else if strings.Contains(line, "教室:") || strings.Contains(line, "教室：") || strings.HasPrefix(line, "教室") {
			classInfo.Classroom = extractValueFromLine(line)
		} else if strings.Contains(line, "周次") || strings.Contains(line, "周(") {
			classInfo.WeekNumbers = extractWeekNumbers(line)
		} else if className == "" {
			// 如果还没设置课程名称，这可能是课程名称
			className = line
		}
	}

	// 如果仍然没有课程名称，使用第一个非空行作为课程名称
	if className == "" && len(lines) > 0 {
		for _, line := range lines {
			cleanLine := cleanText(line)
			if cleanLine != "" && cleanLine != "&nbsp;" &&
				!strings.Contains(cleanLine, "老师:") && !strings.Contains(cleanLine, "教室:") && !strings.Contains(cleanLine, "周次") {
				className = cleanLine
				break
			}
		}
	}

	classInfo.Name = className
	classInfo.DayOfWeek = dayOfWeek
	classInfo.Period = period

	return classInfo
}

// parseSimpleClassInfo 解析简要课程信息
func parseSimpleClassInfo(content string, dayOfWeek int, period string) models.Class {
	var classInfo models.Class

	// 处理换行符
	content = strings.ReplaceAll(content, "<br/>", "<br>")
	content = strings.ReplaceAll(content, "<br />", "<br>")
	content = strings.ReplaceAll(content, "<br >", "<br>")
	parts := strings.Split(content, "<br>")

	var lines []string
	for _, part := range parts {
		cleanPart := cleanText(part)
		if cleanPart != "" && cleanPart != "&nbsp;" {
			lines = append(lines, cleanPart)
		}
	}

	className := ""
	for _, line := range lines {
		line = cleanText(line)
		if line == "" || line == "&nbsp;" {
			continue
		}

		// 检查是否包含教室信息
		if strings.Contains(line, "教室:") || strings.Contains(line, "教室：") || strings.HasPrefix(line, "教室") {
			classInfo.Classroom = extractValueFromLine(line)
		} else if strings.Contains(line, "周次") || strings.Contains(line, "周(") {
			classInfo.WeekNumbers = extractWeekNumbers(line)
		} else if className == "" {
			// 如果还没设置课程名称，这可能是课程名称
			className = line
		}
	}

	// 如果仍然没有课程名称，使用第一个非空行作为课程名称
	if className == "" && len(lines) > 0 {
		for _, line := range lines {
			cleanLine := cleanText(line)
			if cleanLine != "" && cleanLine != "&nbsp;" &&
				!strings.Contains(cleanLine, "教室:") && !strings.Contains(cleanLine, "周次") {
				className = cleanLine
				break
			}
		}
	}

	classInfo.Name = className
	classInfo.DayOfWeek = dayOfWeek
	classInfo.Period = period

	return classInfo
}

// extractValueFromLine 从带标签的行中提取值，例如"教室: 东教一南502" -> "东教一南502"
func extractValueFromLine(line string) string {
	// 查找冒号或中文冒号的位置
	colonIndex := strings.Index(line, ":")
	if colonIndex == -1 {
		colonIndex = strings.Index(line, "：")
	}

	if colonIndex != -1 && len(line) > colonIndex+1 {
		return cleanText(line[colonIndex+1:])
	}

	// 如果没有找到冒号，尝试其他分隔方式
	parts := strings.FieldsFunc(line, func(c rune) bool {
		return c == ':' || c == '：'
	})

	if len(parts) > 1 {
		return cleanText(parts[1])
	}

	return cleanText(line)
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
		// 匹配独立的数字（周次）
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
	}

	// 去重并排序
	weeks = removeDuplicatesAndSort(weeks)

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

	sort.Ints(result)

	return result
}

// cleanText 清理文本
func cleanText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "&nbsp;", "")
	text = strings.Trim(text, ":： ")
	return text
}

// intSliceToStringSlice 将整数切片转换为字符串切片，用于生成唯一键
func intSliceToStringSlice(slice []int) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = strconv.Itoa(v)
	}
	return result
}
