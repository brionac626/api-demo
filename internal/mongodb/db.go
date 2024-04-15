package mongodb

import (
	"context"

	"github.com/brionac626/api-demo/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDBClient(ctx context.Context) (*mongo.Client, error) {
	cfg := config.GetConfig()

	opts := options.Client().
		SetAuth(options.Credential{
			Username: cfg.MongoDB.Username,
			Password: cfg.MongoDB.Password,
		}).
		SetMaxConnIdleTime(cfg.MongoDB.MaxConnIdleTime).
		SetMinPoolSize(cfg.MongoDB.MinPoolSize).
		SetMaxPoolSize(cfg.MongoDB.MaxPoolSize).
		SetMaxConnecting(cfg.MongoDB.MaxConnecting).
		ApplyURI(cfg.MongoDB.Host)

	mc, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongoDB.Timeout)
	defer cancel()
	if err := mc.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return mc, nil
}
