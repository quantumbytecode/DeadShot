package main

import (
	"database/sql"
	"deadshot/internal/app/handlers"
	"deadshot/internal/app/services"
	"deadshot/internal/config"
	"deadshot/internal/persistence"
	"net/http"
	"os"

	_ "modernc.org/sqlite"

	"github.com/sirupsen/logrus"
)

func main() {
	mux := http.NewServeMux()

	db, dberr := sql.Open("sqlite", "internal/persistence/data/deadshot.db")
	if dberr != nil {
		logrus.Error("Could not open deadshot.db")
		os.Exit(1)
	}

	dm := persistence.NewDBManager(db)

	dm.SetupDB()

	srv := services.NewSqlliteDeadshotService(dm)

	handler := handlers.NewDeadshotHandler(srv)

	mux.HandleFunc("/", handler.CaptureLog)
	mux.HandleFunc("/all", handler.GetAllLogs)
	mux.HandleFunc("/get", handler.GetLogByID)
	mux.HandleFunc("/update", handler.UpdateLog)
	mux.HandleFunc("/delete", handler.DeleteLog)
	mux.HandleFunc("/replay", handler.ReplayLog)

	appconfig, conferr := config.LoadConfigs()
	if conferr != nil {
		os.Exit(1)
	}

	server := http.Server{
		Addr:    appconfig.Host + ":" + appconfig.Port,
		Handler: mux,
	}

	logrus.Info("Server started on " + appconfig.Host + ":" + appconfig.Port)

	logrus.Info(server.ListenAndServe())
}
