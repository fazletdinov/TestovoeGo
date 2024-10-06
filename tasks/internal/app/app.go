package app

import (
	"log"
	"os"

	"tasks/config"
	"tasks/internal/database"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Application struct {
	DB         *bun.DB
	Env        *config.Config
	Log        *zerolog.Logger
	GRPCServer *GRPC
}

func App() Application {
	app := &Application{}
	Env, errEnv := config.InitConfig()
	if errEnv != nil {
		log.Fatalf("ошибка загрузки ENV - %v", errEnv)
	}
	PostgresClient := database.InitDatabse(Env)
	log := setupLogger(Env.Env)
	app.GRPCServer = NewGRPC(log, Env, PostgresClient)
	app.Env = Env
	app.Log = log
	app.DB = PostgresClient
	return *app
}

func setupLogger(env string) *zerolog.Logger {
	zerolog.TimeFieldFormat = "02/Jan/2006 - 15:04:05 -0700"
	switch env {
	case envLocal:
		logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
		if err != nil {
			logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
			logger.Fatal().Err(err).Msg("Не удалось открыть файл.")
		}
		multi := zerolog.MultiLevelWriter(os.Stdout, logFile)
		logger := zerolog.New(multi).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
		return &logger
	case envDev:
		logger := zerolog.New(os.Stdout).
			Level(zerolog.DebugLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
		return &logger
	case envProd:
		logger := zerolog.New(os.Stdout).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
		return &logger
	}
	return nil
}
