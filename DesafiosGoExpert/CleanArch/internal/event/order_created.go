package event

import (
	"time"

	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/entity"
)

type OrderCreated struct {
	Name    string
	Payload interface{}
	Order   entity.Order
	Orders  []entity.Order
}

type GetOrder struct {
	Order interface{}
}

func NewOrderCreated() *OrderCreated {
	return &OrderCreated{
		Name: "OrderCreated",
	}
}

func (e *GetOrder) GetOrderById(id string) interface{} {
	return e.Order
}

func (e *OrderCreated) GetOrderById(id string) entity.Order {
	return e.Order
}
func (e *OrderCreated) GetOrders() []entity.Order {
	return e.Orders
}
func (e *OrderCreated) GetName() string {
	return e.Name
}

func (e *OrderCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderCreated) GetDateTime() time.Time {
	return time.Now()
}
