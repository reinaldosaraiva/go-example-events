package events

import (
	"errors"
	"sync"
)

type EventDispatcher struct {
	handlers map[string][]EventHandlersInterface
}

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlersInterface),
	}
}

func (e *EventDispatcher) Register(eventName string, handler EventHandlersInterface) error {
	if _, ok := e.handlers[eventName]; ok {
		for _, h := range e.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	e.handlers[eventName] = append(e.handlers[eventName], handler)
	return nil
}

//clear
func (e *EventDispatcher) Clear() error {
	e.handlers = make(map[string][]EventHandlersInterface)
	return nil
}

//has
func (e *EventDispatcher) Has(eventName string, handler EventHandlersInterface) bool {
	if _, ok := e.handlers[eventName]; ok {
		for _, h := range e.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}
//remove
func (e *EventDispatcher) Remove(eventName string, handler EventHandlersInterface) error {
	if _, ok := e.handlers[eventName]; ok {
		for i, h := range e.handlers[eventName] {
			if h == handler {
				e.handlers[eventName] = append(e.handlers[eventName][:i], e.handlers[eventName][i+1:]...)
			}
		}
	}
	return nil
}
func (e *EventDispatcher) Dispatch(event EventsInterface) error {
	if _, ok := e.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range e.handlers[event.GetName()] {
			wg.Add(1)
			go handler.Handle(event,wg)
		}
		wg.Wait()
	}
	return nil
}