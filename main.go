package main

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/yisleyen/todo-app/middleware"
	"github.com/yisleyen/todo-app/models"
)

func main() {
	listTodo()
}

func createTodo() {
	var todo models.Todos

	todo.Name = "Günlük işler planlanacak"

	id, err := middleware.CreateTodo(todo)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Todo added successfully %v", id)
}

func listTodo() {
	var todos []models.Todos

	todos, err := middleware.GetAllTodos()

	if err != nil {
		panic(err)
	}

	for _, todo := range todos {
		fmt.Printf("%d %s\n", todo.Id, todo.Name)
	}
}

func updateTodo() {
	var todo models.Todos

	todo.Name = "Günlük işler planlanacak"

	id, err := middleware.UpdateTodo(2, todo)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v record updated", id)
}

func deteleTodo() {
	id, err := middleware.DeleteTodo(2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v record deleted", id)
}
