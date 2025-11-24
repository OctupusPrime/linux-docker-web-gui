package main

import (
	testHandler "linux-docker-web-gui/internal/test/handler"
	testRepo "linux-docker-web-gui/internal/test/repository"
	testService "linux-docker-web-gui/internal/test/service"
	webAppHandler "linux-docker-web-gui/internal/web-app/handler"
	"linux-docker-web-gui/pkg/db"
	"log"
	"net/http"
	"time"
)

func main() {
	database, err := db.New("/Users/mihailsokil/Desktop/my-projects/linux-docker-web-gui/temp/database.db")
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
	tHandler.RegisterRoutes(mux)

	webAppHandler := webAppHandler.NewHandler("/Users/mihailsokil/Desktop/my-projects/linux-docker-web-gui/frontend/dist", "index.html")
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
