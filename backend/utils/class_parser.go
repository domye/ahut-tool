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

// 预编译正则表达式以提高性能
var (
	periodRegexp      = regexp.MustCompile(`第([0-9,,-]+)节`)
	fontTagRegexp     = regexp.MustCompile(`(?i)<font[^>]*title\s*=\s*['"]([^'"]*)['"][^>]*>(.*?)</font>`)
	fullWeekRegexp    = regexp.MustCompile(`([\d,-]+)\(周\)`)
	rangeWeekRegexp   = regexp.MustCompile(`(\d+)-(\d+)\(周\)`)
	numWeekRegexp     = regexp.MustCompile(`\d+`)
	courseSplitRegexp = regexp.MustCompile(`-{5,}`) // 用于分割多门课程的分隔符
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
				if courseSplitRegexp.MatchString(content) {
					// 分割多门课程
					for _, coursePart := range courseSplitRegexp.Split(content, -1) {
						processCoursePart(coursePart, dayOfWeek, period, &classes, existingCoursesMap)
					}
				} else {
					// 单门课程
					processCoursePart(content, dayOfWeek, period, &classes, existingCoursesMap)
				}
			})
		})
	})

	return classes, nil
}

// processCoursePart 处理单个课程部分
func processCoursePart(coursePart string, dayOfWeek int, period string, classes *[]models.Class, existingCoursesMap map[string]*models.Class) {
	classInfo := parseDetailedClassInfo(coursePart, dayOfWeek, period)
	if classInfo.Name == "" {
		return
	}

	// 生成唯一键值用于去重（不包含周次）
	key := generateCourseKey(classInfo)

	// 更新或添加课程
	updateOrAddCourse(classInfo, key, classes, existingCoursesMap)
}

// generateCourseKey 生成课程唯一键
func generateCourseKey(classInfo models.Class) string {
	return fmt.Sprintf("%s_%d_%s_%s_%s",
		classInfo.Name, classInfo.DayOfWeek, classInfo.Period, classInfo.Teacher, classInfo.Classroom)
}

// updateOrAddCourse 更新现有课程或添加新课程
func updateOrAddCourse(classInfo models.Class, key string, classes *[]models.Class, existingCoursesMap map[string]*models.Class) {
	if existingClass, exists := existingCoursesMap[key]; exists {
		// 合并周次信息
		existingClass.WeekNumbers = mergeWeekRanges(existingClass.WeekNumbers, classInfo.WeekNumbers)
	} else { // 添加新课程
		*classes = append(*classes, classInfo)
		existingCoursesMap[key] = &(*classes)[len(*classes)-1]
	}
}

// extractPeriod 从节次文本中提取节次信息
func extractPeriod(periodText string) string {
	// 匹配"第X,Y节"或"第X-Y节"的模式
	matches := periodRegexp.FindStringSubmatch(periodText)
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
	content := preprocessContent(contentHtml)

	// 提取font标签信息
	extractFontInfo(content, &classInfo)

	// 提取课程名称
	extractClassName(content, &classInfo)

	classInfo.DayOfWeek = dayOfWeek
	classInfo.Period = period

	return classInfo
}

// preprocessContent 预处理内容
func preprocessContent(contentHtml string) string {
	return strings.NewReplacer(
		"<br/>", "<br>",
		"<br />", "<br>",
		"<br >", "<br>",
		"\t", "",
	).Replace(contentHtml)
}

// extractFontInfo 提取font标签信息
func extractFontInfo(content string, classInfo *models.Class) {
	fontMatches := fontTagRegexp.FindAllStringSubmatch(content, -1)

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
}

// extractClassName 提取课程名称
func extractClassName(content string, classInfo *models.Class) {
	for _, part := range strings.Split(content, "<br>") {
		cleanPart := cleanText(part)
		// 清理HTML标签
		cleanPart = removeHtmlTags(cleanPart)
		// 跳过分隔线
		if isValidClassName(cleanPart) {
			classInfo.Name = cleanPart
			break
		}
	}
}

// isValidClassName 判断是否为有效的课程名称
func isValidClassName(cleanPart string) bool {
	return cleanPart != "" && cleanPart != "&nbsp;" && !strings.Contains(cleanPart, "------")
}

