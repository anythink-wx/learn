package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {

	group, ctx := errgroup.WithContext(context.Background())
	//handle 关闭通道
	closeCh := make(chan bool)

	//linux signal 信号的注册
	c := make(chan os.Signal, 0)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.HandleFunc("/", sayhelloName)
	mux.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "close")
		closeCh <- true
	})
	service := http.Server{Addr: ":8000", Handler: mux}

	group.Go(func() error {
		return service.ListenAndServe()
	})

	//接收到任何通道就 返回错误
	group.Go(func() error {
		select {
		case <-closeCh:
			return errors.New("server is going down: close by handle")
		case s := <-c:
			return errors.New("server is going down:" + s.String())
		}
	})

	//errgroup 接收到错误会返回 done 消息， 此处关闭服务器
	go func() {
		select {
		case <-ctx.Done():
			timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			if err := service.Shutdown(timeoutCtx); err != nil {
				fmt.Println(err)
			}
		}
	}()

	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}
