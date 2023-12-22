package events

import (
	"sync"
	"time"

	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/entity"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
	GetOrderById(id string) entity.Order
	GetOrders() []entity.Order
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear()
}
