/*
 * Copyright (c) 2021.
 * Project:play_server$
 * File:user.go
 * Author:anythink
 * Version:1.0
 * LastUpdate:2021/6/1 下午5:32
 */

package biz

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"play_server/internal/repo"
)

type UserBiz struct {
	Log  *glog.Logger
	Repo *repo.UserRepo
}

func NewUserBiz() *UserBiz {

	return &UserBiz{
		Log:  g.Log("userBiz"),
		Repo: repo.NewUserRepo(),
	}
}
