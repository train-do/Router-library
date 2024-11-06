package service

import (
	"database/sql"
	"fmt"

	"github.com/train-do/Router-library/model"
	"github.com/train-do/Router-library/repository"
)

type ServiceUser struct {
	Db *sql.DB
}

func (s *ServiceUser) CreateUser(u *model.User) error {
	repoUser := repository.RepoUser{
		User: *u,
	}
	if err := repoUser.InsertUser(s.Db); err != nil {
		fmt.Println("Inser User :", err)
		return err
	}
	return nil
}
func (s *ServiceUser) Login(u *model.User) error {
	repoUser := repository.RepoUser{
		User: *u,
	}
	if err := repoUser.Login(s.Db); err != nil {
		fmt.Println("Login User :", err)
		return err
	}
	*u = repoUser.User
	// fmt.Println(u.Id, "++++++++")
	return nil
}
func (s *ServiceUser) GetById(id string) error {
	repoUser := repository.RepoUser{
		User: model.User{
			Id: id,
		},
	}
	if err := repoUser.FindById(s.Db); err != nil {
		fmt.Println("FindById User :", err)
		return err
	}
	return nil
}
func (s *ServiceUser) FindAllUser() ([]model.User, error) {
	repoUser := repository.RepoUser{}
	users, err := repoUser.FindAllUser(s.Db)
	if err != nil {
		fmt.Println("FindById User :", err)
		return []model.User{}, err
	}
	return users, nil
}
