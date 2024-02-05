package main

import (
	"baseapi/api"
	"baseapi/conf"
	"baseapi/handler"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/viper"

	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"

	"github.com/labstack/echo/v4"
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
	runServer()
}

func runServer() {
	db := struct{}{}
	handler := handler.NewHandler(logger, db)

	e := echo.New()

	middleWareL := []strictecho.StrictEchoMiddlewareFunc{}

	api.RegisterHandlers(e, api.NewStrictHandler(handler, middleWareL))

	readConf := conf.NewAppConfiguratoin()
	e.Logger.Fatal(e.Start(fmt.Sprint(readConf.Server)))
}
