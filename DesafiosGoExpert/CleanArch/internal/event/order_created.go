package event

import (
	"time"
)

type OrderCreated struct {
	Name    string
	Payload interface{}
}
type GetOrder struct {
	Name    string
	Payload interface{}
}
type GetOrdersList struct {
	Name    string
	Payload interface{}
}

func NewOrderCreated() *OrderCreated {
	return &OrderCreated{
		Name: "OrderCreated",
	}
}
func NewGetOrder() *GetOrder {
	return &GetOrder{
		Name: "GetOrder",
	}
}
func NewGetOrdersList() *GetOrdersList {
	return &GetOrdersList{
		Name: "GetOrdersList",
	}
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

//

func (e *GetOrder) GetName() string {
	return e.Name
}

func (e *GetOrder) GetPayload() interface{} {
	return e.Payload
}

func (e *GetOrder) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *GetOrder) GetDateTime() time.Time {
	return time.Now()
}

//

func (e *GetOrdersList) GetName() string {
	return e.Name
}

func (e *GetOrdersList) GetPayload() interface{} {
	return e.Payload
}

func (e *GetOrdersList) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *GetOrdersList) GetDateTime() time.Time {
	return time.Now()
}
