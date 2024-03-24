package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type failedResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type successfulResponse struct {
	Data interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, successfulResponse{
		Data: data,
	})
}

func FailureResponse(c *gin.Context, status int, err error) {
	c.JSON(status, failedResponse{
		Status:  status,
		Code:    http.StatusText(status),
		Message: err.Error(),
	})
}
