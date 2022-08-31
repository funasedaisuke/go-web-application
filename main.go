package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/funasedaisuke/go-web-application/config"
)

func main(){

	// if len(os.Args) !=2{
	// 	log.Printf("need port number\n")
	// 	os.Exit(1)


	// }
	// p := os.Args[1]

	//context.Background()は空のコンテキストを生成する。 
	//コンテキストを使う事で関数の外部からサーバーのプロセスを中断できる
	if err := run(context.Background());err != nil{
		log.Printf("failed  to terminate server:%v",err)
		os.Exit(1)
	}
}

func run(ctx context.Context)error{
	// ctx,stop := signal.NotifyContext(ctx,os.Interrupt,syscall.SIGTERM)
	// defer stop()
	cfg,err := config.New()
	if err != nil{
		return err
	}
	l,err:=net.Listen("tcp",fmt.Sprintf(":%d",cfg.Port))
	if err != nil{
		log.Fatalf("failed to listen port %d:%v",cfg.Port,err)
	}
	url :=fmt.Sprintf("http://%s",l.Addr().String())
	log.Printf("start with %s",url)
	mux := NewMux()
	s := NewServer(l,mux)
	return s.Run(ctx)

	// s := &http.Server{
	// 	Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	// 	time.Sleep(5 * time.Second)
	// 	fmt.Fprintf(w,"hello!:%s",r.URL.Path[1:])
	// }),
	// }
	// eg,ctx := errgroup.WithContext(ctx)
	// //別ゴールーチンでサーバーを立ち上げる
	// //errgroupは戻り値をエラーで返すことができる
	// //serverを起動しつつ、contextの終了通知を待機する
	// eg.Go(func() error{
	// 	if err := s.Serve(l); err != nil && err != http.ErrServerClosed{
	// 	log.Printf("failed to close server:%v",err)
	// 	return err
	// 	}
	// 	return nil
	// })
	// //キャンセル処理が来たら、shutdownする
	// <-ctx.Done()
	// if err := s.Shutdown(context.Background());err !=nil{
	// 	log.Printf("failed  to shutdown:%v",err)	
	// }

	// return eg.Wait()
}