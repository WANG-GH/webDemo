package model

import "github.com/jinzhu/gorm"
type Record struct {
	Record_id uint32  `gorm:"primary_key" json:"record_id"`
	User_id        uint32  `json:"user_id"`
	Program_id uint32  `json:"program_id"`
	Status string `json:"status"`
}

func (record *Record)TableName() string{
	return "record"
}

func (record *Record)Create(db *gorm.DB) error{
	return db.Create(&record).Error
}