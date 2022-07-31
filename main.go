package main

import (
	"k8s-httpgo/metrics"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	slog := setupLogger()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		hostname, err := os.Hostname()
		if err != nil {
			slog.Panic(err)
			c.JSON(http.StatusInternalServerError, nil)
		}
		timestamp := time.Now()
		metrics.IndexerSuccessfulWrite.With(prometheus.Labels{"hostname": hostname, "timestamp": strconv.Itoa(int(time.Now().Unix()))}).Inc()
		c.JSON(http.StatusOK, gin.H{
			"hostname":  hostname,
			"timestamp": timestamp,
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "OK",
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	err := r.Run()
	if err != nil {
		return
	}
}

func setupLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	slog := logger.Sugar()
	slog.Info("Logger init")
	return slog
}
