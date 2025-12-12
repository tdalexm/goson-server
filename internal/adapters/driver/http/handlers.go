package driverhttp

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tdalexm/goson-server/internal/domain"
	"github.com/tdalexm/goson-server/internal/services"
)

type Handler struct {
	ListSR        services.ListService
	ListFilterSR  services.ListFilterService
	GetSR         services.GetService
	CreateSR      services.CreateService
	UpdateSR      services.UpdateService
	UpdateFieldSR services.UpdateFieldsService
	DeleteSR      services.DeleteService
}

func (h *Handler) List(c *gin.Context) {
	collection := c.Param("collection")

	var result []domain.Record
	var err error

	field := c.Query("field")

	if field == "id" {
		ReturnErrorResponse(c, domain.AppError{
			Code: domain.ErrSearchByID,
			Msg:  "Cannot filter by ID. Please use the following endpoint '/:collection/:id'.",
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
		result, err = h.ListFilterSR.Execute(collection, filter)
	} else {
		result, err = h.ListSR.Execute(collection)
	}

	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusNoContent, result)
	}

	sort := strings.ToLower(c.Query("sort"))
	if sort == "desc" {
		slices.Reverse(result)
		c.JSON(http.StatusOK, result)
		return
	}

	c.JSON(200, result)
}

func (h *Handler) Get(c *gin.Context) {
	collection := c.Param("collection")
	id := c.Param("id")
	result, err := h.GetSR.Execute(collection, id)
	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) Create(c *gin.Context) {
	collection := c.Param("collection")
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

	id, err := h.CreateSR.Execute(collection, record)
	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  fmt.Sprintf("Record with ID '%s' added to %s collection", id, collection),
		"location": fmt.Sprintf("/%s/%s", collection, id),
	})
}

func (h *Handler) Update(c *gin.Context) {
	collection := c.Param("collection")
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

	if c.Request.Method == "PATCH" {
		updatedID, err = h.UpdateFieldSR.Execute(collection, id, record)
	} else {
		updatedID, err = h.UpdateSR.Execute(collection, id, record)
	}

	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Updated record with ID '%s'", updatedID),
	})
}

func (h *Handler) Delete(c *gin.Context) {
	collection := c.Param("collection")
	id := c.Param("id")

	deletedID, err := h.DeleteSR.Execute(collection, id)
	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Deleted record with ID '%s'", deletedID),
	})
}
