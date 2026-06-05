package models

import (
	"time"

	"gorm.io/gorm"
)

type TimePeriodType string

const (
	TimePeriodTypeDay   TimePeriodType = "day"
	TimePeriodTypeNight TimePeriodType = "night"
)

type TimePeriodStatus string

const (
	TimePeriodStatusActive   TimePeriodStatus = "active"
	TimePeriodStatusInactive TimePeriodStatus = "inactive"
)

type TimePeriod struct {
	ID          uint             `gorm:"primaryKey" json:"id"`
	PeriodNo    string           `gorm:"uniqueIndex;size:50" json:"period_no"`
	PeriodName  string           `gorm:"size:100" json:"period_name"`
	PeriodType  TimePeriodType   `gorm:"size:20" json:"period_type"`
	TimeFrom    string           `gorm:"size:10" json:"time_from"`
	TimeTo      string           `gorm:"size:10" json:"time_to"`
	DayLimit    float64          `json:"day_limit"`
	NightLimit  float64          `json:"night_limit"`
	Status      TimePeriodStatus `gorm:"size:20;default:active" json:"status"`
	Description string           `gorm:"type:text" json:"description"`
	BatchNo     string           `gorm:"size:50" json:"batch_no"`
	Remark      string           `gorm:"type:text" json:"remark"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `gorm:"index" json:"-"`
}

func (TimePeriod) TableName() string {
	return "time_periods"
}
