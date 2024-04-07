package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("testEventHandler already registered")
var ErrHandlerClearError = errors.New("handlers not cleared")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

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

func (ed *EventDispatcher) Clear() error {
	ed.handlers = make(map[string][]EventHandlerInterface)
	if len(ed.handlers) != 0 {
		return ErrHandlerClearError
	}
	return nil
}
