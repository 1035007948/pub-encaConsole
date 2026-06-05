package models

import (
	"time"

	"gorm.io/gorm"
)

type RectificationMeasureStatus string

const (
	RectificationMeasureStatusDraft      RectificationMeasureStatus = "draft"
	RectificationMeasureStatusPending    RectificationMeasureStatus = "pending"
	RectificationMeasureStatusImplementing RectificationMeasureStatus = "implementing"
	RectificationMeasureStatusCompleted  RectificationMeasureStatus = "completed"
	RectificationMeasureStatusVerified   RectificationMeasureStatus = "verified"
	RectificationMeasureStatusArchived   RectificationMeasureStatus = "archived"
)

type RectificationMeasure struct {
	ID                uint                        `gorm:"primaryKey" json:"id"`
	MeasureNo         string                      `gorm:"uniqueIndex;size:50" json:"measure_no"`
	MeasureName       string                      `gorm:"size:200" json:"measure_name"`
	Description       string                      `gorm:"type:text" json:"description"`
	ComplaintID       uint                        `json:"complaint_id"`
	ComplaintNo       string                      `gorm:"size:50" json:"complaint_no"`
	EnterpriseName    string                      `gorm:"size:200" json:"enterprise_name"`
	ResponsibleUser   string                      `gorm:"size:100" json:"responsible_user"`
	Deadline          *time.Time                  `json:"deadline"`
	CompletedAt       *time.Time                  `json:"completed_at"`
	Status            RectificationMeasureStatus  `gorm:"size:20;default:draft" json:"status"`
	Effectiveness     string                      `gorm:"size:50" json:"effectiveness"`
	BatchNo           string                      `gorm:"size:50" json:"batch_no"`
	Remark            string                      `gorm:"type:text" json:"remark"`
	CreatedAt         time.Time                   `json:"created_at"`
	UpdatedAt         time.Time                   `json:"updated_at"`
	DeletedAt         gorm.DeletedAt              `gorm:"index" json:"-"`
}

func (RectificationMeasure) TableName() string {
	return "rectification_measures"
}