// extractWeekNumbers 从文本中提取周次信息
func extractWeekNumbers(text string) string {
	// 尝试匹配完整周次格式，如"1-16(周)"或"1-11,13-14(周)"或"1-12,14(周)"
	if result := extractFullWeekFormat(text); result != "" {
		return result
	}

	// 尝试匹配范围格式，如"X-Y(周)"
	if result := extractRangeWeekFormat(text); result != "" {
		return result
	}

	// 尝试匹配单独的数字
	return extractNumberWeekFormat(text)
}

// extractFullWeekFormat 提取完整周次格式
func extractFullWeekFormat(text string) string {
	fullMatches := fullWeekRegexp.FindStringSubmatch(text)
	if len(fullMatches) >= 2 {
		return fullMatches[1] // 提取周次部分，如"1-11,13-14"或"1-12,14"
	}
	return ""
}

// extractRangeWeekFormat 提取范围周次格式
func extractRangeWeekFormat(text string) string {
	matches := rangeWeekRegexp.FindAllStringSubmatch(text, -1)
	if len(matches) == 0 {
		return ""
	}

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

	uniqueWeeks := uniqueSortedIntsWithoutAlloc(allWeeks)
	return buildWeekRanges(uniqueWeeks)
}

// extractNumberWeekFormat 提取数字周次格式
func extractNumberWeekFormat(text string) string {
	numbers := numWeekRegexp.FindAllString(text, -1)
	if len(numbers) == 0 {
		return ""
	}

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

	sort.Ints(weeks)
	return buildWeekRanges(weeks)
}

// uniqueSortedIntsWithoutAlloc 去重并排序整数切片，优化内存分配
func uniqueSortedIntsWithoutAlloc(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// 使用哈希集合去重，提供O(1)的查找时间
	seen := make(map[int]bool)
	for _, num := range nums {
		seen[num] = true
	}

	// 将唯一的数字复制到结果切片
	result := make([]int, 0, len(seen))
	for num := range seen {
		result = append(result, num)
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

	// 使用strings.Builder优化字符串连接性能
	var builder strings.Builder
	start := weeks[0]
	end := weeks[0]

	for i := 1; i < len(weeks); i++ {
		if weeks[i] == end+1 {
			// 连续的周次，更新结束周次
			end = weeks[i]
		} else {
			// 不连续的周次，保存当前范围
			if start == end {
				builder.WriteString(strconv.Itoa(start))
			} else {
				builder.WriteString(fmt.Sprintf("%d-%d", start, end))
			}
			builder.WriteByte(',')
			start = weeks[i]
			end = weeks[i]
		}
	}

	// 添加最后一个范围
	if start == end {
		builder.WriteString(strconv.Itoa(start))
	} else {
		builder.WriteString(fmt.Sprintf("%d-%d", start, end))
	}

	return builder.String()
}

// cleanText 已移至 common.go
// removeHtmlTags 已移至 common.go

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

	// 使用哈希集合去重，合并两个周次数组
	seen := make(map[int]bool)
	for _, week := range weeks1 {
		seen[week] = true
	}
	for _, week := range weeks2 {
		seen[week] = true
	}

	// 将唯一的周次复制到结果切片
	merged := make([]int, 0, len(seen))
	for week := range seen {
		merged = append(merged, week)
	}

	// 排序并构建周次范围字符串
	sort.Ints(merged)
	return buildWeekRanges(merged)
}

// parseWeekRange 解析周次范围字符串，返回所有周次的数组
func parseWeekRange(weekRange string) []int {
	// 预分配容量以减少内存重新分配
	weeks := make([]int, 0, len(weekRange)/2)

	// 分割多个范围，如"1-9,11-12"
	ranges := strings.Split(weekRange, ",")

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		// 检查是否是范围，如"1-9"
		hyphenIndex := strings.Index(r, "-")
		if hyphenIndex != -1 {
			startStr := r[:hyphenIndex]
			endStr := r[hyphenIndex+1:]
			start, err1 := strconv.Atoi(strings.TrimSpace(startStr))
			end, err2 := strconv.Atoi(strings.TrimSpace(endStr))
			if err1 == nil && err2 == nil {
				// 预计算范围大小并预先扩展切片
				for w := start; w <= end; w++ {
					weeks = append(weeks, w)
				}
			}
		} else {
			// 单个周次
			if week, err := strconv.Atoi(r); err == nil {
				weeks = append(weeks, week)
			}
		}
	}

	return weeks
}
