package frontend_test

import (
	"context"
	"fmt"
	"os"

	"strings"
	"testing"
	"time"

	"github.com/case-management-suite/api/rest"
	"github.com/case-management-suite/caseservice"
	common "github.com/case-management-suite/common/config"
	"github.com/case-management-suite/models"
	"github.com/case-management-suite/rulesengineservice"
	mocks "github.com/case-management-suite/testutil"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestServerActions(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false})
	log.Debug().Msg("STARTING INTEGRATION TEST")
	appConfig := common.NewLocalAppConfig()
	appConfig.CasesStorage.LogSQL = true

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
	app := rest.CreateLiteTestAPIServer(appConfig)
	if err := app.Err(); err != nil {
		log.Fatal().Err(err)
		t.FailNow()
	}
	// In a typical application, we could just use app.Run() here. Since we
	// don't want this example to run forever, we'll use the more-explicit Start
	// and Stop.
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := casesService.Start(startCtx); err != nil {
		log.Fatal().Err(err)
	}

	if err := rulesService.Start(startCtx); err != nil {
		log.Fatal().Err(err)
	}

	if err := app.Start(startCtx); err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Msg("INITIALIZATION COMPLETE")

	t.Run("TestNewCase", testNewCase)
	t.Run("TestStartCase", testStartCase)

	t.Run("TestCloseCase", testCloseCase)
	t.Run("TestGetCaseActions", testGetCaseActions)

	t.Run("TestMoveCaseToEvaluation", testMoveCaseToEvaluation)

	t.Cleanup(func() {
		log.Debug().Msg("ENDING INTEGRATION TEST")
		stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := app.Stop(stopCtx); err != nil {
			log.Debug().Err(err).Msg("Failed to stop REST server")
		}
		if err := casesService.Stop(stopCtx); err != nil {
			log.Debug().Err(err).Msg("Failed to stop Cases Service")
		}
		if err := rulesService.Stop(stopCtx); err != nil {
			log.Debug().Err(err).Msg("Failed to stop Rules Service server")
		}

		os.Remove(appConfig.CasesStorage.Address)
		log.Debug().Msg("CLEANUP COMPLETE")
	})
}

func testNewCase(t *testing.T) {
	// t.Parallel()
	resp, err := mocks.PutRequest("http://localhost:8080/case", strings.NewReader("any thing"))
	HandleError(err, "Failed to create a new case", t)
	respString := string(resp)

	if len(respString) <= 0 {
		t.Fatalf("The response should not be empty")
	}

	resp2, err := mocks.GetRequest("http://localhost:8080/case")
	HandleError(err, "Failed to retrieve the new case", t)
	respString2 := string(resp2)
	if len(respString2) <= 0 {
		t.Fatalf("The response should not be empty")
	}

}

func HandleError(err error, msg string, t *testing.T) {
	if err != nil {
		log.Error().Err(err).Stack().Msg(msg)
		t.FailNow()
	}
}

func FailTest(msg string, t *testing.T) {
	log.Error().Msg(msg)
	t.FailNow()
}

func WithRetry[T any](retries int, cond func(T) bool, fn func() (T, error)) (T, error) {
	for retries > 0 {
		result, _ := fn()
		// HandleError(err, "Finding the record for the executed action should not fail", t)

		if cond(result) {
			break
		}
		time.Sleep(1 * time.Second)
		retries -= 1
	}
	return fn()
}

func FindById(id string) func() (models.CaseRecord, error) {
	return func() (models.CaseRecord, error) {
		rec, err := mocks.FindCase(id)
		return rec, err
	}
}

func ByStatus(status string) func(models.CaseRecord) bool {
	return func(rec models.CaseRecord) bool {
		return rec.Status == status
	}
}

func FindByIdWithRetries(retries int, id string, cond func(models.CaseRecord) bool) (models.CaseRecord, error) {
	return WithRetry(retries, cond, FindById(id))
}

