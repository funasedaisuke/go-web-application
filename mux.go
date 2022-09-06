package main

import (
	"net/http"

	"github.com/funasedaisuke/go-web-application/handler"
	"github.com/funasedaisuke/go-web-application/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux() http.Handler{
	mux := chi.NewRouter()
	mux.HandleFunc("/health",func(w http.ResponseWriter,r *http.Request){
		w.Header().Set("Content-Type","aplication/json; charset=utf-8")
		_,_=w.Write([]byte(`{"STATUS":"OK"}`))
	})
	v := validator.New()
	lt := &handler.AddTask{Store: store.Tasks,Validator: v}
	mux.Get("/tasks",lt.ServeHTTP)
	return mux
}