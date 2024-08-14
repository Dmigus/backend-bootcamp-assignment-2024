// Package main содержит main для запуска программы, а также обработку переменных окружения и аргументов командной строки
package main

import (
	"backend-bootcamp-assignment-2024/internal/apps/renting"
	"errors"

	"github.com/samber/lo"
	"github.com/vrischmann/envconfig"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

type globResult struct {
	fx.Out
	RentingConfig *renting.Config
	Logger        *zap.Logger
}

func initGlobalModule(lc fx.Lifecycle) (globResult, error) {
	rentingConfig := renting.Config{}
	err := parseConfigs(&rentingConfig)
	if err != nil {
		return globResult{}, err
	}
	logger, err := getLogger()
	if err != nil {
		return globResult{}, err
	}
	lc.Append(fx.StopHook(logger.Sync))
	return globResult{Logger: logger, RentingConfig: &rentingConfig}, nil
}

func parseConfigs(configs ...any) error {
	errs := lo.Map(configs, func(conf any, _ int) error {
		return envconfig.Init(conf)
	})
	return errors.Join(errs...)
}

func main() {
	fx.New(fx.Provide(initGlobalModule),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		renting.Module,
	).Run()
}
