package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}

// 初始化验证器并设置中文翻译
func initValidator() (*validator.Validate, ut.Translator) {
	// 初始化翻译器
	zhT := zh.New()
	uni := ut.New(zhT, zhT)
	trans, _ := uni.GetTranslator("zh")

	// 初始化验证器
	validate := validator.New()
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} 不能为空", true) // {0} 是字段名
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field()) // fe.Field() 获取字段名
		return t
	})

	return validate, trans
}

func ValidateUserCreate(c *gin.Context, req *User) error {
	validate := validator.New()
	return validate.Struct(req)
}
