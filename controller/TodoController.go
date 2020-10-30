package controller

import (
	"Todos/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 url ---> controller ---> logic ---> model
 请求      控制器			  业务逻辑	 模型成的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	// 1. 从请求中拿数据
	var todo models.Todo
	c.BindJSON(&todo)

	// 2. 存入数据库
	// err = DB.Create(&todo).Error

	// 3. 返回响应
	if err := models.CreateTodo(&todo); err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)

		//c.JSON(http.StatusOK, gin.H{
		//	"code":	200,
		//	"msg": 	"success",
		//	"data": todo,
		//})
	}
}

func GetTodoList(c *gin.Context) {
	todoList,err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&todo)
	if err = models.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}else {
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}