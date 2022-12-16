# slackmsg
[![Build Status](https://github.com/spider-pigs/slackmsg/actions/workflows/test.yml/badge.svg)](https://github.com/spider-pigs/slackmsg/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/spider-pigs/slackmsg)](https://goreportcard.com/report/github.com/spider-pigs/slackmsg) [![GoDoc](https://godoc.org/github.com/spider-pigs/slackmsg?status.svg)](https://godoc.org/github.com/spider-pigs/slackmsg)

A simple golang library to send a slack message.

## Install

```Go
import "github.com/spider-pigs/slackmsg"
```

## Usage

```Go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spider-pigs/slackmsg"
)

func main() {
	// create message
	msg := slackmsg.New()
	msg.Text = "I am a test message!"
	attachment := slackmsg.Attachment{
		Text: "And here’s an attachment!",
	}
	msg.AddAttachment(attachment)

	// dump JSON to verify?
	fmt.Println(msg.ToJSON())

	// send message
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	webhook := "https://hooks.slack.com/services/..."
	err := msg.Send(ctx, webhook)
	if err != nil {
		fmt.Println("failed to send message:", err)
		return
	}
	fmt.Println("successfully sent message")
}
```
