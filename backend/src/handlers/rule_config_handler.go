package handlers

import (
	"net/http"
	"strconv"

	"noise-complaint-backend/src/repositories"

	"github.com/gin-gonic/gin"
)

type RuleConfigHandler struct {
	repo *repositories.RuleConfigRepository
}

func NewRuleConfigHandler() *RuleConfigHandler {
	return &RuleConfigHandler{repo: repositories.NewRuleConfigRepository()}
}

func (h *RuleConfigHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	filters := make(map[string]interface{})
	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}
	if ruleType := c.Query("rule_type"); ruleType != "" {
		filters["rule_type"] = ruleType
	}

	rules, total, err := h.repo.FindAll(page, pageSize, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"items": rules,
	})
}

func (h *RuleConfigHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	rule, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rule)
}

func (h *RuleConfigHandler) Create(c *gin.Context) {
	var rule struct {
		RuleNo      string  `json:"rule_no" binding:"required"`
		RuleName    string  `json:"rule_name" binding:"required"`
		RuleType    string  `json:"rule_type" binding:"required"`
		Description string  `json:"description"`
		Conditions  string  `json:"conditions"`
		Threshold   float64 `json:"threshold"`
		Action      string  `json:"action"`
		Priority    int     `json:"priority"`
		Remark      string  `json:"remark"`
	}

	if err := c.ShouldBindJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ruleConfig := &struct {
		RuleNo      string
		RuleName    string
		RuleType    string
		Description string
		Conditions  string
		Threshold   float64
		Action      string
		Priority    int
		Status      string
		Remark      string
	}{
		RuleNo:      rule.RuleNo,
		RuleName:    rule.RuleName,
		RuleType:    rule.RuleType,
		Description: rule.Description,
		Conditions:  rule.Conditions,
		Threshold:   rule.Threshold,
		Action:      rule.Action,
		Priority:    rule.Priority,
		Status:      "active",
		Remark:      rule.Remark,
	}

	c.JSON(http.StatusCreated, ruleConfig)
}

func (h *RuleConfigHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	rule, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var req struct {
		RuleName    string  `json:"rule_name"`
		Description string  `json:"description"`
		Conditions  string  `json:"conditions"`
		Threshold   float64 `json:"threshold"`
		Action      string  `json:"action"`
		Priority    int     `json:"priority"`
		Status      string  `json:"status"`
		Remark      string  `json:"remark"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.RuleName != "" {
		rule.RuleName = req.RuleName
	}
	if req.Description != "" {
		rule.Description = req.Description
	}
	if req.Conditions != "" {
		rule.Conditions = req.Conditions
	}
	rule.Threshold = req.Threshold
	if req.Action != "" {
		rule.Action = req.Action
	}
	rule.Priority = req.Priority
	if req.Status != "" {
		rule.Status = "active"
		if req.Status == "inactive" {
			rule.Status = "inactive"
		}
	}
	if req.Remark != "" {
		rule.Remark = req.Remark
	}

	err = h.repo.Update(rule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rule)
}
