package v1

import (
	"fmt"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{
		"err": "ok",
	})
	return
}

func (user *User) CreateByEmail(c *gin.Context) {
	param := service.CreateUserRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	fmt.Println(param)
	err := svc.CreateByEmail(&param)
	if err != nil {
		fmt.Printf("svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"err": "ok",
	})
	return
}

func (user *User) CheckCreateByEmail(c *gin.Context) {
	param := service.ResetDataByEmailRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	fmt.Println(param)
	state, err := svc.CheckCreateByEmail(&param)
	if err != nil {
		response.ToResponse(gin.H{
			"err": "can not create",
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

func (user *User) Update(c *gin.Context) {
	param := service.UpdateRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	fmt.Printf("id = %v, name = %v", param.Userid, param.Username)
	svc := service.New(c.Request.Context())
	user_, err := svc.UpdateUser(&param)
	if err != nil {
		fmt.Printf("svc.UpdateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateFail)
		return
	}
	token, err := app.GenerateToken(user_.UserName, user_.Email, int(user_.Privilege), int(user_.ID))
	if err != nil {
		fmt.Printf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
		"err":   "ok",
	})
	return
}

func (user *User) Login(c *gin.Context) {
	param := service.LoginRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	user_, err := svc.GetStatus(&param)
	fmt.Printf("接受的密码:%v \n", param.Password)
	if err != nil {
		// 用户未存在
		fmt.Printf("svc.LoginFail err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserNotExist)
		return
	}
	if user_.Password != param.Password {
		//密码错误
		fmt.Printf("svc.LoginFail err: 密码错误")
		response.ToErrorResponse(errcode.ErrorPasswdWrong)
		return
	}
	token, err := app.GenerateToken(user_.UserName, user_.Email, int(user_.Privilege), int(user_.ID))
	if err != nil {
		fmt.Printf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
		"err":   "ok",
	})
	return
}

func (user *User) GetStatus(c *gin.Context) {
	// param := service.GetStatusRequest{}
	// c.ShouldBind(&param)
	// response := app.NewResponse(c)
	// svc := service.New(c.Request.Context())
	// var (
	// 	token string
	// 	ecode = errcode.Success
	// )
	// if s, exist := c.GetQuery("token"); exist {
	// 	token = s
	// } else {
	// 	token = c.GetHeader("token")
	// }
	// claim, _ := app.ParseToken(token)
	// svc.GetStatus(claim.Name)
}
