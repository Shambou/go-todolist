package main

import (
	"os"
	"strconv"

	"github.com/Shambou/todolist/internal/config"
	transportHttp "github.com/Shambou/todolist/internal/transport/http"
	log "github.com/sirupsen/logrus"
)

var app config.AppConfig

// var session *scs.SessionManager

func Run() error {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)

	app.InProduction = os.Getenv("ENVIRONMENT") == "production"

	useCache, err := strconv.ParseBool(os.Getenv("USE_CACHE"))
	if err != nil {
		useCache = false
	}

	app.UseCache = useCache

	httpHandler := transportHttp.New(&app)
	if err := httpHandler.Serve(); err != nil {
		log.Println("Failed to run http server: ", err)
		return err
	}

	return nil
}

func main() {
	log.Info("Starting todo api server")

	if err := Run(); err != nil {
		log.Println(err)
		log.Fatal("Error starting application")
	}
}
