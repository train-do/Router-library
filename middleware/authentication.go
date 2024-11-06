package middleware

import (
	"log"
	"net/http"

	"github.com/train-do/Router-library/database"
	"github.com/train-do/Router-library/service"
)

func Authentication(next http.Handler) http.Handler {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("MASUK MIDDLEWARE")
		cookie, err := r.Cookie("access_token")
		if err != nil || err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		accessToken := cookie.Value
		serviceUser := service.ServiceUser{Db: db}
		if err = serviceUser.GetById(accessToken); err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
		defer db.Close()
	})
}
