package models

import (
	"time"

	"gorm.io/gorm"
)

type AnomalyEventStatus string

const (
	AnomalyEventStatusOpen      AnomalyEventStatus = "open"
	AnomalyEventStatusProcessing AnomalyEventStatus = "processing"
	AnomalyEventStatusResolved  AnomalyEventStatus = "resolved"
	AnomalyEventStatusClosed    AnomalyEventStatus = "closed"
)

type AnomalyEventSeverity string

const (
	AnomalyEventSeverityLow      AnomalyEventSeverity = "low"
	AnomalyEventSeverityMedium   AnomalyEventSeverity = "medium"
	AnomalyEventSeverityHigh     AnomalyEventSeverity = "high"
	AnomalyEventSeverityCritical AnomalyEventSeverity = "critical"
)

type AnomalyEvent struct {
	ID               uint                 `gorm:"primaryKey" json:"id"`
	EventNo          string               `gorm:"uniqueIndex;size:50" json:"event_no"`
	EventName        string               `gorm:"size:200" json:"event_name"`
	EventType        string               `gorm:"size:50" json:"event_type"`
	Severity         AnomalyEventSeverity `gorm:"size:20" json:"severity"`
	EntityType       string               `gorm:"size:50" json:"entity_type"`
	EntityID         uint                 `json:"entity_id"`
	EntityNo         string               `gorm:"size:50" json:"entity_no"`
	TriggerField     string               `gorm:"size:100" json:"trigger_field"`
	TriggerValue     string               `gorm:"size:200" json:"trigger_value"`
	ThresholdValue   string               `gorm:"size:200" json:"threshold_value"`
	Description      string               `gorm:"type:text" json:"description"`
	Status           AnomalyEventStatus   `gorm:"size:20;default:open" json:"status"`
	ResponsibleUser  string               `gorm:"size:100" json:"responsible_user"`
	Deadline         *time.Time           `json:"deadline"`
	ResolvedAt       *time.Time           `json:"resolved_at"`
	ResolutionNote   string               `gorm:"type:text" json:"resolution_note"`
	RuleID           uint                 `json:"rule_id"`
	RuleNo           string               `gorm:"size:50" json:"rule_no"`
	BatchNo          string               `gorm:"size:50" json:"batch_no"`
	Remark           string               `gorm:"type:text" json:"remark"`
	CreatedAt        time.Time            `json:"created_at"`
	UpdatedAt        time.Time            `json:"updated_at"`
	DeletedAt        gorm.DeletedAt       `gorm:"index" json:"-"`
}

func (AnomalyEvent) TableName() string {
	return "anomaly_events"
}
