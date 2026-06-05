package services

import (
	"fmt"
	"time"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/models"
	"noise-complaint-backend/src/repositories"
)

type AnomalyService struct {
	repo      *repositories.AnomalyEventRepository
	auditRepo *repositories.AuditLogRepository
}

func NewAnomalyService() *AnomalyService {
	return &AnomalyService{
		repo:      repositories.NewAnomalyEventRepository(),
		auditRepo: repositories.NewAuditLogRepository(),
	}
}

func (s *AnomalyService) Create(event *models.AnomalyEvent, operator string) error {
	err := s.repo.Create(event)
	if err != nil {
		return err
	}

	s.createAuditLog("create", "anomaly_event", event.ID, event.EventNo, operator)
	return nil
}

func (s *AnomalyService) GetByID(id uint) (*models.AnomalyEvent, error) {
	return s.repo.FindByID(id)
}

func (s *AnomalyService) GetList(page, pageSize int, filters map[string]interface{}) (*dto.AnomalyEventListResponse, error) {
	events, total, err := s.repo.FindAll(page, pageSize, filters)
	if err != nil {
		return nil, err
	}

	items := make([]dto.AnomalyEventResponse, len(events))
	for i, e := range events {
		items[i] = dto.AnomalyEventResponse{
			ID:              e.ID,
			EventNo:         e.EventNo,
			EventName:       e.EventName,
			EventType:       e.EventType,
			Severity:        string(e.Severity),
			EntityType:      e.EntityType,
			EntityID:        e.EntityID,
			EntityNo:        e.EntityNo,
			TriggerField:    e.TriggerField,
			TriggerValue:    e.TriggerValue,
			ThresholdValue:  e.ThresholdValue,
			Description:     e.Description,
			Status:          string(e.Status),
			ResponsibleUser: e.ResponsibleUser,
			Deadline:        e.Deadline,
			ResolvedAt:      e.ResolvedAt,
			ResolutionNote:  e.ResolutionNote,
			RuleID:          e.RuleID,
			RuleNo:          e.RuleNo,
			BatchNo:         e.BatchNo,
			Remark:          e.Remark,
			CreatedAt:       e.CreatedAt,
			UpdatedAt:       e.UpdatedAt,
		}
	}

	return &dto.AnomalyEventListResponse{
		Total: int(total),
		Items: items,
	}, nil
}

func (s *AnomalyService) Resolve(id uint, req *dto.AnomalyEventResolveRequest, operator string) error {
	event, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	now := time.Now()
	event.Status = models.AnomalyEventStatusResolved
	event.ResolvedAt = &now
	event.ResolutionNote = req.ResolutionNote

	err = s.repo.Update(event)
	if err != nil {
		return err
	}

	s.createAuditLog("resolve", "anomaly_event", id, event.EventNo, operator)
	return nil
}

func (s *AnomalyService) CreateTimePeriodAnomaly(readingID uint, readingNo, triggerField, triggerValue, thresholdValue string, operator string) error {
	event := &models.AnomalyEvent{
		EventNo:        fmt.Sprintf("ANOM-%d", time.Now().UnixNano()),
		EventName:      "采样时段不合规异常",
		EventType:      "time_period_violation",
		Severity:       models.AnomalyEventSeverityHigh,
		EntityType:     "noise_reading",
		EntityID:       readingID,
		EntityNo:       readingNo,
		TriggerField:   triggerField,
		TriggerValue:   triggerValue,
		ThresholdValue: thresholdValue,
		Description:    fmt.Sprintf("噪声读数 %s 的采样时段不合规，触发字段: %s", readingNo, triggerField),
		Status:         models.AnomalyEventStatusOpen,
		ResponsibleUser: operator,
	}

	deadline := time.Now().Add(24 * time.Hour)
	event.Deadline = &deadline

	return s.Create(event, operator)
}

func (s *AnomalyService) createAuditLog(action, entityType string, entityID uint, entityNo, operator string) {
	log := &models.AuditLog{
		LogNo:      fmt.Sprintf("LOG-%d-%d", time.Now().Unix(), entityID),
		Action:     action,
		EntityType: entityType,
		EntityID:   entityID,
		EntityNo:   entityNo,
		Operator:   operator,
	}
	s.auditRepo.Create(log)
}
