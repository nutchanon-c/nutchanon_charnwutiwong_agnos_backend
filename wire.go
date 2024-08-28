//go:build wireinject
// +build wireinject

package main

import (
	"agnos/backend/src/api"
	"agnos/backend/src/usecase/password"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeServer() (*gin.Engine, error) {
	wire.Build(
		password.NewPasswordUsecase,
		api.NewAPIHandler,
		ProvideServer,
	)
	return &gin.Engine{}, nil
}

func ProvideServer(
	apiHandler *api.APIHandler,
) *gin.Engine {
	r := gin.Default()
	r.POST("/api/strong_password_steps", apiHandler.StrongPasswordSteps)
	return r
}
