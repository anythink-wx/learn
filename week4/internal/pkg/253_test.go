/*
 * Copyright (c) 2021.
 * Project:play_server$
 * File:253_test.go
 * Author:anythink
 * Version:1.0
 * LastUpdate:2021/6/1 下午5:32
 */

package pkg

import (
	"fmt"
	"testing"
)

func TestNewSdk235Login(t *testing.T) {
	sdk235login := NewSdk235Login()
	mobile, err := sdk235login.GetMobile("1233")
	fmt.Println(mobile, err)
}
