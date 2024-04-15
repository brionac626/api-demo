package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
}

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name  string
		want  Config
		setup func()
	}{
		{
			name: "default Config",
			want: Config{
				AppName:  "api-demo",
				Env:      "local",
				LogLevel: "debug",
				APIKey:   "default-api-key",
				Server: server{
					PublicPort: ":3000",
				},
			},
		},
		{
			name: "test Config",
			want: Config{
				AppName:  "api-demo",
				Env:      "local",
				LogLevel: "debug",
				APIKey:   "test-api-key",
				Server: server{
					PublicPort: ":3000",
				},
			},
			setup: func() {
				InitConfig("../../deployment/test/Config.yaml")
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
					Env:      "local",
					LogLevel: "debug",
					APIKey:   "test-api-key",
					Server: server{
						PublicPort: ":3000",
					},
				},
			},
			want: Config{
				AppName:  "api-demo",
				Env:      "local",
				LogLevel: "debug",
				APIKey:   "test-api-key",
				Server: server{
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
