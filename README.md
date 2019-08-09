# slackmsg
[![Build Status](https://travis-ci.org/spider-pigs/slackmsg.svg?branch=master)](https://travis-ci.org/spider-pigs/slackmsg) [![Go Report Card](https://goreportcard.com/badge/github.com/spider-pigs/slackmsg)](https://goreportcard.com/report/github.com/spider-pigs/slackmsg) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/47f6b9abfbb04e888cdc9759df7f28e1)](https://www.codacy.com/app/spider-pigs/slackmsg?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=spider-pigs/slackmsg&amp;utm_campaign=Badge_Grade) [![GoDoc](https://godoc.org/github.com/spider-pigs/slackmsg?status.svg)](https://godoc.org/github.com/spider-pigs/slackmsg)

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
