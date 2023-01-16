package main_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/case-management-suite/api"
	"github.com/case-management-suite/api/controllers"
	"github.com/case-management-suite/caseservice"
	"github.com/case-management-suite/common/config"
	"github.com/case-management-suite/common/service"
	"github.com/case-management-suite/queue"
	"github.com/case-management-suite/rulesengineservice"
	"github.com/case-management-suite/scheduler"
	"github.com/case-management-suite/testutil"
)

func TestSmoke(t *testing.T) {
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = app.Start(ctx)
	testutil.AssertNilError(err, t)

	err = app.Stop(ctx)
	testutil.AssertNilError(err, t)
}
