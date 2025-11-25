package main

import (
	testHandler "linux-docker-web-gui/internal/test/handler"
	testRepo "linux-docker-web-gui/internal/test/repository"
	testService "linux-docker-web-gui/internal/test/service"
	webAppHandler "linux-docker-web-gui/internal/web-app/handler"
	"linux-docker-web-gui/pkg/db"
	"linux-docker-web-gui/pkg/middleware"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := db.New(os.Getenv("DATABASE_PATH"))
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	if err := database.Migrate(); err != nil {
		log.Fatal("Failed to migrate db:", err)
	}
	defer database.Close()

	mux := http.NewServeMux()

	tRepo := testRepo.NewSQLiteRepository(database.DB)
	tService := testService.NewService(tRepo)
	tHandler := testHandler.NewHandler(tService)
	tHandler.RegisterRoutes(mux, middleware.Logger)

	webAppHandler := webAppHandler.NewHandler(os.Getenv("FRONTEND_PATH"), "index.html")
	webAppHandler.RegisterRoutes(mux)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Println("Server starting on port 8080...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
