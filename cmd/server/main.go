package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"smart-cart/internal/cart"
	"smart-cart/internal/db"
	"smart-cart/internal/router"

	"smart-cart/internal/metrics"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	database, err := db.NewPostgres(dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := cart.NewRepository(database)
	service := cart.NewService(repo)
	handler := cart.NewHandler(service)
	r := router.SetupRouter(handler)
	r.Use(metrics.PrometheusMiddleware())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logJSON := map[string]interface{}{
			"time":     param.TimeStamp.Format(time.RFC3339),
			"method":   param.Method,
			"path":     param.Path,
			"status":   param.StatusCode,
			"latency":  param.Latency.Seconds(),
			"clientIP": param.ClientIP,
		}
		b, _ := json.Marshal(logJSON)
		return string(b) + "\n"
	}))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/fail", func(c *gin.Context) {
		c.JSON(500, gin.H{"status": "ok"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	portNum, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Invalid port")
	}

	log.Printf("Starting server on port %d...", portNum)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13, // Используем TLS 1.3
	}

	metrics.Init()

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           r,
		TLSConfig:         tlsConfig,
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServeTLS("certs/server-dev.crt", "certs/server-dev.key"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
