package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
	"github.com/tirzasrwn/fishing/internal/config"
	"github.com/tirzasrwn/fishing/internal/driver"
	"github.com/tirzasrwn/fishing/internal/handler"
	"github.com/tirzasrwn/fishing/internal/helpers"
	"github.com/tirzasrwn/fishing/internal/models"
	"github.com/tirzasrwn/fishing/internal/render"
)

var (
	app      config.AppConfig
	infoLog  *log.Logger
	errorLog *log.Logger
)

func main() {
	fmt.Println("Hello, world!")
	gob.Register(models.Account{})

	err := initializeAppConfig()
	if err != nil {
		fmt.Println("failed to load config")
		return
	}

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// Connect to database.
	log.Println("Connecting to database ...")
	db, err := driver.ConnectSQL(fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		app.DBHost, app.DBPort, app.DBName, app.DBUser, app.DBPassword))
	if err != nil {
		fmt.Println("Cannot connect to the database!")
		return
	}
	log.Println("Connected to database.")
	defer db.SQL.Close()

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Cannot create template cache!", err)
		return
	}

	app.TemplateCache = tc
	repo := handler.NewRepo(&app, db)
	handler.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", app.Port),
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func initializeAppConfig() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.AllowEmptyEnv(false)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	app.Port = viper.GetInt("PORT")
	app.DBUser = viper.GetString("DB_USER")
	app.DBName = viper.GetString("DB_NAME")
	app.DBPassword = viper.GetString("DB_PASSWORD")
	app.DBHost = viper.GetString("DB_HOST")
	app.DBPort = viper.GetInt("DB_PORT")

	log.Println("[INIT] configuration loaded")

	return nil
}
