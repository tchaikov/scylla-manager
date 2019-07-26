// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/rclone/models"
)

// NewCoreStatsParams creates a new CoreStatsParams object
// with the default values initialized.
func NewCoreStatsParams() *CoreStatsParams {
	var ()
	return &CoreStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCoreStatsParamsWithTimeout creates a new CoreStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCoreStatsParamsWithTimeout(timeout time.Duration) *CoreStatsParams {
	var ()
	return &CoreStatsParams{

		timeout: timeout,
	}
}

// NewCoreStatsParamsWithContext creates a new CoreStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCoreStatsParamsWithContext(ctx context.Context) *CoreStatsParams {
	var ()
	return &CoreStatsParams{

		Context: ctx,
	}
}

// NewCoreStatsParamsWithHTTPClient creates a new CoreStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCoreStatsParamsWithHTTPClient(client *http.Client) *CoreStatsParams {
	var ()
	return &CoreStatsParams{
		HTTPClient: client,
	}
}

/*CoreStatsParams contains all the parameters to send to the API endpoint
for the core stats operation typically these are written to a http.Request
*/
type CoreStatsParams struct {

	/*StatsParams
	  Stats parameters

	*/
	StatsParams *models.StatsParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the core stats params
func (o *CoreStatsParams) WithTimeout(timeout time.Duration) *CoreStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the core stats params
func (o *CoreStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the core stats params
func (o *CoreStatsParams) WithContext(ctx context.Context) *CoreStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the core stats params
func (o *CoreStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the core stats params
func (o *CoreStatsParams) WithHTTPClient(client *http.Client) *CoreStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the core stats params
func (o *CoreStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStatsParams adds the statsParams to the core stats params
func (o *CoreStatsParams) WithStatsParams(statsParams *models.StatsParams) *CoreStatsParams {
	o.SetStatsParams(statsParams)
	return o
}

// SetStatsParams adds the statsParams to the core stats params
func (o *CoreStatsParams) SetStatsParams(statsParams *models.StatsParams) {
	o.StatsParams = statsParams
}

// WriteToRequest writes these params to a swagger request
func (o *CoreStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.StatsParams != nil {
		if err := r.SetBodyParam(o.StatsParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}