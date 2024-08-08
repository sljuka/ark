// Code generated by go-swagger; DO NOT EDIT.

package ark_service

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
)

// NewArkServiceListVtxosParams creates a new ArkServiceListVtxosParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewArkServiceListVtxosParams() *ArkServiceListVtxosParams {
	return &ArkServiceListVtxosParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewArkServiceListVtxosParamsWithTimeout creates a new ArkServiceListVtxosParams object
// with the ability to set a timeout on a request.
func NewArkServiceListVtxosParamsWithTimeout(timeout time.Duration) *ArkServiceListVtxosParams {
	return &ArkServiceListVtxosParams{
		timeout: timeout,
	}
}

// NewArkServiceListVtxosParamsWithContext creates a new ArkServiceListVtxosParams object
// with the ability to set a context for a request.
func NewArkServiceListVtxosParamsWithContext(ctx context.Context) *ArkServiceListVtxosParams {
	return &ArkServiceListVtxosParams{
		Context: ctx,
	}
}

// NewArkServiceListVtxosParamsWithHTTPClient creates a new ArkServiceListVtxosParams object
// with the ability to set a custom HTTPClient for a request.
func NewArkServiceListVtxosParamsWithHTTPClient(client *http.Client) *ArkServiceListVtxosParams {
	return &ArkServiceListVtxosParams{
		HTTPClient: client,
	}
}

/*
ArkServiceListVtxosParams contains all the parameters to send to the API endpoint

	for the ark service list vtxos operation.

	Typically these are written to a http.Request.
*/
type ArkServiceListVtxosParams struct {

	// Address.
	Address string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ark service list vtxos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArkServiceListVtxosParams) WithDefaults() *ArkServiceListVtxosParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ark service list vtxos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ArkServiceListVtxosParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ark service list vtxos params
func (o *ArkServiceListVtxosParams) WithTimeout(timeout time.Duration) *ArkServiceListVtxosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ark service list vtxos params
func (o *ArkServiceListVtxosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ark service list vtxos params
func (o *ArkServiceListVtxosParams) WithContext(ctx context.Context) *ArkServiceListVtxosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ark service list vtxos params
func (o *ArkServiceListVtxosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ark service list vtxos params
func (o *ArkServiceListVtxosParams) WithHTTPClient(client *http.Client) *ArkServiceListVtxosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ark service list vtxos params
func (o *ArkServiceListVtxosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the ark service list vtxos params
func (o *ArkServiceListVtxosParams) WithAddress(address string) *ArkServiceListVtxosParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the ark service list vtxos params
func (o *ArkServiceListVtxosParams) SetAddress(address string) {
	o.Address = address
}

// WriteToRequest writes these params to a swagger request
func (o *ArkServiceListVtxosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.Address); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}