package server

import (
	"context"
	golog "log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/henglory/Demo_ServiceStubby/config"
	"github.com/henglory/Demo_ServiceStubby/service"
)

type errorResponse struct {
	ResponseCode int64  `json:"responseCode"`
	Reason       string `json:"reason"`
	RawRequest   string `json:"rawRequest"`
}

type Server struct {
	srv *http.Server
	s   service.Service
}

func NewServer(s service.Service) *Server {
	server := &Server{
		s: s,
	}
	return server
}

func (server *Server) Start() {
	go server.ginStart()
}

func (server *Server) Close() {
	server.srv.Close()
}

type readiness struct {
	Success bool `json:"success"`
}

func (server Server) ginStart() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/api/actionA", func(c *gin.Context) {
		doA(server.s, c)
	})

	r.POST("/api/actionB", func(c *gin.Context) {
		doB(server.s, c)
	})

	r.POST("/api/actionC", func(c *gin.Context) {
		doC(server.s, c)
	})

	server.srv = &http.Server{
		Addr:    config.ServicePort,
		Handler: r,
	}

	if err := server.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		golog.Fatalf("listen: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := server.srv.Shutdown(ctx); err != nil {
		golog.Fatal("Server Shutdown:", err)
	}
	defer cancel()
}
