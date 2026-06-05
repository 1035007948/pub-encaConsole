package services

import (
	"fmt"
	"time"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/models"
	"noise-complaint-backend/src/repositories"
)

type CalculationService struct {
	complaintRepo     *repositories.ComplaintRepository
	samplingPointRepo *repositories.SamplingPointRepository
	readingRepo       *repositories.NoiseReadingRepository
	evidenceRepo      *repositories.EvidenceAttachmentRepository
	timePeriodRepo    *repositories.TimePeriodRepository
}

func NewCalculationService() *CalculationService {
	return &CalculationService{
		complaintRepo:     repositories.NewComplaintRepository(),
		samplingPointRepo: repositories.NewSamplingPointRepository(),
		readingRepo:       repositories.NewNoiseReadingRepository(),
		evidenceRepo:      repositories.NewEvidenceAttachmentRepository(),
		timePeriodRepo:    repositories.NewTimePeriodRepository(),
	}
}

func (s *CalculationService) CalculatePriority(req *dto.PriorityCalculateRequest) (*dto.PriorityCalculateResponse, error) {
	priority := 0
	level := "normal"
	var explanations []string

	levelPriority := map[string]int{
		"urgent": 40,
		"high":   30,
		"medium": 20,
		"low":    10,
	}

	if p, ok := levelPriority[req.ComplaintLevel]; ok {
		priority += p
		explanations = append(explanations, fmt.Sprintf("投诉等级 %s 贡献 %d 分", req.ComplaintLevel, p))
	}

	if !req.SamplingTimeCompliant {
		priority += 20
		explanations = append(explanations, "采样时段不合规贡献 20 分")
	}

	if req.ReadingDeviation > 0 {
		deviationScore := int(req.ReadingDeviation * 2)
		if deviationScore > 20 {
			deviationScore = 20
		}
		priority += deviationScore
		explanations = append(explanations, fmt.Sprintf("读数偏差 %.2f dB 贡献 %d 分", req.ReadingDeviation, deviationScore))
	}

	completenessScore := int(req.EvidenceCompleteness * 10)
	priority += completenessScore
	explanations = append(explanations, fmt.Sprintf("证据完整度 %.2f%% 贡献 %d 分", req.EvidenceCompleteness*100, completenessScore))

	if priority >= 70 {
		level = "urgent"
	} else if priority >= 50 {
		level = "high"
	} else if priority >= 30 {
		level = "medium"
	}

	explanation := fmt.Sprintf("总优先级 %d 分，等级: %s。%s", priority, level, explanations[0])
	for i := 1; i < len(explanations); i++ {
		explanation += "；" + explanations[i]
	}

	return &dto.PriorityCalculateResponse{
		Priority:    priority,
		Level:       level,
		Explanation: explanation,
	}, nil
}

func (s *CalculationService) CalculateCompleteness(req *dto.CompletenessCalculateRequest) (*dto.CompletenessCalculateResponse, error) {
	requiredFields := []string{
		"sampling_point",
		"noise_reading",
		"evidence_attachment",
	}

	var missingFields []string

	if req.SamplingPointCount == 0 {
		missingFields = append(missingFields, "sampling_point")
	}
	if req.ReadingCount == 0 {
		missingFields = append(missingFields, "noise_reading")
	}
	if req.EvidenceCount == 0 {
		missingFields = append(missingFields, "evidence_attachment")
	}

	completeness := float64(len(requiredFields)-len(missingFields)) / float64(len(requiredFields))
	isComplete := len(missingFields) == 0

	explanation := fmt.Sprintf("证据完整度 %.2f%%", completeness*100)
	if len(missingFields) > 0 {
		explanation += fmt.Sprintf("，缺少: %v", missingFields)
	} else {
		explanation += "，所有必要字段已填写"
	}

	return &dto.CompletenessCalculateResponse{
		Completeness:   completeness,
		RequiredFields: requiredFields,
		MissingFields:  missingFields,
		IsComplete:     isComplete,
		Explanation:    explanation,
	}, nil
}

func (s *CalculationService) CheckCompliance(req *dto.ComplianceCheckRequest) (*dto.ComplianceCheckResponse, error) {
	var period *models.TimePeriod
	var err error

	if req.TimePeriodID > 0 {
		period, err = s.timePeriodRepo.FindByID(req.TimePeriodID)
		if err != nil {
			return nil, err
		}
	} else if req.PeriodType != "" {
		period, err = s.timePeriodRepo.FindByType(models.TimePeriodType(req.PeriodType))
		if err != nil {
			return nil, err
		}
	} else {
		periods, err := s.timePeriodRepo.FindActive()
		if err != nil {
			return nil, err
		}
		if len(periods) == 0 {
			return &dto.ComplianceCheckResponse{
				IsCompliant:  false,
				ViolationMsg: "未找到有效的时段分类配置",
			}, nil
		}
		period = &periods[0]
	}

	isCompliant := s.isTimeInRange(req.MeasurementTime, period.TimeFrom, period.TimeTo)

	violationMsg := ""
	if !isCompliant {
		violationMsg = fmt.Sprintf("测量时间 %s 不在时段范围 %s-%s 内", req.MeasurementTime, period.TimeFrom, period.TimeTo)
	}

	return &dto.ComplianceCheckResponse{
		IsCompliant:  isCompliant,
		PeriodType:   string(period.PeriodType),
		TimeFrom:     period.TimeFrom,
		TimeTo:       period.TimeTo,
		ViolationMsg: violationMsg,
	}, nil
}

func (s *CalculationService) isTimeInRange(timeStr, rangeFrom, rangeTo string) bool {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return false
	}
	from, err := time.Parse("15:04", rangeFrom)
	if err != nil {
		return false
	}
	to, err := time.Parse("15:04", rangeTo)
	if err != nil {
		return false
	}

	if to.Before(from) {
		return t.After(from) || t.Before(to) || t.Equal(from) || t.Equal(to)
	}

	return (t.After(from) || t.Equal(from)) && (t.Before(to) || t.Equal(to))
}
