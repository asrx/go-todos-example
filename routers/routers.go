package routers

import (
	"Todos/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine  {
	r := gin.Default()

	// 加载静态文件
	r.Static("/static","static")
	// 加载模版文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项

		// 添加
		v1Group.POST("/todo", controller.CreateATodo)

		// 查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	r.Run()

	return r
}