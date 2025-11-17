package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog/log"
)

func parseProjectConfig(prjDir string) (ProjectConfig, bool) {
	prjConfigFile := filepath.Join(prjDir, ".linear-cli-config.toml")
	log.Debug().Any("prjConfigFile", prjConfigFile).Msg("ParseProjectConfig")

	var prjConfig ProjectConfig

	// check if project config file exists
	if _, err := os.Stat(prjConfigFile); errors.Is(err, os.ErrNotExist) {
		//prjConfigFile does not exist
		return prjConfig, false
	}

	// read in config file
	dat, err := os.ReadFile(prjConfigFile)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Invalid Project Config file at %s", prjConfigFile))
	}

	_, err = toml.Decode(string(dat[:]), &prjConfig)

	log.Debug().Any("prjConfig", prjConfig).Msg("processed project config")

	return prjConfig, true
}

func ParseProjectConfig() ProjectConfig {
	cwdPath, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to get CWD")
	}

	prjConfig, found := parseProjectConfig(cwdPath)

	if !found {
		prjDir := cwdPath

		for _ = range 3 {
			prjDir = filepath.Join(prjDir, "..")
			prjConfig, found := parseProjectConfig(prjDir)
			if found {
				return prjConfig
			}
		}
	}

	return prjConfig
}
