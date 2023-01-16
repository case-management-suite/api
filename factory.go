package api

import (
	"fmt"
	"reflect"

	"github.com/case-management-suite/api/apicommon"
	"github.com/case-management-suite/api/controllers"
	"github.com/case-management-suite/caseservice"
	"github.com/case-management-suite/common/config"
	"github.com/case-management-suite/common/factory"
	"github.com/case-management-suite/common/server"
	"github.com/case-management-suite/rulesengineservice"
	"github.com/case-management-suite/scheduler"
)

type APIFactory func(appConfig config.AppConfig, controller controllers.CaseControllerAPI) server.Server[apicommon.APIService]

type APIServiceServer = server.Server[apicommon.APIService]

type APIServer = server.Server[API]

type APIFactories struct {
	factory.FactorySet
	WorkSchedulerFactories          scheduler.WorkSchedulerFactories
	CaseServiceClientFactory        caseservice.CaseServiceClientFactory
	ControllerFactory               controllers.ControllerFactory
	APIFactory                      APIFactory
	RulesEngineServiceClientFactory rulesengineservice.RulesEngineServiceClientFactory
}

func CastCaseServiceClient(srv server.Server[caseservice.CaseServiceClient]) caseservice.CaseService {
	return srv.Server
}

func (f APIFactories) BuildAPI(appConfig config.AppConfig) (*APIServer, error) {
	if err := factory.ValidateFactorySet(f); err != nil {
		return nil, fmt.Errorf("factory: %s -> %w;", reflect.TypeOf(f).Name(), err)
	}
	workScheduler, err := f.WorkSchedulerFactories.BuildWorkScheduler(appConfig)
	if err != nil {
		return nil, err
	}
	caseServiceClient := f.CaseServiceClientFactory(appConfig)

	api := f.APIFactory(
		appConfig,
		f.ControllerFactory(
			CastCaseServiceClient(caseServiceClient),
			f.RulesEngineServiceClientFactory(*workScheduler),
		),
	)
	return NewAPIServer(appConfig, &api, &caseServiceClient), nil
}
