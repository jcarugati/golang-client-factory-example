package configuration

import (
	"fmt"
	"../constants"
	"github.com/mercadolibre/go-meli-toolkit/gomelipass"
	"gopkg.in/yaml.v2"
	logger "github.com/sirupsen/logrus"
	"os"
	"github.com/kelseyhightower/envconfig"
)

func Init() {
	Config := GetConfigurationInstance() // gets configuration Instance

	loadYaml(Config, createConfigPath()) //loads yaml file with configurations acording to its context

	loadEnv(Config) //loads the yaml information on the config singleton
}

func loadYaml(cfg *configuration, path string) {
	f, err := os.Open(path)
	if err != nil {
		logger.Error(fmt.Sprintf("Error opening config yaml. PATH: %s", path), err, "method:config.loadYaml")
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		logger.Error("Error loading config yaml", err, "method:config.loadYaml")
	}
	logger.Info(fmt.Sprintf("Configuration YAML loaded correctly. PATH: %s", path),"method:loadYaml")
}

func loadEnv(cfg *configuration) {
	err := envconfig.Process("", cfg)
	if err != nil {
		logger.Error("Error loading env variables", err, "method:config.loadEnv")
	}
}

func createConfigPath() string {
	var configPath string
	if gomelipass.GetEnv("ENVIRONMENT") == "QA" || gomelipass.GetEnv("SECRET_ENVIRONMENT") == "QA" {
		configPath = constants.QA_CONFIG_PATH
	}else{
		configPath = constants.PROD_CONFIG_PATH
	}
	return configPath
}