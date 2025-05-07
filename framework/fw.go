package framework

import (
	"time"

	"go.uber.org/zap"
)

type FW interface {
	GetConfig() any
	GetDB(name string) any
	GetLogger() *zap.SugaredLogger
	TimeFunction(functionName string, fn func())
	GetSingleton(name string) any
}

type Framework struct {
	config     any
	databases  map[string]any
	logger     *zap.SugaredLogger
	singletons map[string]any
}

func New(config any, logger *zap.SugaredLogger, dbs map[string]any, singletons map[string]any) FW {
	logger.Info("Initialising Framework")
	return &Framework{
		config:     config,
		databases:  dbs,
		logger:     logger,
		singletons: singletons,
	}
}

func (fw *Framework) GetConfig() any {
	return fw.config
}

func (fw *Framework) GetLogger() *zap.SugaredLogger {
	return fw.logger
}

func (fw *Framework) GetDB(name string) any {
	return fw.databases[name]
}

func (fw *Framework) GetSingleton(name string) any {
	return fw.singletons[name]
}

func (fw *Framework) TimeFunction(functionName string, fn func()) {
	start := time.Now()
	fn()
	fw.logger.Infof("%s ran for %s", functionName, time.Since(start))
}
