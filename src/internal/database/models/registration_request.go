package models

import (
	"time"

	"gorm.io/gorm"
)

type RegistrationRequest struct {
	ID            int64          `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Name          string         `gorm:"column:name;size:255;not null" json:"name"`
	Email         string         `gorm:"column:email;size:255;not null;uniqueIndex" json:"email"`
	Password      string         `gorm:"column:password;not null" json:"-"`
	Nip           *string        `gorm:"column:nip;size:50" json:"nip"`
	Status        string         `gorm:"column:status;size:20;not null;default:'pending'" json:"status"`
	ReviewedBy    *int64         `gorm:"column:reviewed_by" json:"reviewed_by"`
	RejectionNote *string        `gorm:"column:rejection_note;size:500" json:"rejection_note"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index:idx_reg_req_deleted_at" json:"deleted_at"`

	Reviewer *User                      `gorm:"foreignKey:ReviewedBy;references:ID" json:"reviewer,omitempty"`
	School   *RegistrationRequestSchool `gorm:"foreignKey:RegistrationRequestID;references:ID" json:"school,omitempty"`
}

func (*RegistrationRequest) TableName() string {
	return "registration_requests"
}

type RegistrationRequestSchool struct {
	ID                    int64     `gorm:"primaryKey;autoIncrement:true" json:"id"`
	RegistrationRequestID int64     `gorm:"column:registration_request_id;not null;uniqueIndex" json:"registration_request_id"`
	ExistingSchoolID      *int64    `gorm:"column:existing_school_id" json:"existing_school_id"`
	ProposedSchoolName    *string   `gorm:"column:proposed_school_name;size:255" json:"proposed_school_name"`
	Status                string    `gorm:"column:status;size:20;not null;default:'pending'" json:"status"`
	CreatedAt             time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	RegistrationRequest *RegistrationRequest `gorm:"foreignKey:RegistrationRequestID;references:ID" json:"-"`
	ExistingSchool      *MSchool             `gorm:"foreignKey:ExistingSchoolID;references:ID" json:"existing_school,omitempty"`
}

func (*RegistrationRequestSchool) TableName() string {
	return "registration_request_schools"
}
