package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/handlers"
	"noise-complaint-backend/src/routes"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	routes.SetupRoutes(r)
	return r
}

func TestHealthCheck(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/health", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestComplaintCreate(t *testing.T) {
	router := setupTestRouter()

	reqBody := dto.ComplaintCreateRequest{
		ComplaintNo:     "CMP-TEST-0001",
		Title:           "测试投诉单",
		Description:     "这是一个测试投诉单",
		Level:           "medium",
		ComplainantName: "测试人员",
		ComplainantTel:  "138****0000",
		EnterpriseName:  "测试企业",
		EnterpriseAddr:  "测试地址",
		ResponsibleUser: "测试处理员",
		BatchNo:         "BATCH-TEST-001",
	}

	body, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/complaints", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestComplaintList(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/complaints?page=1&page_size=10", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if _, ok := response["total"]; !ok {
		t.Error("Response should contain 'total' field")
	}

	if _, ok := response["items"]; !ok {
		t.Error("Response should contain 'items' field")
	}
}

func TestTimePeriodBatchImport(t *testing.T) {
	router := setupTestRouter()

	reqBody := dto.TimePeriodBatchImportRequest{
		Items: []dto.TimePeriodCreateRequest{
			{
				PeriodNo:    "TP-TEST-0001",
				PeriodName:  "测试时段1",
				PeriodType:  "day",
				TimeFrom:    "08:00",
				TimeTo:      "18:00",
				DayLimit:    65.0,
				NightLimit:  55.0,
			},
			{
				PeriodNo:    "TP-TEST-0002",
				PeriodName:  "测试时段2",
				PeriodType:  "night",
				TimeFrom:    "18:00",
				TimeTo:      "08:00",
				DayLimit:    55.0,
				NightLimit:  45.0,
			},
		},
	}

	body, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/time-periods/batch-import", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestCalculatePriority(t *testing.T) {
	router := setupTestRouter()

	reqBody := dto.PriorityCalculateRequest{
		ComplaintLevel:         "high",
		SamplingTimeCompliant:  true,
		ReadingDeviation:       5.0,
		EvidenceCompleteness:   0.8,
	}

	body, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/calculate/priority", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response dto.PriorityCalculateResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Priority <= 0 {
		t.Error("Priority should be greater than 0")
	}

	if response.Level == "" {
		t.Error("Level should not be empty")
	}
}

func TestCalculateCompleteness(t *testing.T) {
	router := setupTestRouter()

	reqBody := dto.CompletenessCalculateRequest{
		ComplaintID:        1,
		SamplingPointCount: 2,
		ReadingCount:       3,
		EvidenceCount:      1,
	}

	body, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/calculate/completeness", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response dto.CompletenessCalculateResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Completeness < 0 || response.Completeness > 1 {
		t.Error("Completeness should be between 0 and 1")
	}
}

func TestCheckCompliance(t *testing.T) {
	router := setupTestRouter()

	reqBody := dto.ComplianceCheckRequest{
		MeasurementTime: "14:30",
		PeriodType:      "day",
	}

	body, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/calculate/compliance", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestStatisticsDashboard(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/statistics/dashboard", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response dto.StatisticsDashboardResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.TotalComplaints < 0 {
		t.Error("TotalComplaints should not be negative")
	}
}
