package events

import (
	"sync"
	"time"
)

type EventsInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlersInterface interface {
	Handle(event EventsInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlersInterface) error
	Dispatch(event EventsInterface) error
	Remove(eventName string, handler EventHandlersInterface) error
	Has(eventName string, handler EventHandlersInterface) bool
	Clear() error
}