package service

import (
	"errors"
	//"fmt"
	"webDemo/internal/model"
)

type CreateProgramRequest struct {
	Program_id   uint32 `form:"program_id"`
	Program_name string `form:"program_name" binding:"max=100"`
	Content      string `form:"content" binding:"required,min=10,max=500"`
	Ptype        string `form:"ptype" binding:"required,min=2,max=50"`
	Answer       string `form:"answer" binding:"required,min=1,max=100"`
	Difficulty   string `form:"difficulty" binding:"required,min=1,max=50"`
}

func (svc *Service) CreateProgram(param *CreateProgramRequest) error {
	return svc.dao.CreateProgram(param.Program_name, param.Content, param.Ptype, param.Answer, param.Difficulty)
}

type SubmitProgramRequest struct {
	User_id    uint32 `form:"user_id"`
	Program_id uint32 `form:"program_id"`
	Answer     string `form:"answer" binding:"required,min=1,max=50000"`
}

type ReturnProgramListRequest struct {
	Program_id   uint32 `form:"program_id"`
	Program_name string `form:"program_name" binding:"max=100"`
	Answer       string `form:"answer" binding:"required,min=1,max=100"`
	Difficulty   string `form:"difficulty" binding:"required,min=1,max=50"`
}

func (svc *Service) ReturnProgramList(param *ReturnProgramListRequest) ([]model.Program, error) {
	return svc.dao.ReturnProgramList(param.Program_id, param.Program_name, param.Answer, param.Difficulty)
}

type ReturnProgramDetailRequest struct {
	Program_id   uint32 `form:"program_id"`
	Program_name string `form:"program_name" binding:"max=100"`
}

func (svc *Service) ReturnProgramDetail(param *ReturnProgramDetailRequest) ([]model.Program, error) {
	return svc.dao.ReturnProgramDetail(param.Program_id, param.Program_name)
}

func (svc *Service) SubmitProgram(param *SubmitProgramRequest) (bool, error) {
	program, err := svc.dao.ReturnProgramDetail(param.Program_id, "")
	if(len(program) < 1){
		return false, errors.New("no program")
	}
	if err != nil {
		return false, err
	}
	if program[0].Answer == param.Answer {
		err = svc.dao.CreateRecord(param.Program_id, param.User_id, "pass")
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		err = svc.dao.CreateRecord(param.Program_id, param.User_id, "not pass")
		if err != nil {
			return false, err
		}
		return false, nil
	}
}	
