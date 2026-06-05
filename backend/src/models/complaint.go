package models

import (
	"time"

	"gorm.io/gorm"
)

type ComplaintStatus string

const (
	ComplaintStatusDraft      ComplaintStatus = "draft"
	ComplaintStatusPending    ComplaintStatus = "pending"
	ComplaintStatusReviewing  ComplaintStatus = "reviewing"
	ComplaintStatusSupplement ComplaintStatus = "supplement"
	ComplaintStatusConfirmed  ComplaintStatus = "confirmed"
	ComplaintStatusArchived   ComplaintStatus = "archived"
	ComplaintStatusRejected   ComplaintStatus = "rejected"
)

type ComplaintLevel string

const (
	ComplaintLevelLow    ComplaintLevel = "low"
	ComplaintLevelMedium ComplaintLevel = "medium"
	ComplaintLevelHigh   ComplaintLevel = "high"
	ComplaintLevelUrgent ComplaintLevel = "urgent"
)

type Complaint struct {
	ID              uint            `gorm:"primaryKey" json:"id"`
	ComplaintNo     string          `gorm:"uniqueIndex;size:50" json:"complaint_no"`
	Title           string          `gorm:"size:200" json:"title"`
	Description     string          `gorm:"type:text" json:"description"`
	Status          ComplaintStatus `gorm:"size:20;default:draft" json:"status"`
	Level           ComplaintLevel  `gorm:"size:20;default:medium" json:"level"`
	ComplainantName string          `gorm:"size:100" json:"complainant_name"`
	ComplainantTel  string          `gorm:"size:20" json:"complainant_tel"`
	EnterpriseName  string          `gorm:"size:200" json:"enterprise_name"`
	EnterpriseAddr  string          `gorm:"size:300" json:"enterprise_addr"`
	ResponsibleUser string          `gorm:"size:100" json:"responsible_user"`
	BatchNo         string          `gorm:"size:50" json:"batch_no"`
	Priority        int             `gorm:"default:0" json:"priority"`
	Remark          string          `gorm:"type:text" json:"remark"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       gorm.DeletedAt  `gorm:"index" json:"-"`
}

func (Complaint) TableName() string {
	return "complaints"
}
