package handler

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/train-do/Router-library/model"
	"github.com/train-do/Router-library/service"
)

var templates *template.Template

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			templates = template.New("")
			_, err := templates.ParseFiles("pages/login.html")
			if err != nil {
				log.Fatalf("Error parsing login template: %v", err)
			}
			_, err = templates.ParseGlob("templates/*.html")
			if err != nil {
				log.Fatalf("Error parsing templates: %v", err)
			}
			err = templates.ExecuteTemplate(w, "login.html", nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else if r.Method == http.MethodPost {
			email := r.FormValue("email")
			password := r.FormValue("password")
			user := model.User{
				Email:    email,
				Password: password,
			}
			serviceUser := service.ServiceUser{Db: db}
			err := serviceUser.Login(&user)
			if err != nil {
				unauthorized := model.Response{
					StatusCode: http.StatusUnauthorized,
					Message:    "Login Failed",
					Data:       nil,
				}
				json.NewEncoder(w).Encode(unauthorized)
				return
			}
			cookie := http.Cookie{
				Name:   "access_token",
				Value:  user.Id,
				Path:   "/",
				Domain: "localhost",
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/todo/all", http.StatusSeeOther)
		}
	}
}
func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			templates = template.New("")
			_, err := templates.ParseFiles("pages/register.html")
			if err != nil {
				log.Fatalf("Error parsing register template: %v", err)
			}
			_, err = templates.ParseGlob("templates/*.html")
			if err != nil {
				log.Fatalf("Error parsing templates: %v", err)
			}
			err = templates.ExecuteTemplate(w, "register.html", nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else if r.Method == http.MethodPost {
			name := r.FormValue("name")
			email := r.FormValue("email")
			password := r.FormValue("password")
			user := model.User{
				Name:     name,
				Email:    email,
				Password: password,
			}
			// fmt.Printf("%+v\n", user)
			serviceUser := service.ServiceUser{Db: db}
			err := serviceUser.CreateUser(&user)
			if err != nil {
				badResponse := model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "Bad Request",
					Data:       nil,
				}
				json.NewEncoder(w).Encode(badResponse)
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}
func GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serviceUser := service.ServiceUser{Db: db}
		users, err := serviceUser.FindAllUser()
		if err != nil {
			ise := model.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "Internal Server Error",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(ise)
			return
		}
		data := struct {
			Users []model.User
		}{
			Users: users,
		}
		templates = template.New("")
		_, err = templates.ParseFiles("pages/users.html")
		if err != nil {
			log.Fatalf("Error parsing users template: %v", err)
		}
		_, err = templates.ParseGlob("templates/*.html")
		if err != nil {
			log.Fatalf("Error parsing templates: %v", err)
		}
		err = templates.ExecuteTemplate(w, "users.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
