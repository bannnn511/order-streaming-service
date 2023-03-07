package usecases

import (
	"context"
	"order-streaming-services/internal/orders/domain"
)

type usecase struct {
	serviceKafka OrderServiceKafka
}

var _ UserCase = (*usecase)(nil)

func NewUseCase(serviceKafka OrderServiceKafka) *usecase {
	return &usecase{serviceKafka: serviceKafka}
}

func (u usecase) PlaceOrder(ctx context.Context, bean *domain.OrderBean) error {
	return u.serviceKafka.Publish(ctx, bean)
}
