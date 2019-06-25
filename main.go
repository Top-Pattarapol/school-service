package main

// var db *sql.DB

func main() {

	database.connectDataBase()
	// defer db.Close()
	// r := gin.Default()
	// setPath(r)
	// r.Run(getPort())
}

// func setPath(r *gin.Engine) {
// 	r.GET("/api/todos", getTodos)
// 	r.GET("/api/todos/:id", getTodosById)
// 	r.POST("/api/todos/", postTodos)
// 	r.DELETE("/api/todos/:id", deleteTodosById)
// }

// func getPort() string {
// 	var port = os.Getenv("PORT") // ----> (A)
// 	if port == "" {
// 		port = "1234"
// 		fmt.Println("No Port In Heroku" + port)
// 	}
// 	return ":" + port // ----> (B)
// }

// func ParamToInt(c *gin.Context, key string) int {
// 	param := c.Param(key)

// 	value, err := strconv.Atoi(param)
// 	if err != nil {
// 		return 0
// 	}
// 	return value
// }

// func deleteTodosById(c *gin.Context) {

// 	id := ParamToInt(c, "id")

// 	stmt, err := db.Prepare(`DELETE FROM todos WHERE id=$1`)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	_, err = stmt.Query(id)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "success",
// 	})
// }

// func postTodos(c *gin.Context) {

// 	t := &Todo{}

// 	if err := c.BindJSON(t); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	query := `
// 	INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id
// 	`

// 	var id int
// 	row := db.QueryRow(query, t.Title, t.Status)
// 	err := row.Scan(&id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	t.Id = id
// 	c.JSON(http.StatusCreated, t)

// }

// func getTodosById(c *gin.Context) {

// 	id := ParamToInt(c, "id")

// 	stmt, err := db.Prepare(`Select id, title, status FROM todos WHERE id=$1 ORDER BY id ASC`)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	row := stmt.QueryRow(id)

// 	t := Todo{}
// 	err = row.Scan(&t.Id, &t.Title, &t.Status)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, t)
// }

// func getTodos(c *gin.Context) {

// 	stmt, err := db.Prepare(`Select id, title, status FROM todos ORDER BY id ASC`)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	rows, err := stmt.Query()

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	todos := []Todo{}

// 	for rows.Next() {
// 		t := Todo{}
// 		err := rows.Scan(&t.Id, &t.Title, &t.Status)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		todos = append(todos, t)
// 	}

// 	fmt.Println(todos)
// 	c.JSON(http.StatusOK, todos)
// }
