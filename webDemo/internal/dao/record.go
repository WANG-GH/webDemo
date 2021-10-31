package dao

import "webDemo/internal/model"

func (r *Dao) CreateRecord(program_id uint32, user_id uint32,status string) error {
	record := model.Record{
		User_id: user_id,
		Program_id:   program_id,
		Status: status,
	}
	return record.Create(r.engine)
}