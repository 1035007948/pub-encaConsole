package models

import (
	"time"

	"gorm.io/gorm"
)

type StatusTransition struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	EntityType      string         `gorm:"size:50;index" json:"entity_type"`
	EntityID        uint           `gorm:"index" json:"entity_id"`
	EntityNo        string         `gorm:"size:50" json:"entity_no"`
	FromStatus      string         `gorm:"size:20" json:"from_status"`
	ToStatus        string         `gorm:"size:20" json:"to_status"`
	Action          string         `gorm:"size:50" json:"action"`
	Reason          string         `gorm:"type:text" json:"reason"`
	Operator        string         `gorm:"size:100" json:"operator"`
	BatchNo         string         `gorm:"size:50" json:"batch_no"`
	Remark          string         `gorm:"type:text" json:"remark"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

func (StatusTransition) TableName() string {
	return "status_transitions"
}
