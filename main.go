package main

import (
	"fmt"
	"github.com/degary/web-service/internal/model"
	"github.com/degary/web-service/internal/routers"
	"github.com/degary/web-service/pkg/logger"
	"github.com/degary/web-service/pkg/setting"
	"github.com/degary/web-service/global"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init()  {
	err := setupSetting()
	if err != nil{
		log.Fatalf("init.setupSetting err:%v\n",err)
	}
	err = setupDBEngine()
	if err != nil{
		log.Fatalf("init.setupDBEnine err: %v\n",err)
	}
	err = setupLogger()
	if err != nil{
		log.Fatalf("init.setupLogger err: %v\n",err)
	}
}
func setupSetting() error {
	setting,err := setting.NewSetting()
	if err  != nil {
		return err
	}
	err = setting.ReadSection("Server",&global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database",&global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App",&global.AppSeting)
	if err != nil {
		return err
	}
	global.ServerSetting.WriteTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	return nil
}

func setupDBEngine() error{
	var err error
	global.DBEngin,err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil{
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger= logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSeting.LogSavePath + "/" + global.AppSeting.LogFileName + global.AppSeting.LogFileExt,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	},"",log.LstdFlags).WithCaller(2)
	return nil
}

// @title 博客系统
// @version 1.0
// @description Go 编程之旅
// @TermsOfService https://degary.com
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":"+ global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}
	fmt.Printf("server,%#v\n",global.ServerSetting)
	global.Logger.Infof("%s: go-programming-tour-book/%s","degary","web-service")
	global.Logger.Infof("%s: go-programming-tour-book/%s","degary1","web-service1")
	s.ListenAndServe()

}
