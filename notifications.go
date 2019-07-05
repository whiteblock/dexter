package dexter

import (
)

/*
- Provide a channel that multiple writers can put requests on to.
- Hand these requests out to workers that can handle them.
*/

// Notification is something that has a Send method.
type Notification interface {
	Send()
}

// Send will make the HTTP request described by the webhook.
func (w Webhook) Send() {
}


// StartNotifications listens for notification requests on a channel that it returns to the caller.
func StartNotifications(in chan Notification) {

}
