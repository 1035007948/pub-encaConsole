package dto

import "time"

type StatisticsDashboardResponse struct {
	TotalComplaints        int     `json:"total_complaints"`
	PendingComplaints      int     `json:"pending_complaints"`
	CompletedComplaints    int     `json:"completed_complaints"`
	TotalSamplingPoints    int     `json:"total_sampling_points"`
	TotalNoiseReadings     int     `json:"total_noise_readings"`
	ExceededReadings       int     `json:"exceeded_readings"`
	AverageLeq             float64 `json:"average_leq"`
	EvidenceCompleteness   float64 `json:"evidence_completeness"`
	RectificationRate      float64 `json:"rectification_rate"`
	RetestPassRate         float64 `json:"retest_pass_rate"`
	OpenAnomalies          int     `json:"open_anomalies"`
}

type CompletenessStatistics struct {
	Date                 string  `json:"date"`
	BatchNo              string  `json:"batch_no"`
	ResponsibleUser      string  `json:"responsible_user"`
	TotalRecords         int     `json:"total_records"`
	CompleteRecords      int     `json:"complete_records"`
	IncompleteRecords    int     `json:"incomplete_records"`
	CompletenessRate     float64 `json:"completeness_rate"`
	MissingFields        []string `json:"missing_fields"`
}

type CompletenessStatisticsResponse struct {
	Total int                         `json:"total"`
	Items []CompletenessStatistics    `json:"items"`
}

type RectificationRateStatistics struct {
	Date             string  `json:"date"`
	BatchNo          string  `json:"batch_no"`
	ResponsibleUser  string  `json:"responsible_user"`
	TotalMeasures    int     `json:"total_measures"`
	CompletedMeasures int    `json:"completed_measures"`
	VerifiedMeasures  int    `json:"verified_measures"`
	RectificationRate float64 `json:"rectification_rate"`
}

type RectificationRateResponse struct {
	Total int                           `json:"total"`
	Items []RectificationRateStatistics `json:"items"`
}

type RetestPassRateStatistics struct {
	Date           string  `json:"date"`
	BatchNo        string  `json:"batch_no"`
	ResponsibleUser string `json:"responsible_user"`
	TotalRetests   int     `json:"total_retests"`
	PassedRetests  int     `json:"passed_retests"`
	FailedRetests  int     `json:"failed_retests"`
	PassRate       float64 `json:"pass_rate"`
}

type RetestPassRateResponse struct {
	Total int                        `json:"total"`
	Items []RetestPassRateStatistics `json:"items"`
}
