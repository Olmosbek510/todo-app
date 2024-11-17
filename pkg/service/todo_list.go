package service

import (
	"github.com/Olmosbek510/todo-app"
	"github.com/Olmosbek510/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func (t *TodoListService) Update(userId int, listId int, newListBody todo.UpdateListInput) error {
	if err := newListBody.Validate(); err != nil {
		return err
	}
	return t.repo.Update(userId, listId, newListBody)
}

func (t *TodoListService) DeleteById(userId, listId int) error {
	return t.repo.DeleteById(userId, listId)
}

func (t *TodoListService) GetById(userId, id int) (todo.TodoList, error) {
	return t.repo.GetById(userId, id)
}

func (t *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return t.repo.GetAll(userId)
}

func (t *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return t.repo.Create(userId, list)
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}
