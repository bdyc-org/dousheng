package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendErrResponse(c *gin.Context, statusCode int64, err error) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  err.Error(),
	})
}
