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

type Server struct {
	RouterConfig routers.RouterConfig
	DB           *database.DB
}

func New(db *database.DB) *Server {
	mode := gin.ReleaseMode
	if len(os.Getenv("DEV")) > 0 {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
	gin.ForceConsoleColor()

	return &Server{DB: db}
}

func (s *Server) Start() {
	srv := &http.Server{
		Addr:      "127.0.0.1:8000", // TODO: Add the address config
		Handler:   routers.Routers(container.NewContainer(s.DB), s.RouterConfig),
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
