package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Program struct {
	Program_id   uint32    `gorm:"primary_key" json:"program_id"`
	CreatedAt    time.Time `json:"create_at"`
	UpdatedAt    time.Time `json:"update_at"`
	Program_name string    `json:"program_name"`
	Content      string    `json:"content"`
	Ptype        string    `json:"ptype"`
	Answer       string    `json:"answer"`
	Difficulty   string    `json:"difficulty"`
}

func (program *Program) TableName() string {
	return "program"
}

func (p *Program) CreateProgram(db *gorm.DB) error { //创建题目
	return db.Create(&p).Error
}

func (p *Program) GetContent(db *gorm.DB) error { //读取题目内容
	return db.First(&p).Error
}

func (p *Program) DeleteProgram(db *gorm.DB) error { //删除题目
	return db.Where("program_id = ?", p.Program_id).Delete(&p).Error
}
func (p *Program) ReturnProgramList(db *gorm.DB) ([]Program, error) {
	var programs []Program
	result := db.Find(&programs)
	if result.Error != nil {
		return nil, result.Error
	}
	return programs, nil
}

func (p *Program) ReturnProgramDetail(db *gorm.DB) ([]Program, error) {
	var programs []Program
	fmt.Printf("id = %v", p.Program_id)
	result := db.Where("program_id = ?", p.Program_id).Find(&programs)
	if result.Error != nil {
		return nil, result.Error
	}
	return programs, nil
}
