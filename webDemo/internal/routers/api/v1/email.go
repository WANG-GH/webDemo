package v1

import (
	"fmt"
	//"strconv"

	// "strings"
	"webDemo/internal/service"
	"webDemo/pkg/app"

	"github.com/gin-gonic/gin"
)

type Email struct{}

func NewEmail() Email {
	return Email{}
}

func (email *Email) Create(c *gin.Context) {
	param := service.CreateEmailRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	fmt.Print(param.Email_num)
	err := svc.CreateEmail(&param)
	if err != nil {
		fmt.Printf("svc.CreateEmail err: %v", err)
		response.ToResponse(gin.H{"err": err})
		return
	}
	response.ToResponse(gin.H{"err": "ok"})
	return
}

func (email *Email) ResetData(c *gin.Context) {
	param := service.ResetDataByEmailRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	state, err := svc.ResetDataByEmail(&param)
	if err != nil {
		response.ToResponse(gin.H{
			"err": "can not modify the password",
		})
		return
	}
	if state == 1 {
		response.ToResponse(gin.H{
			"err": "ok",
		})
	} else if state == 2 {
		response.ToResponse(gin.H{
			"err": "email not found",
		})
	} else {
		response.ToResponse(gin.H{
			"err": "wrong verify code",
		})
	}
	return
}