func testStartCase(t *testing.T) {
	// t.Parallel()
	response, err := mocks.CreateCase()
	HandleError(err, "Create case should not fail", t)

	log.Debug().Str("UUID", response.UUID).Msg("Received UUID from new case:")

	if len(response.UUID) == 0 {
		t.Fatalf("The case ID should not be empty")
	}

	caseRecord, err := mocks.FindCase(response.UUID)
	HandleError(err, "Finding the record for the executed action should not fail", t)

	if caseRecord.ID != response.UUID {
		FailTest("The response ID should match the requested ID", t)
	}

	_, err = mocks.ExecuteAction(response.UUID, models.CaseActions.Start)
	HandleError(err, "Executing the action case should not fail", t)

	caseRecord, err = FindByIdWithRetries(4, response.UUID, ByStatus(models.CaseStatusList.Started))
	HandleError(err, "Finding the record for the executed action should not fail", t)

	if caseRecord.ID != response.UUID {
		FailTest("The response ID should match the requested ID", t)
	}

	if caseRecord.Status != models.CaseStatusList.Started {
		FailTest(fmt.Sprintf("The case status should have changed to Started, but was %v", caseRecord.Status), t)
	}
}

func testCloseCase(t *testing.T) {
	// t.Parallel()
	response, err := mocks.CreateCase()
	HandleError(err, "Create case should not fail", t)

	_, err = mocks.ExecuteAction(response.UUID, models.CaseActions.Close)
	HandleError(err, "Executing the action case should not fail", t)

	caseRecord, err := FindByIdWithRetries(4, response.UUID, ByStatus(models.CaseStatusList.Closed))
	HandleError(err, "Finding the record for the executed action should not fail", t)

	if caseRecord.ID != response.UUID {
		t.Fatalf("The response ID should match the requested ID")
	}

	if caseRecord.Status != models.CaseStatusList.Closed {
		t.Fatalf("The case status should have changed to closed, but was %v", caseRecord.Status)
	}

}

func testMoveCaseToEvaluation(t *testing.T) {
	// t.Parallel()
	response, err := mocks.CreateCase()
	HandleError(err, "Create case should not fail", t)

	_, err = mocks.ExecuteAction(response.UUID, "MOVE_TO_EVALUATION")
	HandleError(err, "Executing the action case should not fail", t)

	caseRecord, err := FindByIdWithRetries(4, response.UUID, ByStatus("IN_EVALUATION"))
	HandleError(err, "Finding the record for the executed action should not fail", t)

	if caseRecord.ID != response.UUID {
		t.Fatalf("The response ID should match the requested ID")
	}

	if caseRecord.Status != "IN_EVALUATION" {
		t.Fatalf("The case status should have changed to IN_EVALUATION, but was %v", caseRecord.Status)
	}

}

func testGetCaseActions(t *testing.T) {
	// t.Parallel()
	response, err := mocks.CreateCase()
	HandleError(err, "Create case should not fail", t)

	_, err = mocks.ExecuteAction(response.UUID, "START")
	HandleError(err, "Executing the action case should not fail", t)

	_, err = mocks.ExecuteAction(response.UUID, "MOVE_TO_EVALUATION")
	HandleError(err, "Executing the action case should not fail", t)

	_, err = mocks.ExecuteAction(response.UUID, "CLOSE")
	HandleError(err, "Executing the action case should not fail", t)

	caseRecord, err := FindByIdWithRetries(4, response.UUID, ByStatus("CLOSED"))
	HandleError(err, "failed to find the updated record", t)

	if caseRecord.ID != response.UUID {
		t.Fatalf("The response ID should match the requested ID")
	}

	if caseRecord.Status != "CLOSED" {
		t.Fatalf("The case status should have changed to CLOSED, but was %v", caseRecord.Status)
	}

	// caseActions, err := mocks.FindCaseActions(response.UUID)
	// HandleError(err, "Finding the case actions should not fail", t)

	// if len(caseActions) != 3 {
	// 	t.Logf("The number of actions is incorrect. Expected %v, but got %v", 3, len(caseActions))
	// }

	// for _, action := range caseActions {
	// 	if action.CaseRecordID != response.UUID {
	// 		t.Log("Case ID did not match")
	// 		t.FailNow()
	// 	}
	// }

}
