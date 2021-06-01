/*
 * Copyright (c) 2021.
 * Project:play_server$
 * File:autoToken.go
 * Author:anythink
 * Version:1.0
 * LastUpdate:2021/6/1 下午5:32
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthTokenRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {

			c.String(http.StatusUnauthorized, "", "StatusUnauthorized")
			c.Abort()
			return
		}

		fields := strings.Fields(header)

		if len(fields) != 1 {
			c.String(http.StatusUnauthorized, "", "ErrAuthCheckLen")
			c.Abort()
			return
		}

		//userSession, logicErr := db.AccountRdb.CheckUserToken(fields[0])
		//if logicErr != nil {
		//	c.String(http.StatusUnauthorized,"","ErrAuthCheckLen")
		//	ctl.ErrWithException(c, logicErr)
		//	c.Abort()
		//	return
		//}
		//
		//c.Set("uid", userSession.Id)
		//c.Set("channel", userSession.Channel)
		//c.Set("clientid", fields[0])
		//c.Set("userSession", userSession)

		c.Next()

	}
}
