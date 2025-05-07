package framework

import (
	"context"
	"sync"

	"go.uber.org/zap"
)

type Module interface {
	Start(ctx context.Context)
	Name() string
}

type App struct {
	logger  *zap.SugaredLogger
	modules map[string]Module
}

func NewApp(logger *zap.SugaredLogger, modules ...Module) *App {
	moduleMap := map[string]Module{}
	for _, m := range modules {
		moduleMap[m.Name()] = m
	}
	return &App{
		modules: moduleMap,
		logger:  logger,
	}
}

func (a *App) Run(ctx context.Context) {
	a.logger.Info("Starting modules")
	for _, m := range a.modules {
		a.logger.Infof("Starting module: %s", m.Name())
		go m.Start(ctx)
	}
	<-ctx.Done()
	a.logger.Info("Stopping apps")
}

func (a *App) RunOnce(ctx context.Context) {
	a.logger.Info("Starting modules")
	var wg sync.WaitGroup
	for _, m := range a.modules {
		wg.Add(1)
		a.logger.Infof("Starting module: %s", m.Name())
		go func(im Module) {
			im.Start(ctx)
			wg.Done()
		}(m)
	}
	wg.Wait()
	a.logger.Info("Stopping apps")
}
