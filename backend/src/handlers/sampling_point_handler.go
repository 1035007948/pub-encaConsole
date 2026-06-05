package handlers

import (
	"net/http"
	"strconv"

	"noise-complaint-backend/src/dto"
	"noise-complaint-backend/src/services"

	"github.com/gin-gonic/gin"
)

type SamplingPointHandler struct {
	service *services.SamplingPointService
}

func NewSamplingPointHandler() *SamplingPointHandler {
	return &SamplingPointHandler{service: services.NewSamplingPointService()}
}

func (h *SamplingPointHandler) Create(c *gin.Context) {
	var req dto.SamplingPointCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	operator := c.GetHeader("X-Operator")
	if operator == "" {
		operator = "system"
	}

	point, err := h.service.Create(&req, operator)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, point)
}

func (h *SamplingPointHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	point, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, point)
}

func (h *SamplingPointHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if complaintID := c.Query("complaint_id"); complaintID != "" {
		filters["complaint_id"] = complaintID
	}
	if responsibleUser := c.Query("responsible_user"); responsibleUser != "" {
		filters["responsible_user"] = responsibleUser
	}

	response, err := h.service.GetList(page, pageSize, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *SamplingPointHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.SamplingPointUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	operator := c.GetHeader("X-Operator")
	if operator == "" {
		operator = "system"
	}

	point, err := h.service.Update(uint(id), &req, operator)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, point)
}

func (h *SamplingPointHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	operator := c.GetHeader("X-Operator")
	if operator == "" {
		operator = "system"
	}

	err = h.service.Delete(uint(id), operator)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}

func (h *SamplingPointHandler) Transition(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.SamplingPointTransitionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	operator := c.GetHeader("X-Operator")
	if operator == "" {
		operator = "system"
	}

	point, err := h.service.Transition(uint(id), &req, operator)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, point)
}

func (h *SamplingPointHandler) ValidateConsistency(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	isValid, missingFields, err := h.service.ValidateConsistency(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"is_valid":       isValid,
		"missing_fields": missingFields,
	})
}
