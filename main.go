package main

import (
	"fmt"
	"os"

	"github.com/Top-Pattarapol/school-service/service"

	"github.com/Top-Pattarapol/school-service/database"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()
	defer database.Close()

	r := gin.Default()
	setPath(r)
	r.Run(getPort())
}

func setPath(r *gin.Engine) {
	r.GET("/api/todos", service.GetTodos)
	r.GET("/api/todos/:id", service.GetTodosById)
	r.POST("/api/todos/", service.PostTodos)
	r.DELETE("/api/todos/:id", service.DeleteTodosById)
	r.PUT("/api/todos/:id", service.UpdateTodo)
}

func getPort() string {
	var port = os.Getenv("PORT") // ----> (A)
	if port == "" {
		port = "1234"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port // ----> (B)
}
