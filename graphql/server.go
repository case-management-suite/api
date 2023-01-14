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
	"github.com/case-management-suite/common/ctxutils"
	"github.com/rs/zerolog/log"
)

type GraphQLService struct {
	Port       int
	Controller controllers.CaseControllerAPI
	Server     *handler.Server
}

func (c *GraphQLService) Start(ctx context.Context) error {
	ctx = ctxutils.DecorateContext(ctx, ctxutils.ContextDecoration{Name: "GraphQLService"})
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Controller: c.Controller}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Info().Int("port", c.Port).Msg(fmt.Sprintf("connect to http://localhost:%d/ for GraphQL playground", c.Port))
	portStr := fmt.Sprintf(":%d", c.Port)

	c.Server = srv
	return http.ListenAndServe(portStr, nil)
}

func (c GraphQLService) Stop() {

}

func NewGraphQLService(appConfig config.AppConfig, controller controllers.CaseControllerAPI) GraphQLService {
	return GraphQLService{Port: appConfig.GraphQLConfig.Port, Controller: controller}
}
