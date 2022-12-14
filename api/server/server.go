package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/a3510377/control-panel/container"
	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/routers"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func New() *Server { return &Server{} }

func (s *Server) Start(db *database.DB) {
	gin.SetMode("release")
	gin.ForceConsoleColor()

	router := routers.Routers(container.NewContainer(db))

	srv := &http.Server{
		Addr:      "127.0.0.1:8000", // TODO: Add the address config
		Handler:   router,
		TLSConfig: nil, // TODO: Add the TLS config
	}

	go func() {
		log.Println("Server starting...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
