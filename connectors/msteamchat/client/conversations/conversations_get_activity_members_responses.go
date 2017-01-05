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

// ConversationsGetActivityMembersReader is a Reader for the ConversationsGetActivityMembers structure.
type ConversationsGetActivityMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConversationsGetActivityMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewConversationsGetActivityMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewConversationsGetActivityMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewConversationsGetActivityMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewConversationsGetActivityMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewConversationsGetActivityMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewConversationsGetActivityMembersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewConversationsGetActivityMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewConversationsGetActivityMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConversationsGetActivityMembersOK creates a ConversationsGetActivityMembersOK with default headers values
func NewConversationsGetActivityMembersOK() *ConversationsGetActivityMembersOK {
	return &ConversationsGetActivityMembersOK{}
}

/*ConversationsGetActivityMembersOK handles this case with default header values.

An array of ChannelAccount objects
*/
type ConversationsGetActivityMembersOK struct {
	Payload []*models.ChannelAccount
}

func (o *ConversationsGetActivityMembersOK) Error() string {
	return fmt.Sprintf("[GET /v3/conversations/{conversationId}/activities/{activityId}/members][%d] conversationsGetActivityMembersOK  %+v", 200, o.Payload)
}

func (o *ConversationsGetActivityMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationsGetActivityMembersBadRequest creates a ConversationsGetActivityMembersBadRequest with default headers values
func NewConversationsGetActivityMembersBadRequest() *ConversationsGetActivityMembersBadRequest {
	return &ConversationsGetActivityMembersBadRequest{}
}

/*ConversationsGetActivityMembersBadRequest handles this case with default header values.

The request was malformed or otherwise incorrect. Inspect the message for a more detailed description.
*/
type ConversationsGetActivityMembersBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ConversationsGetActivityMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /v3/conversations/{conversationId}/activities/{activityId}/members][%d] conversationsGetActivityMembersBadRequest  %+v", 400, o.Payload)
}

func (o *ConversationsGetActivityMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationsGetActivityMembersUnauthorized creates a ConversationsGetActivityMembersUnauthorized with default headers values
func NewConversationsGetActivityMembersUnauthorized() *ConversationsGetActivityMembersUnauthorized {
	return &ConversationsGetActivityMembersUnauthorized{}
}

/*ConversationsGetActivityMembersUnauthorized handles this case with default header values.

The bot is not authorized to make this request. Please check your Microsoft App ID and Microsoft App password.
*/
type ConversationsGetActivityMembersUnauthorized struct {
}

func (o *ConversationsGetActivityMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v3/conversations/{conversationId}/activities/{activityId}/members][%d] conversationsGetActivityMembersUnauthorized ", 401)
}

func (o *ConversationsGetActivityMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConversationsGetActivityMembersForbidden creates a ConversationsGetActivityMembersForbidden with default headers values
func NewConversationsGetActivityMembersForbidden() *ConversationsGetActivityMembersForbidden {
	return &ConversationsGetActivityMembersForbidden{}
}

/*ConversationsGetActivityMembersForbidden handles this case with default header values.

The request was a valid request, but the server is refusing to respond to it. The user might be logged in but does not have the necessary permissions for the resource.
*/
type ConversationsGetActivityMembersForbidden struct {
}

func (o *ConversationsGetActivityMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /v3/conversations/{conversationId}/activities/{activityId}/members][%d] conversationsGetActivityMembersForbidden ", 403)
}

func (o *ConversationsGetActivityMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConversationsGetActivityMembersNotFound creates a ConversationsGetActivityMembersNotFound with default headers values
func NewConversationsGetActivityMembersNotFound() *ConversationsGetActivityMembersNotFound {
	return &ConversationsGetActivityMembersNotFound{}
}

/*ConversationsGetActivityMembersNotFound handles this case with default header values.

The resource was not found.
*/
type ConversationsGetActivityMembersNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ConversationsGetActivityMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /v3/conversations/{conversationId}/activities/{activityId}/members][%d] conversationsGetActivityMembersNotFound  %+v", 404, o.Payload)
}

func (o *ConversationsGetActivityMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationsGetActivityMembersMethodNotAllowed creates a ConversationsGetActivityMembersMethodNotAllowed with default headers values
func NewConversationsGetActivityMembersMethodNotAllowed() *ConversationsGetActivityMembersMethodNotAllowed {
	return &ConversationsGetActivityMembersMethodNotAllowed{}
}

/*ConversationsGetActivityMembersMethodNotAllowed handles this case with default header values.

The method and URI you are trying to use is not allowed on this service.  For example, not all services support the DELETE or PUT verbs on an activity URI.
*/
type ConversationsGetActivityMembersMethodNotAllowed struct {
	Payload *models.ErrorResponse
}

func (o *ConversationsGetActivityMembersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v3/conversations/{conversationId}/activities/{activityId}/members][%d] conversationsGetActivityMembersMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ConversationsGetActivityMembersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationsGetActivityMembersInternalServerError creates a ConversationsGetActivityMembersInternalServerError with default headers values
func NewConversationsGetActivityMembersInternalServerError() *ConversationsGetActivityMembersInternalServerError {
	return &ConversationsGetActivityMembersInternalServerError{}
}

/*ConversationsGetActivityMembersInternalServerError handles this case with default header values.

An internal server has occurred. Inspect the message for a more detailed description.
*/
type ConversationsGetActivityMembersInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ConversationsGetActivityMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v3/conversations/{conversationId}/activities/{activityId}/members][%d] conversationsGetActivityMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *ConversationsGetActivityMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConversationsGetActivityMembersServiceUnavailable creates a ConversationsGetActivityMembersServiceUnavailable with default headers values
func NewConversationsGetActivityMembersServiceUnavailable() *ConversationsGetActivityMembersServiceUnavailable {
	return &ConversationsGetActivityMembersServiceUnavailable{}
}

/*ConversationsGetActivityMembersServiceUnavailable handles this case with default header values.

The service you are trying to communciate with is unavailable.
*/
type ConversationsGetActivityMembersServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *ConversationsGetActivityMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v3/conversations/{conversationId}/activities/{activityId}/members][%d] conversationsGetActivityMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ConversationsGetActivityMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
