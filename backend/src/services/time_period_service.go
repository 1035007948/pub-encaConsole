package services

import (
	"fmt"
	"strings"
	"time"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/models"
	"noise-complaint-backend/src/repositories"
)

type TimePeriodService struct {
	repo      *repositories.TimePeriodRepository
	auditRepo *repositories.AuditLogRepository
}

func NewTimePeriodService() *TimePeriodService {
	return &TimePeriodService{
		repo:      repositories.NewTimePeriodRepository(),
		auditRepo: repositories.NewAuditLogRepository(),
	}
}

func (s *TimePeriodService) Create(req *dto.TimePeriodCreateRequest, operator string) (*models.TimePeriod, error) {
	period := &models.TimePeriod{
		PeriodNo:    req.PeriodNo,
		PeriodName:  req.PeriodName,
		PeriodType:  models.TimePeriodType(req.PeriodType),
		TimeFrom:    req.TimeFrom,
		TimeTo:      req.TimeTo,
		DayLimit:    req.DayLimit,
		NightLimit:  req.NightLimit,
		Description: req.Description,
		Status:      models.TimePeriodStatusActive,
		BatchNo:     req.BatchNo,
		Remark:      req.Remark,
	}

	err := s.repo.Create(period)
	if err != nil {
		return nil, err
	}

	s.createAuditLog("create", "time_period", period.ID, period.PeriodNo, operator)

	return period, nil
}

func (s *TimePeriodService) Update(id uint, req *dto.TimePeriodUpdateRequest, operator string) (*models.TimePeriod, error) {
	period, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.PeriodName != "" {
		period.PeriodName = req.PeriodName
	}
	if req.TimeFrom != "" {
		period.TimeFrom = req.TimeFrom
	}
	if req.TimeTo != "" {
		period.TimeTo = req.TimeTo
	}
	period.DayLimit = req.DayLimit
	period.NightLimit = req.NightLimit
	if req.Description != "" {
		period.Description = req.Description
	}
	if req.Status != "" {
		period.Status = models.TimePeriodStatus(req.Status)
	}
	if req.BatchNo != "" {
		period.BatchNo = req.BatchNo
	}
	if req.Remark != "" {
		period.Remark = req.Remark
	}

	err = s.repo.Update(period)
	if err != nil {
		return nil, err
	}

	s.createAuditLog("update", "time_period", period.ID, period.PeriodNo, operator)

	return period, nil
}

func (s *TimePeriodService) Delete(id uint, operator string) error {
	period, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	err = s.repo.Delete(id)
	if err != nil {
		return err
	}

	s.createAuditLog("delete", "time_period", period.ID, period.PeriodNo, operator)

	return nil
}

func (s *TimePeriodService) GetByID(id uint) (*models.TimePeriod, error) {
	return s.repo.FindByID(id)
}

func (s *TimePeriodService) GetList(page, pageSize int, filters map[string]interface{}) (*dto.TimePeriodListResponse, error) {
	periods, total, err := s.repo.FindAll(page, pageSize, filters)
	if err != nil {
		return nil, err
	}

	items := make([]dto.TimePeriodResponse, len(periods))
	for i, p := range periods {
		items[i] = dto.TimePeriodResponse{
			ID:          p.ID,
			PeriodNo:    p.PeriodNo,
			PeriodName:  p.PeriodName,
			PeriodType:  string(p.PeriodType),
			TimeFrom:    p.TimeFrom,
			TimeTo:      p.TimeTo,
			DayLimit:    p.DayLimit,
			NightLimit:  p.NightLimit,
			Status:      string(p.Status),
			Description: p.Description,
			BatchNo:     p.BatchNo,
			Remark:      p.Remark,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		}
	}

	return &dto.TimePeriodListResponse{
		Total: int(total),
		Items: items,
	}, nil
}

func (s *TimePeriodService) BatchImport(req *dto.TimePeriodBatchImportRequest, operator string) (*dto.TimePeriodBatchImportResponse, error) {
	response := &dto.TimePeriodBatchImportResponse{
		Total:    len(req.Items),
		Success:  0,
		Failed:   0,
		Errors:   []string{},
		Warnings: []string{},
	}

	var validPeriods []models.TimePeriod
	for i, item := range req.Items {
		if item.PeriodNo == "" {
			response.Failed++
			response.Errors = append(response.Errors, fmt.Sprintf("Row %d: period_no is required", i+1))
			continue
		}

		existing, _ := s.repo.FindByNo(item.PeriodNo)
		if existing != nil {
			response.Warnings = append(response.Warnings, fmt.Sprintf("Row %d: period_no %s already exists, will be skipped", i+1, item.PeriodNo))
			continue
		}

		if !s.isValidTimeRange(item.TimeFrom, item.TimeTo) {
			response.Warnings = append(response.Warnings, fmt.Sprintf("Row %d: time range %s-%s may be invalid", i+1, item.TimeFrom, item.TimeTo))
		}

		period := models.TimePeriod{
			PeriodNo:    item.PeriodNo,
			PeriodName:  item.PeriodName,
			PeriodType:  models.TimePeriodType(item.PeriodType),
			TimeFrom:    item.TimeFrom,
			TimeTo:      item.TimeTo,
			DayLimit:    item.DayLimit,
			NightLimit:  item.NightLimit,
			Description: item.Description,
			Status:      models.TimePeriodStatusActive,
			BatchNo:     item.BatchNo,
			Remark:      item.Remark,
		}
		validPeriods = append(validPeriods, period)
	}

	if len(validPeriods) > 0 {
		err := s.repo.BatchCreate(validPeriods)
		if err != nil {
			return nil, err
		}
		response.Success = len(validPeriods)
	}

	return response, nil
}

func (s *TimePeriodService) isValidTimeRange(timeFrom, timeTo string) bool {
	fromParts := strings.Split(timeFrom, ":")
	toParts := strings.Split(timeTo, ":")
	return len(fromParts) == 2 && len(toParts) == 2
}

func (s *TimePeriodService) createAuditLog(action, entityType string, entityID uint, entityNo, operator string) {
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
