package config

import (
	"log/slog"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/spf13/viper"
)

const (
	// EnvUnitTest is an environment variable for internal unit test
	EnvUnitTest = "UNIT_TEST"

	_serviceName            = "api-demo"
	_serviceLogLevel        = "debug"
	_serviceAPIKey          = "default-api-key"
	_servicePublicPort      = ":3000"
	_mongodbDB              = "articles"
	_mongodbTimeout         = 5 * time.Second
	_mongodbMaxConnIdleTime = 1 * time.Minute
	_mongodbMinPoolSize     = 0
	_mongodbMaxPoolSize     = 20
	_mongodbMaxConnecting   = 2
)

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

// Set default values for the service global configuration
func init() {
	viper.SetDefault("app_name", _serviceName)
	viper.SetDefault("env", "local")
	viper.SetDefault("log_level", _serviceLogLevel)
	viper.SetDefault("api_key", _serviceAPIKey)
	viper.SetDefault("server", Server{PublicPort: _servicePublicPort})
	viper.SetDefault("mongodb",
		Mongodb{
			Host:            "localhost:27017",
			DB:              _mongodbDB,
			Timeout:         _mongodbTimeout,
			MaxConnIdleTime: _mongodbMaxConnIdleTime,
			MinPoolSize:     _mongodbMinPoolSize,
			MaxPoolSize:     _mongodbMaxPoolSize,
			MaxConnecting:   _mongodbMaxConnecting,
		},
	)
}

// Config is the api-demo configuration structure
type Config struct {
	AppName  string  `mapstructure:"app_name"`
	Env      string  `mapstructure:"env"`
	LogLevel string  `mapstructure:"log_level"`
	APIKey   string  `mapstructure:"api_key"`
	Server   Server  `mapstructure:"server"`
	MongoDB  Mongodb `mapstructure:"mongodb"`
}

// server is the server configuration structure for start the api-demo http service
type Server struct {
	PublicPort string `mapstructure:"public_port"`
}

type Mongodb struct {
	Host            string        `mapstructure:"host"`
	DB              string        `mapstructure:"db"`
	Collection      string        `mapstructure:"collection"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	Timeout         time.Duration `mapstructure:"timeout"`
	MaxConnIdleTime time.Duration `mapstructure:"max_conn_idle_time"`
	MinPoolSize     uint64        `mapstructure:"min_pool_size"`
	MaxPoolSize     uint64        `mapstructure:"max_pool_size"`
	MaxConnecting   uint64        `mapstructure:"max_connecting"`
}

// GetPublicPort get the public port that api-demo service is listening at
func (s *Server) GetPublicPort() string {
	return s.PublicPort
}

// InjectTestConfig inject a config for internal unit test
func InjectTestConfig(config Config) {
	env, exists := os.LookupEnv(EnvUnitTest)
	if !exists {
		slog.Warn("can't find environment variable",
			slog.String("env", EnvUnitTest))

		return
	}

	if b, err := strconv.ParseBool(env); !b || err != nil {
		slog.Warn("isn't unit test env or can't parse bool value",
			slog.Bool("unit_test_flag", b),
			slog.Any("err", err))

		return
	}

	slog.Debug("inject unit test config")

	_config.Store(config)
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
