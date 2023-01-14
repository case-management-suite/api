package test

import (
	"context"
	"net/http"

	"os"
	"testing"
	"time"

	"github.com/Khan/genqlient/graphql"
	graphql_server "github.com/case-management-suite/api/graphql"
	"github.com/case-management-suite/caseservice"
	"github.com/case-management-suite/common/config"
	"github.com/case-management-suite/models"
	"github.com/case-management-suite/rulesengineservice"
	"github.com/case-management-suite/testutil"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestNewCase(t *testing.T) {
	runApp(t, func(t *testing.T) {
		t.Run("testNewCase", testNewCase)
	})
}

func runApp(t *testing.T, fn func(t *testing.T)) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false})
	appConfig := config.NewLocalAppConfig()

	casesService := caseservice.NewCaseServiceGRPCServer(appConfig)
	if err := casesService.Err(); err != nil {
		log.Fatal().Err(err)
		t.FailNow()
	}

	rulesService := rulesengineservice.NewRulesServiceCServer(appConfig)
	if err := rulesService.Err(); err != nil {
		log.Fatal().Err(err)
		t.FailNow()
	}

	app := graphql_server.CreateLiteGraphQLAPIServer(appConfig)

	ctx, done := context.WithTimeout(context.Background(), 30*time.Second)
	defer done()

	if err := casesService.Start(ctx); err != nil {
		log.Fatal().Err(err)
	}

	if err := rulesService.Start(ctx); err != nil {
		log.Fatal().Err(err)
	}

	err := app.Start(ctx)
	testutil.AssertNilError(err, t)

	fn(t)

	defer app.Stop(ctx)
}

func testNewCase(t *testing.T) {
	ctx := context.Background()
	client := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)
	resp, err := NewCase(ctx, client)
	testutil.AssertNilError(err, t)
	// graphqlClient := NewClient("https://api.github.com/graphql", &httpClient)
	id := resp.GetNewCase()

	resp2, err := FindCase(ctx, client, id)
	testutil.AssertNilError(err, t)
	testutil.AssertEq(id, resp2.Case.ID, t)
	testutil.AssertEq(models.CaseStatusList.NewCase, resp2.Case.Status, t)

	_, err = Execute(ctx, client, id, models.CaseActions.Start)
	testutil.AssertNilError(err, t)

	time.Sleep(2 * time.Second)

	resp3, err := FindCase(ctx, client, id)
	testutil.AssertNilError(err, t)
	testutil.AssertEq(id, resp3.Case.ID, t)
	testutil.AssertEq(models.CaseStatusList.Started, resp3.Case.Status, t)
}
