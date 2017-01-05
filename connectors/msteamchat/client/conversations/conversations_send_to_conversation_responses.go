package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/uva-its/gopherbot/connectors/msteamchat/models"
)

// ConversationsSendToConversationReader is a Reader for the ConversationsSendToConversation structure.
type ConversationsSendToConversationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConversationsSendToConversationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewConversationsSendToConversationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewConversationsSendToConversationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewConversationsSendToConversationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewConversationsSendToConversationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewConversationsSendToConversationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewConversationsSendToConversationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewConversationsSendToConversationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewConversationsSendToConversationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewConversationsSendToConversationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConversationsSendToConversationOK creates a ConversationsSendToConversationOK with default headers values
func NewConversationsSendToConversationOK() *ConversationsSendToConversationOK {
	return &ConversationsSendToConversationOK{}
}

/*ConversationsSendToConversationOK handles this case with default header values.

An object will be returned containing the ID for the resource.
*/
type ConversationsSendToConversationOK struct {
	Payload *models.ResourceResponse
}

func (o *ConversationsSendToConversationOK) Error() string {
	return fmt.Sprintf("[POST /v3/conversations/{conversationId}/activities][%d] conversationsSendToConversationOK  %+v", 200, o.Payload)
}

func (o *ConversationsSendToConversationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationsSendToConversationCreated creates a ConversationsSendToConversationCreated with default headers values
func NewConversationsSendToConversationCreated() *ConversationsSendToConversationCreated {
	return &ConversationsSendToConversationCreated{}
}

/*ConversationsSendToConversationCreated handles this case with default header values.

A ResourceResponse object will be returned containing the ID for the resource.
*/
type ConversationsSendToConversationCreated struct {
	Payload *models.ResourceResponse
}

func (o *ConversationsSendToConversationCreated) Error() string {
	return fmt.Sprintf("[POST /v3/conversations/{conversationId}/activities][%d] conversationsSendToConversationCreated  %+v", 201, o.Payload)
}

func (o *ConversationsSendToConversationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationsSendToConversationAccepted creates a ConversationsSendToConversationAccepted with default headers values
func NewConversationsSendToConversationAccepted() *ConversationsSendToConversationAccepted {
	return &ConversationsSendToConversationAccepted{}
}

/*ConversationsSendToConversationAccepted handles this case with default header values.

An object will be returned containing the ID for the resource.
*/
type ConversationsSendToConversationAccepted struct {
	Payload *models.ResourceResponse
}

func (o *ConversationsSendToConversationAccepted) Error() string {
	return fmt.Sprintf("[POST /v3/conversations/{conversationId}/activities][%d] conversationsSendToConversationAccepted  %+v", 202, o.Payload)
}

func (o *ConversationsSendToConversationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationsSendToConversationBadRequest creates a ConversationsSendToConversationBadRequest with default headers values
func NewConversationsSendToConversationBadRequest() *ConversationsSendToConversationBadRequest {
	return &ConversationsSendToConversationBadRequest{}
}

/*ConversationsSendToConversationBadRequest handles this case with default header values.

The request was malformed or otherwise incorrect. Inspect the message for a more detailed description.
*/
type ConversationsSendToConversationBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ConversationsSendToConversationBadRequest) Error() string {
	return fmt.Sprintf("[POST /v3/conversations/{conversationId}/activities][%d] conversationsSendToConversationBadRequest  %+v", 400, o.Payload)
}

func (o *ConversationsSendToConversationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationsSendToConversationUnauthorized creates a ConversationsSendToConversationUnauthorized with default headers values
func NewConversationsSendToConversationUnauthorized() *ConversationsSendToConversationUnauthorized {
	return &ConversationsSendToConversationUnauthorized{}
}

/*ConversationsSendToConversationUnauthorized handles this case with default header values.

The bot is not authorized to make this request. Please check your Microsoft App ID and Microsoft App password.
*/
type ConversationsSendToConversationUnauthorized struct {
}

func (o *ConversationsSendToConversationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v3/conversations/{conversationId}/activities][%d] conversationsSendToConversationUnauthorized ", 401)
}

func (o *ConversationsSendToConversationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConversationsSendToConversationForbidden creates a ConversationsSendToConversationForbidden with default headers values
func NewConversationsSendToConversationForbidden() *ConversationsSendToConversationForbidden {
	return &ConversationsSendToConversationForbidden{}
}

/*ConversationsSendToConversationForbidden handles this case with default header values.

The request was a valid request, but the server is refusing to respond to it. The user might be logged in but does not have the necessary permissions for the resource.
*/
type ConversationsSendToConversationForbidden struct {
}

func (o *ConversationsSendToConversationForbidden) Error() string {
	return fmt.Sprintf("[POST /v3/conversations/{conversationId}/activities][%d] conversationsSendToConversationForbidden ", 403)
}

func (o *ConversationsSendToConversationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConversationsSendToConversationNotFound creates a ConversationsSendToConversationNotFound with default headers values
func NewConversationsSendToConversationNotFound() *ConversationsSendToConversationNotFound {
	return &ConversationsSendToConversationNotFound{}
}

/*ConversationsSendToConversationNotFound handles this case with default header values.

The resource was not found.
*/
type ConversationsSendToConversationNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ConversationsSendToConversationNotFound) Error() string {
	return fmt.Sprintf("[POST /v3/conversations/{conversationId}/activities][%d] conversationsSendToConversationNotFound  %+v", 404, o.Payload)
}

func (o *ConversationsSendToConversationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationsSendToConversationInternalServerError creates a ConversationsSendToConversationInternalServerError with default headers values
func NewConversationsSendToConversationInternalServerError() *ConversationsSendToConversationInternalServerError {
	return &ConversationsSendToConversationInternalServerError{}
}

/*ConversationsSendToConversationInternalServerError handles this case with default header values.

An internal server has occurred. Inspect the message for a more detailed description.
*/
type ConversationsSendToConversationInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ConversationsSendToConversationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v3/conversations/{conversationId}/activities][%d] conversationsSendToConversationInternalServerError  %+v", 500, o.Payload)
}

func (o *ConversationsSendToConversationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationsSendToConversationServiceUnavailable creates a ConversationsSendToConversationServiceUnavailable with default headers values
func NewConversationsSendToConversationServiceUnavailable() *ConversationsSendToConversationServiceUnavailable {
	return &ConversationsSendToConversationServiceUnavailable{}
}

/*ConversationsSendToConversationServiceUnavailable handles this case with default header values.

The service you are trying to communciate with is unavailable.
*/
type ConversationsSendToConversationServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *ConversationsSendToConversationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v3/conversations/{conversationId}/activities][%d] conversationsSendToConversationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ConversationsSendToConversationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
