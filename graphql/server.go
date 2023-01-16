package graphql

import (
	// "dig"
	"context"
	"fmt"

	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/case-management-suite/api/controllers"
	"github.com/case-management-suite/api/graphql/graph"
	"github.com/case-management-suite/common/config"
	"github.com/case-management-suite/common/server"
)

type GraphQLService struct {
	Port       int
	Host       string
	Controller controllers.CaseControllerAPI
	Server     *handler.Server
	server.ServerUtils
}

func (c *GraphQLService) Start(ctx context.Context) error {
	// ctx = ctxutils.DecorateContext(ctx, ctxutils.ContextDecoration{Name: "GraphQLService"})
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Controller: c.Controller}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	c.Logger.Info().Int("port", c.Port).Msg(fmt.Sprintf("connect to http://localhost:%d/ for GraphQL playground", c.Port))
	portStr := fmt.Sprintf(":%d", c.Port)

	c.Server = srv
	go func() {
		http.ListenAndServe(portStr, nil)
	}()
	return nil
}

func (c GraphQLService) Stop(_ context.Context) error {
	return nil
}

func (c GraphQLService) GetName() string {
	return "graphql_api"
}

func (s GraphQLService) GetServerConfig() *server.ServerConfig {
	c := server.ServerConfig{
		Type: server.HttpServerType,
		Port: s.Port,
	}
	return &c
}

func NewGraphQLService(appConfig config.AppConfig, controller controllers.CaseControllerAPI, su server.ServerUtils) GraphQLService {
	return GraphQLService{Port: appConfig.GraphQLConfig.Port, Controller: controller, ServerUtils: su}
}

var _ server.Serveable = &GraphQLService{}
