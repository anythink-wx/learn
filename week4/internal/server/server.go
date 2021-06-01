/*
 * Copyright (c) 2021.
 * Project:play_server$
 * File:server.go
 * Author:anythink
 * Version:1.0
 * LastUpdate:2021/6/1 下午5:32
 */

package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"log"
	"net/http"
	"os"
	"os/signal"
	"play_server/internal/data"
	"play_server/internal/service"
	"syscall"
	"time"
)

type Service struct {
	engine *gin.Engine
}

func (s *Service) Run() {

	s.engine = gin.New()
	s.engine.Use(gin.Recovery())
	s.engine.Use(gin.Logger())
	//engine.Use(middleware.Timer())
	//engine.Use(middleware.ApiLog())
	data.Db.Init()

	s.Router()
	srv := &http.Server{
		Addr:    ":" + g.Cfg().GetString("server.port", "8180"),
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func (s *Service) Router() {

	//路由组: v1
	v1 := s.engine.Group("/v1")
	{
		userGroup := v1.Group("/user")
		{
			userGroup.POST("/mobileLogin", service.User.MobileLogin)
		}
	}

}
