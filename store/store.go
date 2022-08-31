package store

import (
	"errors"

	"github.com/funasedaisuke/go-web-application/entity"
)

var (
	Tasks = &TaskStore{Tasks: map[int]*entity.Task{}}
	ErrorNotFound = errors.New("not found")
)

type TaskStore struct{
	LastID entity.TaskID
	Tasks map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task)(int,error){
	ts.LastID++
	t.ID =ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID,nil
}

func  (ts *TaskStore) All()entity.Tasks{
	tasks := make([]*entity.Task,len(ts.Tasks))
	for i,t *= range ts.Tatasks{
		tasks[i-1]=t
	}
	return tasks
}