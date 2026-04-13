package events

import "log"

func OnUserCreated(e Event) {
	log.Printf("event: user created: %v", e.Payload)
	// TODO: send welcome email via EmailService
}

func OnItemCreated(e Event) {
	log.Printf("event: item created: %v", e.Payload)
	// TODO: notify followers
}

func RegisterAll(bus *Bus) {
	bus.Subscribe(UserCreated, OnUserCreated)
	bus.Subscribe(ItemCreated, OnItemCreated)
}
