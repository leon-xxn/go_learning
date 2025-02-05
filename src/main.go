package src

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// 模拟慢请求
func sleep(ctx *gin.Context) {
	t := ctx.Query("t")
	s, err := strconv.Atoi(t)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误: " + t})
		return
	}

	time.Sleep(time.Duration(s) * time.Second)
	ctx.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("sleep %d s", s)})
}

const (
	stateHealth   = "health"
	stateUnHealth = "unhealth"
)

var state = stateHealth

func health(ctx *gin.Context) {
	status := http.StatusOK
	if state == stateUnHealth {
		status = http.StatusServiceUnavailable
	}
	ctx.JSON(status, gin.H{"data": state})
}

func main() {
	e := gin.Default()
	e.GET("/health", health)
	e.GET("/sleep", sleep)
	server := &http.Server{
		Addr:    ":8080",
		Handler: e,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("server run err :%+v", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Kill)
	<-quit
	log.Println("shutting down server")

	state = stateUnHealth
	log.Println("shutting down state:", state)

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("server foced to shutdown ", err)
	}
}
