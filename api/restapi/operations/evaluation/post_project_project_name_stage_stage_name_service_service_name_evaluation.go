// Code generated by go-swagger; DO NOT EDIT.

package evaluation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/keptn/keptn/api/models"
)

// PostProjectProjectNameStageStageNameServiceServiceNameEvaluationHandlerFunc turns a function with the right signature into a post project project name stage stage name service service name evaluation handler
type PostProjectProjectNameStageStageNameServiceServiceNameEvaluationHandlerFunc func(PostProjectProjectNameStageStageNameServiceServiceNameEvaluationParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostProjectProjectNameStageStageNameServiceServiceNameEvaluationHandlerFunc) Handle(params PostProjectProjectNameStageStageNameServiceServiceNameEvaluationParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostProjectProjectNameStageStageNameServiceServiceNameEvaluationHandler interface for that can handle valid post project project name stage stage name service service name evaluation params
type PostProjectProjectNameStageStageNameServiceServiceNameEvaluationHandler interface {
	Handle(PostProjectProjectNameStageStageNameServiceServiceNameEvaluationParams, *models.Principal) middleware.Responder
}

// NewPostProjectProjectNameStageStageNameServiceServiceNameEvaluation creates a new http.Handler for the post project project name stage stage name service service name evaluation operation
func NewPostProjectProjectNameStageStageNameServiceServiceNameEvaluation(ctx *middleware.Context, handler PostProjectProjectNameStageStageNameServiceServiceNameEvaluationHandler) *PostProjectProjectNameStageStageNameServiceServiceNameEvaluation {
	return &PostProjectProjectNameStageStageNameServiceServiceNameEvaluation{Context: ctx, Handler: handler}
}

/*PostProjectProjectNameStageStageNameServiceServiceNameEvaluation swagger:route POST /project/{projectName}/stage/{stageName}/service/{serviceName}/evaluation evaluation postProjectProjectNameStageStageNameServiceServiceNameEvaluation

Trigger a new evaluation

*/
type PostProjectProjectNameStageStageNameServiceServiceNameEvaluation struct {
	Context *middleware.Context
	Handler PostProjectProjectNameStageStageNameServiceServiceNameEvaluationHandler
}

func (o *PostProjectProjectNameStageStageNameServiceServiceNameEvaluation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostProjectProjectNameStageStageNameServiceServiceNameEvaluationParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}