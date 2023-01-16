package apicommon

import (
	"context"

	"github.com/case-management-suite/common/server"
)

type APIService interface {
	Start(context.Context) error
	Stop(context.Context) error
	GetName() string
	GetServerConfig() *server.ServerConfig
}
