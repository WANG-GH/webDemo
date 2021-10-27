package dao

import (
	"fmt"
	"webDemo/internal/model"
)

func (p *Dao) CreateProgram(program_id uint32, program_name string, content string, ptype string, answer string, difficulty string) error {
	program := model.Program{
		Program_id:   program_id,
		Program_name: program_name,
		Content:      content,
		Ptype:        ptype,
		Answer:       answer,
		Difficulty:   difficulty,
	}
	return program.CreateProgram(p.engine)
}

func (p *Dao) ReturnProgramList(program_id uint32, program_name string, answer string, difficulty string) ([]model.Program, error) {
	program := model.Program{
		Program_id:   program_id,
		Program_name: program_name,
		Answer:       answer,
		Difficulty:   difficulty,
	}
	return program.ReturnProgramList(p.engine)
}

func (p *Dao) ReturnProgramDetail(program_id uint32, program_name string, content string, answer string, difficulty string) ([]model.Program, error) {
	program := model.Program{
		Program_id:   program_id,
		Program_name: program_name,
		Content:      content,
		Answer:       answer,
		Difficulty:   difficulty,
	}
	fmt.Printf("id = %d", program.Program_id)
	return program.ReturnProgramDetail(p.engine)
}
