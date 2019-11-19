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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetClusterClusterIDBackupsParams creates a new GetClusterClusterIDBackupsParams object
// with the default values initialized.
func NewGetClusterClusterIDBackupsParams() *GetClusterClusterIDBackupsParams {
	var ()
	return &GetClusterClusterIDBackupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterClusterIDBackupsParamsWithTimeout creates a new GetClusterClusterIDBackupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterClusterIDBackupsParamsWithTimeout(timeout time.Duration) *GetClusterClusterIDBackupsParams {
	var ()
	return &GetClusterClusterIDBackupsParams{

		timeout: timeout,
	}
}

// NewGetClusterClusterIDBackupsParamsWithContext creates a new GetClusterClusterIDBackupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterClusterIDBackupsParamsWithContext(ctx context.Context) *GetClusterClusterIDBackupsParams {
	var ()
	return &GetClusterClusterIDBackupsParams{

		Context: ctx,
	}
}

// NewGetClusterClusterIDBackupsParamsWithHTTPClient creates a new GetClusterClusterIDBackupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterClusterIDBackupsParamsWithHTTPClient(client *http.Client) *GetClusterClusterIDBackupsParams {
	var ()
	return &GetClusterClusterIDBackupsParams{
		HTTPClient: client,
	}
}

/*GetClusterClusterIDBackupsParams contains all the parameters to send to the API endpoint
for the get cluster cluster ID backups operation typically these are written to a http.Request
*/
type GetClusterClusterIDBackupsParams struct {

	/*ClusterID*/
	ClusterID string
	/*ClusterID*/
	QueryClusterID *string
	/*Host*/
	Host string
	/*Keyspace*/
	Keyspace []string
	/*Locations*/
	Locations []string
	/*MaxDate*/
	MaxDate *strfmt.DateTime
	/*MinDate*/
	MinDate *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) WithTimeout(timeout time.Duration) *GetClusterClusterIDBackupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) WithContext(ctx context.Context) *GetClusterClusterIDBackupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) WithHTTPClient(client *http.Client) *GetClusterClusterIDBackupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) WithClusterID(clusterID string) *GetClusterClusterIDBackupsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithQueryClusterID adds the clusterID to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) WithQueryClusterID(clusterID *string) *GetClusterClusterIDBackupsParams {
	o.SetQueryClusterID(clusterID)
	return o
}

// SetQueryClusterID adds the clusterId to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) SetQueryClusterID(clusterID *string) {
	o.QueryClusterID = clusterID
}

// WithHost adds the host to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) WithHost(host string) *GetClusterClusterIDBackupsParams {
	o.SetHost(host)
	return o
}

// SetHost adds the host to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) SetHost(host string) {
	o.Host = host
}

// WithKeyspace adds the keyspace to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) WithKeyspace(keyspace []string) *GetClusterClusterIDBackupsParams {
	o.SetKeyspace(keyspace)
	return o
}

// SetKeyspace adds the keyspace to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) SetKeyspace(keyspace []string) {
	o.Keyspace = keyspace
}

// WithLocations adds the locations to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) WithLocations(locations []string) *GetClusterClusterIDBackupsParams {
	o.SetLocations(locations)
	return o
}

// SetLocations adds the locations to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) SetLocations(locations []string) {
	o.Locations = locations
}

// WithMaxDate adds the maxDate to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) WithMaxDate(maxDate *strfmt.DateTime) *GetClusterClusterIDBackupsParams {
	o.SetMaxDate(maxDate)
	return o
}

// SetMaxDate adds the maxDate to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) SetMaxDate(maxDate *strfmt.DateTime) {
	o.MaxDate = maxDate
}

// WithMinDate adds the minDate to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) WithMinDate(minDate *strfmt.DateTime) *GetClusterClusterIDBackupsParams {
	o.SetMinDate(minDate)
	return o
}

// SetMinDate adds the minDate to the get cluster cluster ID backups params
func (o *GetClusterClusterIDBackupsParams) SetMinDate(minDate *strfmt.DateTime) {
	o.MinDate = minDate
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterClusterIDBackupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.QueryClusterID != nil {

		// query param cluster_id
		var qrClusterID string
		if o.QueryClusterID != nil {
			qrClusterID = *o.QueryClusterID
		}
		qClusterID := qrClusterID
		if qClusterID != "" {
			if err := r.SetQueryParam("cluster_id", qClusterID); err != nil {
				return err
			}
		}

	}

	// query param host
	qrHost := o.Host
	qHost := qrHost
	if qHost != "" {
		if err := r.SetQueryParam("host", qHost); err != nil {
			return err
		}
	}

	valuesKeyspace := o.Keyspace

	joinedKeyspace := swag.JoinByFormat(valuesKeyspace, "")
	// query array param keyspace
	if err := r.SetQueryParam("keyspace", joinedKeyspace...); err != nil {
		return err
	}

	valuesLocations := o.Locations

	joinedLocations := swag.JoinByFormat(valuesLocations, "")
	// query array param locations
	if err := r.SetQueryParam("locations", joinedLocations...); err != nil {
		return err
	}

	if o.MaxDate != nil {

		// query param max_date
		var qrMaxDate strfmt.DateTime
		if o.MaxDate != nil {
			qrMaxDate = *o.MaxDate
		}
		qMaxDate := qrMaxDate.String()
		if qMaxDate != "" {
			if err := r.SetQueryParam("max_date", qMaxDate); err != nil {
				return err
			}
		}

	}

	if o.MinDate != nil {

		// query param min_date
		var qrMinDate strfmt.DateTime
		if o.MinDate != nil {
			qrMinDate = *o.MinDate
		}
		qMinDate := qrMinDate.String()
		if qMinDate != "" {
			if err := r.SetQueryParam("min_date", qMinDate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}