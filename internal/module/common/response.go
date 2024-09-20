package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseSuccess(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": msg,
	})
	return
}

func ResponseBad(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  400,
		"message": msg,
	})
	return
}
