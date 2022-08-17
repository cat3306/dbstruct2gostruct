package model

type StudentGrade struct {
	Id            int64  `gorm:"column:id" json:"id"`                         // 主键id
	Class         int64  `gorm:"column:class" json:"class"`                   // 班级
	ExamCode      string `gorm:"column:exam_code" json:"exam_code"`           // 考号
	Name          string `gorm:"column:name" json:"name"`                     // 姓名
	LanguageScore int64  `gorm:"column:language_score" json:"language_score"` // 语文分数
	MathScore     int64  `gorm:"column:math_score" json:"math_score"`         // 数学分数
	EnglishScore  int64  `gorm:"column:english_score" json:"english_score"`   // 英语分数
	TotalScore    int64  `gorm:"column:total_score" json:"total_score"`       // 总分数
}

func (StudentGrade) TableName() string {
	return "student_grade"
}
