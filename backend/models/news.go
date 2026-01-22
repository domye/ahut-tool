package models

// News 新闻信息结构体
type news struct {
	Title   string `json:"title"`   // 新闻标题
	Content string `json:"content"` // 新闻内容
	Date    string `json:"date"`    // 新闻发布日期
	Url     string `json:"url"`     // 新闻链接
}
