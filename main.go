package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lmp/logger"
	"lmp/models"
	"lmp/routes"
	"lmp/settings"

	"github.com/facebookgo/pidfile"
	"go.uber.org/zap"
)

func main() {
	fmt.Println(models.Logo)

	pidfile.SetPidfilePath(os.Args[0] + ".pid")
	pidfile.Write()

	if err := settings.Init(); err != nil {
		fmt.Println("Init settings failed, err:", err)
		return
	}

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.AppConfig.Mode); err != nil {
		fmt.Println("Init logger failed, err:", err)
		return
	}
	defer zap.L().Sync()

	/*
		if err := influxdb.Init(settings.Conf.InfluxdbConfig); err != nil {
			fmt.Println("Init influxdb failed, err:", err)
			return
		}
	*/

	bpfscan := &models.BpfScan{}
	if err := bpfscan.Init(); err != nil {
		fmt.Println("Init bpfscan failed, err:", err)
	}
	bpfscan.Run()
	bpfscan.Watch()

	r := routes.SetupRouter(settings.Conf.AppConfig.Mode)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.AppConfig.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen failed :", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("shutdown server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
