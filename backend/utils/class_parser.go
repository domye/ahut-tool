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

	// 创建一个映射来跟踪已存在的课程，防止重复，并支持合并周次信息
	existingCoursesMap := make(map[string]*models.Class)

	// 遍历课表表格中的每一行（跳过表头）
	doc.Find("#kbtable tr").Each(func(rowIndex int, row *goquery.Selection) {
		// 跳过表头
		if rowIndex == 0 {
			return
		}

		// 获取当前行的节次信息
		periodText := strings.TrimSpace(row.Find("th").First().Text())
		// 提取节次信息，例如"第1,2节" -> "1-2"
		period := extractPeriod(periodText)

		// 遍历该行的所有单元格（课程信息）
		row.Find("td").Each(func(colIndex int, cell *goquery.Selection) {
			// 星期几（colIndex + 1，因为表格第一列是节次）
			dayOfWeek := colIndex + 1

			// 只收集该单元格中所有显示的div元素（详细信息版本）
			// 不再解析kbcontent1，避免重复
			// 跳过sykb2元素，避免解析空div
			cell.Find("div.kbcontent:not(.sykb2)").Each(func(i int, div *goquery.Selection) {
				content, _ := div.Html()
				if content == "" || strings.TrimSpace(div.Text()) == "&nbsp;" || strings.TrimSpace(div.Text()) == "" {
					return
				}

				// 检查是否包含多门课程（用"---------------------"分隔）
				if strings.Contains(content, "---------------------") {
					// 分割多门课程
					for _, coursePart := range strings.Split(content, "---------------------") {
						classInfo := parseDetailedClassInfo(coursePart, dayOfWeek, period)
						if classInfo.Name != "" {
							// 生成唯一键值用于去重（不包含周次）
							key := fmt.Sprintf("%s_%d_%s_%s_%s", classInfo.Name, classInfo.DayOfWeek, classInfo.Period, classInfo.Teacher, classInfo.Classroom)

							// 检查是否已存在相同课程
							if existingClass, exists := existingCoursesMap[key]; exists {
								// 合并周次信息
								existingClass.WeekNumbers = mergeWeekRanges(existingClass.WeekNumbers, classInfo.WeekNumbers)
							} else {
								// 添加新课程
								classes = append(classes, classInfo)
								existingCoursesMap[key] = &classes[len(classes)-1]
							}
						}
					}
				} else {
					// 单门课程
					classInfo := parseDetailedClassInfo(content, dayOfWeek, period)
					if classInfo.Name != "" {
						// 生成唯一键值用于去重（不包含周次）
						key := fmt.Sprintf("%s_%d_%s_%s_%s", classInfo.Name, classInfo.DayOfWeek, classInfo.Period, classInfo.Teacher, classInfo.Classroom)

						// 检查是否已存在相同课程
						if existingClass, exists := existingCoursesMap[key]; exists {
							// 合并周次信息
							existingClass.WeekNumbers = mergeWeekRanges(existingClass.WeekNumbers, classInfo.WeekNumbers)
						} else {
							// 添加新课程
							classes = append(classes, classInfo)
							existingCoursesMap[key] = &classes[len(classes)-1]
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

	// 处理换行符和清理无用标签
	content := strings.NewReplacer(
		"<br/>", "<br>",
		"<br />", "<br>",
		"<br >", "<br>",
		"\t", "",
	).Replace(contentHtml)

	// 使用正则表达式提取font标签及其内容
	fontRe := regexp.MustCompile(`(?i)<font[^>]*title\s*=\s*['"]([^'"]*)['"][^>]*>(.*?)</font>`)
	fontMatches := fontRe.FindAllStringSubmatch(content, -1)

	// 根据索引位置提取信息: 0是教师，1是周次，2是教室
	for i, match := range fontMatches {
		if len(match) < 3 {
			continue
		}

		title := match[1]
		text := match[2]

		switch {
		case i == 0 && title == "老师":
			classInfo.Teacher = cleanText(text)
		case i == 1 && (title == "周次" || title == "周次(节次)"):
			classInfo.WeekNumbers = extractWeekNumbers(text)
		case i == 2 && title == "教室":
			classInfo.Classroom = cleanText(text)
		}
	}

	// 处理剩余内容，提取课程名称
	for _, part := range strings.Split(content, "<br>") {
		cleanPart := cleanText(part)
		// 清理HTML标签
		cleanPart = removeHtmlTags(cleanPart)
		// 跳过分隔线
		if cleanPart != "" && cleanPart != "&nbsp;" && !strings.Contains(cleanPart, "------") {
			classInfo.Name = cleanPart
			break
		}
	}

	classInfo.DayOfWeek = dayOfWeek
	classInfo.Period = period

	return classInfo
}

// extractWeekNumbers 从文本中提取周次信息
func extractWeekNumbers(text string) string {
	// 匹配周次信息，如"1-16(周)"或"1-11,13-14(周)"或"1-12,14(周)"
	fullRe := regexp.MustCompile(`([\d,-]+)\(周\)`)
	fullMatches := fullRe.FindStringSubmatch(text)

	if len(fullMatches) >= 2 {
		// 提取周次部分，如"1-11,13-14"或"1-12,14"
		return fullMatches[1]
	}

	// 如果没有匹配到完整格式，尝试匹配"X-Y(周)"模式
	re := regexp.MustCompile(`(\d+)-(\d+)\(周\)`)
	matches := re.FindAllStringSubmatch(text, -1)

	if len(matches) > 0 {
		// 提取所有周次范围
		var allWeeks []int
		for _, match := range matches {
			if len(match) >= 3 {
				start, _ := strconv.Atoi(match[1])
				end, _ := strconv.Atoi(match[2])
				for w := start; w <= end; w++ {
					allWeeks = append(allWeeks, w)
				}
			}
		}

		// 去重并排序
		uniqueWeeks := uniqueSortedInts(allWeeks)

		// 将周次转换为范围字符串
		return buildWeekRanges(uniqueWeeks)
	}

	// 如果没有匹配到范围，尝试匹配单独的数字
	numRe := regexp.MustCompile(`\d+`)
	numbers := numRe.FindAllString(text, -1)

	if len(numbers) > 0 {
		var weeks []int
		seen := make(map[int]bool)

		for _, numStr := range numbers {
			if num, err := strconv.Atoi(numStr); err == nil && num >= 1 && num <= 30 && !seen[num] {
				seen[num] = true
				weeks = append(weeks, num)
			}
		}

		if len(weeks) == 0 {
			return ""
		}

		// 排序并构建范围
		sort.Ints(weeks)
		return buildWeekRanges(weeks)
	}

	return ""
}

// uniqueSortedInts 去重并排序整数切片
func uniqueSortedInts(nums []int) []int {
	seen := make(map[int]bool)
	var result []int

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	sort.Ints(result)
	return result
}

// buildWeekRanges 构建周次范围字符串
func buildWeekRanges(weeks []int) string {
	if len(weeks) == 0 {
		return ""
	}

	if len(weeks) == 1 {
		return strconv.Itoa(weeks[0])
	}

	var result []string
	start := weeks[0]
	end := weeks[0]

	for i := 1; i < len(weeks); i++ {
		if weeks[i] == end+1 {
			// 连续的周次，更新结束周次
			end = weeks[i]
		} else {
			// 不连续的周次，保存当前范围
			if start == end {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, fmt.Sprintf("%d-%d", start, end))
			}
			start = weeks[i]
			end = weeks[i]
		}
	}

	// 添加最后一个范围
	if start == end {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, fmt.Sprintf("%d-%d", start, end))
	}

	return strings.Join(result, ",")
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
	// 使用正则表达式去除HTML标签
	re := regexp.MustCompile(`<[^>]*>`)
	text = re.ReplaceAllString(text, "")
	// 去除多余的"P"标记（通常表示实践课）
	text = strings.ReplaceAll(text, " P", "")
	text = strings.ReplaceAll(text, "P", "")
	// 清理多余的空格
	text = strings.TrimSpace(text)
	return text
}

// mergeWeekRanges 合并两个周次范围字符串
func mergeWeekRanges(range1, range2 string) string {
	if range1 == "" {
		return range2
	}
	if range2 == "" {
		return range1
	}

	// 解析第一个周次范围
	weeks1 := parseWeekRange(range1)
	// 解析第二个周次范围
	weeks2 := parseWeekRange(range2)

	// 合并两个周次数组
	merged := append(weeks1, weeks2...)

	// 去重并排序
	uniqueWeeks := uniqueSortedInts(merged)

	// 构建周次范围字符串
	return buildWeekRanges(uniqueWeeks)
}

// parseWeekRange 解析周次范围字符串，返回所有周次的数组
func parseWeekRange(weekRange string) []int {
	var weeks []int

	// 分割多个范围，如"1-9,11-12"
	ranges := strings.Split(weekRange, ",")

	for _, r := range ranges {
		// 检查是否是范围，如"1-9"
		if strings.Contains(r, "-") {
			parts := strings.Split(r, "-")
			if len(parts) == 2 {
				start, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
				end, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
				if err1 == nil && err2 == nil {
					for w := start; w <= end; w++ {
						weeks = append(weeks, w)
					}
				}
			}
		} else {
			// 单个周次
			if week, err := strconv.Atoi(strings.TrimSpace(r)); err == nil {
				weeks = append(weeks, week)
			}
		}
	}

	return weeks
}
