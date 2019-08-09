package slackmsg_test

import (
	"context"
	"testing"
	"time"

	"github.com/spider-pigs/slackmsg"
)

func TestSimpleMsg(t *testing.T) {
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

	json, err := msg.ToJSON()
	if err != nil {
		t.Error("error should not be set")
	}
	if len(json) == 0 {
		t.Error("JSON string should not be empty")
	}
}

func TestActionMsg(t *testing.T) {
	// create message
	msg := slackmsg.New()
	msg.Text = "Would you like to play a game?"
	attachment := slackmsg.Attachment{
		Text:       "Choose a game to play",
		Fallback:   "You are unable to choose a game",
		CallbackID: "wopr_game",
		Color:      "#3AA3E3",
		Type:       "default",
		Actions: []slackmsg.Action{
			slackmsg.Action{
				Name:  "game",
				Text:  "Chess",
				Type:  "button",
				Value: "chess",
			},
			slackmsg.Action{
				Name:  "game",
				Text:  "Falken's Maze",
				Type:  "button",
				Value: "maze",
			},
			slackmsg.Action{
				Name:  "game",
				Text:  "Thermonuclear War",
				Style: "danger",
				Type:  "button",
				Value: "war",
				Confirm: slackmsg.Confirm{
					Title:       "Are you sure?",
					Text:        "Wouldn't you prefer a good game of chess?",
					OkText:      "Yes",
					DismissText: "No",
				},
			},
		},
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

	json, err := msg.ToJSON()
	if err != nil {
		t.Error("error should not be set")
	}
	if len(json) == 0 {
		t.Error("JSON string should not be empty")
	}
}
