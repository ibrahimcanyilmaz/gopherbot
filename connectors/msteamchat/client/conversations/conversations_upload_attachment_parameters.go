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

	"github.com/uva-its/gopherbot/connectors/msteamchat/models"
)

// NewConversationsUploadAttachmentParams creates a new ConversationsUploadAttachmentParams object
// with the default values initialized.
func NewConversationsUploadAttachmentParams() *ConversationsUploadAttachmentParams {
	var ()
	return &ConversationsUploadAttachmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConversationsUploadAttachmentParamsWithTimeout creates a new ConversationsUploadAttachmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConversationsUploadAttachmentParamsWithTimeout(timeout time.Duration) *ConversationsUploadAttachmentParams {
	var ()
	return &ConversationsUploadAttachmentParams{

		timeout: timeout,
	}
}

// NewConversationsUploadAttachmentParamsWithContext creates a new ConversationsUploadAttachmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewConversationsUploadAttachmentParamsWithContext(ctx context.Context) *ConversationsUploadAttachmentParams {
	var ()
	return &ConversationsUploadAttachmentParams{

		Context: ctx,
	}
}

/*ConversationsUploadAttachmentParams contains all the parameters to send to the API endpoint
for the conversations upload attachment operation typically these are written to a http.Request
*/
type ConversationsUploadAttachmentParams struct {

	/*AttachmentUpload
	  Attachment data

	*/
	AttachmentUpload *models.AttachmentData
	/*ConversationID
	  Conversation ID

	*/
	ConversationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the conversations upload attachment params
func (o *ConversationsUploadAttachmentParams) WithTimeout(timeout time.Duration) *ConversationsUploadAttachmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the conversations upload attachment params
func (o *ConversationsUploadAttachmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the conversations upload attachment params
func (o *ConversationsUploadAttachmentParams) WithContext(ctx context.Context) *ConversationsUploadAttachmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the conversations upload attachment params
func (o *ConversationsUploadAttachmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithAttachmentUpload adds the attachmentUpload to the conversations upload attachment params
func (o *ConversationsUploadAttachmentParams) WithAttachmentUpload(attachmentUpload *models.AttachmentData) *ConversationsUploadAttachmentParams {
	o.SetAttachmentUpload(attachmentUpload)
	return o
}

// SetAttachmentUpload adds the attachmentUpload to the conversations upload attachment params
func (o *ConversationsUploadAttachmentParams) SetAttachmentUpload(attachmentUpload *models.AttachmentData) {
	o.AttachmentUpload = attachmentUpload
}

// WithConversationID adds the conversationID to the conversations upload attachment params
func (o *ConversationsUploadAttachmentParams) WithConversationID(conversationID string) *ConversationsUploadAttachmentParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the conversations upload attachment params
func (o *ConversationsUploadAttachmentParams) SetConversationID(conversationID string) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *ConversationsUploadAttachmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.AttachmentUpload == nil {
		o.AttachmentUpload = new(models.AttachmentData)
	}

	if err := r.SetBodyParam(o.AttachmentUpload); err != nil {
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
