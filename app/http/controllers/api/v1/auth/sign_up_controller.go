package auth

import (
	v1 "goblog/app/http/controllers/api/v1"
	"goblog/app/models/user"
	"goblog/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	v1.BaseApiController
}

func (sc *SignUpController) IsPhoneExist(c *gin.Context) {

	request := requests.SignUpPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignUpPhoneExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})

}

func (sc *SignUpController) IsEmailExist(c *gin.Context) {
	request := requests.SignUpEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignUpEmailExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
