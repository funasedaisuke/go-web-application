package main

import "net/http"


func NewMux() http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/health",func(w http.ResponseWriter,r *http.Request){
		w.Header().Set("Content-Type","aaplication/json; charset=utf-8")
		_,_=w.Write([]byte(`{"STATUS":"OK"}`))
	})
	return mux
}