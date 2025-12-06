package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tdalexm/goson-server/internal/domain"
	"github.com/tdalexm/goson-server/internal/services"
)

type Handler struct {
	listSR       services.ListService
	listFilterSR services.ListFilterService
	getSR        services.GetService
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
