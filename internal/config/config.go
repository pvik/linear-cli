package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog/log"
)

func ParseConfig(configFile string) AppConfig {
	log.Debug().Any("configFile", configFile).Msg("ParseConfig")

	// read in config file
	dat, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Config file not found. Please create config file at %s", configFile))
	}

	var appConfig AppConfig

	_, err = toml.Decode(string(dat[:]), &appConfig)
	// _, err := env.UnmarshalFromEnviron(&appConfig)
	// if err != nil {
	// 	log.Fatal().AnErr("error reading env vars", err)
	// }

	log.Debug().Any("appConfig", appConfig).Msg("processed app config")

	return appConfig
}
