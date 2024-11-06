package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/train-do/Router-library/model"
)

type RepoUser struct {
	User model.User
}

func (r *RepoUser) InsertUser(db *sql.DB) error {
	query := `insert into "User" (id, name, email, password) values ($1, $2, $3, $4) returning id;`
	err := db.QueryRow(query, uuid.New().String(), r.User.Name, r.User.Email, r.User.Password).Scan(&r.User.Id)
	if err != nil {
		fmt.Println("Insert User Error: ", err)
		return err
	}
	return nil
}
func (r *RepoUser) Login(db *sql.DB) error {
	query := `select id, name from "User" where email=$1 and password=$2;`
	err := db.QueryRow(query, r.User.Email, r.User.Password).Scan(&r.User.Id, &r.User.Name)
	if err != nil {
		fmt.Println("Login User Error: ", err)
		return err
	}
	// fmt.Println(r.User.Id, "------")
	return nil
}
func (r *RepoUser) FindById(db *sql.DB) error {
	query := `select id from "User" where id=$1;`
	err := db.QueryRow(query, r.User.Id).Scan(&r.User.Id)
	if err != nil {
		fmt.Println("FindById User Error: ", err)
		return err
	}
	return nil
}
func (r *RepoUser) FindAllUser(db *sql.DB) ([]model.User, error) {
	var users []model.User
	query := `select * from "User";`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("FindById User Error: ", err)
		return []model.User{}, err
	}
	for rows.Next() {
		rows.Scan(&r.User.Id, &r.User.Name, &r.User.Email, &r.User.Password)
		if err != nil {
			fmt.Println("FindById User Error: ", err)
			return []model.User{}, err
		}
		users = append(users, r.User)
	}
	return users, nil
}
