package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/mochizukikotaro/go-repository-pattern/handler"
)

func main() {
	http.HandleFunc("/users", handler.FindAll)
	http.HandleFunc("/user/", handler.FindByID)
	http.HandleFunc("/notes", handler.NoteFindAll)
	http.ListenAndServe(":8080", nil)
}
