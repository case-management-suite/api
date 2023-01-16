package api

import (
	"context"
	"time"

	"github.com/case-management-suite/caseservice"
	"github.com/case-management-suite/common/config"
	"github.com/case-management-suite/common/server"
)

type API struct {
	API        *APIServiceServer
	CaseClient *caseservice.CaseClient
	server.ServerUtils
}

func (API) GetName() string {
	return "API Server"
}

func (API) GetServerConfig() *server.ServerConfig {
	return &server.ServerConfig{
		Type: server.GroupOfServersType,
	}
}

func (s API) Start(ctx context.Context) error {
	errchan1 := s.RunAsync(s.API.Start, ctx, 30*time.Second)
	errchan2 := s.RunAsync(s.CaseClient.Start, ctx, 30*time.Second)
	defer close(errchan1)
	defer close(errchan2)

	if err := <-errchan1; err != nil {
		return err
	}

	if err := <-errchan2; err != nil {
		return err
	}

	return nil
}

func (s API) Stop(ctx context.Context) error {
	defer s.API.Stop(ctx)
	defer s.CaseClient.Stop(ctx)
	return nil
}

var _ server.Serveable = API{}

func NewAPIServer(appConfig config.AppConfig, api *APIServiceServer, caseClient *caseservice.CaseClient) *APIServer {
	s := server.NewServer(func(su server.ServerUtils) API {
		return API{API: api, CaseClient: caseClient, ServerUtils: su}
	}, appConfig)
	return &s
}
