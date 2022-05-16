package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignUpPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" validate:"phone"`
}

type SignUpEmailExistRequest struct {
	Email string `json:"email,omitempty" validate:"email"`
}

func ValidateSignUpPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:10"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手機號碼必填",
			"digits:手機號碼必須為 10 碼",
		},
	}

	opts := govalidator.Options{
		TagIdentifier: "validate",
		Data:          data,
		Rules:         rules,
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}

func ValidateSignUpEmailExist(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{"required", "email"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:電子郵件必填",
			"email:電子郵件格式錯誤",
		},
	}

	opts := govalidator.Options{
		TagIdentifier: "validate",
		Data:          data,
		Rules:         rules,
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
