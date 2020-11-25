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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/panther-labs/panther/api/gateway/analysis/models"
)

// GetDataModelReader is a Reader for the GetDataModel structure.
type GetDataModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDataModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDataModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDataModelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDataModelOK creates a GetDataModelOK with default headers values
func NewGetDataModelOK() *GetDataModelOK {
	return &GetDataModelOK{}
}

/*GetDataModelOK handles this case with default header values.

OK
*/
type GetDataModelOK struct {
	Payload *models.DataModel
}

func (o *GetDataModelOK) Error() string {
	return fmt.Sprintf("[GET /datamodel][%d] getDataModelOK  %+v", 200, o.Payload)
}

func (o *GetDataModelOK) GetPayload() *models.DataModel {
	return o.Payload
}

func (o *GetDataModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataModelBadRequest creates a GetDataModelBadRequest with default headers values
func NewGetDataModelBadRequest() *GetDataModelBadRequest {
	return &GetDataModelBadRequest{}
}

/*GetDataModelBadRequest handles this case with default header values.

Bad request
*/
type GetDataModelBadRequest struct {
	Payload *models.Error
}

func (o *GetDataModelBadRequest) Error() string {
	return fmt.Sprintf("[GET /datamodel][%d] getDataModelBadRequest  %+v", 400, o.Payload)
}

func (o *GetDataModelBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDataModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataModelNotFound creates a GetDataModelNotFound with default headers values
func NewGetDataModelNotFound() *GetDataModelNotFound {
	return &GetDataModelNotFound{}
}

/*GetDataModelNotFound handles this case with default header values.

DataModel does not exist
*/
type GetDataModelNotFound struct {
}

func (o *GetDataModelNotFound) Error() string {
	return fmt.Sprintf("[GET /datamodel][%d] getDataModelNotFound ", 404)
}

func (o *GetDataModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDataModelInternalServerError creates a GetDataModelInternalServerError with default headers values
func NewGetDataModelInternalServerError() *GetDataModelInternalServerError {
	return &GetDataModelInternalServerError{}
}

/*GetDataModelInternalServerError handles this case with default header values.

Internal server error
*/
type GetDataModelInternalServerError struct {
}

func (o *GetDataModelInternalServerError) Error() string {
	return fmt.Sprintf("[GET /datamodel][%d] getDataModelInternalServerError ", 500)
}

func (o *GetDataModelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}