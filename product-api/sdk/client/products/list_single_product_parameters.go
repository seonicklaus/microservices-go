// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListSingleProductParams creates a new ListSingleProductParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSingleProductParams() *ListSingleProductParams {
	return &ListSingleProductParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSingleProductParamsWithTimeout creates a new ListSingleProductParams object
// with the ability to set a timeout on a request.
func NewListSingleProductParamsWithTimeout(timeout time.Duration) *ListSingleProductParams {
	return &ListSingleProductParams{
		timeout: timeout,
	}
}

// NewListSingleProductParamsWithContext creates a new ListSingleProductParams object
// with the ability to set a context for a request.
func NewListSingleProductParamsWithContext(ctx context.Context) *ListSingleProductParams {
	return &ListSingleProductParams{
		Context: ctx,
	}
}

// NewListSingleProductParamsWithHTTPClient creates a new ListSingleProductParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSingleProductParamsWithHTTPClient(client *http.Client) *ListSingleProductParams {
	return &ListSingleProductParams{
		HTTPClient: client,
	}
}

/* ListSingleProductParams contains all the parameters to send to the API endpoint
   for the list single product operation.

   Typically these are written to a http.Request.
*/
type ListSingleProductParams struct {

	/* ID.

	   The ID of the product for which the operation relates

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list single product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSingleProductParams) WithDefaults() *ListSingleProductParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list single product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSingleProductParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list single product params
func (o *ListSingleProductParams) WithTimeout(timeout time.Duration) *ListSingleProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list single product params
func (o *ListSingleProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list single product params
func (o *ListSingleProductParams) WithContext(ctx context.Context) *ListSingleProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list single product params
func (o *ListSingleProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list single product params
func (o *ListSingleProductParams) WithHTTPClient(client *http.Client) *ListSingleProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list single product params
func (o *ListSingleProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list single product params
func (o *ListSingleProductParams) WithID(id int64) *ListSingleProductParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list single product params
func (o *ListSingleProductParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListSingleProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
