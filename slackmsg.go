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
	Attachments []Attachment `json:"attachments,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
	Text        string       `json:"text,omitempty"`
	Username    string       `json:"username,omitempty"`
}

// Attachment type
type Attachment struct {
	Actions    []Action `json:"actions,omitempty"`
	AuthorName string   `json:"author_name,omitempty"`
	AuthorIcon string   `json:"author_icon,omitempty"`
	AuthorLink string   `json:"author_link,omitempty"`
	CallbackID string   `json:"callback_id,omitempty"`
	Color      string   `json:"color,omitempty"`
	Fallback   string   `json:"fallback,omitempty"`
	Fields     []Field  `json:"fields,omitempty"`
	Footer     string   `json:"footer,omitempty"`
	FooterIcon string   `json:"footer_icon,omitempty"`
	ImageURL   string   `json:"image_url,omitempty"`
	MarkdownIn []string `json:"mrkdwn_in,omitempty"`
	Pretext    string   `json:"pretext,omitempty"`
	Text       string   `json:"text,omitempty"`
	ThumbURL   string   `json:"thumb_url,omitempty"`
	Timestamp  int64    `json:"ts,omitempty"`
	Title      string   `json:"title,omitempty"`
	TitleLink  string   `json:"title_link,omitempty"`
	Type       string   `json:"attachment_type,omitempty"`
}

// Action type
type Action struct {
	Confirm Confirm `json:"confirm,omitempty"`
	Name    string  `json:"name,omitempty"`
	Style   string  `json:"style,omitempty"`
	Text    string  `json:"text,omitempty"`
	Type    string  `json:"type,omitempty"`
	Value   string  `json:"value,omitempty"`
}

// Confirm type
type Confirm struct {
	DismissText string `json:"dismiss_text,omitempty"`
	OkText      string `json:"ok_text,omitempty"`
	Text        string `json:"text,omitempty"`
	Title       string `json:"title,omitempty"`
}

// Field type
type Field struct {
	Title string `json:"title,omitempty"`
	Short bool   `json:"short,omitempty"`
	Value string `json:"value,omitempty"`
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

// ToJSON returns the message as a json string
func (msg *Message) ToJSON() (string, error) {
	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
