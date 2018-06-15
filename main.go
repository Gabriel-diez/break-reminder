package main

import (
	"github.com/deckarep/gosx-notifier"
	"log"
)

func main() {
	n := gosxnotifier.NewNotification("Test notification")

	if err := n.Push(); err != nil {
		log.Fatal(err)
	}
}