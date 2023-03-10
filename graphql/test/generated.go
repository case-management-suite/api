// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// ExecuteResponse is returned by Execute on success.
type ExecuteResponse struct {
	ExecuteAction string `json:"ExecuteAction"`
}

// GetExecuteAction returns ExecuteResponse.ExecuteAction, and is useful for accessing the field via an interface.
func (v *ExecuteResponse) GetExecuteAction() string { return v.ExecuteAction }

// FindCaseCaseCaseRecord includes the requested fields of the GraphQL type CaseRecord.
type FindCaseCaseCaseRecord struct {
	ID        string    `json:"ID"`
	Status    string    `json:"Status"`
	CreatedAt time.Time `json:"CreatedAt"`
}

// GetID returns FindCaseCaseCaseRecord.ID, and is useful for accessing the field via an interface.
func (v *FindCaseCaseCaseRecord) GetID() string { return v.ID }

// GetStatus returns FindCaseCaseCaseRecord.Status, and is useful for accessing the field via an interface.
func (v *FindCaseCaseCaseRecord) GetStatus() string { return v.Status }

// GetCreatedAt returns FindCaseCaseCaseRecord.CreatedAt, and is useful for accessing the field via an interface.
func (v *FindCaseCaseCaseRecord) GetCreatedAt() time.Time { return v.CreatedAt }

// FindCaseResponse is returned by FindCase on success.
type FindCaseResponse struct {
	Case FindCaseCaseCaseRecord `json:"case"`
}

// GetCase returns FindCaseResponse.Case, and is useful for accessing the field via an interface.
func (v *FindCaseResponse) GetCase() FindCaseCaseCaseRecord { return v.Case }

// GetCasesCasesCaseRecord includes the requested fields of the GraphQL type CaseRecord.
type GetCasesCasesCaseRecord struct {
	ID        string    `json:"ID"`
	Status    string    `json:"Status"`
	CreatedAt time.Time `json:"CreatedAt"`
}

// GetID returns GetCasesCasesCaseRecord.ID, and is useful for accessing the field via an interface.
func (v *GetCasesCasesCaseRecord) GetID() string { return v.ID }

// GetStatus returns GetCasesCasesCaseRecord.Status, and is useful for accessing the field via an interface.
func (v *GetCasesCasesCaseRecord) GetStatus() string { return v.Status }

// GetCreatedAt returns GetCasesCasesCaseRecord.CreatedAt, and is useful for accessing the field via an interface.
func (v *GetCasesCasesCaseRecord) GetCreatedAt() time.Time { return v.CreatedAt }

// GetCasesResponse is returned by GetCases on success.
type GetCasesResponse struct {
	Cases []GetCasesCasesCaseRecord `json:"cases"`
}

// GetCases returns GetCasesResponse.Cases, and is useful for accessing the field via an interface.
func (v *GetCasesResponse) GetCases() []GetCasesCasesCaseRecord { return v.Cases }

// NewCaseResponse is returned by NewCase on success.
type NewCaseResponse struct {
	NewCase string `json:"NewCase"`
}

// GetNewCase returns NewCaseResponse.NewCase, and is useful for accessing the field via an interface.
func (v *NewCaseResponse) GetNewCase() string { return v.NewCase }

// __ExecuteInput is used internally by genqlient
type __ExecuteInput struct {
	Case_id string `json:"case_id"`
	Action  string `json:"action"`
}

// GetCase_id returns __ExecuteInput.Case_id, and is useful for accessing the field via an interface.
func (v *__ExecuteInput) GetCase_id() string { return v.Case_id }

// GetAction returns __ExecuteInput.Action, and is useful for accessing the field via an interface.
func (v *__ExecuteInput) GetAction() string { return v.Action }

// __FindCaseInput is used internally by genqlient
type __FindCaseInput struct {
	Case_id string `json:"case_id"`
}

// GetCase_id returns __FindCaseInput.Case_id, and is useful for accessing the field via an interface.
func (v *__FindCaseInput) GetCase_id() string { return v.Case_id }

func Execute(
	ctx context.Context,
	client graphql.Client,
	case_id string,
	action string,
) (*ExecuteResponse, error) {
	req := &graphql.Request{
		OpName: "Execute",
		Query: `
mutation Execute ($case_id: String!, $action: String!) {
	ExecuteAction(id: $case_id, action: $action)
}
`,
		Variables: &__ExecuteInput{
			Case_id: case_id,
			Action:  action,
		},
	}
	var err error

	var data ExecuteResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func FindCase(
	ctx context.Context,
	client graphql.Client,
	case_id string,
) (*FindCaseResponse, error) {
	req := &graphql.Request{
		OpName: "FindCase",
		Query: `
query FindCase ($case_id: String!) {
	case(id: $case_id) {
		ID
		Status
		CreatedAt
	}
}
`,
		Variables: &__FindCaseInput{
			Case_id: case_id,
		},
	}
	var err error

	var data FindCaseResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetCases(
	ctx context.Context,
	client graphql.Client,
) (*GetCasesResponse, error) {
	req := &graphql.Request{
		OpName: "GetCases",
		Query: `
query GetCases {
	cases {
		ID
		Status
		CreatedAt
	}
}
`,
	}
	var err error

	var data GetCasesResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func NewCase(
	ctx context.Context,
	client graphql.Client,
) (*NewCaseResponse, error) {
	req := &graphql.Request{
		OpName: "NewCase",
		Query: `
mutation NewCase {
	NewCase
}
`,
	}
	var err error

	var data NewCaseResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
