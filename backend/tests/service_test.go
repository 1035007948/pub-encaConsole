package tests

import (
	"testing"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/services"
)

func TestPriorityCalculation(t *testing.T) {
	service := services.NewCalculationService()

	tests := []struct {
		name      string
		request   dto.PriorityCalculateRequest
		minPriority int
	}{
		{
			name: "High priority complaint",
			request: dto.PriorityCalculateRequest{
				ComplaintLevel:         "urgent",
				SamplingTimeCompliant:  false,
				ReadingDeviation:       15.0,
				EvidenceCompleteness:   1.0,
			},
			minPriority: 60,
		},
		{
			name: "Low priority complaint",
			request: dto.PriorityCalculateRequest{
				ComplaintLevel:         "low",
				SamplingTimeCompliant:  true,
				ReadingDeviation:       0,
				EvidenceCompleteness:   0.5,
			},
			minPriority: 10,
		},
		{
			name: "Medium priority complaint",
			request: dto.PriorityCalculateRequest{
				ComplaintLevel:         "medium",
				SamplingTimeCompliant:  true,
				ReadingDeviation:       5.0,
				EvidenceCompleteness:   0.8,
			},
			minPriority: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := service.CalculatePriority(&tt.request)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if response.Priority < tt.minPriority {
				t.Errorf("Expected priority >= %d, got %d", tt.minPriority, response.Priority)
			}

			if response.Level == "" {
				t.Error("Level should not be empty")
			}

			if response.Explanation == "" {
				t.Error("Explanation should not be empty")
			}
		})
	}
}

func TestCompletenessCalculation(t *testing.T) {
	service := services.NewCalculationService()

	tests := []struct {
		name             string
		request          dto.CompletenessCalculateRequest
		expectedComplete bool
	}{
		{
			name: "Complete evidence",
			request: dto.CompletenessCalculateRequest{
				ComplaintID:        1,
				SamplingPointCount: 1,
				ReadingCount:       1,
				EvidenceCount:      1,
			},
			expectedComplete: true,
		},
		{
			name: "Incomplete evidence",
			request: dto.CompletenessCalculateRequest{
				ComplaintID:        2,
				SamplingPointCount: 0,
				ReadingCount:       1,
				EvidenceCount:      0,
			},
			expectedComplete: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := service.CalculateCompleteness(&tt.request)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if response.IsComplete != tt.expectedComplete {
				t.Errorf("Expected IsComplete=%v, got %v", tt.expectedComplete, response.IsComplete)
			}

			if len(response.RequiredFields) == 0 {
				t.Error("RequiredFields should not be empty")
			}
		})
	}
}

func TestComplianceCheck(t *testing.T) {
	service := services.NewCalculationService()

	tests := []struct {
		name              string
		request           dto.ComplianceCheckRequest
		expectCompliant   bool
	}{
		{
			name: "Day time measurement",
			request: dto.ComplianceCheckRequest{
				MeasurementTime: "14:30",
				PeriodType:      "day",
			},
			expectCompliant: true,
		},
		{
			name: "Night time measurement",
			request: dto.ComplianceCheckRequest{
				MeasurementTime: "23:30",
				PeriodType:      "night",
			},
			expectCompliant: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := service.CheckCompliance(&tt.request)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if response.PeriodType == "" {
				t.Error("PeriodType should not be empty")
			}
		})
	}
}
