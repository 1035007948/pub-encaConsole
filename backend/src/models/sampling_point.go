package models

import (
	"time"

	"gorm.io/gorm"
)

type SamplingPointStatus string

const (
	SamplingPointStatusDraft     SamplingPointStatus = "draft"
	SamplingPointStatusScheduled SamplingPointStatus = "scheduled"
	SamplingPointStatusSampling  SamplingPointStatus = "sampling"
	SamplingPointStatusCompleted SamplingPointStatus = "completed"
	SamplingPointStatusArchived  SamplingPointStatus = "archived"
)

type SamplingPoint struct {
	ID                uint                `gorm:"primaryKey" json:"id"`
	PointNo           string              `gorm:"uniqueIndex;size:50" json:"point_no"`
	PointName         string              `gorm:"size:200" json:"point_name"`
	Address           string              `gorm:"size:300" json:"address"`
	Longitude         float64             `json:"longitude"`
	Latitude          float64             `json:"latitude"`
	Status            SamplingPointStatus `gorm:"size:20;default:draft" json:"status"`
	ComplaintID       uint                `json:"complaint_id"`
	ComplaintNo       string              `gorm:"size:50" json:"complaint_no"`
	ResponsibleUser   string              `gorm:"size:100" json:"responsible_user"`
	ScheduledDate     *time.Time          `json:"scheduled_date"`
	ScheduledTimeFrom string              `gorm:"size:10" json:"scheduled_time_from"`
	ScheduledTimeTo   string              `gorm:"size:10" json:"scheduled_time_to"`
	BatchNo           string              `gorm:"size:50" json:"batch_no"`
	Remark            string              `gorm:"type:text" json:"remark"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	DeletedAt         gorm.DeletedAt      `gorm:"index" json:"-"`
}

func (SamplingPoint) TableName() string {
	return "sampling_points"
}
