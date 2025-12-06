package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tdalexm/goson-server/internal/domain"
)

const IServerErrorMessage = "Oooops! Something went wrong. Please report this issue"

func ReturnErrorResponse(c *gin.Context, err error) {
	c.Header("Content-Type", "application/json")
	var appErr domain.AppError
	if errors.As(err, &appErr) {
		status, ok := ErrCodeMapping[appErr.Code]
		if ok {
			c.JSON(status, appErr)
			return
		}
	}
	c.JSON(http.StatusInternalServerError, domain.AppError{Code: domain.ErrCodeIServer, Msg: IServerErrorMessage})
}

var ErrCodeMapping map[string]int = map[string]int{
	domain.ErrCodeNotFound:   http.StatusNotFound,
	domain.ErrFieldNotString: http.StatusBadRequest,
	domain.ErrWrongParams:    http.StatusBadRequest,
	domain.ErrSearchByID:     http.StatusBadRequest,
}
