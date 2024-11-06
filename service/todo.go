package service

import (
	"database/sql"
	"fmt"

	"github.com/train-do/Router-library/model"
	"github.com/train-do/Router-library/repository"
)

type ServiceTodo struct {
	Db *sql.DB
}

func (s *ServiceTodo) FindAllTodo(userId string) ([]model.Todo, error) {
	repoTodo := repository.RepoTodo{}
	todos, err := repoTodo.FindAllTodo(s.Db, userId)
	if err != nil {
		fmt.Println("Insert Todo :", err)
		return []model.Todo{}, err
	}
	return todos, nil
}
func (s *ServiceTodo) InsertTodo(t *model.Todo) error {
	repoTodo := repository.RepoTodo{
		Todo: *t,
	}
	if err := repoTodo.InsertTodo(s.Db); err != nil {
		fmt.Println("Insert Todo :", err)
		return err
	}
	return nil
}
func (s *ServiceTodo) UpdateTodo(u *model.Todo) error {
	return nil
}
func (s *ServiceTodo) DeleteTodo(u *model.Todo) error {
	return nil
}
