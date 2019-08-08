package slackmsg

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Message type
type Message struct {
	Text        string       `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Attachment type
type Attachment struct {
	Title      string `json:"title,omitempty"`
	TitleLink  string `json:"title_link,omitempty"`
	Text       string `json:"text,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	Color      string `json:"color,omitempty"`
	Footer     string `json:"footer,omitempty"`
}

// New constructs a new message
func New() Message {
	attachments := make([]Attachment, 0)
	return Message{Attachments: attachments}
}

// AddAttachment adds an attachement
func (msg *Message) AddAttachment(attachment Attachment) {
	msg.Attachments = append(msg.Attachments, attachment)
}

// Send sends the message
func (msg *Message) Send(ctx context.Context, uri string) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	req = req.WithContext(ctx)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		text := fmt.Sprintf("got HTTP status code %d from %s", resp.StatusCode, uri)
		return errors.New(text)
	}

	return nil
}
