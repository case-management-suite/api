package frontend

import (
	"github.com/case-management-suite/api/controllers"
	"github.com/case-management-suite/caseservice"
	"github.com/case-management-suite/common/config"
	"github.com/case-management-suite/queue"
	"github.com/case-management-suite/rulesengineservice"
	"github.com/case-management-suite/scheduler"
	"go.uber.org/fx"
)

func FxGetClientServices(rconfig config.RulesServiceConfig) fx.Option {
	return fx.Options(
		scheduler.WorkSchedulerModule,
		caseservice.NewFXCaseServiceClient(),
		rulesengineservice.RulesServiceClientModule,
		fx.Provide(
			queue.QueueServiceFactory(rconfig.QueueType),
			rconfig.QueueConfig,
			controllers.NewCaseControllerImpl,
		),
	)
}
