package services

import (
	"errors"
	"fmt"
	"time"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/models"
	"noise-complaint-backend/src/repositories"
)

type SamplingPointService struct {
	repo           *repositories.SamplingPointRepository
	complaintRepo  *repositories.ComplaintRepository
	transitionRepo *repositories.StatusTransitionRepository
	auditRepo      *repositories.AuditLogRepository
}

func NewSamplingPointService() *SamplingPointService {
	return &SamplingPointService{
		repo:           repositories.NewSamplingPointRepository(),
		complaintRepo:  repositories.NewComplaintRepository(),
		transitionRepo: repositories.NewStatusTransitionRepository(),
		auditRepo:      repositories.NewAuditLogRepository(),
	}
}

func (s *SamplingPointService) Create(req *dto.SamplingPointCreateRequest, operator string) (*models.SamplingPoint, error) {
	if req.ComplaintID > 0 {
		complaint, err := s.complaintRepo.FindByID(req.ComplaintID)
		if err != nil {
			return nil, errors.New("complaint not found")
		}
		if complaint.Status == models.ComplaintStatusDraft {
			return nil, errors.New("cannot create sampling point for draft complaint")
		}
	}

	point := &models.SamplingPoint{
		PointNo:         req.PointNo,
		PointName:       req.PointName,
		Address:         req.Address,
		Longitude:       req.Longitude,
		Latitude:        req.Latitude,
		Status:          models.SamplingPointStatusDraft,
		ComplaintID:     req.ComplaintID,
		ComplaintNo:     req.ComplaintNo,
		ResponsibleUser: req.ResponsibleUser,
		BatchNo:         req.BatchNo,
		Remark:          req.Remark,
	}

	if req.ScheduledDate != "" {
		t, err := time.Parse("2006-01-02", req.ScheduledDate)
		if err == nil {
			point.ScheduledDate = &t
		}
	}
	point.ScheduledTimeFrom = req.ScheduledTimeFrom
	point.ScheduledTimeTo = req.ScheduledTimeTo

	err := s.repo.Create(point)
	if err != nil {
		return nil, err
	}

	s.createAuditLog("create", "sampling_point", point.ID, point.PointNo, operator)

	return point, nil
}

func (s *SamplingPointService) Update(id uint, req *dto.SamplingPointUpdateRequest, operator string) (*models.SamplingPoint, error) {
	point, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.PointName != "" {
		point.PointName = req.PointName
	}
	if req.Address != "" {
		point.Address = req.Address
	}
	point.Longitude = req.Longitude
	point.Latitude = req.Latitude
	if req.ResponsibleUser != "" {
		point.ResponsibleUser = req.ResponsibleUser
	}
	if req.ScheduledDate != "" {
		t, err := time.Parse("2006-01-02", req.ScheduledDate)
		if err == nil {
			point.ScheduledDate = &t
		}
	}
	point.ScheduledTimeFrom = req.ScheduledTimeFrom
	point.ScheduledTimeTo = req.ScheduledTimeTo
	if req.BatchNo != "" {
		point.BatchNo = req.BatchNo
	}
	if req.Remark != "" {
		point.Remark = req.Remark
	}

	err = s.repo.Update(point)
	if err != nil {
		return nil, err
	}

	s.createAuditLog("update", "sampling_point", point.ID, point.PointNo, operator)

	return point, nil
}

func (s *SamplingPointService) Delete(id uint, operator string) error {
	point, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	err = s.repo.Delete(id)
	if err != nil {
		return err
	}

	s.createAuditLog("delete", "sampling_point", point.ID, point.PointNo, operator)

	return nil
}

func (s *SamplingPointService) GetByID(id uint) (*models.SamplingPoint, error) {
	return s.repo.FindByID(id)
}

func (s *SamplingPointService) GetList(page, pageSize int, filters map[string]interface{}) (*dto.SamplingPointListResponse, error) {
	points, total, err := s.repo.FindAll(page, pageSize, filters)
	if err != nil {
		return nil, err
	}

	items := make([]dto.SamplingPointResponse, len(points))
	for i, p := range points {
		items[i] = dto.SamplingPointResponse{
			ID:                p.ID,
			PointNo:           p.PointNo,
			PointName:         p.PointName,
			Address:           p.Address,
			Longitude:         p.Longitude,
			Latitude:          p.Latitude,
			Status:            string(p.Status),
			ComplaintID:       p.ComplaintID,
			ComplaintNo:       p.ComplaintNo,
			ResponsibleUser:   p.ResponsibleUser,
			ScheduledDate:     p.ScheduledDate,
			ScheduledTimeFrom: p.ScheduledTimeFrom,
			ScheduledTimeTo:   p.ScheduledTimeTo,
			BatchNo:           p.BatchNo,
			Remark:            p.Remark,
			CreatedAt:         p.CreatedAt,
			UpdatedAt:         p.UpdatedAt,
		}
	}

	return &dto.SamplingPointListResponse{
		Total: int(total),
		Items: items,
	}, nil
}

func (s *SamplingPointService) Transition(id uint, req *dto.SamplingPointTransitionRequest, operator string) (*models.SamplingPoint, error) {
	point, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	fromStatus := point.Status
	toStatus := models.SamplingPointStatus(req.ToStatus)

	if !s.isValidTransition(fromStatus, toStatus) {
		return nil, errors.New("invalid status transition")
	}

	err = s.repo.UpdateStatus(id, toStatus)
	if err != nil {
		return nil, err
	}

	transition := &models.StatusTransition{
		EntityType: "sampling_point",
		EntityID:   id,
		EntityNo:   point.PointNo,
		FromStatus: string(fromStatus),
		ToStatus:   string(toStatus),
		Action:     fmt.Sprintf("transition from %s to %s", fromStatus, toStatus),
		Reason:     req.Reason,
		Operator:   operator,
	}
	s.transitionRepo.Create(transition)

	s.createAuditLog("transition", "sampling_point", id, point.PointNo, operator)

	point.Status = toStatus
	return point, nil
}

func (s *SamplingPointService) isValidTransition(from, to models.SamplingPointStatus) bool {
	validTransitions := map[models.SamplingPointStatus][]models.SamplingPointStatus{
		models.SamplingPointStatusDraft:     {models.SamplingPointStatusScheduled},
		models.SamplingPointStatusScheduled: {models.SamplingPointStatusSampling, models.SamplingPointStatusDraft},
		models.SamplingPointStatusSampling:  {models.SamplingPointStatusCompleted, models.SamplingPointStatusScheduled},
		models.SamplingPointStatusCompleted: {models.SamplingPointStatusArchived},
	}

	allowed, exists := validTransitions[from]
	if !exists {
		return false
	}

	for _, status := range allowed {
		if status == to {
			return true
		}
	}

	return false
}

func (s *SamplingPointService) ValidateConsistency(pointID uint) (bool, []string, error) {
	point, err := s.repo.FindByID(pointID)
	if err != nil {
		return false, nil, err
	}

	var missingFields []string

	if point.PointName == "" {
		missingFields = append(missingFields, "point_name")
	}
	if point.Address == "" {
		missingFields = append(missingFields, "address")
	}
	if point.ComplaintID == 0 {
		missingFields = append(missingFields, "complaint_id")
	}
	if point.ScheduledDate == nil {
		missingFields = append(missingFields, "scheduled_date")
	}

	return len(missingFields) == 0, missingFields, nil
}

func (s *SamplingPointService) createAuditLog(action, entityType string, entityID uint, entityNo, operator string) {
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
