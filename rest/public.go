package rest

import (
	"context"

	"github.com/case-management-suite/api/apicommon"
	"github.com/case-management-suite/api/controllers"
	"github.com/case-management-suite/api/rest/restapi"
	"github.com/case-management-suite/common/config"
	"github.com/case-management-suite/common/server"
)

type RestServer struct {
	Server *restapi.Server
}

func (s RestServer) Start(ctx context.Context) error {
	go func() {
		s.Server.Serve()
	}()
	return nil
}

func (s RestServer) Stop(ctx context.Context) error {
	return s.Server.Shutdown()
}

func (s RestServer) GetName() string {
	return "rest_server"
}

func (s RestServer) GetServerConfig() *server.ServerConfig {
	return &server.ServerConfig{
		Type: server.HttpServerType,
		Host: s.Server.Host,
		Port: s.Server.Port,
	}
}

func RestAPIFactory(appConfig config.AppConfig, controller controllers.CaseControllerAPI) server.Server[apicommon.APIService] {
	return server.NewServer(func(su server.ServerUtils) apicommon.APIService {
		caseMgmntAPI := NewCaseMgmtAPI(controller, su.Logger)
		srv := restapi.NewServer(&caseMgmntAPI)
		return RestServer{Server: srv}
	},
		appConfig)
}

var _ server.Serveable = RestServer{}
