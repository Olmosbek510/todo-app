package repository

import (
	"github.com/Olmosbek510/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type (
	TodoList interface {
		Create(id int, list todo.TodoList) (int, error)
		GetAll(id int) ([]todo.TodoList, error)
		GetById(userId, listId int) (todo.TodoList, error)
		DeleteById(userId, listId int) error
		Update(userId, listId int, newListBody todo.UpdateListInput) error
	}
)

type TodoItem interface {
	Create(listId int, todoItem todo.TodoItem) (int, error)
	GetAll(userId, lisId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId int, itemId int, itemInput todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
