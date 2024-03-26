package main

import (
	"fmt"
	"net/http"
	"tinyk/api/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.Router(mux)

	err := http.ListenAndServe("localhost:4000", mux)

	fmt.Println(err)
}
