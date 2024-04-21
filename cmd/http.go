package cmd

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/brionac626/api-demo/internal/config"
	"github.com/brionac626/api-demo/provider"
	"github.com/spf13/cobra"
)

var configPath string

var apiCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a http server for article api service",
	Run: func(cmd *cobra.Command, args []string) {
		if err:=config.InitConfig(configPath);err!=nil{
			slog.Error("init config", err)
			os.Exit(1)
		}

		
		ctx := context.TODO()
		mc, err := provider.ProvideMongoClient(ctx)
		if err != nil {
			slog.Error("mongodb client", err)
			os.Exit(1)
		}

		repo := provider.ProvideArticleRepository(mc)
		handler := provider.ProvideArticleHandler(repo)
		server := provider.ProvideServer(handler)

		go func() {
			if err := server.Start(config.GetConfig().Server.PublicPort); err != nil {
				slog.Error("start server failed", slog.String("err", err.Error()))
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		slog.Info("shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			slog.Error("shutdown server failed", slog.String("err", err.Error()))
			return
		}

		slog.Info("shutdown server successful")
	},
}

var boAPICmd = &cobra.Command{
	Use:   "back-office",
	Short: "Start the http server for the back office service",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("app back-office test")
	},
}
