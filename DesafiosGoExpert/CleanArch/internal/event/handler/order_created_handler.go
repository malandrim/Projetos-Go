package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/pkg/events"
	"github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}
type GetOrderHandler struct {
	RabbitMQChannel *amqp.Channel
}

type GetOrdersListHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}
func NewGetOrderHandler(rabbitMQChannel *amqp.Channel) *GetOrderHandler {
	return &GetOrderHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}
func NewGetOrdersListHandler(rabbitMQChannel *amqp.Channel) *GetOrdersListHandler {
	return &GetOrdersListHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order Created %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct",
		"",
		false,
		false,
		msgRabbitmq,
	)
}

func (h *GetOrderHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Get Order %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct",
		"",
		false,
		false,
		msgRabbitmq,
	)
}
func (h *GetOrdersListHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Get Orders List %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct",
		"",
		false,
		false,
		msgRabbitmq,
	)
}
