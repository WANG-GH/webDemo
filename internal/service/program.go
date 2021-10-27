package service

import "webDemo/internal/model"

type CreateProgramRequest struct {
	Program_id   uint32 `form:"program_id" binding:"oneof=0 1"`
	Program_name string `form:"program_name" binding:"max=100"`
	Content      string `form:"content" binding:"required,min=10,max=500"`
	Ptype        string `form:"ptype" binding:"required,min=2,max=50"`
	Answer       string `form:"answer" binding:"required,min=1,max=100"`
	Difficulty   string `form:"difficulty" binding:"required,min=1,max=50"`
}

func (svc *Service) CreateProgram(param *CreateProgramRequest) error {
	return svc.dao.CreateProgram(param.Program_id, param.Program_name, param.Content, param.Ptype, param.Answer, param.Difficulty)
}

type ReturnProgramListRequest struct {
	Program_id   uint32 `form:"program_id" binding:"oneof=0 1"`
	Program_name string `form:"program_name" binding:"max=100"`
	Answer       string `form:"answer" binding:"required,min=1,max=100"`
	Difficulty   string `form:"difficulty" binding:"required,min=1,max=50"`
}

func (svc *Service) ReturnProgramList(param *ReturnProgramListRequest) ([]model.Program,error)  {
	return svc.dao.ReturnProgramList(param.Program_id, param.Program_name, param.Answer, param.Difficulty)
}

type ReturnProgramDetailRequest struct {
	Program_id   uint32 `form:"program_id" binding:"oneof=0 1"`
	Program_name string `form:"program_name" binding:"max=100"`
	Content      string `form:"content" binding:"required,min=10,max=500"`
	Answer       string `form:"answer" binding:"required,min=1,max=100"`
	Difficulty   string `form:"difficulty" binding:"required,min=1,max=50"`
}

func (svc *Service) ReturnProgramDetail(param *ReturnProgramDetailRequest) ([]model.Program,error) {
	return svc.dao.ReturnProgramDetail(param.Program_id, param.Program_name, param.Content, param.Answer, param.Difficulty)
}
