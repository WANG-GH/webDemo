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

type Record1 struct{}

func NewRecord1() Record1 {
	return Record1{}
}

func (record *Record1) ReturnRecord(c *gin.Context) {
	param := service.ReturnRecord{}
	c.ShouldBind(&param)
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	id, err := strconv.Atoi(c.Param("user_id"))

	fmt.Printf("paramid = %v, id = %v", c.Param("user_id"), id)
	param.User_id = uint32(id)

	records, err := svc.ReturnRecord(&param)
	if err != nil {
		fmt.Printf("svc.ReturnRecord err: %v", err)
		response.ToErrorResponse(errcode.ErrorReturnRecord)
		return
	}
	if len(records) == 0 {
		c.JSON(200, gin.H{
			"err": "no program",
		})
		return
	}
	c.JSON(200, gin.H{
		"records": records,
		"err":     "ok",
	})
	return
}
