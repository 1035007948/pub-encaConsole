package models

import (
	"time"

	"gorm.io/gorm"
)

type RetestRecordStatus string

const (
	RetestRecordStatusPending   RetestRecordStatus = "pending"
	RetestRecordStatusTesting   RetestRecordStatus = "testing"
	RetestRecordStatusPassed    RetestRecordStatus = "passed"
	RetestRecordStatusFailed    RetestRecordStatus = "failed"
	RetestRecordStatusArchived  RetestRecordStatus = "archived"
)

type RetestRecord struct {
	ID               uint               `gorm:"primaryKey" json:"id"`
	RetestNo         string             `gorm:"uniqueIndex;size:50" json:"retest_no"`
	ComplaintID      uint               `json:"complaint_id"`
	ComplaintNo      string             `gorm:"size:50" json:"complaint_no"`
	SamplingPointID  uint               `json:"sampling_point_id"`
	PointNo          string             `gorm:"size:50" json:"point_no"`
	MeasureID        uint               `json:"measure_id"`
	MeasureNo        string             `gorm:"size:50" json:"measure_no"`
	RetestDate       time.Time          `json:"retest_date"`
	OriginalLeq      float64            `json:"original_leq"`
	RetestLeq        float64            `json:"retest_leq"`
	ReductionValue   float64            `json:"reduction_value"`
	IsPassed         bool               `json:"is_passed"`
	Status           RetestRecordStatus `gorm:"size:20;default:pending" json:"status"`
	Conclusion       string             `gorm:"type:text" json:"conclusion"`
	ResponsibleUser  string             `gorm:"size:100" json:"responsible_user"`
	BatchNo          string             `gorm:"size:50" json:"batch_no"`
	Remark           string             `gorm:"type:text" json:"remark"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	DeletedAt        gorm.DeletedAt     `gorm:"index" json:"-"`
}

func (RetestRecord) TableName() string {
	return "retest_records"
}
