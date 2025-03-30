package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dbiagi/shopping-bag/internal/config"
	"github.com/dbiagi/shopping-bag/internal/http/handler"
	"github.com/dbiagi/shopping-bag/internal/repository"
	"github.com/dbiagi/shopping-bag/internal/util"
	"github.com/gorilla/mux"
)

var (
	signalsToListenTo = []os.Signal{
		syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM,
	}
)

type (
	Server struct {
		config.Configuration
		*http.Server
		*mux.Router
	}

	handlers struct {
		handler.CartHandler
		handler.HealthCheckHandler
	}
)

func NewServer(c config.Configuration) *Server {
	return &Server{
		Configuration: c,
	}
}

func (s *Server) Start() {
	server, router := createServer(s.Configuration.WebConfig)

	handlers := createHandlers(s.Configuration)

	registerRoutesAndMiddlewares(router, handlers)
	configureGracefullShutdown(server, s.Configuration.WebConfig)
}

func (s *Server) ForceShutdown() {
	s.Server.Shutdown(context.Background())
}

func createServer(webConfig config.WebConfig) (*http.Server, *mux.Router) {
	router := mux.NewRouter()
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", webConfig.Port),
		Handler:      router,
		IdleTimeout:  webConfig.IdleTimeout,
		ReadTimeout:  webConfig.ReadTimeout,
		WriteTimeout: webConfig.WriteTimeout,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err.Error() != "http: Server closed" {
			slog.Error("Error starting server.", slog.String("error", err.Error()))
		}
	}()

	return srv, router
}

func registerRoutesAndMiddlewares(router *mux.Router, h handlers) {
	router.Use(util.TraceIdMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/health", h.HealthCheckHandler.Health).Methods("GET")
}

func configureGracefullShutdown(server *http.Server, webConfig config.WebConfig) {
	slog.Info("Configuring graceful shutdown.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, signalsToListenTo...)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), webConfig.ShutdownTimeout)
	defer cancel()

	server.Shutdown(ctx)
	slog.Info("Shutting down server")
	os.Exit(0)
}

func createHandlers(c config.Configuration) handlers {
	dynamodb, err := config.CreateDynamoDBConnection(c.AWSConfig)

	if err != nil {
		slog.Error(fmt.Sprintf("Error creating the dynamodb connection. Error %s.", *err))
		panic(err)
	}

	cartRepository := repository.NewCartRepository(dynamodb)

	return handlers{
		HealthCheckHandler: handler.NewHealthCheckHandler(),
		CartHandler:        handler.NewCartHandler(cartRepository),
	}
}
