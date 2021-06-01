/*
 * Copyright (c) 2021.
 * Project:play_server$
 * File:user.go
 * Author:anythink
 * Version:1.0
 * LastUpdate:2021/6/1 下午5:32
 */

package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"play_server/internal/biz"
)

type userService struct {
	log *glog.Logger
	biz *biz.UserBiz
}

var User = &userService{log: g.Log("user"), biz: biz.NewUserBiz()}

func (u *userService) MobileLogin(ctx *gin.Context) {
	u.log.Info("login hello")
}
