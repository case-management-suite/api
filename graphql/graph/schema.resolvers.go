package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/case-management-suite/api/graphql/graph/model"
	"github.com/case-management-suite/models"
	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewCase is the resolver for the NewCase field.
func (r *mutationResolver) NewCase(ctx context.Context) (string, error) {
	return r.Controller.NewCase(ctx)
}

// ExecuteAction is the resolver for the ExecuteAction field.
func (r *mutationResolver) ExecuteAction(ctx context.Context, id string, action string) (string, error) {
	return r.Controller.ExecuteAction(id, action, ctx)
}

// Cases is the resolver for the cases field.
func (r *queryResolver) Cases(ctx context.Context) ([]*model.CaseRecord, error) {
	spec := GetCaseRecordSpec(ctx)
	log.Info().Msg(fmt.Sprintf("Spec: %v", spec.GetIncludedFields()))

	cases, err := r.Controller.GetCases(spec)
	if err != nil {
		return nil, err
	}

	casesJSON, err := json.Marshal(cases)
	if err != nil {
		return nil, err
	}

	response := []*model.CaseRecord{}

	err = json.Unmarshal(casesJSON, &response)

	return response, err
}

// Case is the resolver for the case field.
func (r *queryResolver) Case(ctx context.Context, id string) (*model.CaseRecord, error) {
	spec := GetCaseRecordSpec(ctx)
	log.Info().Msg(fmt.Sprintf("Spec: %v", spec.GetIncludedFields()))
	result, err := r.Controller.FindCase(id, spec)
	if err != nil {
		return nil, err
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	response := &model.CaseRecord{}
	err = json.Unmarshal(resultJson, &response)
	return response, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func GetCaseRecordSpec(ctx context.Context) models.CaseRecordSpec {
	reqCtx := graphql.GetOperationContext(ctx)
	fieldSelections := graphql.GetFieldContext(ctx).Field.Selections
	var sels []string

	spec := models.NewCaseRecordSpec(false)

	for _, sel := range fieldSelections {
		switch sel := sel.(type) {
		case *ast.Field:
			sels = append(sels, fmt.Sprintf("%s as %s", sel.Name, sel.Alias))
			err := spec.Set(sel.Name, true)
			if err != nil {
				log.Logger.Warn().Msg("Trying to set an invalid field in the spec")
			}
		case *ast.InlineFragment:
			sels = append(sels, fmt.Sprintf("inline fragment on %s", sel.TypeCondition))
		case *ast.FragmentSpread:
			fragment := reqCtx.Doc.Fragments.ForName(sel.Name)
			sels = append(sels, fmt.Sprintf("named fragment %s on %s", sel.Name, fragment.TypeCondition))
		}
	}
	return spec
}
