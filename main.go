package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	err := http.ListenAndServe(":18081",
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"hello!:%s",r.URL.Path[1:])
	}),
	)
if err != nil{
	fmt.Printf("failed to terminate server: %v",err)
	os.Exit(1)
}

}
