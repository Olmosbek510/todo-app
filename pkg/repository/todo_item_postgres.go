package repository

import (
	"fmt"
	"github.com/Olmosbek510/todo-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func (t *TodoItemPostgres) Update(userId int, itemId int, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`update %s ti set %s from %s li, %s ul
									where ti.id = li.item_id and li.list_id = ul.list_id and ul.user_id = $%d and ti.id = $%d
    `, todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)

	args = append(args, userId, itemId)

	logrus.Debug("updateQuery:", query)
	logrus.Debug("args", args)

	_, err := t.db.Exec(query, args...)
	return err
}

func (t *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`
	DELETE FROM %s ti
	USING %s li, %s ul
	WHERE ti.id = li.item_id
  		AND li.list_id = ul.list_id
  		AND ul.user_id = $1
  		AND ti.id = $2;
    `, todoItemsTable, listsItemsTable, usersListsTable)
	logrus.Printf("Generated query: %s", query)
	logrus.Printf("Args: userId=%d, itemId=%d", userId, itemId)
	_, err := t.db.Exec(query, userId, itemId)
	return err
}

func (t *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	todoItemQuery := fmt.Sprintf(`
	SELECT ti.id, ti.title, ti.description, ti.done
	FROM %s ti
         JOIN %s li on ti.id = li.item_id
         JOIN %s ul on ul.list_id = li.list_id AND ti.id = $1 AND ul.user_id = $2
`, todoItemsTable, listsItemsTable, usersListsTable)
	var item todo.TodoItem
	if err := t.db.Get(&item, todoItemQuery, itemId, userId); err != nil {
		return item, err
	}
	return item, nil
}

func (t *TodoItemPostgres) GetAll(userId int, listId int) ([]todo.TodoItem, error) {
	todoItemsQuery := fmt.Sprintf(`
	SELECT ti.id, ti.title, ti.description, ti.done
	FROM %s ti
         JOIN %s li on ti.id = li.item_id
         JOIN %s ul on li.list_id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2
	`, todoItemsTable, listsItemsTable, usersListsTable)
	var items []todo.TodoItem
	if err := t.db.Select(&items, todoItemsQuery, userId, listId); err != nil {
		return items, err
	}
	return items, nil
}

func (t *TodoItemPostgres) Create(listId int, todoItem todo.TodoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int

	createItemQuery := fmt.Sprintf(`INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id`, todoItemsTable)
	row := t.db.QueryRow(createItemQuery, todoItem.Title, todoItem.Description)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListsItemsQuery := fmt.Sprintf(`INSERT INTO %s (item_id, list_id) VALUES ($1, $2)`, listsItemsTable)
	_, err = t.db.Exec(createListsItemsQuery, itemId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}
