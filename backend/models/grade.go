package models

// Grade 成绩信息结构体
type Grade struct {
	Index        int     // 序号
	Semester     string  // 开课学期
	CourseID     string  // 课程编号
	CourseName   string  // 课程名称
	GroupName    string  // 分组名
	Score        string  // 成绩
	ScoreFlag    string  // 成绩标识
	Credit       float64 // 学分
	TotalHours   int     // 总学时
	GPA          float64 // 绩点
	RetakeSem    string  // 补重学期
	ExamMode     string  // 考核方式
	ExamType     string  // 考试性质
	CourseAttr   string  // 课程属性
	CourseNature string  // 课程性质
	GEType       string  // 通选课类别
}

// GradeSummary 成绩汇总信息
type GradeSummary struct {
	CourseCount int     // 所修门数
	TotalCredit float64 // 所修总学分
	AvgGPA      float64 // 平均学分绩点
	AvgScore    float64 // 平均成绩
}
