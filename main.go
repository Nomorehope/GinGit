package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var tasks = []string{"task1", "task2", "task3"}

func tasksHandler(ctx *gin.Context) {
	switch ctx.Request.Method {
	case "GET":
		ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
	case "POST":
		var newTask string
		if err := ctx.BindJSON(&newTask); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неправильный ввод"})
			return
		}
		tasks = append(tasks, newTask)
		ctx.JSON(http.StatusOK, gin.H{"message": "Задача добавлена"})
	default:
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Ошибка расопзнавания метода"})
	}
}

func main() {

	r := gin.Default()
	r.GET("/tasks", tasksHandler)
	r.POST("/tasks", tasksHandler)
	s
	r.Run(":8080")
}
