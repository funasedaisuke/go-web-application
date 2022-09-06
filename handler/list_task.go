package handler

import (
	"net/http"

	"github.com/funasedaisuke/go-web-application/entity"
	"github.com/funasedaisuke/go-web-application/store"
)

type ListTask struct {
	Store *store.TaskStore
}

type task struct{
	ID entity.TaskID         `json:"id"`
	Title string             `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (tl *ListTask)ServeHTTP(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	tasks := tl.Store.All()
	rsp :=[]task{}
	for _ ,t := range tasks{
		rsp = append(rsp,task{
		ID: t.ID,
		Title: t.Title,
		Status: t.Status,
	})
	}
	RespondJSON(ctx,w,rsp,http.StatusOK)
}