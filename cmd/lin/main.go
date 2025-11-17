package main

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"pvik/linear-cli/internal/config"
)

func init() {
	app_config_path := filepath.Join(xdg.ConfigHome, "linear-cli")

	// check if application XDG dir exists
	_, err := os.Stat(app_config_path)
	if err != nil {
		// dir doesn't exist create
		err_create := os.Mkdir(app_config_path, 0755)
		if err_create != nil {
			log.Fatal().Err(err_create).Msg("Unable to create application config directory")
		}
	}

	// logger settings
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	logLevel := zerolog.InfoLevel
	logLevelEnvStr := os.Getenv("LOG_LEVEL")
	if logLevelEnvStr != "" {
		logLevelEnv, err := zerolog.ParseLevel(logLevelEnvStr)
		if err == nil {
			logLevel = logLevelEnv
		}
	}
	zerolog.SetGlobalLevel(logLevel)
}

func main() {

	log.Debug().
		Msg("Starting linear-cli...")

	configFilePath, err := xdg.ConfigFile("linear-cli/config.toml")
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Any("using config file at:", configFilePath).Msg("b")

	appConfig := config.ParseConfig(configFilePath)
	log.Debug().Any("appConfig", appConfig).Msg("APP Config")

	prjConfig := config.ParseProjectConfig()
	log.Debug().Any("prjConfig", prjConfig).Msg("Project Config")

	app := App{LinearAPIToken: appConfig.LinearAPIToken, ProjectConfig: prjConfig}
	app.ParseCLIParams()
}
