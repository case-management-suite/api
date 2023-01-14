package controllers

import (
	"context"

	"github.com/case-management-suite/caseservice"
	"github.com/case-management-suite/models"
	"github.com/case-management-suite/rulesengineservice"
	"github.com/rs/zerolog/log"
)

type CaseControllerImpl struct {
	CaseService caseservice.CaseService
	Rules       rulesengineservice.RulesServiceClient
}

func NewCaseControllerImpl(caseService caseservice.CaseService, Rules rulesengineservice.RulesServiceClient) CaseControllerAPI {
	return &CaseControllerImpl{CaseService: caseService, Rules: Rules}
}

func (c *CaseControllerImpl) NewCase(context context.Context) (string, error) {
	id, err := c.CaseService.NewCase()
	if err != nil {
		return "", err
	}

	log.Debug().Str("UUID_IN_CONTROLLER", id)

	return id, nil
}

func (c *CaseControllerImpl) GetCases(spec models.CaseRecordSpec) ([]models.CaseRecord, error) {
	return c.CaseService.GetCases(spec)
}

func (c *CaseControllerImpl) FindCase(id string, spec models.CaseRecordSpec) (models.CaseRecord, error) {
	caser, err := c.CaseService.FindCase(id, spec)
	return caser, err
}

func (c *CaseControllerImpl) ExecuteAction(id string, action string, context context.Context) (string, error) {
	record, err := c.CaseService.FindCase(id, models.NewCaseRecordSpec(true))
	if err != nil {
		return "ACK", err
	}
	err = c.Rules.ExecuteAction(record, action, context)
	if err != nil {
		return "ACK", err
	}

	uuid, err := c.CaseService.SaveCaseAction(models.CaseAction{ID: id, CaseRecord: record, Action: action})

	if err != nil {
		log.Debug().Err(err).Str("action", action).Str("CaseActionUUID", id).Msg("Failed to save the case context")
		return "", err
	}

	return uuid, nil
}

func (c *CaseControllerImpl) FindCaseActions(caseId models.Identifier, spec models.CaseActionSpec) ([]models.CaseAction, error) {
	return c.CaseService.GetActionRecords(caseId, models.NewCaseActionSpec(true))
}
