package rest_test

import (
	net "net/http"
	"testing"

	"github.com/case-management-suite/api/mocks"
	"github.com/case-management-suite/api/rest"
	"github.com/case-management-suite/api/rest/restapi/ops"
	"github.com/case-management-suite/common/logger"
	models "github.com/case-management-suite/models"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog/log"
)

func TestNewCaseAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	controller := mocks.NewMockCaseControllerAPI(ctrl)

	id := models.NewCaseRecordUUID()

	controller.EXPECT().GetCases(gomock.Any()).AnyTimes().Return([]models.CaseRecord{{ID: id, Status: "my status"}}, nil)

	casetMgmtAPI := rest.NewCaseMgmtAPI(controller, logger.NewTestLogger())

	responder := casetMgmtAPI.FindCasesHandler.Handle(ops.FindCasesParams{})

	if responder == nil {
		t.Fail()
	}
}

func TestStartCaseAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	controller := mocks.NewMockCaseControllerAPI(ctrl)

	id := models.NewCaseRecordUUID()

	controller.EXPECT().ExecuteAction(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(id, nil)

	casetMgmtAPI := rest.NewCaseMgmtAPI(controller, logger.NewTestLogger())

	// var responder cases_api.FindCasesDefault
	responder := casetMgmtAPI.ExecuteActionHandler.Handle(ops.ExecuteActionParams{Action: "START", ID: id})

	if responder == nil {
		t.Fail()
	}

	producer := mocks.NewMockProducer(ctrl)
	producer.EXPECT().Produce(gomock.Any(), gomock.Any())

	responder.WriteResponse(MockResponseWriter{}, producer)
}

func TestFindCaseAPI(t *testing.T) {
	ctrl := gomock.NewController(t)

	controller := mocks.NewMockCaseControllerAPI(ctrl)

	id := models.NewCaseRecordUUID()

	controller.EXPECT().FindCase(gomock.Eq(id), gomock.Any()).AnyTimes().Return(models.CaseRecord{ID: id, Status: "STARTED"}, nil)

	casetMgmtAPI := rest.NewCaseMgmtAPI(controller, logger.NewTestLogger())

	// var responder cases_api.FindCasesDefault
	responder := casetMgmtAPI.FindCaseHandler.Handle(ops.FindCaseParams{ID: id})

	if responder == nil {
		t.Fail()
	}

	producer := mocks.NewMockProducer(ctrl)
	producer.EXPECT().Produce(gomock.Any(), gomock.Any()).Do(func(_ interface{}, caseModel models.CaseRecord) {
		if caseModel.Status != "STARTED" {
			log.Error().Str("ACTION", caseModel.Status).Msg("The action should be started")
			t.Fail()
		}
	})

	responder.WriteResponse(MockResponseWriter{}, producer)
}

type MockResponseWriter struct{}

func (mrw MockResponseWriter) Header() net.Header {
	return map[string][]string{}
}

func (mrw MockResponseWriter) Write([]byte) (int, error) {
	return 1, nil
}

func (mrw MockResponseWriter) WriteHeader(statusCode int) {

}
