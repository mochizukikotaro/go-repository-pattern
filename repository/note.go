package repository

import (
	"database/sql"
	"fmt"

	"github.com/mochizukikotaro/go-repository-pattern/model"

)

type NoteRepository interface {
	FindAll() []model.Note
}

type noteRepository struct {
	DB *sql.DB
}

func NewNoteRepository(db *sql.DB) NoteRepository {
	return &noteRepository{db}
}

func (noteRepository *noteRepository) FindAll() []model.Note {
	db := noteRepository.DB
	defer db.Close()
	q := `select * from notes`
	rows, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	notes := []model.Note{}
	u := model.Note{}
	for rows.Next() {
		_ = rows.Scan(&u.ID, &u.Content, &u.UserID)
		notes = append(notes, u)
	}
	return notes
}