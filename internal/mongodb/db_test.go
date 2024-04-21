package mongodb

import (
	"context"
	"flag"
	"log"
	"os"
	"testing"
	"time"

	"github.com/brionac626/api-demo/internal/config"
)

func TestMain(m *testing.M) {
	flag.Parse()
	if testing.Short() {
		log.Print("skipping mongodb client integration test in short mode")
		return
	}

	os.Setenv(config.EnvUnitTest, "1")
	config.InjectTestConfig(
		config.Config{
			MongoDB: config.Mongodb{
				Host:            "mongodb://localhost:27017",
				DB:              "articles",
				Username:        "articles-repo",
				Password:        "repo-password",
				Timeout:         5 * time.Second,
				MaxConnIdleTime: 1 * time.Minute,
				MinPoolSize:     0,
				MaxPoolSize:     20,
				MaxConnecting:   2,
			},
		},
	)
}

// TestInitMongoDBClient it's the integration test will try to connect to the local MongoDB
// use -short flag to skip this test
func TestInitMongoDBClient(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "local-test",
			args:    args{ctx: context.Background()},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := InitMongoDBClient(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitMongoDBClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
