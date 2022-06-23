package middleware

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yisleyen/todo-app/models"
)

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Fatalf("Error opening database")
	}

	return db
}

func GetAllTodos() ([]models.Todos, error) {
	todos, err := getAllTodos()

	if err != nil {
		log.Fatalf("Unable to get all todos. %v", err)
	}

	return todos, err
}

func getAllTodos() ([]models.Todos, error) {
	db := createConnection()

	defer db.Close()

	var todos []models.Todos

	sqlStatement := `SELECT * FROM todos`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var todo models.Todos

		err = rows.Scan(&todo.Id, &todo.Name)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		todos = append(todos, todo)
	}

	return todos, err
}

func CreateTodo(todo models.Todos) (int64, error) {
	id, err := createTodo(todo)

	if err != nil {
		log.Fatalf("Unable to create todos. %v", err)
	}

	return id, err
}

func createTodo(todo models.Todos) (int64, error) {
	db := createConnection()

	sqlStatement := `INSERT INTO todos (name) VALUES ($1) RETURNING Id`

	var id int64

	err := db.QueryRow(sqlStatement, todo.Name).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer db.Close()

	return id, err
}

func UpdateTodo(id int64, todo models.Todos) (int64, error) {
	id, err := updateTodo(id, todo)

	if err != nil {
		log.Fatalf("Unable to update todos. %v", err)
	}

	return id, err
}

func updateTodo(id int64, todo models.Todos) (int64, error) {
	db := createConnection()

	sqlStatement := `UPDATE todos SET name=$2 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, todo.Name)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	defer db.Close()

	return rowsAffected, err
}

func DeleteTodo(id int64) (int64, error) {
	id, err := deleteTodo(id)

	if err != nil {
		log.Fatalf("Unable to delete todo. %v", err)
	}

	return id, err
}

func deleteTodo(id int64) (int64, error) {
	db := createConnection()

	sqlStatement := `DELETE FROM todos WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	defer db.Close()

	return rowsAffected, err
}
