package model

import (
	//"time"
	"github.com/jinzhu/gorm"
)

type Email struct {
	Email_num   string `gorm:"primary_key" json:"email_num"`
	Verify_code uint32 `json:"verify_code"`
	//CreatedAt    time.Time `json:"create_at"`
	//UpdatedAt    time.Time `json:"update_at"`
}

func (email *Email) TableName() string {
	return "email"
}

//先测一下能不能法，以及重复之后能不能删了重建e
func (e *Email) CreateEmail(db *gorm.DB) error {
	e_copy := *e
	db.Where("Email_num = ?", e.Email_num).Delete(&e)
	return db.Create(&e_copy).Error
}

//验证一次就删除
func (e *Email) CheckVerifyCode(db *gorm.DB) (int, error) {
	result := db.Where("Email_num = ?", e.Email_num).First(&e)
	if result.Error != nil {
		return 0, result.Error
	}
	ver_code := e.Verify_code
	db.Where("Email_num = ?", e.Email_num).Delete(&e)
	return int(ver_code), nil
}
