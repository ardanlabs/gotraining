// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0
package main

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/satori/go.uuid"
)

var (
	// ErrNilMsg indicates a nil Msg.
	ErrNilMsg = errors.New("me: message.Msg cannot be nil")
)

// Msg represents information transmitted via NATS in its nats.Msg.Data field
type Message struct {
	TransactionID string    `json:"transactionId"`   // Every msg has unique Id
	CreatedAt     time.Time `json:"createdAt"`       // Every msg gets timestamp
	Status        int       `json:"status"`          // 0 - pass, 1 - fail
	Text          string    `json:"text, omitempty"` // Optional message text
	JSON          string    `json:"json, omitempty"` // Optional message payload
}

// MarshalJSON populates CreatedAt and TransactionID fields.
func (m *Message) MarshalJSON() ([]byte, error) {

	if m == nil {
		return nil, ErrNilMsg
	}

	type alias Message

	a := alias(*m)

	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now().UTC()
	}

	if len(a.TransactionID) == 0 {
		a.TransactionID = uuid.NewV4().String()
	}

	return json.Marshal(a)
}
