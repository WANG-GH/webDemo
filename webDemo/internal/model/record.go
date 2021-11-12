package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Record struct {
	Record_id  uint32 `gorm:"primary_key" json:"record_id"`
	User_id    uint32 `json:"user_id"`
	Program_id uint32 `json:"program_id"`
	Status     string `json:"status"`
}

func (record *Record) TableName() string {
	return "record"
}

func (record *Record) Create(db *gorm.DB) error {
	return db.Create(&record).Error
}

func (record *Record) ReturnRecord(db *gorm.DB) ([]Record, error) {
	var records []Record
	fmt.Printf("model id = %d", record.User_id)
	result := db.Where("user_id = ?", record.User_id).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}
