package graphql_test

import (
	"testing"

	"github.com/case-management-suite/api/graphql"
	"github.com/case-management-suite/common/config"
	"github.com/case-management-suite/testutil"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestCaseServiceGRPCServerModule(t *testing.T) {
	appConfig := config.NewLocalTestAppConfig()
	testutil.AppFx(t, fx.Options(graphql.FxServerOpts(appConfig)), func(a *fxtest.App) {})
}
