/*
 * Copyright (c) 2021.
 * Project:play_server$
 * File:user.go
 * Author:anythink
 * Version:1.0
 * LastUpdate:2021/6/1 下午5:57
 */

package model

import (
	"fmt"
	"github.com/gogf/gf/crypto/gcrc32"
	"github.com/gogf/gf/frame/g"
	"gorm.io/gorm"
)

type Player struct {
	Id       uint64 `gorm:"column:id"`
	Username string
	Verify   int
	Mobile   string
}

func T(prefix string, str interface{}) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		encrypt := gcrc32.Encrypt(str)
		index := uint32(0)
		if g.Cfg().GetBool("mysql.account.split", true) {
			index = encrypt % 50
		}
		name := fmt.Sprintf("%s_%d", prefix, index)
		return tx.Table(name)
	}
}
