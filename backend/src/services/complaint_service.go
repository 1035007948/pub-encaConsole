package services

import (
	"errors"
	"fmt"
	"time"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/models"
	"noise-complaint-backend/src/repositories"
)

type ComplaintService struct {
	repo          *repositories.ComplaintRepository
	transitionRepo *repositories.StatusTransitionRepository
	auditRepo     *repositories.AuditLogRepository
}

func NewComplaintService() *ComplaintService {
	return &ComplaintService{
		repo:          repositories.NewComplaintRepository(),
		transitionRepo: repositories.NewStatusTransitionRepository(),
		auditRepo:     repositories.NewAuditLogRepository(),
	}
}

func (s *ComplaintService) Create(req *dto.ComplaintCreateRequest, operator string) (*models.Complaint, error) {
	complaint := &models.Complaint{
		ComplaintNo:     req.ComplaintNo,
		Title:           req.Title,
		Description:     req.Description,
		Status:          models.ComplaintStatusDraft,
		Level:           models.ComplaintLevel(req.Level),
		ComplainantName: req.ComplainantName,
		ComplainantTel:  req.ComplainantTel,
		EnterpriseName:  req.EnterpriseName,
		EnterpriseAddr:  req.EnterpriseAddr,
		ResponsibleUser: req.ResponsibleUser,
		BatchNo:         req.BatchNo,
		Remark:          req.Remark,
	}

	if complaint.Level == "" {
		complaint.Level = models.ComplaintLevelMedium
	}

	err := s.repo.Create(complaint)
	if err != nil {
		return nil, err
	}

	s.createAuditLog("create", "complaint", complaint.ID, complaint.ComplaintNo, "", "", operator)

	return complaint, nil
}

func (s *ComplaintService) Update(id uint, req *dto.ComplaintUpdateRequest, operator string) (*models.Complaint, error) {
	complaint, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		complaint.Title = req.Title
	}
	if req.Description != "" {
		complaint.Description = req.Description
	}
	if req.Level != "" {
		complaint.Level = models.ComplaintLevel(req.Level)
	}
	if req.ComplainantName != "" {
		complaint.ComplainantName = req.ComplainantName
	}
	if req.ComplainantTel != "" {
		complaint.ComplainantTel = req.ComplainantTel
	}
	if req.EnterpriseName != "" {
		complaint.EnterpriseName = req.EnterpriseName
	}
	if req.EnterpriseAddr != "" {
		complaint.EnterpriseAddr = req.EnterpriseAddr
	}
	if req.ResponsibleUser != "" {
		complaint.ResponsibleUser = req.ResponsibleUser
	}
	if req.BatchNo != "" {
		complaint.BatchNo = req.BatchNo
	}
	if req.Remark != "" {
		complaint.Remark = req.Remark
	}

	err = s.repo.Update(complaint)
	if err != nil {
		return nil, err
	}

	s.createAuditLog("update", "complaint", complaint.ID, complaint.ComplaintNo, "", "", operator)

	return complaint, nil
}

func (s *ComplaintService) Delete(id uint, operator string) error {
	complaint, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	err = s.repo.Delete(id)
	if err != nil {
		return err
	}

	s.createAuditLog("delete", "complaint", complaint.ID, complaint.ComplaintNo, "", "", operator)

	return nil
}

func (s *ComplaintService) GetByID(id uint) (*models.Complaint, error) {
	return s.repo.FindByID(id)
}

func (s *ComplaintService) GetList(page, pageSize int, filters map[string]interface{}) (*dto.ComplaintListResponse, error) {
	complaints, total, err := s.repo.FindAll(page, pageSize, filters)
	if err != nil {
		return nil, err
	}

	items := make([]dto.ComplaintResponse, len(complaints))
	for i, c := range complaints {
		items[i] = dto.ComplaintResponse{
			ID:              c.ID,
			ComplaintNo:     c.ComplaintNo,
			Title:           c.Title,
			Description:     c.Description,
			Status:          string(c.Status),
			Level:           string(c.Level),
			ComplainantName: c.ComplainantName,
			ComplainantTel:  c.ComplainantTel,
			EnterpriseName:  c.EnterpriseName,
			EnterpriseAddr:  c.EnterpriseAddr,
			ResponsibleUser: c.ResponsibleUser,
			BatchNo:         c.BatchNo,
			Priority:        c.Priority,
			Remark:          c.Remark,
			CreatedAt:       c.CreatedAt,
			UpdatedAt:       c.UpdatedAt,
		}
	}

	return &dto.ComplaintListResponse{
		Total: int(total),
		Items: items,
	}, nil
}

func (s *ComplaintService) Transition(id uint, req *dto.ComplaintTransitionRequest, operator string) (*models.Complaint, error) {
	complaint, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	fromStatus := complaint.Status
	toStatus := models.ComplaintStatus(req.ToStatus)

	if !s.isValidTransition(fromStatus, toStatus) {
		return nil, errors.New("invalid status transition")
	}

	if toStatus == models.ComplaintStatusRejected && req.Reason == "" {
		return nil, errors.New("reason is required when rejecting")
	}

	err = s.repo.UpdateStatus(id, toStatus)
	if err != nil {
		return nil, err
	}

	transition := &models.StatusTransition{
		EntityType: "complaint",
		EntityID:   id,
		EntityNo:   complaint.ComplaintNo,
		FromStatus: string(fromStatus),
		ToStatus:   string(toStatus),
		Action:     fmt.Sprintf("transition from %s to %s", fromStatus, toStatus),
		Reason:     req.Reason,
		Operator:   operator,
	}
	s.transitionRepo.Create(transition)

	s.createAuditLog("transition", "complaint", id, complaint.ComplaintNo, string(fromStatus), string(toStatus), operator)

	complaint.Status = toStatus
	return complaint, nil
}

func (s *ComplaintService) isValidTransition(from, to models.ComplaintStatus) bool {
	validTransitions := map[models.ComplaintStatus][]models.ComplaintStatus{
		models.ComplaintStatusDraft:      {models.ComplaintStatusPending, models.ComplaintStatusArchived},
		models.ComplaintStatusPending:    {models.ComplaintStatusReviewing, models.ComplaintStatusSupplement, models.ComplaintStatusRejected},
		models.ComplaintStatusReviewing:  {models.ComplaintStatusConfirmed, models.ComplaintStatusSupplement, models.ComplaintStatusRejected},
		models.ComplaintStatusSupplement: {models.ComplaintStatusPending, models.ComplaintStatusReviewing},
		models.ComplaintStatusConfirmed:  {models.ComplaintStatusArchived},
		models.ComplaintStatusRejected:   {models.ComplaintStatusPending},
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

func (s *ComplaintService) UpdatePriority(id uint, priority int) error {
	return s.repo.UpdatePriority(id, priority)
}

func (s *ComplaintService) createAuditLog(action, entityType string, entityID uint, entityNo, oldValue, newValue, operator string) {
	log := &models.AuditLog{
		LogNo:      fmt.Sprintf("LOG-%d-%d", time.Now().Unix(), entityID),
		Action:     action,
		EntityType: entityType,
		EntityID:   entityID,
		EntityNo:   entityNo,
		Operator:   operator,
		OldValue:   oldValue,
		NewValue:   newValue,
	}
	s.auditRepo.Create(log)
}
