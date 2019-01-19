package repository

import (
	"database/sql"
	"fmt"

	"github.com/mochizukikotaro/go-repository-pattern/model"

)

type UserRepository interface {
	FindAll() []model.User
	FindByID(ID string) model.User
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) FindAll() []model.User {
	db := userRepository.DB
	defer db.Close()
	q := `select * from users`
	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	users := []model.User{}
	u := model.User{}
	for rows.Next() {
		_ = rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}
	return users
}

func (userRepository *userRepository) FindByID(ID string) model.User {
	db := userRepository.DB
	defer db.Close()
	var u model.User
	selectQuery := `select * from users where id = ?`
	row := db.QueryRow(selectQuery, ID)
	row.Scan(&u.ID, &u.Name)
	fmt.Printf("n: %v\n", u)
	return u
}
