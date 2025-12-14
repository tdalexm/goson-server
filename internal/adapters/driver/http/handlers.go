package driverhttp

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tdalexm/goson-server/internal/adapters/driver/http/serializer"
	"github.com/tdalexm/goson-server/internal/domain"
	portsdriver "github.com/tdalexm/goson-server/internal/ports/driver"
)

type Handler struct {
	listSR         portsdriver.ListService
	listFilterSR   portsdriver.ListFilterService
	getSR          portsdriver.GetService
	createSR       portsdriver.CreateService
	updateSR       portsdriver.UpdateService
	updateFieldsSR portsdriver.UpdateFieldsService
	deleteSR       portsdriver.DeleteService
	serializer     *serializer.JSONAPISerializer
}

func NewHandler(
	list portsdriver.ListService,
	listFilter portsdriver.ListFilterService,
	get portsdriver.GetService,
	create portsdriver.CreateService,
	update portsdriver.UpdateService,
	updateFields portsdriver.UpdateFieldsService,
	del portsdriver.DeleteService,
	baseURL string,
) *Handler {
	return &Handler{
		listSR:         list,
		listFilterSR:   listFilter,
		getSR:          get,
		createSR:       create,
		updateSR:       update,
		updateFieldsSR: updateFields,
		deleteSR:       del,
		serializer:     serializer.NewJSONAPISerializer(baseURL),
	}
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
		result, err = h.listFilterSR.Execute(collection, filter)
	} else {
		result, err = h.listSR.Execute(collection)
	}

	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	total := len(result)
	if total == 0 {
		c.JSON(http.StatusNoContent, result)
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	queryParams := c.Request.URL.Query()
	sort := strings.ToLower(c.Query("sort"))
	if sort == "desc" {
		slices.Reverse(result)
	}

	start := (page - 1) * limit
	end := start + limit
	paginatedData := result[start:end]
	responseData := h.serializer.SerializeCollection(collection, paginatedData, total, page, limit, queryParams)

	c.PureJSON(200, responseData)
}

func (h *Handler) Get(c *gin.Context) {
	collection := c.Param("collection")
	id := c.Param("id")
	result, err := h.getSR.Execute(collection, id)
	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	responseData := h.serializer.SerializeResource(collection, result)
	c.PureJSON(http.StatusOK, responseData)
}

func (h *Handler) Create(c *gin.Context) {
	collectionType := c.Param("collection")
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

	result, err := h.createSR.Execute(collectionType, record)
	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	responseData := h.serializer.SerializeResource(collectionType, result)

	c.JSON(http.StatusCreated, responseData)
}

func (h *Handler) Update(c *gin.Context) {
	collectionType := c.Param("collection")
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

	var result domain.Record
	var err error

	if c.Request.Method == "PATCH" {
		result, err = h.updateFieldsSR.Execute(collectionType, id, record)
	} else {
		result, err = h.updateSR.Execute(collectionType, id, record)
	}

	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	responseData := h.serializer.SerializeResource(collectionType, result)

	c.JSON(http.StatusOK, responseData)
}

func (h *Handler) Delete(c *gin.Context) {
	collectionType := c.Param("collection")
	id := c.Param("id")

	deletedID, err := h.deleteSR.Execute(collectionType, id)
	if err != nil {
		ReturnErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Deleted record with ID '%s'", deletedID),
	})
}
