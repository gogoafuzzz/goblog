package auth

import (
	"fmt"
	v1 "goblog/app/http/controllers/api/v1"
	"goblog/app/models/user"
	"goblog/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseApiController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignUpPhoneExistRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}

	errs := requests.ValidateSignUpPhoneExist(&request, c)

	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": errs,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})

}
