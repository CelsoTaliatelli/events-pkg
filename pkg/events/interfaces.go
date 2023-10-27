package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayLoad() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventHandlerInterface)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventNamestring, andler EventHandlerInterface) bool
	Clear() error
}
