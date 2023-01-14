package controllers

import (
	"context"

	"github.com/case-management-suite/models"
)

type CaseControllerAPI interface {
	NewCase(context context.Context) (string, error)
	GetCases(models.CaseRecordSpec) ([]models.CaseRecord, error)
	FindCase(id string, spec models.CaseRecordSpec) (models.CaseRecord, error)
	ExecuteAction(id string, action string, context context.Context) (string, error)
	FindCaseActions(caseId models.Identifier, spec models.CaseActionSpec) ([]models.CaseAction, error)
}
