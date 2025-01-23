// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"my-go-project/client/models"
)

// AddProductReader is a Reader for the AddProduct structure.
type AddProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddProductCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewAddProductUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewAddProductNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /products] addProduct", response, response.Code())
	}
}

// NewAddProductCreated creates a AddProductCreated with default headers values
func NewAddProductCreated() *AddProductCreated {
	return &AddProductCreated{}
}

/*
AddProductCreated describes a response with status code 201, with default header values.

Data structure representing a single product
*/
type AddProductCreated struct {
	Payload *models.Product
}

// IsSuccess returns true when this add product created response has a 2xx status code
func (o *AddProductCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add product created response has a 3xx status code
func (o *AddProductCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add product created response has a 4xx status code
func (o *AddProductCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add product created response has a 5xx status code
func (o *AddProductCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add product created response a status code equal to that given
func (o *AddProductCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the add product created response
func (o *AddProductCreated) Code() int {
	return 201
}

func (o *AddProductCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /products][%d] addProductCreated %s", 201, payload)
}

func (o *AddProductCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /products][%d] addProductCreated %s", 201, payload)
}

func (o *AddProductCreated) GetPayload() *models.Product {
	return o.Payload
}

func (o *AddProductCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProductUnprocessableEntity creates a AddProductUnprocessableEntity with default headers values
func NewAddProductUnprocessableEntity() *AddProductUnprocessableEntity {
	return &AddProductUnprocessableEntity{}
}

/*
AddProductUnprocessableEntity describes a response with status code 422, with default header values.

Validation errors as an array of strings
*/
type AddProductUnprocessableEntity struct {
	Payload *AddProductUnprocessableEntityBody
}

// IsSuccess returns true when this add product unprocessable entity response has a 2xx status code
func (o *AddProductUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add product unprocessable entity response has a 3xx status code
func (o *AddProductUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add product unprocessable entity response has a 4xx status code
func (o *AddProductUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this add product unprocessable entity response has a 5xx status code
func (o *AddProductUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this add product unprocessable entity response a status code equal to that given
func (o *AddProductUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the add product unprocessable entity response
func (o *AddProductUnprocessableEntity) Code() int {
	return 422
}

func (o *AddProductUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /products][%d] addProductUnprocessableEntity %s", 422, payload)
}

func (o *AddProductUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /products][%d] addProductUnprocessableEntity %s", 422, payload)
}

func (o *AddProductUnprocessableEntity) GetPayload() *AddProductUnprocessableEntityBody {
	return o.Payload
}

func (o *AddProductUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddProductUnprocessableEntityBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProductNotImplemented creates a AddProductNotImplemented with default headers values
func NewAddProductNotImplemented() *AddProductNotImplemented {
	return &AddProductNotImplemented{}
}

/*
AddProductNotImplemented describes a response with status code 501, with default header values.

Generic error message as a string
*/
type AddProductNotImplemented struct {
	Payload *AddProductNotImplementedBody
}

// IsSuccess returns true when this add product not implemented response has a 2xx status code
func (o *AddProductNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add product not implemented response has a 3xx status code
func (o *AddProductNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add product not implemented response has a 4xx status code
func (o *AddProductNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this add product not implemented response has a 5xx status code
func (o *AddProductNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this add product not implemented response a status code equal to that given
func (o *AddProductNotImplemented) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the add product not implemented response
func (o *AddProductNotImplemented) Code() int {
	return 501
}

func (o *AddProductNotImplemented) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /products][%d] addProductNotImplemented %s", 501, payload)
}

func (o *AddProductNotImplemented) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /products][%d] addProductNotImplemented %s", 501, payload)
}

func (o *AddProductNotImplemented) GetPayload() *AddProductNotImplementedBody {
	return o.Payload
}

func (o *AddProductNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddProductNotImplementedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AddProductNotImplementedBody add product not implemented body
swagger:model AddProductNotImplementedBody
*/
type AddProductNotImplementedBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add product not implemented body
func (o *AddProductNotImplementedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add product not implemented body based on context it is used
func (o *AddProductNotImplementedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProductNotImplementedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProductNotImplementedBody) UnmarshalBinary(b []byte) error {
	var res AddProductNotImplementedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AddProductUnprocessableEntityBody add product unprocessable entity body
swagger:model AddProductUnprocessableEntityBody
*/
type AddProductUnprocessableEntityBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add product unprocessable entity body
func (o *AddProductUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add product unprocessable entity body based on context it is used
func (o *AddProductUnprocessableEntityBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProductUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProductUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res AddProductUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
