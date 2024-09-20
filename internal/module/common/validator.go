package common

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	err = validate.Struct(data)
	if err != nil {
		switch ve := err.(type) {
		case validator.ValidationErrors:
			for _, v := range ve {
				return v.Translate(trans), 500
			}
		case *validator.InvalidValidationError:
			return fmt.Sprintf("无效的验证错误: %v", err), 500
		default:
			return fmt.Sprintf("未知错误: %v", err), 500
		}
	}
	return "", 200
}
