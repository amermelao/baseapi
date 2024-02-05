package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"truckvault/api"
	"truckvault/conf"
	"truckvault/handler"
	"truckvault/handler/persist/database"
	"truckvault/handler/utilsmiddleware"

	"github.com/spf13/viper"
	"golang.org/x/exp/slog"

	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var configFilePath string
var logger *slog.Logger

var makeTables bool

func init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	flag.StringVar(&configFilePath, "conf", "", "configuration file path")
	flag.BoolVar(&makeTables, "tables", false, "make tables in db")
	flag.Parse()

	viper.SetConfigType("toml")

	if configFilePath != "" {

		logger.Info(
			"reading configuration from file",
			slog.String("path", configFilePath),
		)

		confFile, err := os.Open(configFilePath)
		if err != nil { // Handle errors reading the config file
			logger.Error(fmt.Sprintf("fatal error config file: %s", err))
			os.Exit(1)
		}
		defer confFile.Close()

		err = viper.ReadConfig(confFile)
		if err != nil { // Handle errors reading the config file
			logger.Error(fmt.Sprintf("fatal error config file: %s", err))
			os.Exit(1)
		}

	} else {
		logger.Info("looking for config.toml file")

		viper.SetConfigName("config.toml") // name of config file (without extension)
		viper.AddConfigPath("./etc/")      // path to look for the config file in
		viper.AddConfigPath(".")           // optionally look for config in the working directory
		err := viper.ReadInConfig()        // Find and read the config file
		if err != nil {                    // Handle errors reading the config file
			logger.Error(fmt.Sprintf("fatal error config file: %s", err))
			os.Exit(1)
		}
	}
}

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config api/config.yaml api/definition.yaml
func main() {
	if makeTables {
		makeDBTables()
	} else {
		runServer()
	}
}

func makeDBTables() {
	readConf := conf.NewAppConfiguratoin().DB
	logger.Info("making db tables", "conf", readConf)
	connectionPersist, err := database.Connection(readConf)
	if err != nil {
		logger.Error("bad connection", slog.String("error", err.Error()))
		return
	}

	if err := connectionPersist.MakeTables(); err != nil {
		logger.Error("fail to generate tables: %w", err)
	}
}

func runServer() {
	readConf := conf.NewAppConfiguratoin()
	logger.Info("starting", "conf", readConf)
	connectionPersist, err := database.Connection(readConf.DB)
	if err != nil {
		logger.Error("bad connection", slog.String("error", err.Error()))
		return
	}
	handler := handler.NewHandler(logger, connectionPersist)

	e := echo.New()

	middleWareL := []strictecho.StrictEchoMiddlewareFunc{}

	api.RegisterHandlers(e, api.NewStrictHandler(handler, middleWareL))

	e.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: readConf.Server.AllowOrigins,
				AllowMethods: []string{http.MethodGet, http.MethodOptions},
			},
		),
		middleware.Logger(),
		middleware.Recover(),
	)

	if readConf.ValidJWT.Enabled() {
		e.Use(utilsmiddleware.Authorization(logger, readConf.ValidJWT))
	} else {
		logger.Warn("auth is disabled")
	}

	e.Logger.Fatal(e.Start(fmt.Sprint(readConf.Server)))
}
