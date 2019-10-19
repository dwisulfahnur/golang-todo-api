package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/dwisulfahnur/todo-api/db"
	TaskController "github.com/dwisulfahnur/todo-api/controllers"
)

func main() {
	log.Println("Starting server..")

	db.Init();

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("/", TaskController.GetTasks)
			tasks.GET(":id", TaskController.GetTask)
			tasks.POST("/", TaskController.CreateTask)
			tasks.PUT("/:id", TaskController.UpdateTask)
			tasks.DELETE("/:id", TaskController.DeleteTask)
		}
	}

	r.Run()
}