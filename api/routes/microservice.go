package routes

import (
	"net/http"
	"tinyk/api/controllers"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/url", main)
}

func main(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		controllers.GetUrl(w, r)
	case "POST":
		controllers.CreateUrl(w, r)
	case "DELETE":
		controllers.DeleteUrl(w, r)
	default:
		http.Error(w, r.Method+" method not allowed", http.StatusMethodNotAllowed)
	}
}
