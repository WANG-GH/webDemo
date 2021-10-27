package v1

import (
	"fmt"
	"net/http"
	"webDemo/global"
	"webDemo/internal/service"
	"webDemo/pkg/app"
	"webDemo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (user *User) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	//c.JSON(200, gin.H{"message": "pong"})
	//fmt.Println("in Get")
	return
}

func (user *User) Delete(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)

	return
}

func (user *User) Create(c *gin.Context) {
	param := service.CreateUserRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	fmt.Println(param)
	err := svc.CreateUser(&param)
	if err != nil {
		fmt.Printf("svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"err":"ok",
	})
	return
}

func (user *User) Update(c *gin.Context) {
	param := service.UpdateRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	err := svc.UpdateUser(&param)
	if err != nil {
		fmt.Printf("svc.UpdateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateFail)
		return
	}
	response.ToResponse(gin.H{"update": "ok"})
	return
}

func (user *User) Login(c *gin.Context) {
	param := service.LoginRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	passwd, err := svc.LoginUser(&param)
	fmt.Printf("接受的密码:%v \n", param.Password)
	if err != nil {
		// 用户未存在
		fmt.Printf("svc.LoginFail err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserNotExist)
		return
	}
	if passwd != param.Password {
		//密码错误
		fmt.Printf("svc.LoginFail err: 密码错误")
		response.ToErrorResponse(errcode.ErrorPasswdWrong)
		return
	}
	token, err := app.GenerateToken(global.JWTSetting.Secret, global.JWTSetting.Issuer)
	if err != nil {
        fmt.Printf("app.GenerateToken err: %v", err)
        response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
        return
    }

	response.ToResponse(gin.H{
		"err":"ok",
		"token": token})
	return
}
