package config

import (
	"fmt"
	"strings"

	"github.com/sonnht1409/scanning/service/common"
	"github.com/spf13/viper"
)

// Config ...
var (
	Values Config

	log = common.GetLogger("scanning-config", "INFO")
)

// Config of config
type Config struct {
	DB struct {
		Address  string `mapstructure:"address"`
		Port     int    `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"db"`

	Env    string `mapstructure:"env"`
	Worker struct {
		Topic   string `mapstructure:"topic"`
		Channel string `mapstructure:"channel"`
	} `mapstructure:"worker"`

	Application struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"application"`

	Redis struct {
		URL        string `mapstructure:"url"`
		DB         int    `mapstructure:"db"`
		MaxRetries int    `mapstructure:"max_retries"`
	}
	AccessToken string `mapstructure:"access_token"`
}

// InitConfig ...
func init() {
	// Initialize viper default instance with API base config.
	config := viper.New()
	config.SetConfigName("config")        // Name of config file (without extension).
	config.AddConfigPath(".")             // Look for config in current directory
	config.AddConfigPath("./config")      // Optionally look for config in the working directory.
	config.AddConfigPath("../config/")    // Look for config needed for tests.
	config.AddConfigPath("../../config/") // Look for config needed for tests.
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	config.AutomaticEnv()
	// Initialize map that contains viper configuration objects.
	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = config.Unmarshal(&Values)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	log.Infof("Current Config: %+v", Values)
}
