package service

import (
	"context"

	"github.com/funasedaisuke/go-web-application/entity"
	"github.com/funasedaisuke/go-web-application/store"
)

type TaskAdder  interface{
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}


type TaskLister  interface{
	ListTasks(ctx context.Context, db store.Queryer)(entity.Tasks, error)
}

type UserRegister interface {
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}