package handlers

import (
	"net/http"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/services"

	"github.com/gin-gonic/gin"
)

type CalculationHandler struct {
	service *services.CalculationService
}

func NewCalculationHandler() *CalculationHandler {
	return &CalculationHandler{service: services.NewCalculationService()}
}

func (h *CalculationHandler) CalculatePriority(c *gin.Context) {
	var req dto.PriorityCalculateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.CalculatePriority(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *CalculationHandler) CalculateCompleteness(c *gin.Context) {
	var req dto.CompletenessCalculateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.CalculateCompleteness(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *CalculationHandler) CheckCompliance(c *gin.Context) {
	var req dto.ComplianceCheckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.CheckCompliance(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
