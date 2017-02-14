package bootenvs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/rackn/rocket-skates/models"
)

// PatchBootenvHandlerFunc turns a function with the right signature into a patch bootenv handler
type PatchBootenvHandlerFunc func(PatchBootenvParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchBootenvHandlerFunc) Handle(params PatchBootenvParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PatchBootenvHandler interface for that can handle valid patch bootenv params
type PatchBootenvHandler interface {
	Handle(PatchBootenvParams, *models.Principal) middleware.Responder
}

// NewPatchBootenv creates a new http.Handler for the patch bootenv operation
func NewPatchBootenv(ctx *middleware.Context, handler PatchBootenvHandler) *PatchBootenv {
	return &PatchBootenv{Context: ctx, Handler: handler}
}

/*PatchBootenv swagger:route PATCH /bootenvs/{name} Bootenvs patchBootenv

Patch Bootenv

*/
type PatchBootenv struct {
	Context *middleware.Context
	Handler PatchBootenvHandler
}

func (o *PatchBootenv) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPatchBootenvParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
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
