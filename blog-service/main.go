package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jlulxy/go-programming-tour-book/blog-service/internal/model"

	"github.com/jlulxy/go-programming-tour-book/blog-service/global"

	"github.com/jlulxy/go-programming-tour-book/blog-service/pkg/setting"

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
}
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
