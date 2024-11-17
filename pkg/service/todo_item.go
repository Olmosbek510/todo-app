package service

import (
	"github.com/Olmosbek510/todo-app"
	"github.com/Olmosbek510/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func (t *TodoItemService) Update(userId, itemId int, itemInput todo.UpdateItemInput) error {
	return t.repo.Update(userId, itemId, itemInput)
}

func (t *TodoItemService) Delete(userId, itemId int) error {
	return t.repo.Delete(userId, itemId)
}

func (t *TodoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return t.repo.GetById(userId, itemId)
}

func (t *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return t.repo.GetAll(userId, listId)
}

func (t *TodoItemService) Create(userId int, listId int, todoItem todo.TodoItem) (int, error) {
	_, err := t.listRepo.GetById(userId, listId)
	if err != nil {
		// the listRepo does not exist or does not belong to user
		return 0, err
	}
	return t.repo.Create(listId, todoItem)
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}
