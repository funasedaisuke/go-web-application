package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)


type ErrResponse struct {
	Message string `json:"message"`
	Details []string `json;"details,omitempty"` 
}

func RespondJSON(ctx context.Context, w http.ResponseWriter, body any ,status int){
	w.Header().Set("content-Type","application/json; charset -utf-8")
	// Marshal 元帥 json encodeをかえす

	bodyBytes,err:=json.Marshal(body)
	if err != nil{
	rsp := ErrResponse{
		Message: http.StatudText(http.StatusInternslServerError),
	}
	if err := json.NewEncoder(w).Encode(rsp); err != nil{
		fmt.Printf("write error response error: %v",err)
	}
	return
}
w.WriteHeader(status)
if _,err := fmt.Fprintf(w,"%s",bodyBytes); err != nil {
	fmt.Printf("write response error: %v",err)
}
}