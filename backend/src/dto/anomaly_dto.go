package dto

import "time"

type AnomalyEventResponse struct {
	ID              uint       `json:"id"`
	EventNo         string     `json:"event_no"`
	EventName       string     `json:"event_name"`
	EventType       string     `json:"event_type"`
	Severity        string     `json:"severity"`
	EntityType      string     `json:"entity_type"`
	EntityID        uint       `json:"entity_id"`
	EntityNo        string     `json:"entity_no"`
	TriggerField    string     `json:"trigger_field"`
	TriggerValue    string     `json:"trigger_value"`
	ThresholdValue  string     `json:"threshold_value"`
	Description     string     `json:"description"`
	Status          string     `json:"status"`
	ResponsibleUser string     `json:"responsible_user"`
	Deadline        *time.Time `json:"deadline"`
	ResolvedAt      *time.Time `json:"resolved_at"`
	ResolutionNote  string     `json:"resolution_note"`
	RuleID          uint       `json:"rule_id"`
	RuleNo          string     `json:"rule_no"`
	BatchNo         string     `json:"batch_no"`
	Remark          string     `json:"remark"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type AnomalyEventListResponse struct {
	Total int                    `json:"total"`
	Items []AnomalyEventResponse `json:"items"`
}

type AnomalyEventResolveRequest struct {
	ResolutionNote string `json:"resolution_note" binding:"required"`
}
