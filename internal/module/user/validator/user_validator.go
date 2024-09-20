package validator

import (
	"github.com/blackaay/gin-start/internal/module/common"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `validate:"required" label:"姓名"`
	Password string `validate:"required" label:"密码"`
}

func ValidateUserCreate(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		common.ResponseBad(c, "Failed to parse request body")
		return false
	}
	msg, code := common.Validate(data)
	if code != 200 {
		common.ResponseBad(c, msg)
		return false
	}
	return true
}
