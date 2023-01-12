package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/tinoquang/service-boilerplate/pkg/api/server"
	"github.com/tinoquang/service-boilerplate/pkg/config"
	"github.com/tinoquang/service-boilerplate/pkg/ctxlogger"
	"github.com/tinoquang/service-boilerplate/pkg/services"
)

var (
	ldBuildDate = "<Unset build date>"
	ldGitCommit = "<Unset git commit>"
)

func main() {
	_ = godotenv.Load()

	version := flag.Bool("v", false, "prints version")
	flag.Parse()
	if *version {
		fmt.Printf("built at %s with commit %s\n", ldBuildDate, ldGitCommit)
		os.Exit(0)
	}

	cfg := config.New()

	var logger *zap.Logger
	if cfg.Env == config.Prod {
		encoder := zap.NewProductionEncoderConfig()
		encoder.EncodeTime = zapcore.ISO8601TimeEncoder
		logConfig := zap.Config{
			Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
			Development: false,
			Sampling: &zap.SamplingConfig{
				Initial:    100,
				Thereafter: 100,
			},
			Encoding:         "json",
			EncoderConfig:    encoder,
			OutputPaths:      []string{"stderr"},
			ErrorOutputPaths: []string{"stderr"},
		}

		logger, _ = logConfig.Build()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	logger = logger.With(zap.String("env", cfg.Env.String()), zap.String("commitHash", ldGitCommit))
	ctxlogger.SetDefaultLogger(logger)
	svc := services.New(cfg)
	s := server.New(cfg, logger, svc)

	// create a context with graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
		cancel()
	}()

	s.Start(ctx)

	_ = logger.Sync()
}
