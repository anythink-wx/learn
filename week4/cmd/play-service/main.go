/*
 * Copyright (c) 2021.
 * Project:play_server$
 * File:main.go
 * Author:anythink
 * Version:1.0
 * LastUpdate:2021/6/1 下午5:32
 */

package main

import "play_server/internal/server"

func main() {

	serv := server.Service{}
	serv.Run()
}
