package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tdalexm/goson-server/internal/domain"
	"github.com/tdalexm/goson-server/internal/services"
)

type Handler struct {
	listSR        services.ListService
	listFilterSR  services.ListFilterService
	getSR         services.GetService
	createSR      services.CreateService
	updateSR      services.UpdateService
	updateFieldSR services.UpdateFieldsService
}

func (h *Handler) List(c *gin.Context) {
	resource := c.Param("resource")

	var result []domain.Record
	var err error

	field := c.Query("field")

	if field == "id" {
		ReturnErrorResponse(c, domain.AppError{
			Code: domain.ErrSearchByID,
			Msg:  "Cannot filter by ID. Please use the following endpoint '/:resource/:id'.",
		})
		return
	}

	if field != "" {
		value := c.Query("value")
		contains := c.Query("contains")
		if value == "" && contains == "" {
			ReturnErrorResponse(c, domain.AppError{
				Code: domain.ErrWrongParams,
				Msg:  "value or contains must be specified when filtering by field.",
			})
			return
		}
		filter := domain.Filter{
			Field:    field,
			Value:    value,
			Contains: contains,
		}
		result, err = h.listFilterSR.Execute(resource, filter)
	} else {
		result, err = h.listSR.Execute(resource)
	}

	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	sort := strings.ToLower(c.Query("sort"))
	if sort == "desc" {
		slices.Reverse(result)
		c.JSON(200, result)
		return
	}

	c.JSON(200, result)
}

func (h *Handler) Get(c *gin.Context) {
	resource := c.Param("resource")
	id := c.Param("id")
	result, err := h.getSR.Execute(resource, id)
	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	c.JSON(200, result)
}

func (h *Handler) Create(c *gin.Context) {
	resource := c.Param("resource")
	var record domain.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		if err.Error() == "EOF" {
			ReturnErrorResponse(c, domain.NewAppError(
				domain.ErrValidation,
				"Request body cannot be empty",
			))
			return
		}
		ReturnErrorResponse(c, domain.NewAppError(
			domain.ErrValidation,
			fmt.Sprintf("Invalid JSON format: %v", err),
		))
		return
	}

	if len(record) == 0 {
		ReturnErrorResponse(c, domain.NewAppError(
			domain.ErrValidation,
			"Request body cannot be empty",
		))
		return
	}

	id, err := h.createSR.Execute(resource, record)
	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  fmt.Sprintf("Record with ID '%s' added to %s resource", id, resource),
		"location": fmt.Sprintf("/%s/%s", resource, id),
	})
}

func (h *Handler) Update(c *gin.Context) {
	resource := c.Param("resource")
	id := c.Param("id")

	var record domain.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		if err.Error() == "EOF" {
			ReturnErrorResponse(c, domain.NewAppError(
				domain.ErrValidation,
				"Request body cannot be empty",
			))
			return
		}
		ReturnErrorResponse(c, domain.NewAppError(
			domain.ErrValidation,
			fmt.Sprintf("Invalid JSON format: %v", err),
		))
		return
	}

	var updatedID string
	var err error

	log.Println(c.Request.Method)
	if c.Request.Method == "PATCH" {
		updatedID, err = h.updateFieldSR.Execute(resource, id, record)
	} else {
		updatedID, err = h.updateSR.Execute(resource, id, record)
	}

	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Updated record with ID '%s'", updatedID),
	})
}
