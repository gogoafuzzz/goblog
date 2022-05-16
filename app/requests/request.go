package requests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidateFunc func(data interface{}, c *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidateFunc) bool {
	// ShouldBind 解析請求參數，支援 JSON & Form 表單 & Query
	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "請求解析錯誤，請確認請求格式是否正確。上傳文件請使用 multipart/form-data 格式, 參數請使用 JSON 格式。",
			"error":   err.Error(),
		})
		fmt.Println(err.Error())
		return false
	}

	errs := handler(obj, c)

	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": errs,
		})
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	opts := govalidator.Options{
		TagIdentifier: "validate",
		Data:          data,
		Rules:         rules,
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
