package services

import (
	"fmt"
	"time"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/models"
	"noise-complaint-backend/src/repositories"
)

type NoiseReadingService struct {
	repo      *repositories.NoiseReadingRepository
	auditRepo *repositories.AuditLogRepository
}

func NewNoiseReadingService() *NoiseReadingService {
	return &NoiseReadingService{
		repo:      repositories.NewNoiseReadingRepository(),
		auditRepo: repositories.NewAuditLogRepository(),
	}
}

func (s *NoiseReadingService) Create(req *dto.NoiseReadingCreateRequest, operator string) (*models.NoiseReading, error) {
	measurementDate, err := time.Parse("2006-01-02", req.MeasurementDate)
	if err != nil {
		return nil, err
	}

	reading := &models.NoiseReading{
		ReadingNo:       req.ReadingNo,
		SamplingPointID: req.SamplingPointID,
		PointNo:         req.PointNo,
		ComplaintID:     req.ComplaintID,
		ComplaintNo:     req.ComplaintNo,
		TimePeriodID:    req.TimePeriodID,
		PeriodName:      req.PeriodName,
		MeasurementDate: measurementDate,
		MeasurementTime: req.MeasurementTime,
		Leq:             req.Leq,
		Lmax:            req.Lmax,
		Lmin:            req.Lmin,
		L10:             req.L10,
		L90:             req.L90,
		StandardLimit:   req.StandardLimit,
		Status:          models.NoiseReadingStatusDraft,
		ResponsibleUser: req.ResponsibleUser,
		BatchNo:         req.BatchNo,
		Remark:          req.Remark,
	}

	if reading.StandardLimit > 0 {
		reading.ExceedValue = reading.Leq - reading.StandardLimit
		reading.IsExceeded = reading.Leq > reading.StandardLimit
	}

	err = s.repo.Create(reading)
	if err != nil {
		return nil, err
	}

	s.createAuditLog("create", "noise_reading", reading.ID, reading.ReadingNo, operator)

	return reading, nil
}

func (s *NoiseReadingService) Update(id uint, req *dto.NoiseReadingUpdateRequest, operator string) (*models.NoiseReading, error) {
	reading, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.PointNo != "" {
		reading.PointNo = req.PointNo
	}
	if req.PeriodName != "" {
		reading.PeriodName = req.PeriodName
	}
	if req.MeasurementDate != "" {
		t, err := time.Parse("2006-01-02", req.MeasurementDate)
		if err == nil {
			reading.MeasurementDate = t
		}
	}
	if req.MeasurementTime != "" {
		reading.MeasurementTime = req.MeasurementTime
	}
	reading.Leq = req.Leq
	reading.Lmax = req.Lmax
	reading.Lmin = req.Lmin
	reading.L10 = req.L10
	reading.L90 = req.L90
	reading.StandardLimit = req.StandardLimit

	if reading.StandardLimit > 0 {
		reading.ExceedValue = reading.Leq - reading.StandardLimit
		reading.IsExceeded = reading.Leq > reading.StandardLimit
	}

	if req.ResponsibleUser != "" {
		reading.ResponsibleUser = req.ResponsibleUser
	}
	if req.BatchNo != "" {
		reading.BatchNo = req.BatchNo
	}
	if req.Remark != "" {
		reading.Remark = req.Remark
	}

	err = s.repo.Update(reading)
	if err != nil {
		return nil, err
	}

	s.createAuditLog("update", "noise_reading", reading.ID, reading.ReadingNo, operator)

	return reading, nil
}

func (s *NoiseReadingService) Delete(id uint, operator string) error {
	reading, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	err = s.repo.Delete(id)
	if err != nil {
		return err
	}

	s.createAuditLog("delete", "noise_reading", reading.ID, reading.ReadingNo, operator)

	return nil
}

func (s *NoiseReadingService) GetByID(id uint) (*models.NoiseReading, error) {
	return s.repo.FindByID(id)
}

func (s *NoiseReadingService) GetList(page, pageSize int, filters map[string]interface{}) (*dto.NoiseReadingListResponse, error) {
	readings, total, err := s.repo.FindAll(page, pageSize, filters)
	if err != nil {
		return nil, err
	}

	items := make([]dto.NoiseReadingResponse, len(readings))
	for i, r := range readings {
		items[i] = dto.NoiseReadingResponse{
			ID:              r.ID,
			ReadingNo:       r.ReadingNo,
			SamplingPointID: r.SamplingPointID,
			PointNo:         r.PointNo,
			ComplaintID:     r.ComplaintID,
			ComplaintNo:     r.ComplaintNo,
			TimePeriodID:    r.TimePeriodID,
			PeriodName:      r.PeriodName,
			MeasurementDate: r.MeasurementDate,
			MeasurementTime: r.MeasurementTime,
			Leq:             r.Leq,
			Lmax:            r.Lmax,
			Lmin:            r.Lmin,
			L10:             r.L10,
			L90:             r.L90,
			StandardLimit:   r.StandardLimit,
			ExceedValue:     r.ExceedValue,
			IsExceeded:      r.IsExceeded,
			Status:          string(r.Status),
			ResponsibleUser: r.ResponsibleUser,
			BatchNo:         r.BatchNo,
			Remark:          r.Remark,
			CreatedAt:       r.CreatedAt,
			UpdatedAt:       r.UpdatedAt,
		}
	}

	return &dto.NoiseReadingListResponse{
		Total: int(total),
		Items: items,
	}, nil
}

func (s *NoiseReadingService) GetBySamplingPointID(samplingPointID uint) ([]models.NoiseReading, error) {
	return s.repo.FindBySamplingPointID(samplingPointID)
}

func (s *NoiseReadingService) GetByComplaintID(complaintID uint) ([]models.NoiseReading, error) {
	return s.repo.FindByComplaintID(complaintID)
}

func (s *NoiseReadingService) createAuditLog(action, entityType string, entityID uint, entityNo, operator string) {
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
