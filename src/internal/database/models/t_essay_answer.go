package models

import "time"

type TEssayAnswer struct {
	ID                  int64      `gorm:"primaryKey;autoIncrement:true;type:serial" json:"id"`
	UserID              int64      `gorm:"column:user_id" json:"user_id"`
	EssayQuestionID     int64      `gorm:"column:essay_question_id" json:"essay_question_id"`
	Answer              string     `gorm:"type:text" json:"answer"`
	KonteksPenjelasan   string     `gorm:"column:konteks_penjelasan;type:text" json:"konteks_penjelasan"`
	Keruntutan          int        `gorm:"column:keruntutan" json:"keruntutan"`
	Kebenaran           int        `gorm:"column:kebenaran" json:"kebenaran"`
	TeacherNotes        string     `gorm:"column:teacher_notes;type:text" json:"teacher_notes"`
	IsApprovedByTeacher bool       `gorm:"column:is_approved_by_teacher;default:false" json:"is_approved_by_teacher"`
	IsActive            bool       `gorm:"column:isactive;default:true" json:"isactive"`
	CreatedAt           time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt           *time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	UserCreate          int64      `gorm:"column:user_create" json:"user_create"`
	UserUpdate          int64      `gorm:"column:user_update" json:"user_update"`

	// Foreign Key Relationships
	User          *User          `gorm:"foreignKey:UserID;references:ID" json:"user"`
	EssayQuestion *TEssayQuestion `gorm:"foreignKey:EssayQuestionID;references:ID" json:"essay_question"`
}

func (*TEssayAnswer) TableName() string {
	return "t_essay_answer"
}
