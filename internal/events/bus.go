package events

import "sync"

type EventType string
const (
	UserCreated  EventType = "user.created"
	ItemCreated  EventType = "item.created"
	ItemDeleted  EventType = "item.deleted"
)

type Event struct{ Type EventType; Payload interface{} }
type Handler func(Event)

type Bus struct{ mu sync.RWMutex; handlers map[EventType][]Handler }

func NewBus() *Bus { return &Bus{handlers: make(map[EventType][]Handler)} }

func (b *Bus) Subscribe(t EventType, h Handler) {
	b.mu.Lock(); defer b.mu.Unlock()
	b.handlers[t] = append(b.handlers[t], h)
}

func (b *Bus) Publish(e Event) {
	b.mu.RLock()
	hs := append([]Handler{}, b.handlers[e.Type]...)
	b.mu.RUnlock()
	for _, h := range hs { go h(e) }
}
