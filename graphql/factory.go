package graphql

import (
	"github.com/case-management-suite/api/apicommon"
	"github.com/case-management-suite/api/controllers"
	"github.com/case-management-suite/common/config"
	"github.com/case-management-suite/common/server"
)

func GraphQLAPIFactory(appConfig config.AppConfig, controller controllers.CaseControllerAPI) server.Server[apicommon.APIService] {
	return server.NewServer(func(su server.ServerUtils) apicommon.APIService {
		srv := NewGraphQLService(appConfig, controller, su)
		return &srv
	},
		appConfig)
}
