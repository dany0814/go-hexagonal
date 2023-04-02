package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	web "github.com/dany0814/go-hexagonal/internal/adapters/driver"
	"github.com/dany0814/go-hexagonal/internal/core/application"
	"github.com/labstack/echo/v4"
)

type AppService struct {
	UserService application.UserService
}

type Server struct {
	engine          *echo.Echo
	httpAddr        string
	ShutdownTimeout time.Duration
	app             AppService
}

func NewServer(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, app AppService) (context.Context, Server) {
	srv := Server{
		engine:          echo.New(),
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		ShutdownTimeout: shutdownTimeout,
		app:             app,
	}
	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)
	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}

func (s *Server) registerRoutes() {
	// User Routes
	uh := web.NewUserHandler(s.app.UserService)
	s.engine.POST("/user/sigin", func(c echo.Context) error {
		uh.Ctx = c
		return uh.SignInHandler()
	})
}
