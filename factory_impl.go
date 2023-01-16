package api

import (
	"github.com/case-management-suite/api/apicommon"
	"github.com/case-management-suite/api/controllers"
	"github.com/case-management-suite/api/graphql"
	"github.com/case-management-suite/api/rest"
	"github.com/case-management-suite/common/config"
	"github.com/case-management-suite/common/server"
)

func GetAPIFactory(t config.ApiType) APIFactory {
	switch t {
	case config.REST:
		return rest.RestAPIFactory
	case config.GraphQL:
		return graphql.GraphQLAPIFactory
	default:
		return func(appConfig config.AppConfig, controller controllers.CaseControllerAPI) server.Server[apicommon.APIService] {
			panic("Unimplemented APIType")
		}
	}
}
