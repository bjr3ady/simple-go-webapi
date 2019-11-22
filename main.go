package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"runtime"

	logger "github.com/bjr3ady/go-logger"
	models "git.r3ady.com/golang/school-board/models/orm"
	"git.r3ady.com/golang/school-board/pkg/setting"
	"git.r3ady.com/golang/school-board/router"
	
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	logger.DefaultPrefix = setting.RunMode
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duation for which the server graceful wait for existing connections to finish -e.g. 15s or 1m")
	flag.Parse()

	err := models.ConnectDb(setting.Cfg)
	for err != nil {
		logger.Error("Connect to database failed, wait 3 seconds to re-connect... Details:", err)
		time.Sleep(3 * time.Second)
		err = models.ConnectDb(setting.Cfg)
	}
	logger.Info("Database connect success")

	r := router.InitRouter()

	srv := &http.Server{
		Handler:        r,
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		WriteTimeout:   setting.WriteTimeout,
		ReadTimeout:    setting.ReadTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(err)
			os.Exit(1)
			return
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	logger.Info("RESTful Web API service start listening request from:", setting.HTTPPort)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	logger.Fatal("shutting down...")
	os.Exit(0)
}