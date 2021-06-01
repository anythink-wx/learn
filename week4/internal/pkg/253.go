/*
 * Copyright (c) 2021.
 * Project:play_server$
 * File:253.go
 * Author:anythink
 * Version:1.0
 * LastUpdate:2021/6/1 下午5:32
 */

package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
	"net/url"
	"time"
)

type sdk235login struct {
	// 应用APPID
	appId string
	// 应用APPKEY
	appKey string
	apiUrl string
	client *ghttp.Client
}

func NewSdk235Login() *sdk235login {

	return &sdk235login{
		appId:  g.Cfg().GetString("login253.appid"),
		appKey: g.Cfg().GetString("login253.appKey"),
		apiUrl: g.Cfg().GetString("login253.apiUrl", "https://api.253.com/open/flashsdk/mobile-query"),
		client: ghttp.NewClient().SetTimeout(time.Second*4).SetRetry(1, time.Millisecond*500),
	}
}

// PKCS5Unpacking 一键登录，请求获取手机号接口demo
func (s *sdk235login) pKCS5Unpacking(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func (s *sdk235login) SignParam(data url.Values, key string) string {
	message := "appId" + data.Get("appId") + "token" + data.Get("token")
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	signature := hex.EncodeToString(mac.Sum(nil))
	return string(signature)
}

func (s *sdk235login) DecryptPhone(data string, key string) string {
	hash := md5.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])
	block, _ := aes.NewCipher([]byte(hashString[:16]))
	ecb := cipher.NewCBCDecrypter(block, []byte(hashString[16:]))
	source, _ := hex.DecodeString(data)
	decrypted := make([]byte, len(source))
	ecb.CryptBlocks(decrypted, source)
	return string(s.pKCS5Unpacking(decrypted))
}

type apiResult struct {
	Code         string `json:"code"`
	ChargeStatus int    `json:"chargeStatus"`
	Message      string `json:"message"`
	Data         struct {
		TradeNo    string `json:"tradeNo"`
		MobileName string `json:"mobileName"`
		Fanqizha   int    `json:"fanqizha"`
		Tag        string `json:"tag"`
	} `json:"data"`
}

func (s *sdk235login) GetMobile(token string) (mobile string, errs error) {
	formData := url.Values{}
	formData.Set("appId", s.appId)
	formData.Set("token", token)
	formData.Set("sign", s.SignParam(formData, s.appKey))

	resp, err := s.client.PostForm(s.apiUrl, formData)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var res apiResult
	if err := json.Unmarshal(body, &res); err != nil {
		return
	}
	fmt.Printf("置换手机号响应: %v \n ", res)

	if res.Code == "200000" {
		mobile = s.DecryptPhone(res.Data.MobileName, s.appKey)
		fmt.Printf("手机号: %s \n", mobile)
		return
	}
	return "", errors.New("get mobile error " + res.Message)
}
