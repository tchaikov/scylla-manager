// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCacheServiceKeyCacheKeysToSavePostParams creates a new CacheServiceKeyCacheKeysToSavePostParams object
// with the default values initialized.
func NewCacheServiceKeyCacheKeysToSavePostParams() *CacheServiceKeyCacheKeysToSavePostParams {
	var ()
	return &CacheServiceKeyCacheKeysToSavePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceKeyCacheKeysToSavePostParamsWithTimeout creates a new CacheServiceKeyCacheKeysToSavePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceKeyCacheKeysToSavePostParamsWithTimeout(timeout time.Duration) *CacheServiceKeyCacheKeysToSavePostParams {
	var ()
	return &CacheServiceKeyCacheKeysToSavePostParams{

		timeout: timeout,
	}
}

// NewCacheServiceKeyCacheKeysToSavePostParamsWithContext creates a new CacheServiceKeyCacheKeysToSavePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceKeyCacheKeysToSavePostParamsWithContext(ctx context.Context) *CacheServiceKeyCacheKeysToSavePostParams {
	var ()
	return &CacheServiceKeyCacheKeysToSavePostParams{

		Context: ctx,
	}
}

// NewCacheServiceKeyCacheKeysToSavePostParamsWithHTTPClient creates a new CacheServiceKeyCacheKeysToSavePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceKeyCacheKeysToSavePostParamsWithHTTPClient(client *http.Client) *CacheServiceKeyCacheKeysToSavePostParams {
	var ()
	return &CacheServiceKeyCacheKeysToSavePostParams{
		HTTPClient: client,
	}
}

/*CacheServiceKeyCacheKeysToSavePostParams contains all the parameters to send to the API endpoint
for the cache service key cache keys to save post operation typically these are written to a http.Request
*/
type CacheServiceKeyCacheKeysToSavePostParams struct {

	/*Kckts
	  key cache keys to save

	*/
	Kckts int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service key cache keys to save post params
func (o *CacheServiceKeyCacheKeysToSavePostParams) WithTimeout(timeout time.Duration) *CacheServiceKeyCacheKeysToSavePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service key cache keys to save post params
func (o *CacheServiceKeyCacheKeysToSavePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service key cache keys to save post params
func (o *CacheServiceKeyCacheKeysToSavePostParams) WithContext(ctx context.Context) *CacheServiceKeyCacheKeysToSavePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service key cache keys to save post params
func (o *CacheServiceKeyCacheKeysToSavePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service key cache keys to save post params
func (o *CacheServiceKeyCacheKeysToSavePostParams) WithHTTPClient(client *http.Client) *CacheServiceKeyCacheKeysToSavePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service key cache keys to save post params
func (o *CacheServiceKeyCacheKeysToSavePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKckts adds the kckts to the cache service key cache keys to save post params
func (o *CacheServiceKeyCacheKeysToSavePostParams) WithKckts(kckts int32) *CacheServiceKeyCacheKeysToSavePostParams {
	o.SetKckts(kckts)
	return o
}

// SetKckts adds the kckts to the cache service key cache keys to save post params
func (o *CacheServiceKeyCacheKeysToSavePostParams) SetKckts(kckts int32) {
	o.Kckts = kckts
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceKeyCacheKeysToSavePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param kckts
	qrKckts := o.Kckts
	qKckts := swag.FormatInt32(qrKckts)
	if qKckts != "" {
		if err := r.SetQueryParam("kckts", qKckts); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}