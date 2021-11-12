package dao

import (
	"fmt"
	"webDemo/internal/model"
)

func (r *Dao) CreateRecord(program_id uint32, user_id uint32, status string) error {
	record := model.Record{
		User_id:    user_id,
		Program_id: program_id,
		Status:     status,
	}
	return record.Create(r.engine)
}

func (r *Dao) ReturnRecord(record_id uint32, program_id uint32, user_id uint32, status string) ([]model.Record, error) {
	record := model.Record{
		Record_id:  record_id,
		User_id:    user_id,
		Program_id: program_id,
		Status:     status,
	}
	fmt.Printf("dao record %v", record)
	fmt.Printf("dao record_id = %v", record.User_id)
	return record.ReturnRecord(r.engine)
}
