package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"strconv"

	"smart-cart/internal/cart"
	"smart-cart/internal/db"
	"smart-cart/internal/router"

	"github.com/joho/godotenv"
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

	server := &http.Server{
		Addr:      ":" + port,
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	if err := server.ListenAndServeTLS("certs/server-dev.crt", "certs/server-dev.key"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
