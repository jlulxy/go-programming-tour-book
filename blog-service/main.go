package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jlulxy/go-programming-tour-book/blog-service/pkg/logger"

	"github.com/jlulxy/go-programming-tour-book/blog-service/internal/model"

	"github.com/jlulxy/go-programming-tour-book/blog-service/global"

	"github.com/jlulxy/go-programming-tour-book/blog-service/pkg/setting"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/jlulxy/go-programming-tour-book/blog-service/internal/routers"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSeting err :%v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBengine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)

	}
}

// @title 博客系统
// @version 1.0
// @description  学习go和gin的博客系统
// @termsOfService https://github.com/jlulxy/go-programming-tour-book/tree/master/blog-service
func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("DataBase", &global.DataBaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DataBaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
