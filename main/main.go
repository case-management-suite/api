package main

import (
	"log"

	"github.com/case-management-suite/api"
	"github.com/case-management-suite/api/controllers"
	"github.com/case-management-suite/caseservice"
	"github.com/case-management-suite/common/config"
	"github.com/case-management-suite/common/server"
	"github.com/case-management-suite/common/service"
	"github.com/case-management-suite/queue"
	"github.com/case-management-suite/rulesengineservice"
	"github.com/case-management-suite/scheduler"
)

func main() {
	appConfig := config.NewLocalAppConfig()
	factories := api.APIFactories{
		CaseServiceClientFactory:        caseservice.NewCaseServiceClient,
		ControllerFactory:               controllers.GetControllerFactory(),
		APIFactory:                      api.GetAPIFactory(appConfig.API.APIType),
		RulesEngineServiceClientFactory: rulesengineservice.NewRulesServiceClient,
		WorkSchedulerFactories: scheduler.WorkSchedulerFactories{
			WorkSchedulerFactory: scheduler.NewWorkScheduler,
			QueueServiceFactory:  queue.QueueServiceFactory(appConfig.RulesServiceConfig.QueueType),
			// TODO: use prod
			ServiceUtilsFactory: service.NewTestServiceUtils,
		},
	}
	app, err := factories.BuildAPI(appConfig)
	if err != nil {
		log.Fatalf("failed to start the API server: %v", err)
	}
	server.StartServer(*app)
}
