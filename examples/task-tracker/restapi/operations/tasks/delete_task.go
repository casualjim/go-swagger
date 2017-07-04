package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteTaskHandlerFunc turns a function with the right signature into a delete task handler
type DeleteTaskHandlerFunc func(DeleteTaskParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTaskHandlerFunc) Handle(params DeleteTaskParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteTaskHandler interface for that can handle valid delete task params
type DeleteTaskHandler interface {
	Handle(DeleteTaskParams, interface{}) middleware.Responder
}

// NewDeleteTask creates a new http.Handler for the delete task operation
func NewDeleteTask(ctx *middleware.Context, handler DeleteTaskHandler) *DeleteTask {
	return &DeleteTask{Context: ctx, Handler: handler}
}

/*DeleteTask swagger:route DELETE /tasks/{id} tasks deleteTask

Deletes a task.

This is a soft delete and changes the task status to ignored.


*/
type DeleteTask struct {
	Context *middleware.Context
	Handler DeleteTaskHandler
}

func (o *DeleteTask) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteTaskParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
