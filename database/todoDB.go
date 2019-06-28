package database

import (
	"database/sql"

	"github.com/Top-Pattarapol/school-service/model"
)

func CreateTodoTable() error {
	return baseExec(db, `CREATE TABLE IF NOT EXISTS todos( id SERIAL PRIMARY KEY, title TEXT, status TEXT );`)
}

func GetTodos() (*sql.Rows, error) {
	return baseQuery(db, `Select id, title, status FROM todos ORDER BY id ASC`)
}

func GetTodoById(id int) (*sql.Row, error) {
	return baseQueryRow(db, `Select id, title, status FROM todos WHERE id=$1 ORDER BY id ASC`, id)
}

func PostTodos(t *model.Todo) (*sql.Row, error) {
	return baseQueryRow(db, `INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id`, t.Title, t.Status)
}

func DeleteTodoById(id int) error {
	return baseExec(db, `DELETE FROM todos WHERE id=$1`, id)
}

func UpdateTodo(id int, title string, status string) error {
	return baseExec(db, `UPDATE todos SET title=$2, status=$3 WHERE id=$1`, id, title, status)
}
