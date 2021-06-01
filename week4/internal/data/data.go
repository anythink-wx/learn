/*
 * Copyright (c) 2021.
 * Project:play_server$
 * File:data.go
 * Author:anythink
 * Version:1.0
 * LastUpdate:2021/6/1 下午5:32
 */

package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"time"
)

type DataSource struct {
	Db    *gorm.DB
	Redis *redis.Client
}

var Db *DataSource

func (d *DataSource) Init() {
	Db = &DataSource{Db: NewDb("account"), Redis: NewRedis("account")}
}

func NewDb(dataSource string) *gorm.DB {

	dsn := g.Cfg().GetString("mysql." + dataSource + ".sources")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n[db] ", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	})
	if err != nil {
		panic("db connection error:" + err.Error())
	}

	replicase := g.Cfg().GetArray("mysql." + dataSource + ".replicas")
	var dialector []gorm.Dialector

	if replicase != nil {
		for _, dnsName := range replicase {
			dialector = append(dialector, mysql.Open(dnsName.(string)))
		}
	}

	config := dbresolver.Config{
		Replicas: dialector,
		Policy:   dbresolver.RandomPolicy{},
	}

	maxIdleConns := g.Cfg().GetInt("mysql."+dataSource+".maxIdleConns", 100)
	maxOpenConns := g.Cfg().GetInt("mysql."+dataSource+".maxOpenConns", 200)
	register := dbresolver.Register(config).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(maxIdleConns).
		SetMaxOpenConns(maxOpenConns)

	err = db.Use(register)
	if err != nil {
		fmt.Println("db connection dbresolver Register error:", err)
		os.Exit(0)
	}

	return db

}

//NewRedis("redis.account")

func NewRedis(dataSource string) *redis.Client {

	option := &redis.Options{
		Addr:         g.Cfg().GetString("redis."+dataSource+".host", "127.0.0.1"),
		Username:     g.Cfg().GetString("redis."+dataSource+".username", ""),
		Password:     g.Cfg().GetString("redis."+dataSource+".password", ""), // no password set
		DB:           g.Cfg().GetInt("redis."+dataSource+".db", 0),           // use default DB
		ReadTimeout:  time.Duration(g.Cfg().GetInt("redis."+dataSource+".readTimeout", 300)) * time.Millisecond,
		WriteTimeout: time.Duration(g.Cfg().GetInt("redis."+dataSource+".writeTimeout", 300)) * time.Millisecond,
		//// Maximum number of socket connections.
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		//PoolSize: 80,
	}
	rdb := redis.NewClient(option)

	log.Println("redis data source:", dataSource)
	log.Println("redis addr:", option.Addr)
	log.Println("redis db:", option.DB)
	log.Println("redis WriteTimeout:", option.WriteTimeout)
	log.Println("redis ReadTimeout:", option.ReadTimeout)

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(err.Error())
	}

	return rdb
}
