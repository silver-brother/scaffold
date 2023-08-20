package validator

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ValidationErrorResponse 封装验证错误的通用响应结构
type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// BeautifyValidationErrors 将 validator 的验证错误美化为自定义的错误响应
func BeautifyValidationErrors(c *gin.Context, err error) string {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		// 处理验证器错误
		return "Invalid validation request"
	}

	var validationErrors validator.ValidationErrors
	var errorMsg string
	if errors.As(err, &validationErrors) {
		for _, e := range validationErrors {
			errorMsg = fmt.Sprintf("Invalid value for %s.", e.Tag())
			break
		}
		return errorMsg
	}
	return ""
}
