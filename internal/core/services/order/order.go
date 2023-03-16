package order

import (
	"context"
	"shop/internal/core/domain"
	"shop/internal/core/models"
	"shop/internal/core/ports"
)

type OrderService struct {
	orderRepo   ports.OrderRepository
	productRepo ports.ProductRepository
}

var _ ports.OrderService = &OrderService{}

func NewOrderService(orderRepo ports.OrderRepository, productRepo ports.ProductRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo, productRepo: productRepo}
}

func (s *OrderService) Create(ctx context.Context, o domain.Order, userId uint) (uint, error) {
	order, ids := parseToModel(o, userId)
	return s.orderRepo.Create(ctx, order, ids)
}

func (s *OrderService) Update(ctx context.Context, o domain.Order, userId uint) (uint, error) {
	order, ids := parseToModel(o, userId)
	return s.orderRepo.Update(ctx, order, ids)
}

func (s *OrderService) Get(ctx context.Context, id uint) (domain.Order, error) {
	model, err := s.orderRepo.Get(ctx, id)
	if err != nil {
		return domain.Order{}, err
	}
	products := make([]domain.Product, len(model))
	for i, m := range model {
		product, err := s.productRepo.Get(ctx, m.ProductId)
		if err != nil {
			return domain.Order{}, err
		}
		products[i] = product
	}
	order := parseToDomain(model[0], products)
	return order, nil
}

func (s *OrderService) GetOrdersByUserId(ctx context.Context, userId uint) ([]domain.Order, error) {
	modelOrder, err := s.orderRepo.GetOrdersByUserId(ctx, userId)
	if err != nil {
		return []domain.Order{}, err
	}
	orders := make([]domain.Order, len(modelOrder))
	for l, mo := range modelOrder {
		model, err := s.orderRepo.Get(ctx, mo.OrderId)
		if err != nil {
			return []domain.Order{}, err
		}
		products := make([]domain.Product, len(model))
		for i, m := range model {
			product, err := s.productRepo.Get(ctx, m.ProductId)
			if err != nil {
				return []domain.Order{}, err
			}
			products[i] = product
		}
		orders[l] = parseToDomain(model[0], products)
	}
	return orders, nil
}

func (s *OrderService) GetAll(ctx context.Context) ([]domain.Order, error) {
	modelOrder, err := s.orderRepo.GetAll(ctx)
	if err != nil {
		return []domain.Order{}, err
	}
	orders := make([]domain.Order, len(modelOrder))
	for l, mo := range modelOrder {
		model, err := s.orderRepo.Get(ctx, mo.OrderId)
		if err != nil {
			return []domain.Order{}, err
		}
		products := make([]domain.Product, len(model))
		for i, m := range model {
			product, err := s.productRepo.Get(ctx, m.ProductId)
			if err != nil {
				return []domain.Order{}, err
			}
			products[i] = product
		}
		orders[l] = parseToDomain(model[0], products)
	}
	return orders, nil
}

func (s *OrderService) Delete(ctx context.Context, id uint, userId uint) error {
	order := models.Order{
		OrderId: id,
		UserId:  userId,
	}
	return s.orderRepo.Delete(ctx, order)
}

func parseToModel(order domain.Order, userId uint) (models.Order, []uint) {
	productIds := make([]uint, 0)
	for _, p := range order.Products {
		b := checkUniqueId(productIds, p.ID)
		if !b {
			continue
		}
		productIds = append(productIds, p.ID)
	}
	return models.Order{
		OrderId: order.ID,
		UserId:  userId,
	}, productIds
}

func checkUniqueId(productIds []uint, id uint) bool {
	for _, i := range productIds {
		if i == id {
			return false
		}
	}
	return true
}

func parseToDomain(model models.ProductOrder, products []domain.Product) domain.Order {
	return domain.Order{
		ID:       model.OrderId,
		Products: products,
	}
}
