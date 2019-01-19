package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

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

type UserResponse struct {
	Status int
	Result User
}

func FindAll(w http.ResponseWriter, r *http.Request) {
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

func FindByID(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
	ID := strings.Replace(r.URL.Path, "/user/", "", 1)
	fmt.Printf("ID: %v\n", ID)
	db := db.Db()
	defer db.Close()
	var u User
	selectQuery := `select * from users where id = ?`
	row := db.QueryRow(selectQuery, ID)
	row.Scan(&u.ID, &u.Name)
	fmt.Printf("n: %v\n", u)
	// TODO: 失敗しているときは 200 以外を返したいです。
	res, _ := json.Marshal(UserResponse{http.StatusOK, u})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
