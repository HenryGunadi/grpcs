package types

import (
	"context"

	grpcs "github.com/HenryGunadi/grpcs/services/common/genproto/mcdonald"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *grpcs.CreateOrderRequest) error
	GetOrder(ctx context.Context) []*grpcs.CreateOrderRequest
}