package dto

type PriorityCalculateRequest struct {
	ComplaintID         uint    `json:"complaint_id"`
	ComplaintLevel      string  `json:"complaint_level"`
	SamplingTimeCompliant bool   `json:"sampling_time_compliant"`
	ReadingDeviation    float64 `json:"reading_deviation"`
	EvidenceCompleteness float64 `json:"evidence_completeness"`
}

type PriorityCalculateResponse struct {
	Priority    int    `json:"priority"`
	Level       string `json:"level"`
	Explanation string `json:"explanation"`
}

type CompletenessCalculateRequest struct {
	ComplaintID        uint `json:"complaint_id"`
	SamplingPointCount int  `json:"sampling_point_count"`
	ReadingCount       int  `json:"reading_count"`
	EvidenceCount      int  `json:"evidence_count"`
}

type CompletenessCalculateResponse struct {
	Completeness       float64  `json:"completeness"`
	RequiredFields     []string `json:"required_fields"`
	MissingFields      []string `json:"missing_fields"`
	IsComplete         bool     `json:"is_complete"`
	Explanation        string   `json:"explanation"`
}

type ComplianceCheckRequest struct {
	MeasurementTime string `json:"measurement_time" binding:"required"`
	TimePeriodID    uint   `json:"time_period_id"`
	PeriodType      string `json:"period_type"`
}

type ComplianceCheckResponse struct {
	IsCompliant  bool   `json:"is_compliant"`
	PeriodType   string `json:"period_type"`
	TimeFrom     string `json:"time_from"`
	TimeTo       string `json:"time_to"`
	ViolationMsg string `json:"violation_msg"`
}
