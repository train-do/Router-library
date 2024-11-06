package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/train-do/Router-library/model"
	"github.com/train-do/Router-library/service"
)

func GetTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("access_token")
		serviceTodo := service.ServiceTodo{Db: db}
		todos, err := serviceTodo.FindAllTodo(cookie.Value)
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
			Todos []model.Todo
		}{
			Todos: todos,
		}
		templates = template.New("")
		_, err = templates.ParseFiles("pages/todos.html")
		if err != nil {
			log.Fatalf("Error parsing todos template: %v", err)
		}
		_, err = templates.ParseGlob("templates/*.html")
		if err != nil {
			log.Fatalf("Error parsing templates: %v", err)
		}
		err = templates.ExecuteTemplate(w, "todos.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
func CreateTodo(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Println("MASUK CREATE")
			templates = template.New("")
			_, err := templates.ParseFiles("pages/addTodo.html")
			if err != nil {
				log.Fatalf("Error parsing addTodo template: %v", err)
			}
			_, err = templates.ParseGlob("templates/*.html")
			if err != nil {
				log.Fatalf("Error parsing templates: %v", err)
			}
			err = templates.ExecuteTemplate(w, "addTodo.html", nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else if r.Method == http.MethodPost {
			cookie, _ := r.Cookie("access_token")
			var isDone bool
			if r.FormValue("isDone") == "on" {
				isDone = true
			}
			// fmt.Printf("%v %v  %v\n", reflect.TypeOf(isDone), isDone, r.FormValue("isDone"))
			todo := model.Todo{
				UserId: cookie.Value,
				Todo:   r.FormValue("todo"),
				IsDone: isDone,
			}
			serviceTodo := service.ServiceTodo{Db: db}
			err := serviceTodo.InsertTodo(&todo)
			if err != nil {
				unauthorized := model.Response{
					StatusCode: http.StatusUnauthorized,
					Message:    "Bad Request Service",
					Data:       nil,
				}
				json.NewEncoder(w).Encode(unauthorized)
				return
			}
			http.Redirect(w, r, "/todo/all", http.StatusSeeOther)
			// templates = template.New("")
			// _, err = templates.ParseFiles("pages/todos.html")
			// if err != nil {
			// 	log.Fatalf("Error parsing todos template: %v", err)
			// }
			// _, err = templates.ParseGlob("templates/*.html")
			// if err != nil {
			// 	log.Fatalf("Error parsing templates: %v", err)
			// }
			// err = templates.ExecuteTemplate(w, "todos.html", nil)
			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }
			// w.Header().Set("Content-Type", "application/json")
			// w.WriteHeader(http.StatusCreated)
			// json.NewEncoder(w).Encode(todo)
		}
	}
}
