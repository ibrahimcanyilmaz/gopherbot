package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewConversationsGetActivityMembersParams creates a new ConversationsGetActivityMembersParams object
// with the default values initialized.
func NewConversationsGetActivityMembersParams() *ConversationsGetActivityMembersParams {
	var ()
	return &ConversationsGetActivityMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConversationsGetActivityMembersParamsWithTimeout creates a new ConversationsGetActivityMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConversationsGetActivityMembersParamsWithTimeout(timeout time.Duration) *ConversationsGetActivityMembersParams {
	var ()
	return &ConversationsGetActivityMembersParams{

		timeout: timeout,
	}
}

// NewConversationsGetActivityMembersParamsWithContext creates a new ConversationsGetActivityMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewConversationsGetActivityMembersParamsWithContext(ctx context.Context) *ConversationsGetActivityMembersParams {
	var ()
	return &ConversationsGetActivityMembersParams{

		Context: ctx,
	}
}

/*ConversationsGetActivityMembersParams contains all the parameters to send to the API endpoint
for the conversations get activity members operation typically these are written to a http.Request
*/
type ConversationsGetActivityMembersParams struct {

	/*ActivityID
	  Activity ID

	*/
	ActivityID string
	/*ConversationID
	  Conversation ID

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the conversations get activity members params
func (o *ConversationsGetActivityMembersParams) WithTimeout(timeout time.Duration) *ConversationsGetActivityMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the conversations get activity members params
func (o *ConversationsGetActivityMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the conversations get activity members params
func (o *ConversationsGetActivityMembersParams) WithContext(ctx context.Context) *ConversationsGetActivityMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the conversations get activity members params
func (o *ConversationsGetActivityMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithActivityID adds the activityID to the conversations get activity members params
func (o *ConversationsGetActivityMembersParams) WithActivityID(activityID string) *ConversationsGetActivityMembersParams {
	o.SetActivityID(activityID)
	return o
}

// SetActivityID adds the activityId to the conversations get activity members params
func (o *ConversationsGetActivityMembersParams) SetActivityID(activityID string) {
	o.ActivityID = activityID
}

// WithConversationID adds the conversationID to the conversations get activity members params
func (o *ConversationsGetActivityMembersParams) WithConversationID(conversationID string) *ConversationsGetActivityMembersParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the conversations get activity members params
func (o *ConversationsGetActivityMembersParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *ConversationsGetActivityMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param activityId
	if err := r.SetPathParam("activityId", o.ActivityID); err != nil {
		return err
	}

	// path param conversationId
	if err := r.SetPathParam("conversationId", o.ConversationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
