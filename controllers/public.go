package controllers

import (
	"github.com/case-management-suite/caseservice"
	"github.com/case-management-suite/rulesengineservice"
)

type ControllerFactory func(caseservice.CaseService, rulesengineservice.RulesServiceClient) CaseControllerAPI

func GetControllerFactory() ControllerFactory {
	return NewCaseControllerImpl
}
