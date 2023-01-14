// Package classification Petstore API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//	Schemes: http
//	Host: localhost:8080
//	Version: 0.0.1
//	License: MIT http://opensource.org/licenses/MIT
//	Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Extensions:
//	x-meta-value: value
//	x-meta-array:
//	  - value1
//	  - value2
//	x-meta-array-obj:
//	  - name: obj
//	    value: field
//
// swagger:meta
package rest

import (
	"github.com/case-management-suite/models"
)

type CaseServiceREST interface {
	// swagger:route PUT /case createCase
	//
	// # Create a new case
	//
	// This will show all available cases
	//
	//	Consumes:
	//	- application/json
	//
	//	Produces:
	//	- application/json
	//
	// Schemes: http
	//
	// Responses:
	//
	//	default: error
	//	200: UUIDResponse
	NewCase() (models.Identifier, error)

	// swagger:route GET /case findCases
	//
	// # Find all cases for the user
	//
	// This will show all available cases
	//
	//	Consumes:
	//	- application/json
	//
	//	Produces:
	//	- application/json
	//
	// Schemes: http
	//
	// Responses:
	//
	//	default: error
	//	200: []CaseRecord
	GetCases(spec models.CaseRecordSpec) ([]models.CaseRecord, error)

	// swagger:route GET /case/{id} findCase
	//
	// # Find a given case
	//
	// This will show the requested case
	//
	//		Consumes:
	//		- application/json
	//
	//		Produces:
	//		- application/json
	//		Schemes: http
	//
	//
	//
	//	    Parameters:
	//	      + name: id
	//	        in: path
	//	        description: case ID
	//	        required: true
	//	        type: string
	//
	//		Responses:
	//		  default: error
	//		  200: CaseRecord
	FindCase(id string, spec models.CaseRecordSpec) (models.CaseRecord, error)

	// swagger:route POST /case/{id} executeAction
	//
	// # Execute an action on the given case
	//
	// This will execute a case
	//
	//		Consumes:
	//		- application/json
	//
	//		Produces:
	//		- application/json
	//		Schemes: http
	//
	//
	//
	//	    Parameters:
	//	      + name: id
	//	        in: path
	//	        description: case ID
	//	        required: true
	//	        type: string
	//		  + name: action
	//	        in: query
	//	        description: action
	//	        required: true
	//	        type: string
	//
	//		Responses:
	//		  default: error
	//		  200: UUIDResponse
	ExecuteAction(id string, action string) (models.CaseRecord, error)

	// swagger:route GET /case/{id}/actions getActionRecords
	//
	// # Find the action records for a given case
	//
	// This will retrieve actions
	//
	//		Consumes:
	//		- application/json
	//
	//		Produces:
	//		- application/json
	//		Schemes: http
	//
	//
	//
	//	    Parameters:
	//	      + name: id
	//	        in: path
	//	        description: case ID
	//	        required: true
	//	        type: string
	//
	//		Responses:
	//		  default: error
	//		  200: []CaseAction
	GetActionRecords(caseId models.Identifier, spec models.CaseActionSpec) ([]models.CaseAction, error)
}
