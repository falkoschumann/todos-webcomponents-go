package messagehandler

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/domain"
)

func TestAddTodo(t *testing.T) {
	t.Run("saves new todo.", func(t *testing.T) {
		repo := newMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		})
		addTodo := AddTodo(repo)

		status := addTodo(domain.AddTodoCommand{Title: "Buy Unicorn"})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("AddTodo() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		}
		if diff := cmp.Diff(want, repo.Load()); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("does nothing if title is empty.", func(t *testing.T) {
		repo := newMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		})
		addTodo := AddTodo(repo)

		status := addTodo(domain.AddTodoCommand{Title: ""})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("AddTodo() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		}
		if diff := cmp.Diff(want, repo.Load()); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})
}
