package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Top-Pattarapol/school-service/database"
	"github.com/Top-Pattarapol/school-service/model"

	"github.com/gin-gonic/gin"
)

func DeleteTodosById(c *gin.Context) {

	id, err := paramToInt(c, "id")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = database.DeleteTodoById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func PostTodos(c *gin.Context) {

	t := &model.Todo{}

	if err := c.BindJSON(t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var id int
	row, err := database.PostTodos(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	err = row.Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	t.Id = id
	c.JSON(http.StatusCreated, t)

}

func GetTodosById(c *gin.Context) {

	id, err := paramToInt(c, "id")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	row, err := database.GetTodoById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	t := model.Todo{}
	err = row.Scan(&t.Id, &t.Title, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, t)
}

func GetTodos(c *gin.Context) {

	rows, err := database.GetTodos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todos := []model.Todo{}

	for rows.Next() {
		t := model.Todo{}
		err := rows.Scan(&t.Id, &t.Title, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, t)
	}

	fmt.Println(todos)
	c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {

	id, err := paramToInt(c, "id")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t := &model.Todo{}

	if err := c.BindJSON(t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	t.Id = id

	err = database.UpdateTodo(id, t.Title, t.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}

func paramToInt(c *gin.Context, key string) (int, error) {
	param := c.Param(key)
	value, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return value, nil
}
