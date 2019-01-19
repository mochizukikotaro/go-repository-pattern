package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/mochizukikotaro/go-repository-pattern/db"
	"github.com/mochizukikotaro/go-repository-pattern/repository"
	"github.com/mochizukikotaro/go-repository-pattern/model"
)

type Ping struct {
	Status int
	Result string
}

type UsersResponse struct {
	Status int
	Result []model.User
}

type UserResponse struct {
	Status int
	Result model.User
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
	userRepository := repository.NewUserRepository(db.Db())
	users := userRepository.FindAll()
	res, _ := json.Marshal(UsersResponse{http.StatusOK, users})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func FindByID(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
	ID := strings.Replace(r.URL.Path, "/user/", "", 1)
	userRepository := repository.NewUserRepository(db.Db())
	u := userRepository.FindByID(ID)

	// TODO: 失敗しているときは 200 以外を返したいです。
	res, _ := json.Marshal(UserResponse{http.StatusOK, u})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
