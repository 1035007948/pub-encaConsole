package services

import (
	"encoding/json"
	"fmt"
	"time"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/models"
	"noise-complaint-backend/src/repositories"
)

type ArchiveService struct {
	snapshotRepo      *repositories.ArchiveSnapshotRepository
	complaintRepo     *repositories.ComplaintRepository
	samplingPointRepo *repositories.SamplingPointRepository
	readingRepo       *repositories.NoiseReadingRepository
	auditRepo         *repositories.AuditLogRepository
}

func NewArchiveService() *ArchiveService {
	return &ArchiveService{
		snapshotRepo:      repositories.NewArchiveSnapshotRepository(),
		complaintRepo:     repositories.NewComplaintRepository(),
		samplingPointRepo: repositories.NewSamplingPointRepository(),
		readingRepo:       repositories.NewNoiseReadingRepository(),
		auditRepo:         repositories.NewAuditLogRepository(),
	}
}

func (s *ArchiveService) CreateSnapshot(req *dto.ArchiveSnapshotCreateRequest, operator string) (*models.ArchiveSnapshot, error) {
	var data interface{}
	var recordCount int

	switch req.EntityType {
	case "complaint":
		complaints := make([]models.Complaint, 0)
		for _, id := range req.EntityIDs {
			c, err := s.complaintRepo.FindByID(id)
			if err == nil {
				complaints = append(complaints, *c)
			}
		}
		data = complaints
		recordCount = len(complaints)
	case "sampling_point":
		points := make([]models.SamplingPoint, 0)
		for _, id := range req.EntityIDs {
			p, err := s.samplingPointRepo.FindByID(id)
			if err == nil {
				points = append(points, *p)
			}
		}
		data = points
		recordCount = len(points)
	case "noise_reading":
		readings := make([]models.NoiseReading, 0)
		for _, id := range req.EntityIDs {
			r, err := s.readingRepo.FindByID(id)
			if err == nil {
				readings = append(readings, *r)
			}
		}
		data = readings
		recordCount = len(readings)
	default:
		return nil, fmt.Errorf("unsupported entity type: %s", req.EntityType)
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	entityIDsJSON, err := json.Marshal(req.EntityIDs)
	if err != nil {
		return nil, err
	}

	snapshot := &models.ArchiveSnapshot{
		SnapshotNo:   fmt.Sprintf("SNAP-%d", time.Now().UnixNano()),
		SnapshotName: req.SnapshotName,
		SnapshotType: req.SnapshotType,
		EntityType:   req.EntityType,
		EntityIDs:    string(entityIDsJSON),
		Data:         string(dataJSON),
		RecordCount:  recordCount,
		Operator:     operator,
		Remark:       req.Remark,
	}

	err = s.snapshotRepo.Create(snapshot)
	if err != nil {
		return nil, err
	}

	s.createAuditLog("create_snapshot", "archive_snapshot", snapshot.ID, snapshot.SnapshotNo, operator)

	return snapshot, nil
}

func (s *ArchiveService) GetSnapshotList(page, pageSize int) (*dto.ArchiveSnapshotListResponse, error) {
	snapshots, total, err := s.snapshotRepo.FindAll(page, pageSize)
	if err != nil {
		return nil, err
	}

	items := make([]dto.ArchiveSnapshotResponse, len(snapshots))
	for i, snap := range snapshots {
		items[i] = dto.ArchiveSnapshotResponse{
			ID:           snap.ID,
			SnapshotNo:   snap.SnapshotNo,
			SnapshotName: snap.SnapshotName,
			SnapshotType: snap.SnapshotType,
			EntityType:   snap.EntityType,
			EntityIDs:    snap.EntityIDs,
			Data:         snap.Data,
			RecordCount:  snap.RecordCount,
			Operator:     snap.Operator,
			BatchNo:      snap.BatchNo,
			Remark:       snap.Remark,
			CreatedAt:    snap.CreatedAt,
			UpdatedAt:    snap.UpdatedAt,
		}
	}

	return &dto.ArchiveSnapshotListResponse{
		Total: int(total),
		Items: items,
	}, nil
}

func (s *ArchiveService) Export(req *dto.ArchiveExportRequest) (*dto.ArchiveExportResponse, error) {
	var data interface{}
	var recordCount int

	switch req.EntityType {
	case "complaint":
		complaints, _, err := s.complaintRepo.FindAll(1, 1000, map[string]interface{}{
			"batch_no": req.BatchNo,
		})
		if err != nil {
			return nil, err
		}
		data = complaints
		recordCount = len(complaints)
	case "sampling_point":
		points, _, err := s.samplingPointRepo.FindAll(1, 1000, map[string]interface{}{
			"batch_no": req.BatchNo,
		})
		if err != nil {
			return nil, err
		}
		data = points
		recordCount = len(points)
	case "noise_reading":
		readings, _, err := s.readingRepo.FindAll(1, 1000, map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		data = readings
		recordCount = len(readings)
	default:
		return nil, fmt.Errorf("unsupported entity type: %s", req.EntityType)
	}

	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s_export_%s.json", req.EntityType, time.Now().Format("20060102_150405"))

	return &dto.ArchiveExportResponse{
		FileName:    fileName,
		FileSize:    int64(len(dataJSON)),
		RecordCount: recordCount,
		DownloadURL: fmt.Sprintf("/api/archive/download/%s", fileName),
	}, nil
}

func (s *ArchiveService) createAuditLog(action, entityType string, entityID uint, entityNo, operator string) {
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
