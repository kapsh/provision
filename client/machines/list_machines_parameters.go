package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListMachinesParams creates a new ListMachinesParams object
// with the default values initialized.
func NewListMachinesParams() *ListMachinesParams {

	return &ListMachinesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListMachinesParamsWithTimeout creates a new ListMachinesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListMachinesParamsWithTimeout(timeout time.Duration) *ListMachinesParams {

	return &ListMachinesParams{

		timeout: timeout,
	}
}

// NewListMachinesParamsWithContext creates a new ListMachinesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListMachinesParamsWithContext(ctx context.Context) *ListMachinesParams {

	return &ListMachinesParams{

		Context: ctx,
	}
}

// NewListMachinesParamsWithHTTPClient creates a new ListMachinesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListMachinesParamsWithHTTPClient(client *http.Client) *ListMachinesParams {

	return &ListMachinesParams{
		HTTPClient: client,
	}
}

/*ListMachinesParams contains all the parameters to send to the API endpoint
for the list machines operation typically these are written to a http.Request
*/
type ListMachinesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list machines params
func (o *ListMachinesParams) WithTimeout(timeout time.Duration) *ListMachinesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list machines params
func (o *ListMachinesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list machines params
func (o *ListMachinesParams) WithContext(ctx context.Context) *ListMachinesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list machines params
func (o *ListMachinesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list machines params
func (o *ListMachinesParams) WithHTTPClient(client *http.Client) *ListMachinesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list machines params
func (o *ListMachinesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListMachinesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
