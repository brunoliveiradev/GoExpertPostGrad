package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("testEventHandler already registered")
var ErrHandlerClearError = errors.New("handlers not cleared")
var ErrHandlerNotRegistered = errors.New("handler not registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Dispatch sends the event to all registered handlers
func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	registeredHandlers, handlerExists := ed.handlers[event.GetName()]
	if handlerExists {
		// WaitGroup to wait for all handlers to finish
		wg := &sync.WaitGroup{}
		for _, handler := range registeredHandlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

// Register adds a handler to the event name
func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if ed.Has(eventName, handler) {
		return ErrHandlerAlreadyRegistered
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

// Has checks if the handler is already registered for the event name provided
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	registeredHandlers, handlerExists := ed.handlers[eventName]
	if handlerExists {
		for _, registeredHandler := range registeredHandlers {
			if registeredHandler == handler {
				return true
			}
		}
	}
	return false
}

// Remove deletes a handler from the given eventName
func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	registeredHandlers, handlerExists := ed.handlers[eventName]
	if handlerExists {
		for i, registeredHandler := range registeredHandlers {
			if registeredHandler == handler {
				ed.handlers[eventName] = append(registeredHandlers[:i], registeredHandlers[i+1:]...)
				return nil
			}
		}
	}
	return ErrHandlerNotRegistered
}

// Clear removes all handlers
func (ed *EventDispatcher) Clear() error {
	ed.handlers = make(map[string][]EventHandlerInterface)
	if len(ed.handlers) != 0 {
		return ErrHandlerClearError
	}
	return nil
}
