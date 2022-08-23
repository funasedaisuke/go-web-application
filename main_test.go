package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T){
	l,err :=net.Listen("tcp","localhost:0")
	if err != nil{
		t.Fatalf("failed to listen port %v",err)
	}
	ctx, cancel := context.WithCancel(context.Background())

	eg,ctx :=errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx,l)
	})
	in := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), in)
	rsp,err :=http.Get(url)
	if err != nil{
		t.Errorf("failed to get %v",err)
		os.Exit(1)
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll(rsp.Body)
	if err !=nil{
		t.Errorf("failed to read body: %v",err)
	}
	want := fmt.Sprintf("hello!:%s",in)
	if string(got) != want{
		t.Errorf("want %q, but got %q",want,got)
	}
	cancel()
	if err := eg.Wait();err != nil{
		t.Fatal(err)
	}

}