package service

import (
	"fmt"
	"math/rand"
	"time"
	"webDemo/global"
	"webDemo/internal/model"
	"webDemo/pkg/email"
)

//import "webDemo/pkg/errcode"

type CreateUserRequest struct {
	UserName  string `form:"username" binding:"max=100"`
	Password  string `form:"password" binding:"max=100"`
	Privilege uint32 `form:"privilege,default=0" binding:"oneof=0 1"`
	Email     string `form:"email" binding:"max=100"`
}
type LoginRequest struct {
	Name     string `form:"name" binding:"max=100"`
	Password string `form:"password" binding:"max=100"`
}
type UpdateRequest struct {
	Userid    uint32 `form:"userid"`
	Username  string `form:"username" binding:"max=100"`
	Password  string `form:"password" binding:"max=100"`
	Privilege uint32 `form:"privilege,default=0" binding:"oneof=0 1"`
}

func (svc *Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(param.Email, param.UserName, param.Password, param.Privilege)
}

func (svc *Service) LoginUser(param *LoginRequest) (string, error) {
	return svc.dao.GetUserPasswd(param.Name)
}

func (svc *Service) UpdateUser(param *UpdateRequest) (model.User, error) {
	return svc.dao.UpdateUser(param.Username, param.Password, param.Privilege)
}

func (svc *Service) GetStatus(param *LoginRequest) (model.User, error) {
	return svc.dao.GetStatus(param.Name)
}

func (svc *Service) CreateByEmail(param *CreateUserRequest) error {
	rand.Seed(time.Now().Unix())
	verify_code := rand.Uint32() % 10000
	mailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	fmt.Printf("to = %v", param.Email)
	err := mailer.SendMail(param.Email,
		fmt.Sprintf("注册验证"),
		fmt.Sprintf("您的注册验证码为: %d", verify_code),
	)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return svc.dao.CreateEmail(param.Email, verify_code)
}

// 只能创建非特权用户
func (svc *Service) CheckCreateByEmail(param *ResetDataByEmailRequest) (int, error) {
	verify_code, err := svc.dao.CheckVerifyCode(param.Email_num)
	fmt.Printf("%v", param)
	if err != nil {
		return 2, nil
	}
	if verify_code == int(param.Verify_code) {
		err =svc.dao.CreateUser(param.Email_num, param.User_name, param.Password, 0)
		if err != nil {
			return 0, err
		}
		return 1, nil
	} else {
		return 3, nil
	}
}
