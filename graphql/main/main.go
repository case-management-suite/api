package main

import (
	"github.com/case-management-suite/api/graphql"
	"github.com/case-management-suite/common/config"
	"github.com/rs/zerolog"
)

func main() {
	appConfig := config.NewLocalAppConfig()
	appConfig.RulesServiceConfig.QueueConfig.LogLevel = zerolog.DebugLevel
	graphql.CreateLiteGraphQLAPIServer(appConfig).Run()
}
