package services

import (
	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/repositories"
)

type StatisticsService struct {
	complaintRepo        *repositories.ComplaintRepository
	samplingPointRepo    *repositories.SamplingPointRepository
	readingRepo          *repositories.NoiseReadingRepository
	evidenceRepo         *repositories.EvidenceAttachmentRepository
	rectificationRepo    *repositories.RectificationMeasureRepository
	retestRepo           *repositories.RetestRecordRepository
	anomalyRepo          *repositories.AnomalyEventRepository
}

func NewStatisticsService() *StatisticsService {
	return &StatisticsService{
		complaintRepo:     repositories.NewComplaintRepository(),
		samplingPointRepo: repositories.NewSamplingPointRepository(),
		readingRepo:       repositories.NewNoiseReadingRepository(),
		evidenceRepo:      repositories.NewEvidenceAttachmentRepository(),
		rectificationRepo: repositories.NewRectificationMeasureRepository(),
		retestRepo:        repositories.NewRetestRecordRepository(),
		anomalyRepo:       repositories.NewAnomalyEventRepository(),
	}
}

func (s *StatisticsService) GetDashboard() (*dto.StatisticsDashboardResponse, error) {
	totalComplaints, _ := s.complaintRepo.Count()
	pendingComplaints, _ := s.complaintRepo.CountByStatus("pending")
	reviewingComplaints, _ := s.complaintRepo.CountByStatus("reviewing")
	supplementComplaints, _ := s.complaintRepo.CountByStatus("supplement")
	pendingTotal := pendingComplaints + reviewingComplaints + supplementComplaints

	confirmedComplaints, _ := s.complaintRepo.CountByStatus("confirmed")
	archivedComplaints, _ := s.complaintRepo.CountByStatus("archived")
	completedTotal := confirmedComplaints + archivedComplaints

	totalSamplingPoints, _ := s.samplingPointRepo.Count()
	totalNoiseReadings, _ := s.readingRepo.Count()
	exceededReadings, _ := s.readingRepo.CountExceeded()
	averageLeq, _ := s.readingRepo.AverageLeq()

	totalEvidence, _ := s.evidenceRepo.Count()
	evidenceCompleteness := 0.0
	if totalComplaints > 0 {
		evidenceCompleteness = float64(totalEvidence) / float64(totalComplaints) * 100
		if evidenceCompleteness > 100 {
			evidenceCompleteness = 100
		}
	}

	totalMeasures, _ := s.rectificationRepo.Count()
	completedMeasures, _ := s.rectificationRepo.CountByStatus("completed")
	verifiedMeasures, _ := s.rectificationRepo.CountByStatus("verified")
	rectificationRate := 0.0
	if totalMeasures > 0 {
		rectificationRate = float64(completedMeasures+verifiedMeasures) / float64(totalMeasures) * 100
	}

	totalRetests, _ := s.retestRepo.Count()
	passedRetests, _ := s.retestRepo.CountPassed()
	retestPassRate := 0.0
	if totalRetests > 0 {
		retestPassRate = float64(passedRetests) / float64(totalRetests) * 100
	}

	openAnomalies, _ := s.anomalyRepo.CountOpen()

	return &dto.StatisticsDashboardResponse{
		TotalComplaints:     int(totalComplaints),
		PendingComplaints:   int(pendingTotal),
		CompletedComplaints: int(completedTotal),
		TotalSamplingPoints: int(totalSamplingPoints),
		TotalNoiseReadings:  int(totalNoiseReadings),
		ExceededReadings:    int(exceededReadings),
		AverageLeq:          averageLeq,
		EvidenceCompleteness: evidenceCompleteness,
		RectificationRate:   rectificationRate,
		RetestPassRate:      retestPassRate,
		OpenAnomalies:       int(openAnomalies),
	}, nil
}

func (s *StatisticsService) GetCompletenessStatistics(filters map[string]interface{}) (*dto.CompletenessStatisticsResponse, error) {
	complaints, _, err := s.complaintRepo.FindAll(1, 1000, filters)
	if err != nil {
		return nil, err
	}

	var items []dto.CompletenessStatistics
	for _, c := range complaints {
		samplingPoints, _ := s.samplingPointRepo.FindByComplaintID(c.ID)
		readings, _ := s.readingRepo.FindByComplaintID(c.ID)
		evidenceCount, _ := s.evidenceRepo.CountByComplaintID(c.ID)

		totalRecords := 3
		completeRecords := 0
		var missingFields []string

		if len(samplingPoints) > 0 {
			completeRecords++
		} else {
			missingFields = append(missingFields, "sampling_point")
		}

		if len(readings) > 0 {
			completeRecords++
		} else {
			missingFields = append(missingFields, "noise_reading")
		}

		if evidenceCount > 0 {
			completeRecords++
		} else {
			missingFields = append(missingFields, "evidence_attachment")
		}

		completeness := float64(completeRecords) / float64(totalRecords) * 100

		items = append(items, dto.CompletenessStatistics{
			Date:              c.CreatedAt.Format("2006-01-02"),
			BatchNo:           c.BatchNo,
			ResponsibleUser:   c.ResponsibleUser,
			TotalRecords:      totalRecords,
			CompleteRecords:   completeRecords,
			IncompleteRecords: totalRecords - completeRecords,
			CompletenessRate:  completeness,
			MissingFields:     missingFields,
		})
	}

	return &dto.CompletenessStatisticsResponse{
		Total: len(items),
		Items: items,
	}, nil
}

func (s *StatisticsService) GetRectificationRateStatistics(filters map[string]interface{}) (*dto.RectificationRateResponse, error) {
	measures, _, err := s.rectificationRepo.FindAll(1, 1000, filters)
	if err != nil {
		return nil, err
	}

	var items []dto.RectificationRateStatistics
	for _, m := range measures {
		completed := 0
		verified := 0
		if m.Status == "completed" {
			completed = 1
		}
		if m.Status == "verified" {
			verified = 1
		}

		rate := 0.0
		if completed+verified > 0 {
			rate = 100.0
		}

		items = append(items, dto.RectificationRateStatistics{
			Date:              m.CreatedAt.Format("2006-01-02"),
			BatchNo:           m.BatchNo,
			ResponsibleUser:   m.ResponsibleUser,
			TotalMeasures:     1,
			CompletedMeasures: completed,
			VerifiedMeasures:  verified,
			RectificationRate: rate,
		})
	}

	return &dto.RectificationRateResponse{
		Total: len(items),
		Items: items,
	}, nil
}

func (s *StatisticsService) GetRetestPassRateStatistics(filters map[string]interface{}) (*dto.RetestPassRateResponse, error) {
	records, _, err := s.retestRepo.FindAll(1, 1000, filters)
	if err != nil {
		return nil, err
	}

	var items []dto.RetestPassRateStatistics
	for _, r := range records {
		passed := 0
		failed := 0
		if r.IsPassed {
			passed = 1
		} else {
			failed = 1
		}

		rate := 0.0
		if r.IsPassed {
			rate = 100.0
		}

		items = append(items, dto.RetestPassRateStatistics{
			Date:            r.CreatedAt.Format("2006-01-02"),
			BatchNo:         r.BatchNo,
			ResponsibleUser: r.ResponsibleUser,
			TotalRetests:    1,
			PassedRetests:   passed,
			FailedRetests:   failed,
			PassRate:        rate,
		})
	}

	return &dto.RetestPassRateResponse{
		Total: len(items),
		Items: items,
	}, nil
}
