package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	stopSignal := make(chan struct{})

	mux := http.NewServeMux()
	//http server 启动
	mux.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	//http server 关闭
	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		stopSignal <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	//part1
	//若 server.ListenAndServe() 出错，会直接触发 g.cancel()，part2 收到信号后，
	//停止阻塞，调用server.Shutdown结束server，part3亦如是，part1 part2 part3全注销
	g.Go(func() error {
		return server.ListenAndServe()
	})

	//part2
	//收到stopSignal信号后之后，调用server.Shutdown，part1注销，part3 ctx.Done 收到信号停止阻塞退出，part3注销，part1 part2 part3全注销
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup canceled")
		case <-stopSignal:
			log.Println("the server will be closed...")
		}
		log.Println("close the server...")
		return server.Shutdown(ctx)
	})

	//part3 注册 linux signal
	//收到quit之后，停止阻塞退出，part3注销，part2 ctx.Done 收到信号停止，调用server.Shutdown，part1 part2 注销
	g.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.New(fmt.Sprintf("quit:", sig))
		}
	})

	if err := g.Wait(); err != nil {
		log.Println(err)
	}
}
