/*
 * Copyright (c) 2021.
 * Project:play_server$
 * File:user.go
 * Author:anythink
 * Version:1.0
 * LastUpdate:2021/6/1 下午5:32
 */

package repo

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"play_server/internal/data"
)

type UserRepo struct {
	*data.DataSource
	log *glog.Logger
}

func NewUserRepo() *UserRepo {

	return &UserRepo{
		DataSource: data.Db,
		log:        g.Log("userRepo"),
	}
}
