package models

import "time"

type CodeQuestion struct {
	ID           int64      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	SubLessonID  int64      `gorm:"column:sub_lesson_id" json:"sub_lesson_id"`
	CodeQuestion string     `gorm:"column:code_question;type:text" json:"code_question"`
	Image        string     `gorm:"type:text" json:"image"`
	Score        int        `gorm:"column:score" json:"score"`
	Hint         string     `gorm:"type:text" json:"hint"`
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// Foreign Key Relationships
	SubLesson *MSubLesson `gorm:"foreignKey:SubLessonID;references:ID" json:"sub_lesson"`
	CodeAnswers []TCodeAnswer `gorm:"foreignKey:CodeQuestionID;references:ID" json:"code_answers"`
	EssayQuestions []EssayQuestion `gorm:"foreignKey:CodeQuestionID;references:ID" json:"essay_questions"`
	CodeHistoryLogs []TCodeHistoryLogs `gorm:"foreignKey:CodeQuestionID;references:ID" json:"code_history_logs"`
}

func (*CodeQuestion) TableName() string {
	return "t_code_question"
}
