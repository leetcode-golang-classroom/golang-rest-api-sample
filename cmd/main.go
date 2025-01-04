package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/application"
	"github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/config"
	mlog "github.com/leetcode-golang-classroom/golang-rest-api-sample/interanl/logger"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(
		os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		},
	))
	rootContext := context.WithValue(context.Background(), mlog.CtxKey{}, logger)
	app := application.New(rootContext, config.AppConfig)
	ctx, cancel := signal.NotifyContext(rootContext, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		logger.Error("failed to start app", "error", err)
	}
}
