package dto

type MRoleResponseDto struct {
	ID       int64  `json:"id"`
	RoleName string `json:"role_name"`
	IsActive bool   `json:"is_active"`
}

type MClassResponseDto struct {
	ID         int64  `json:"id"`
	TeacherID  int64  `json:"teacher_id"`
	ClassName  string `json:"class_name"`
	SchoolName string `json:"school_name"`
	ClassCode  string `json:"class_code"`
	IsActive   bool   `json:"is_active"`
}
