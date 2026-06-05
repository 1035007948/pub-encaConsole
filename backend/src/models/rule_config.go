package models

import (
	"time"

	"gorm.io/gorm"
)

type RuleConfigStatus string

const (
	RuleConfigStatusActive   RuleConfigStatus = "active"
	RuleConfigStatusInactive RuleConfigStatus = "inactive"
)

type RuleConfig struct {
	ID          uint             `gorm:"primaryKey" json:"id"`
	RuleNo      string           `gorm:"uniqueIndex;size:50" json:"rule_no"`
	RuleName    string           `gorm:"size:200" json:"rule_name"`
	RuleType    string           `gorm:"size:50" json:"rule_type"`
	Description string           `gorm:"type:text" json:"description"`
	Conditions  string           `gorm:"type:text" json:"conditions"`
	Threshold   float64          `json:"threshold"`
	Action      string           `gorm:"size:100" json:"action"`
	Priority    int              `json:"priority"`
	Status      RuleConfigStatus `gorm:"size:20;default:active" json:"status"`
	BatchNo     string           `gorm:"size:50" json:"batch_no"`
	Remark      string           `gorm:"type:text" json:"remark"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `gorm:"index" json:"-"`
}

func (RuleConfig) TableName() string {
	return "rule_configs"
}
