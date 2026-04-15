package routes

import (
	"net/http"

	"user-crud/controllers"
)

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			controllers.GetUsers(w, r)
		case "POST":
			controllers.CreateUser(w, r)
		case "PUT":
			controllers.UpdateUser(w, r)
		case "DELETE":
			controllers.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", 405)
		}
	})

	return mux
}
