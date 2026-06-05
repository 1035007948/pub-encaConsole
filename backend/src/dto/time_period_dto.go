package dto

import "time"

type TimePeriodCreateRequest struct {
	PeriodNo    string  `json:"period_no" binding:"required"`
	PeriodName  string  `json:"period_name" binding:"required"`
	PeriodType  string  `json:"period_type" binding:"required"`
	TimeFrom    string  `json:"time_from" binding:"required"`
	TimeTo      string  `json:"time_to" binding:"required"`
	DayLimit    float64 `json:"day_limit"`
	NightLimit  float64 `json:"night_limit"`
	Description string  `json:"description"`
	BatchNo     string  `json:"batch_no"`
	Remark      string  `json:"remark"`
}

type TimePeriodUpdateRequest struct {
	PeriodName  string  `json:"period_name"`
	TimeFrom    string  `json:"time_from"`
	TimeTo      string  `json:"time_to"`
	DayLimit    float64 `json:"day_limit"`
	NightLimit  float64 `json:"night_limit"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	BatchNo     string  `json:"batch_no"`
	Remark      string  `json:"remark"`
}

type TimePeriodResponse struct {
	ID          uint      `json:"id"`
	PeriodNo    string    `json:"period_no"`
	PeriodName  string    `json:"period_name"`
	PeriodType  string    `json:"period_type"`
	TimeFrom    string    `json:"time_from"`
	TimeTo      string    `json:"time_to"`
	DayLimit    float64   `json:"day_limit"`
	NightLimit  float64   `json:"night_limit"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	BatchNo     string    `json:"batch_no"`
	Remark      string    `json:"remark"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TimePeriodListResponse struct {
	Total int                   `json:"total"`
	Items []TimePeriodResponse `json:"items"`
}

type TimePeriodBatchImportRequest struct {
	Items []TimePeriodCreateRequest `json:"items" binding:"required"`
}

type TimePeriodBatchImportResponse struct {
	Total    int      `json:"total"`
	Success  int      `json:"success"`
	Failed   int      `json:"failed"`
	Errors   []string `json:"errors"`
	Warnings []string `json:"warnings"`
}
