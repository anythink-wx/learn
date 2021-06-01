/*
 * Copyright (c) 2021.
 * Project:play_server$
 * File:user.pb_test.go
 * Author:anythink
 * Version:1.0
 * LastUpdate:2021/6/1 下午5:25
 */

package v1

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net/http"
	"testing"
)

func TestUserMobileLoginRequest_Descriptor(t *testing.T) {

	request := UserMobileLoginRequest{
		Token: "123",
	}

	marshal, err := proto.Marshal(&request)

	fmt.Println("err", err)
	reader := bytes.NewReader(marshal)
	http.Post("http://127.0.0.1:8180/v1/user/mobileLogin", "application/x-protobuf", reader)
}
