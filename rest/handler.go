package rest

import (
	"context"

	"github.com/case-management-suite/api/controllers"
	"github.com/case-management-suite/api/rest/restapi"
	"github.com/case-management-suite/api/rest/restapi/ops"
	"github.com/case-management-suite/models"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/zerolog/log"
)

func NewCaseMgmtAPI(controller controllers.CaseControllerAPI) *ops.CaseMgmtAPI {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Error().Err(err).Msg("Could not load Swagger definition")
	}
	api := ops.NewCaseMgmtAPI(swaggerSpec)

	api.FindCasesHandler = ops.FindCasesHandlerFunc(func(fcp ops.FindCasesParams) middleware.Responder {
		cases, err := controller.GetCases(models.NewCaseRecordSpec(true))

		if err != nil {
			log.Error().Err(err).Msg("Could not retrieve list of cases")
			return ops.NewFindCasesDefault(500).WithMessage(err.Error())
		}
		return ops.NewFindCasesOK().WithPayload(cases)
	})

	api.CreateCaseHandler = ops.CreateCaseHandlerFunc(func(ccp ops.CreateCaseParams) middleware.Responder {
		id, err := controller.NewCase(context.TODO())
		if err != nil {
			log.Error().Err(err).Msg("Could retrieve the requested case")
			return ops.NewCreateCaseDefault(500).WithMessage(err.Error())
		}
		return ops.NewCreateCaseOK().WithPayload(models.UUIDResponse{UUID: id})
	})

	api.FindCaseHandler = ops.FindCaseHandlerFunc(func(fcp ops.FindCaseParams) middleware.Responder {
		case_record, err := controller.FindCase(fcp.ID, models.NewCaseRecordSpec(true))

		if err != nil {
			log.Error().Err(err).Msg("Could not retrieve list of cases")
			return ops.NewFindCaseDefault(500).WithMessage(err.Error())
		}
		return ops.NewFindCaseOK().WithPayload(case_record)
	})

	api.ExecuteActionHandler = ops.ExecuteActionHandlerFunc(func(eap ops.ExecuteActionParams) middleware.Responder {
		result, err := controller.ExecuteAction(eap.ID, eap.Action, context.TODO())

		if err != nil {
			log.Error().Err(err).Msg("Could not retrieve list of cases")
			return ops.NewExecuteActionDefault(500).WithMessage(err.Error())
		}
		return ops.NewExecuteActionOK().WithPayload(models.UUIDResponse{UUID: result})
	})

	api.GetActionRecordsHandler = ops.GetActionRecordsHandlerFunc(func(garp ops.GetActionRecordsParams) middleware.Responder {
		result, err := controller.FindCaseActions(garp.ID, models.NewCaseActionSpec(true))

		if err != nil {
			log.Error().Err(err).Msg("Could not retrieve list of cases")
			return ops.NewGetActionRecordsDefault(500).WithMessage(err.Error())
		}
		return ops.NewGetActionRecordsOK().WithPayload(result)
	})

	return api
}
