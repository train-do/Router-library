package repository

import (
	"database/sql"
	"fmt"

	"github.com/train-do/Router-library/model"
)

type RepoTodo struct {
	Todo model.Todo
}

func (r *RepoTodo) FindAllTodo(db *sql.DB, userId string) ([]model.Todo, error) {
	query := `select * from "Todo" where user_id=$1;`
	todos := []model.Todo{}
	rows, err := db.Query(query, userId)
	if err != nil {
		fmt.Println("FindAllTodo Todo ErrorRow: ", err)
		return []model.Todo{}, err
	}
	for rows.Next() {
		err := rows.Scan(&r.Todo.Id, &r.Todo.UserId, &r.Todo.Todo, &r.Todo.IsDone)
		if err != nil {
			fmt.Println("FindAllTodo Todo ErrorNext: ", err)
			return []model.Todo{}, err
		}
		todos = append(todos, r.Todo)
	}
	return todos, nil
}
func (r *RepoTodo) InsertTodo(db *sql.DB) error {
	query := `insert into "Todo" (user_id, todo, is_done) values ($1, $2, $3) returning id;`
	err := db.QueryRow(query, r.Todo.UserId, r.Todo.Todo, r.Todo.IsDone).Scan(&r.Todo.Id)
	if err != nil {
		fmt.Println("Insert Todo Error: ", err)
		return err
	}
	return nil
}
