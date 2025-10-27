package repository

import (
	"context"
	"fmt"
	"sync"

	"orders-microservice/internal/models"
	"orders-microservice/pkg/api/test"
)

type OrdersRepository struct {
	storage map[string]models.Order
	rwm     sync.RWMutex
}

func New() *OrdersRepository {
	return &OrdersRepository{
		storage: make(map[string]models.Order),
	}
}

func (r *OrdersRepository) Create(ctx context.Context, order models.Order) (string, error) {
	if err := checkNilContext(ctx); err != nil {
		return "", err
	}
	if err := checkContextDone(ctx); err != nil {
		return "", err
	}

	r.rwm.Lock()
	defer r.rwm.Unlock()

	id := generateId()
	err := r.findOrder(id)
	for err == nil {
		id = generateId()
		err = r.findOrder(id)

		if err := checkContextDone(ctx); err != nil {
			return "", err
		}
	}

	r.storage[id] = order

	return id, nil
}

func (r *OrdersRepository) GetOne(ctx context.Context, id string) (models.Order, error) {
	if err := checkNilContext(ctx); err != nil {
		return models.Order{}, err
	}
	if err := checkContextDone(ctx); err != nil {
		return models.Order{}, err
	}

	r.rwm.RLock()
	defer r.rwm.RUnlock()

	if err := r.findOrder(id); err != nil {
		return models.Order{}, fmt.Errorf("could not get: %w", err)
	}

	return r.storage[id], nil
}

func (r *OrdersRepository) GetAll(ctx context.Context) ([]*test.Order, error) {
	if err := checkNilContext(ctx); err != nil {
		return nil, err
	}
	if err := checkContextDone(ctx); err != nil {
		return nil, err
	}

	r.rwm.RLock()
	defer r.rwm.RUnlock()

	orders := make([]*test.Order, 0, len(r.storage))
	for id, order := range r.storage {
		orders = append(orders, &test.Order{
			Id: id,
			Item: order.Item,
			Quantity: order.Quantity,
		})

		if err := checkContextDone(ctx); err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func (r *OrdersRepository) Update(ctx context.Context, id string, newOrder models.Order) error {
	if err := checkNilContext(ctx); err != nil {
		return err
	}
	if err := checkContextDone(ctx); err != nil {
		return err
	}

	r.rwm.Lock()
	defer r.rwm.Unlock()

	if err := r.findOrder(id); err != nil {
		return fmt.Errorf("could not update: %w", err)
	}

	r.storage[id] = newOrder

	return nil
}

func (r *OrdersRepository) Delete(ctx context.Context, id string) error {
	if err := checkNilContext(ctx); err != nil {
		return err
	}
	if err := checkContextDone(ctx); err != nil {
		return err
	}

	r.rwm.Lock()
	defer r.rwm.Unlock()

	if err := r.findOrder(id); err != nil {
		return fmt.Errorf("could not delete: %w", err)
	}

	delete(r.storage, id)

	return nil
}

func (r *OrdersRepository) Close() {
	if r.storage != nil {
		r.storage = nil
	}
}

func (r *OrdersRepository) findOrder(id string) error {
	if _, ok := r.storage[id]; !ok {
		return fmt.Errorf("storage: order with id %s not exist not found in storage", id)
	}

	return nil
}

func checkNilContext(ctx context.Context) error {
	if ctx == nil {
		return fmt.Errorf("storage: nil context")
	}

	return nil
}

func checkContextDone(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
