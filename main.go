package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main(){
	//context.Background()は空のコンテキストを生成する。 
	//コンテキストを使う事で関数の外部からサーバーのプロセスを中断できる
	if err := run(context.Background());err != nil{
		log.Printf("failed  to terminate server:%v",err)
	}
}

func run(ctx context.Context)error{
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"hello!:%s",r.URL.Path[1:])
	}),
	}
	eg,ctx := errgroup.WithContext(ctx)
	//別ゴールーチンでサーバーを立ち上げる
	//errgroupは戻り値をエラーで返すことができる
	//serverを起動しつつ、contextの終了通知を待機する
	eg.Go(func() error{
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed{
		log.Printf("failed to close server:%v",err)
		return err
		}
		return nil
	})
	//キャンセル処理が来たら、shutdownする
	<-ctx.Done()
	if err := s.Shutdown(context.Background());err !=nil{
		log.Printf("failed  to shutdown:%v",err)	
	}

	return eg.Wait()
}