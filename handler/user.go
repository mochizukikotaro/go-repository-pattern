package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mochizukikotaro/go-repository-pattern/db"
	"github.com/mochizukikotaro/go-repository-pattern/model"
	"github.com/mochizukikotaro/go-repository-pattern/repository"
)

type UsersResponse struct {
	Status int
	Result []model.User
}

type UserResponse struct {
	Status int
	Result model.User
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	userRepository := repository.NewUserRepository(db.Db())
	users := userRepository.FindAll()
	res, _ := json.Marshal(UsersResponse{http.StatusOK, users})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func FindByID(w http.ResponseWriter, r *http.Request) {
	ID := strings.Replace(r.URL.Path, "/user/", "", 1)
	userRepository := repository.NewUserRepository(db.Db())
	u := userRepository.FindByID(ID)

	// TODO: 失敗しているときは 200 以外を返したいです。
	res, _ := json.Marshal(UserResponse{http.StatusOK, u})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
