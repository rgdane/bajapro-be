package models

type Leaderboard struct {
	UserID       int64  `gorm:"column:user_id"`
	Name         string `gorm:"column:name"`
	BadgeName    string `gorm:"column:badge_name"`
	BadgeImage   string `gorm:"column:badge_image"`
	ReadingScore int    `gorm:"column:reading_score"`
	CodingScore  int    `gorm:"column:coding_score"`
	EssayScore   int    `gorm:"column:essay_score"`
	TotalScore   int    `gorm:"column:total_score"`
}
