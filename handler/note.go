package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/mochizukikotaro/go-repository-pattern/db"
	"github.com/mochizukikotaro/go-repository-pattern/model"
	"github.com/mochizukikotaro/go-repository-pattern/repository"
)

type NotesResponse struct {
	Status int
	Result []model.Note
}

func NoteFindAll(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
	noteRepository := repository.NewNoteRepository(db.Db())
	notes := noteRepository.FindAll()
	res, _ := json.Marshal(NotesResponse{http.StatusOK, notes})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
