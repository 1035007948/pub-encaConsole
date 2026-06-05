package handlers

import (
	"net/http"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/services"

	"github.com/gin-gonic/gin"
)

type StatisticsHandler struct {
	service *services.StatisticsService
}

func NewStatisticsHandler() *StatisticsHandler {
	return &StatisticsHandler{service: services.NewStatisticsService()}
}

func (h *StatisticsHandler) GetDashboard(c *gin.Context) {
	response, err := h.service.GetDashboard()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *StatisticsHandler) GetCompletenessStatistics(c *gin.Context) {
	filters := make(map[string]interface{})
	if batchNo := c.Query("batch_no"); batchNo != "" {
		filters["batch_no"] = batchNo
	}
	if responsibleUser := c.Query("responsible_user"); responsibleUser != "" {
		filters["responsible_user"] = responsibleUser
	}

	response, err := h.service.GetCompletenessStatistics(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *StatisticsHandler) GetRectificationRateStatistics(c *gin.Context) {
	filters := make(map[string]interface{})
	if batchNo := c.Query("batch_no"); batchNo != "" {
		filters["batch_no"] = batchNo
	}

	response, err := h.service.GetRectificationRateStatistics(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *StatisticsHandler) GetRetestPassRateStatistics(c *gin.Context) {
	filters := make(map[string]interface{})
	if batchNo := c.Query("batch_no"); batchNo != "" {
		filters["batch_no"] = batchNo
	}

	response, err := h.service.GetRetestPassRateStatistics(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
