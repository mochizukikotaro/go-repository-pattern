package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/mochizukikotaro/go-repository-pattern/db"
)

type Ping struct {
	Status int
	Result string
}

type User struct {
	ID   int
	Name string
}

type UsersResponse struct {
	Status int
	Result []User
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ping := Ping{http.StatusOK, "hello repository pattern"}
	res, _ := json.Marshal(ping)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	db := db.Db()
	defer db.Close()
	q := `select * from users`
	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	users := []User{}
	u := User{}
	for rows.Next() {
		_ = rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}
	res, _ := json.Marshal(UsersResponse{http.StatusOK, users})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":8080", nil)
}
