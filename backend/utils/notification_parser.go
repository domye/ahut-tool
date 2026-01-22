package utils

import (
	"strings"

	"ahut-tool/backend/models"
	"github.com/PuerkitoBio/goquery"
)

// ParseNotifications 从HTML中解析学业通知信息
func ParseNotifications(html string) ([]models.News, error) {
	var newsList []models.News

	// 使用goquery解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	// 预分配切片容量以提高性能（假设平均约20条新闻）
	newsList = make([]models.News, 0, 20)

	// 查找新闻列表
	doc.Find(".scienceYield ul li").Each(func(i int, s *goquery.Selection) {
		// 获取新闻链接
		link := s.Find("a")
		if link.Length() == 0 {
			return
		}

		// 提取新闻URL
		href, exists := link.Attr("href")
		if !exists {
			return
		}

		// 补全URL为绝对路径
		if strings.HasPrefix(href, "../") {
			href = "https://jwc.ahut.edu.cn/" + strings.TrimPrefix(href, "../")
		}

		// 提取新闻标题
		title := normalizeText(link.Find("h3").Text())
		if title == "" {
			return
		}

		// 提取新闻日期
		dateText := normalizeText(link.Find(".time").Text())
		// 移除日期前的图片标签文本
		dateText = strings.TrimPrefix(dateText, "图片")

		// 提取新闻内容
		content := normalizeText(link.Find("p").Not(".time").Text())

		news := models.News{
			Title:   title,
			Content: content,
			Date:    dateText,
			Url:     href,
		}

		newsList = append(newsList, news)
	})

	return newsList, nil
}
