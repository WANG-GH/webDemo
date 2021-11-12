package v1

import (
	"fmt"
	"strconv"

	// "strings"
	"webDemo/internal/service"
	"webDemo/pkg/app"
	"webDemo/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Program struct{}
type Record struct{}

func NewProgram() Program {
	return Program{}
}

func NewRecord() Record {
	return Record{}
}

func (program *Program) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

func (program *Program) List(c *gin.Context) {

	response := app.NewResponse(c)

	response.ToResponse(gin.H{})
	return
}

func (program *Program) Delete(c *gin.Context) {

	param := service.DeleteProgramRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	program_id, err := strconv.Atoi(c.Param("program_id"))
	if err != nil{
		response.ToResponse(gin.H{"err": "param wrong"})
		return
	}
	// 确认管理员

	err = svc.DeleteProgram(uint32(program_id))
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteProgram)
		return
	}
	response.ToResponse(gin.H{"err": "ok"})
	return
	
}

func (program *Program) Create(c *gin.Context) {
	param := service.CreateProgramRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	err := svc.CreateProgram(&param)
	if err != nil {
		fmt.Printf("svc.CreateProgram err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateProgramFail)
		return
	}
	response.ToResponse(gin.H{"err": "ok"})
	return
}

func (program *Program) ReturnProgramList(c *gin.Context) {
	param := service.ReturnProgramListRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	programs, err := svc.ReturnProgramList(&param)
	if err != nil {
		fmt.Printf("svc.ReturnProgramList err: %v", err)
		response.ToErrorResponse(errcode.ErrorReturnProgramListFail)
		return
	}
	c.JSON(200, gin.H{
		"err":      "ok",
		"programs": programs,
	})
	return
}

func (program *Program) ReturnProgramDetail(c *gin.Context) {
	param := service.ReturnProgramDetailRequest{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	fmt.Printf("in api: id =  %s", c.Param("program_id"))
	id, err := strconv.Atoi(c.Param("program_id"))
	param.Program_id = uint32(id)

	programs, err := svc.ReturnProgramDetail(&param)
	if err != nil {
		fmt.Printf("svc.ReturnProgramDetail err: %v", err)
		response.ToErrorResponse(errcode.ErrorReturnProgramDetail)
		return
	}
	c.JSON(200, gin.H{
		"id":         programs[0].Program_id,
		"name":       programs[0].Program_name,
		"content":    programs[0].Content,
		"difficulty": programs[0].Difficulty,
		"err":        "ok",
	})
	return
}

func (program *Program) SubmitProgram(c *gin.Context) {
	param := service.SubmitProgramRequest{}
	c.ShouldBind(&param)
	fmt.Printf("answer = %v, userid = %v, proid = %v", param.Answer, param.User_id, param.Program_id)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	pass, err := svc.SubmitProgram(&param)
	// if strings.Contains(err.Error(), "no program") {
	// 	response.ToResponse(gin.H{
	// 		"err": "no program",s
	// 	})
	// 	return
	// }
	if err != nil {
		response.ToResponse(gin.H{
			"err": "can not submit",
		})
		return
	}
	if pass {
		response.ToResponse(gin.H{
			"err":    "ok",
			"status": "pass",
		})
	} else {
		response.ToResponse(gin.H{
			"err":    "ok",
			"status": "not pass",
		})
	}
	return
}

