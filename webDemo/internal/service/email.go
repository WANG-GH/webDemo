package service

import (
	"fmt"
	"math/rand"
	"time"
	"webDemo/global"
	"webDemo/pkg/email"
)

//"context"
//"errors"
//"fmt"
//"io"
//"os"
//"strconv"
//"time"

//"fmt"
//"io/ioutil"
//"webDemo/internal/model"

//"github.com/docker/docker/api/types"
//"github.com/docker/docker/api/types/container"
//"github.com/docker/docker/client"
type CreateEmailRequest struct {
	Email_num string `form:"email_num"`
}

func (svc *Service) CreateEmail(param *CreateEmailRequest) error {
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
	fmt.Printf("to = %v", param.Email_num)
	err := mailer.SendMail(param.Email_num,
		fmt.Sprintf("重置信息验证码邮件"),
		fmt.Sprintf("您的重置信息验证码为: %d", verify_code),
	)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return svc.dao.CreateEmail(param.Email_num, verify_code)
}

type ResetDataByEmailRequest struct {
	Email_num   string `form:"email_num"`
	Verify_code uint32 `form:"verify_code"`
	User_name   string `form:"user_name"`
	Password    string `form:"password" binding:"max=100"`
}

// 更改了更新用户，通过username更新
// 1= success, 2 = not found, 3=wrong verify code
// 通过id
func (svc *Service) ResetDataByEmail(param *ResetDataByEmailRequest) (int, error) {
	verify_code, err := svc.dao.CheckVerifyCode(param.Email_num)
	fmt.Printf("%v", param)
	if err != nil {
		return 2, nil
	}
	if verify_code == int(param.Verify_code) {
		_, err = svc.dao.UpdateUser(param.User_name, param.Password, 0)
		if err != nil {
			return 0, err
		}
		return 1, nil
	} else {
		return 3, nil
	}
}
