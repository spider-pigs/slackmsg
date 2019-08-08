package slackmsg_test

import (
	"context"
	"testing"
	"time"

	"github.com/spider-pigs/slackmsg"
)

func TestMsg(t *testing.T) {
	// create message
	msg := slackmsg.New()
	msg.Text = "I am a test message!"
	attachment := slackmsg.Attachment{
		Text: "And here’s an attachment!",
	}
	msg.AddAttachment(attachment)

	// send message
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	webhook := "https://hooks.slack.com..."
	err := msg.Send(ctx, webhook)
	if err == nil {
		t.Error("error should be set")
	}
}
