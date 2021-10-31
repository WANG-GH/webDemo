package routers

import (
	"webDemo/internal/middleware"
	v1 "webDemo/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	user := v1.NewUser()
	program := v1.NewProgram()

	apiv1 := r.Group("/api/v1")
	//apiv1.POST("/users", user.Create)
	apiv1.POST("/register", user.Create)
	apiv1.POST("/login", user.Login)
	apiv1.Use(middleware.JWT())
	apiv1.Use()
	{
		apiv1.DELETE("/users/:name", user.Delete)
		apiv1.PUT("/users/:name", user.Update)
		apiv1.GET("/users", user.Get)

		apiv1.POST("/programs", program.Create) //上传题目
		apiv1.GET("/programs/:program_id", program.ReturnProgramDetail)//获取题目详情
		apiv1.GET("/programs", program.ReturnProgramList)   //获取题目列表
		
		apiv1.POST("/programs/:program_id", program.SubmitProgram) //提交答案
	}
	return r
}
