package main

import (
	"github.com/deckarep/gosx-notifier"
	"log"
	"time"
)

func main() {
	n := buildNotification()
	for {
		time.Sleep(60 * time.Minute)
		go pushNotification(&n)
	}
}

func pushNotification (n* gosxnotifier.Notification) {
	if err := n.Push(); err != nil {
		log.Fatal(err)
	}
}

func buildNotification () gosxnotifier.Notification {
	n := gosxnotifier.NewNotification("It's break time ! Don't forget to drink and stretch")
	n.Title = "Pause Reminder"
	n.Sender = "com.apple.Stickies"
	n.AppIcon = "gopher.png"
	n.Sound = gosxnotifier.Default

	return *n
}