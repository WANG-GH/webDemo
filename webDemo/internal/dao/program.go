package dao

import (
	"fmt"
	"webDemo/internal/model"
)

func (p *Dao) CreateProgram(program_name string, content string, ptype string, answer string, difficulty string) error {
	program := model.Program{
		Program_name: program_name,
		Content:      content,
		Ptype:        ptype,
		Answer:       answer,
		Difficulty:   difficulty,
	}
	return program.CreateProgram(p.engine)
}

func (p *Dao) DeleteProgram(program_id uint32) error {
	program := model.Program{
		Program_id: program_id,
	}
	return program.DeleteProgram(p.engine)
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

func (p *Dao) ReturnProgramDetail(program_id uint32, program_name string) ([]model.Program, error) {
	program := model.Program{
		Program_id:   program_id,
		Program_name: program_name,
	}
	fmt.Printf("id = %d", program.Program_id)
	return program.ReturnProgramDetail(p.engine)
}
