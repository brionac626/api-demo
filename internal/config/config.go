package config

import (
	"sync/atomic"

	"github.com/spf13/viper"
)

// Set default values for the service global configuration
func init() {
	viper.SetDefault("app_name", "api-demo")
	viper.SetDefault("env", "local")
	viper.SetDefault("log_level", "debug")
	viper.SetDefault("api_key", "default-api-key")
	viper.SetDefault("server", server{PublicPort: ":3000"})
}

var (
	// global configuration
	_config atomic.Value

	// action for initialization or reload configuration
	initOrReloadConfig = func(path string) error {
		viper.SetConfigFile(path)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		var cfg Config
		if err := viper.Unmarshal(&cfg); err != nil {
			return err
		}

		_config.Store(cfg)

		return nil
	}
)

// Config is the api-demo configuration structure
type Config struct {
	AppName  string `mapstructure:"app_name"`
	Env      string `mapstructure:"env"`
	LogLevel string `mapstructure:"log_level"`
	APIKey   string `mapstructure:"api_key"`
	Server   server `mapstructure:"server"`
}

// server is the server configuration structure for start the api-demo http service
type server struct {
	PublicPort string `mapstructure:"public_port"`
}

// GetPublicPort get the public port that api-demo service is listening at
func (s *server) GetPublicPort() string {
	return s.PublicPort
}

// InitConfig initialization the service's global configuration
func InitConfig(path string) error {
	return initOrReloadConfig(path)
}

// GetConfig get a copy of service's global configuration
func GetConfig() Config {
	v := _config.Load()
	cfg, ok := v.(Config)
	if !ok {
		// get default Config here
		viper.ReadInConfig()
		var c Config
		viper.Unmarshal(&c)

		return c
	}

	return cfg
}

// UpdateConfig swaps the global configuration by the given configuration
func UpdateConfig(newCfg Config) {
	_config.Swap(newCfg)
}

// ReloadConfig reload the service's global configuration from the given file path
func ReloadConfig(path string) error {
	return initOrReloadConfig(path)
}
