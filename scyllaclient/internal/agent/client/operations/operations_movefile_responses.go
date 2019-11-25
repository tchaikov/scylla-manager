// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/scyllaclient/internal/agent/models"
)

// OperationsMovefileReader is a Reader for the OperationsMovefile structure.
type OperationsMovefileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperationsMovefileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperationsMovefileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewOperationsMovefileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOperationsMovefileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOperationsMovefileOK creates a OperationsMovefileOK with default headers values
func NewOperationsMovefileOK() *OperationsMovefileOK {
	return &OperationsMovefileOK{}
}

/*OperationsMovefileOK handles this case with default header values.

Job
*/
type OperationsMovefileOK struct {
	Payload *models.Jobid
}

func (o *OperationsMovefileOK) Error() string {
	return fmt.Sprintf("[POST /rclone/operations/movefile][%d] operationsMovefileOK  %+v", 200, o.Payload)
}

func (o *OperationsMovefileOK) GetPayload() *models.Jobid {
	return o.Payload
}

func (o *OperationsMovefileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Jobid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperationsMovefileNotFound creates a OperationsMovefileNotFound with default headers values
func NewOperationsMovefileNotFound() *OperationsMovefileNotFound {
	return &OperationsMovefileNotFound{}
}

/*OperationsMovefileNotFound handles this case with default header values.

Not found
*/
type OperationsMovefileNotFound struct {
	Payload *models.ErrorResponse
}

func (o *OperationsMovefileNotFound) Error() string {
	return fmt.Sprintf("[POST /rclone/operations/movefile][%d] operationsMovefileNotFound  %+v", 404, o.Payload)
}

func (o *OperationsMovefileNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *OperationsMovefileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperationsMovefileInternalServerError creates a OperationsMovefileInternalServerError with default headers values
func NewOperationsMovefileInternalServerError() *OperationsMovefileInternalServerError {
	return &OperationsMovefileInternalServerError{}
}

/*OperationsMovefileInternalServerError handles this case with default header values.

Server error
*/
type OperationsMovefileInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *OperationsMovefileInternalServerError) Error() string {
	return fmt.Sprintf("[POST /rclone/operations/movefile][%d] operationsMovefileInternalServerError  %+v", 500, o.Payload)
}

func (o *OperationsMovefileInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *OperationsMovefileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}