package main

import (
	"fmt"
	"log"
	"net/http"

	"todos_backend_server/adapter/portal"
	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain"
	"todos_backend_server/messagehandler"
)

func main() {
	host, port := portal.ParseFlags()
	repo := provider.NewJsonTodosRepository("./data/todos.json")
	createTodosRouter(repo)
	runServer(host, port)
}

func createTodosRouter(r domain.TodosRepository) {
	http.Handle("/api/todos/add-todo", portal.AddTodo(messagehandler.AddTodo(r)))
	http.Handle("/api/todos/clear-completed", portal.ClearCompleted(messagehandler.ClearCompleted(r)))
	http.Handle("/api/todos/destroy-todo", portal.DestroyTodo(messagehandler.DestroyTodo(r)))
	http.Handle("/api/todos/save-todo", portal.SaveTodo(messagehandler.SaveTodo(r)))
	http.Handle("/api/todos/select-todos", portal.SelectTodos(messagehandler.SelectTodos(r)))
}

func runServer(host string, port uint) {
	log.Println("Server listening on port", port)
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
