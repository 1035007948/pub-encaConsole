package models

import (
	"time"

	"gorm.io/gorm"
)

type AuditLog struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	LogNo       string         `gorm:"uniqueIndex;size:50" json:"log_no"`
	Action      string         `gorm:"size:100" json:"action"`
	EntityType  string         `gorm:"size:50;index" json:"entity_type"`
	EntityID    uint           `gorm:"index" json:"entity_id"`
	EntityNo    string         `gorm:"size:50" json:"entity_no"`
	Operator    string         `gorm:"size:100" json:"operator"`
	OldValue    string         `gorm:"type:text" json:"old_value"`
	NewValue    string         `gorm:"type:text" json:"new_value"`
	IP          string         `gorm:"size:50" json:"ip"`
	UserAgent   string         `gorm:"size:500" json:"user_agent"`
	BatchNo     string         `gorm:"size:50" json:"batch_no"`
	Remark      string         `gorm:"type:text" json:"remark"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (AuditLog) TableName() string {
	return "audit_logs"
}
