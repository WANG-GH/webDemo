package model

import (
	"time"
	"webDemo/pkg/errcode"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint32    `gorm:"primarykey" json:"user_id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	UserName  string    `json:"username"` // 用户名
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Privilege uint32    `json:"privilege"` //normal,admin,super
}

func (u *User) CreateUser(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u *User) Update(db *gorm.DB) error {
	return db.Model(&User{}).Where("id = ?", u.ID).Update(u).Error
}

func (u *User) DeleteById(db *gorm.DB) error {
	return db.Where("id = ?", u.ID).Delete(&u).Error
}

func (u *User) DeleteByName(db *gorm.DB) error {
	return db.Where("user_name = ?", u.UserName).Delete(&u).Error
}

func (u *User) Count(db *gorm.DB) (int, error) {
	count := 0
	if u.UserName != "" {
		db = db.Where("user_name = ?", u.UserName)
	}
	if err := db.Model(&u).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (u *User) GetPasswd(db *gorm.DB) (string, error) {
	var user User
	result := db.Where("user_name = ?", u.UserName).First(&user)

	if result.Error != nil {
		return "", errcode.ErrorUserNotExist
	}
	return user.Password, nil
}

func (u *User) GetStatus(db *gorm.DB) (User, error) {
	var user User
	result := db.Where("user_name = ?", u.UserName).First(&user)

	if result.Error != nil {
		return user, errcode.ErrorUserNotExist
	}
	return user, nil
}