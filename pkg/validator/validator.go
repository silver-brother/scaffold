package validator

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// 自定义验证器
func CustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", mobile)
	}
}

// 自定义验证方法

// 手机号验证
var mobile validator.Func = func(fieldLevel validator.FieldLevel) bool {
	mobile, ok := fieldLevel.Field().Interface().(string)
	if ok {
		ok, _ := regexp.MatchString(`^1[3456789]\d{9}$`, mobile)
		return ok
	}
	return true
}
