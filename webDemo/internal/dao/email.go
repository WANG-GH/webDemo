package dao

import (
	"webDemo/internal/model"
)


func (e *Dao) CreateEmail(email_num string, verify_code uint32) error {
	email := model.Email {
		Email_num:        email_num,
		Verify_code:      verify_code,
	}
	return email.CreateEmail(e.engine)
}

func (e *Dao) CheckVerifyCode(email_num string) (int,error) {
	email := model.Email {
		Email_num:        email_num,
	}
	return email.CheckVerifyCode(e.engine)
}