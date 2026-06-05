package models

import (
	"time"

	"gorm.io/gorm"
)

type NoiseReadingStatus string

const (
	NoiseReadingStatusDraft     NoiseReadingStatus = "draft"
	NoiseReadingStatusPending   NoiseReadingStatus = "pending"
	NoiseReadingStatusReviewing NoiseReadingStatus = "reviewing"
	NoiseReadingStatusConfirmed NoiseReadingStatus = "confirmed"
	NoiseReadingStatusArchived  NoiseReadingStatus = "archived"
)

type NoiseReading struct {
	ID               uint              `gorm:"primaryKey" json:"id"`
	ReadingNo        string            `gorm:"uniqueIndex;size:50" json:"reading_no"`
	SamplingPointID  uint              `json:"sampling_point_id"`
	PointNo          string            `gorm:"size:50" json:"point_no"`
	ComplaintID      uint              `json:"complaint_id"`
	ComplaintNo      string            `gorm:"size:50" json:"complaint_no"`
	TimePeriodID     uint              `json:"time_period_id"`
	PeriodName       string            `gorm:"size:50" json:"period_name"`
	MeasurementDate  time.Time         `json:"measurement_date"`
	MeasurementTime  string            `gorm:"size:10" json:"measurement_time"`
	Leq              float64           `json:"leq"`
	Lmax             float64           `json:"lmax"`
	Lmin             float64           `json:"lmin"`
	L10              float64           `json:"l10"`
	L90              float64           `json:"l90"`
	StandardLimit    float64           `json:"standard_limit"`
	ExceedValue      float64           `json:"exceed_value"`
	IsExceeded       bool              `json:"is_exceeded"`
	Status           NoiseReadingStatus `gorm:"size:20;default:draft" json:"status"`
	ResponsibleUser  string            `gorm:"size:100" json:"responsible_user"`
	BatchNo          string            `gorm:"size:50" json:"batch_no"`
	Remark           string            `gorm:"type:text" json:"remark"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	DeletedAt        gorm.DeletedAt    `gorm:"index" json:"-"`
}

func (NoiseReading) TableName() string {
	return "noise_readings"
}
