package web

import (
	"encoding/json"
	"net/http"

	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/entity"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/internal/usecase"
	"github.com/malandrim/Projetos-Go/DesafiosGoExpert/CleanArch/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// /////
type WebGetOrderHandler struct {
	EventDispatcher events.EventDispatcherInterface
	OrderRepository entity.OrderRepositoryInterface
	GetOrderEvent   events.EventInterface
}

func NewWebGetOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	GetOrderEvent events.EventInterface,
) *WebGetOrderHandler {
	return &WebGetOrderHandler{
		EventDispatcher: EventDispatcher,
		OrderRepository: OrderRepository,
		GetOrderEvent:   GetOrderEvent,
	}
}

func (h *WebGetOrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getOrder := usecase.NewGetOrderUseCase(h.OrderRepository, h.GetOrderEvent, h.EventDispatcher)
	output, err := getOrder.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// //////
type WebGetOrdersListHandler struct {
	EventDispatcher    events.EventDispatcherInterface
	OrderRepository    entity.OrderRepositoryInterface
	GetOrdersListEvent events.EventInterface
}

func NewWebGetOrdersListHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	GetOrdersListEvent events.EventInterface,
) *WebGetOrdersListHandler {
	return &WebGetOrdersListHandler{
		EventDispatcher:    EventDispatcher,
		OrderRepository:    OrderRepository,
		GetOrdersListEvent: GetOrdersListEvent,
	}
}

func (h *WebGetOrdersListHandler) GetOrdersList(w http.ResponseWriter, r *http.Request) {

	OrderList := usecase.NewGetOrdersListUseCase(h.OrderRepository, h.GetOrdersListEvent, h.EventDispatcher)
	output, err := OrderList.OrderRepository.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
