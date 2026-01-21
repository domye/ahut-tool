package models

// Class 课程信息结构体
type Class struct {
	Name        string `json:"name"`        // 课程名称
	Teacher     string `json:"teacher"`     // 教师
	Classroom   string `json:"classroom"`   // 教室
	DayOfWeek   int    `json:"dayOfWeek"`   // 星期几（1-7）
	Period      string `json:"period"`      // 节次（如"1-2"）
	WeekNumbers []int  `json:"weekNumbers"` // 周次列表（如[1,2,3...]）
}

// ClassScheduleResponse 课程表响应结构体
type ClassScheduleResponse struct {
	Classes []Class `json:"classes"`
}
