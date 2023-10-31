package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/mirzaahmedov/simple_todo/internal/config"
	"github.com/mirzaahmedov/simple_todo/internal/http"
	"github.com/mirzaahmedov/simple_todo/internal/store/postgres"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "config/local.toml", "path to the config file")
}
func main() {
	flag.Parse()

	slog.SetDefault(
		slog.New(
			slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		),
	)

	cfg, err := config.Load(configPath)
	if err != nil {
		slog.Error("failed to load config", slog.String("path", configPath), slog.String("error", err.Error()))
		os.Exit(1)
	}

	s := postgres.NewStore(cfg.DatabaseURL)

	err = s.Open()
	if err != nil {
		slog.Error("can not connect to the database", slog.String("database", cfg.DatabaseURL), slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer s.Close()

	r := router.NewHTTPRouter(s, slog.Default())

	r.Init()

	err = r.Listen(cfg.BindingAddress)
	if err != nil {
		slog.Error("error running http server", slog.String("binding_address", cfg.BindingAddress), slog.String("error", err.Error()))
		os.Exit(1)
	}
}
