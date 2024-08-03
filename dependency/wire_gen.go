// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependency

import (
	"os"
	"sagala-todo/pkg/adapters"
	"sagala-todo/pkg/constants"
	"sagala-todo/src/delivery/v1http"
	"sagala-todo/src/usecase"
)

// Injectors from dependency.go:

func InitConfiguration() adapters.Config {
	config := provideConfiguration()
	return config
}

func InitOsSignalChannel() chan os.Signal {
	v := provideOsSignal()
	return v
}

func InitTodoV1HttpHandler(cfg adapters.Config) *v1http.V1Handler {
	v := _wireValue
	v2 := _wireValue2
	v3 := provideSql(cfg, v, v2)
	todoUsecase := usecase.ProvideUsecase(v3, cfg)
	v1Handler := &v1http.V1Handler{
		Config:  cfg,
		Usecase: todoUsecase,
	}
	return v1Handler
}

var (
	_wireValue = []adapters.SqlConfig{
		{
			RegistryName: constants.ConnSqlDefault,
			DriverName:   constants.SqliteDriver,
			MaxLifeTime:  1,
			MaxIdleTime:  5,
			MaxIdleConns: 10,
			MaxOpenConns: 100,
		},
	}
	_wireValue2 = []string{constants.DSNDefault}
)
