package graphql

import (
	"context"

	"github.com/case-management-suite/api/controllers"
	"github.com/case-management-suite/api/frontend"
	"github.com/case-management-suite/common/config"
	"go.uber.org/fx"
)

type GraphQLServerParams struct {
	fx.In
	AppConfig  config.AppConfig
	Controller controllers.CaseControllerAPI
}

func NewGraphQLServer(lc fx.Lifecycle, params GraphQLServerParams) GraphQLService {
	server := NewGraphQLService(params.AppConfig, params.Controller)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Start(ctx)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// return http.
			return nil
		},
	})

	return server
}

func NewQueueConfig(appConfig config.AppConfig) config.QueueConnectionConfig {
	return appConfig.RulesServiceConfig.QueueConfig
}

func FxServerOpts(appConfig config.AppConfig) fx.Option {
	return fx.Module("case_service_graphql",
		config.FxConfig(appConfig),
		frontend.FxGetClientServices(appConfig.RulesServiceConfig),
		fx.Provide(
			NewGraphQLServer,
		),
		fx.Invoke(func(GraphQLService) {}),
	)
}

func CreateLiteGraphQLAPIServer(appConfig config.AppConfig) *fx.App {
	return fx.New(
		FxServerOpts(appConfig),
		// fx.Invoke(func(*handler.Server) {})
	)
}
