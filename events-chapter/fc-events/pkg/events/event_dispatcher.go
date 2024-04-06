package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if ed.isHandlerRegistered(eventName, handler) {
		return ErrHandlerAlreadyRegistered
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

// isHandlerRegistered checks if a handler is already registered for a specific event
func (ed *EventDispatcher) isHandlerRegistered(eventName string, handler EventHandlerInterface) bool {
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
