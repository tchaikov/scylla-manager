// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/rclone/models"
)

// OperationsDeletefileReader is a Reader for the OperationsDeletefile structure.
type OperationsDeletefileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperationsDeletefileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOperationsDeletefileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewOperationsDeletefileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewOperationsDeletefileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOperationsDeletefileOK creates a OperationsDeletefileOK with default headers values
func NewOperationsDeletefileOK() *OperationsDeletefileOK {
	return &OperationsDeletefileOK{}
}

/*OperationsDeletefileOK handles this case with default header values.

Job ID
*/
type OperationsDeletefileOK struct {
	Payload *models.Jobid
}

func (o *OperationsDeletefileOK) Error() string {
	return fmt.Sprintf("[POST /operations/deletefile][%d] operationsDeletefileOK  %+v", 200, o.Payload)
}

func (o *OperationsDeletefileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Jobid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperationsDeletefileNotFound creates a OperationsDeletefileNotFound with default headers values
func NewOperationsDeletefileNotFound() *OperationsDeletefileNotFound {
	return &OperationsDeletefileNotFound{}
}

/*OperationsDeletefileNotFound handles this case with default header values.

Not found
*/
type OperationsDeletefileNotFound struct {
	Payload *models.ErrorResponse
}

func (o *OperationsDeletefileNotFound) Error() string {
	return fmt.Sprintf("[POST /operations/deletefile][%d] operationsDeletefileNotFound  %+v", 404, o.Payload)
}

func (o *OperationsDeletefileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperationsDeletefileInternalServerError creates a OperationsDeletefileInternalServerError with default headers values
func NewOperationsDeletefileInternalServerError() *OperationsDeletefileInternalServerError {
	return &OperationsDeletefileInternalServerError{}
}

/*OperationsDeletefileInternalServerError handles this case with default header values.

Server error
*/
type OperationsDeletefileInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *OperationsDeletefileInternalServerError) Error() string {
	return fmt.Sprintf("[POST /operations/deletefile][%d] operationsDeletefileInternalServerError  %+v", 500, o.Payload)
}

func (o *OperationsDeletefileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}