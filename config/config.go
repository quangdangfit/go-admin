package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Schema struct {
	Env          string `mapstructure:"env"`
	DefaultLimit int    `mapstructure:"default_limit"`
	MaxLimit     int    `mapstructure:"max_limit"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Env      string `mapstructure:"env"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Database int    `mapstructure:"database"`
	} `mapstructure:"redis"`

	Cache struct {
		Enable     bool `mapstructure:"enable"`
		ExpiryTime int  `mapstructure:"expiry_time"`
	} `mapstructure:"cache"`

	JWTAuth struct {
		SigningKey          string `mapstructure:"signing_key"`
		Expired             int    `mapstructure:"expired"`
		SigningRefreshKey   string `mapstructure:"signing_refresh_key"`
		ExpiredRefreshToken int    `mapstructure:"expired_refresh_token"`
	} `mapstructure:"jwt_auth"`
}

var Config Schema

func init() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(".")             // Look for config in current directory
	config.AddConfigPath("config/")       // Optionally look for config in the working directory.
	config.AddConfigPath("../config/")    // Look for config needed for tests.
	config.AddConfigPath("../")           // Look for config needed for tests.
	config.AddConfigPath("../../config/") // Look for config needed for tests.
	config.AddConfigPath("../../")        // Look for config needed for tests.

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	config.AutomaticEnv()

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = config.Unmarshal(&Config)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// fmt.Printf("Current Config: %+v", Config)
}
