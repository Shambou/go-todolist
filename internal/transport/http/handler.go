package http

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Shambou/todolist/internal/config"
	"github.com/Shambou/todolist/internal/repository"
	"github.com/Shambou/todolist/internal/repository/dbrepo"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Router    *mux.Router
	Server    *http.Server
	AppConfig *config.AppConfig
	DB        repository.DatabaseRepo
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// jsonResponse renders json response
func jsonResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	if err := json.NewEncoder(w).Encode(Response{
		Status:  status,
		Message: message,
		Data:    data,
	}); err != nil {
		panic(err)
	}
}

// New creates a new HTTP handler
func New(appConfig *config.AppConfig) *Handler {
	log.Info("Creating new http service")

	h := &Handler{
		AppConfig: appConfig,
		DB:        dbrepo.NewPostgresRepo(appConfig),
	}

	h.Router = mux.NewRouter()
	h.MapRoutes()

	h.Server = &http.Server{
		Addr:    ":8080",
		Handler: h.Router,
		// Good practice to set timeouts to avoid Slow loris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	return h
}

// Healthz returns the health of the service
func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	log.Info("API Health is OK")
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    "Server is running AND ready to accept requests",
	}); err != nil {
		panic(err)
	}
}

// Serve - gracefully serves our newly set up handler function
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := h.Server.Shutdown(ctx)
	if err != nil {
		return err
	}

	log.Println("shutting down gracefully")
	return nil
}
