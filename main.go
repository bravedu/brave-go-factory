package main

import (
	"context"
	"fmt"
	"github.com/bravedu/brave-go-factory/config"
	"github.com/bravedu/brave-go-factory/pkg/xlog"
	api "github.com/bravedu/brave-go-factory/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env != "dev" {
		//cnf := config.ConfInstance(env)
	} else {
		config.ConfInstanceDev(env)
	}
	fmt.Println(config.Conf)
	defer config.Conf.Close()
	r := api.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Conf.YamlDao.ProCnf.ProjectPort),
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			xlog.Errorf("Listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		xlog.Errorf("Server Shutdown:", err)
	}
}
