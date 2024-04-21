package config

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	defaultConfig := Config{
		AppName:  _serviceName,
		Env:      "local",
		LogLevel: _serviceLogLevel,
		APIKey:   _serviceAPIKey,
		Server: Server{
			PublicPort: _servicePublicPort,
		},
		MongoDB: Mongodb{
			Host:            "localhost:27017",
			DB:              _mongodbDB,
			Timeout:         _mongodbTimeout,
			MaxConnIdleTime: _mongodbMaxConnIdleTime,
			MinPoolSize:     _mongodbMinPoolSize,
			MaxPoolSize:     _mongodbMaxPoolSize,
			MaxConnecting:   _mongodbMaxConnecting,
		},
	}

	testConfig := Config{
		AppName:  "api-demo",
		Env:      "unit-test",
		LogLevel: "debug",
		APIKey:   "test-api-key",
		Server: Server{
			PublicPort: ":3000",
		},
		MongoDB: Mongodb{
			Host:            "test:27017",
			DB:              "articles",
			Username:        "test-user",
			Password:        "test-user-password",
			Timeout:         5 * time.Second,
			MaxConnIdleTime: 1 * time.Minute,
			MinPoolSize:     0,
			MaxPoolSize:     20,
			MaxConnecting:   2,
		},
	}

	tests := []struct {
		name  string
		want  Config
		setup func()
	}{
		{
			name: "default Config",
			want: defaultConfig,
		},
		{
			name: "test Config",
			want: testConfig,
			setup: func() {
				InitConfig("../../deployment/test/config.yaml")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			got := GetConfig()
			if !assert.EqualValues(t, tt.want, got) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
				return
			}
			got.Server.GetPublicPort()
			t.Logf("Config %+v\n", got)
		})
	}
}

func TestUpdateConfig(t *testing.T) {
	type args struct {
		newCfg Config
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "swap Config",
			args: args{
				newCfg: Config{
					AppName:  "api-demo",
					Env:      "unit-test",
					LogLevel: "debug",
					APIKey:   "test-api-key",
					Server: Server{
						PublicPort: ":3000",
					},
				},
			},
			want: Config{
				AppName:  "api-demo",
				Env:      "unit-test",
				LogLevel: "debug",
				APIKey:   "test-api-key",
				Server: Server{
					PublicPort: ":3000",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateConfig(tt.args.newCfg)
			got := _config.Load().(Config)
			if !assert.EqualValues(t, tt.want, _config.Load().(Config)) {
				t.Errorf("UpdateConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReloadConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "reload Config",
			args: args{
				path: "../../deployment/test/config_reload.yaml",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReloadConfig(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("ReloadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			Config := _config.Load().(Config)
			t.Logf("Config %+v", Config)
		})
	}
}

func TestInitConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test Config",
			args: args{
				path: "../../deployment/test/Config.yaml",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitConfig(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("InitConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			Config := _config.Load().(Config)
			t.Logf("Config %+v", Config)
		})
	}
	_config = atomic.Value{}
}
