// Code generated by go-swagger; DO NOT EDIT.

package ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ExecuteActionHandlerFunc turns a function with the right signature into a execute action handler
type ExecuteActionHandlerFunc func(ExecuteActionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ExecuteActionHandlerFunc) Handle(params ExecuteActionParams) middleware.Responder {
	return fn(params)
}

// ExecuteActionHandler interface for that can handle valid execute action params
type ExecuteActionHandler interface {
	Handle(ExecuteActionParams) middleware.Responder
}

// NewExecuteAction creates a new http.Handler for the execute action operation
func NewExecuteAction(ctx *middleware.Context, handler ExecuteActionHandler) *ExecuteAction {
	return &ExecuteAction{Context: ctx, Handler: handler}
}

/*
	ExecuteAction swagger:route POST /case/{id} executeAction

# Execute an action on the given case

This will execute a case
*/
type ExecuteAction struct {
	Context *middleware.Context
	Handler ExecuteActionHandler
}

func (o *ExecuteAction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewExecuteActionParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}