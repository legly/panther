// Code generated by go-swagger; DO NOT EDIT.

package operations

/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

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

	"github.com/panther-labs/panther/api/gateway/analysis/models"
)

// NewCreateRuleParams creates a new CreateRuleParams object
// with the default values initialized.
func NewCreateRuleParams() *CreateRuleParams {
	var ()
	return &CreateRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRuleParamsWithTimeout creates a new CreateRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRuleParamsWithTimeout(timeout time.Duration) *CreateRuleParams {
	var ()
	return &CreateRuleParams{

		timeout: timeout,
	}
}

// NewCreateRuleParamsWithContext creates a new CreateRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRuleParamsWithContext(ctx context.Context) *CreateRuleParams {
	var ()
	return &CreateRuleParams{

		Context: ctx,
	}
}

// NewCreateRuleParamsWithHTTPClient creates a new CreateRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRuleParamsWithHTTPClient(client *http.Client) *CreateRuleParams {
	var ()
	return &CreateRuleParams{
		HTTPClient: client,
	}
}

/*CreateRuleParams contains all the parameters to send to the API endpoint
for the create rule operation typically these are written to a http.Request
*/
type CreateRuleParams struct {

	/*Body*/
	Body *models.UpdateRule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create rule params
func (o *CreateRuleParams) WithTimeout(timeout time.Duration) *CreateRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create rule params
func (o *CreateRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create rule params
func (o *CreateRuleParams) WithContext(ctx context.Context) *CreateRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create rule params
func (o *CreateRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create rule params
func (o *CreateRuleParams) WithHTTPClient(client *http.Client) *CreateRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create rule params
func (o *CreateRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create rule params
func (o *CreateRuleParams) WithBody(body *models.UpdateRule) *CreateRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create rule params
func (o *CreateRuleParams) SetBody(body *models.UpdateRule) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}