package orders

import (
	"context"
	"log"
	"time"

	grpcs "github.com/HenryGunadi/grpcs/services/common/genproto/mcdonald"
)

type OrdersService struct {
}

func NewOrdersService() *OrdersService {
	return &OrdersService{}
}

var OrderReqDB = make([]*grpcs.CreateOrderRequest, 0)

func (h *OrdersService) CreateOrder(ctx context.Context, order *grpcs.CreateOrderRequest) error {
	OrderReqDB = append(OrderReqDB, order)

	return nil
}

func (h *OrdersService) GetOrder(ctx context.Context) []*grpcs.CreateOrderRequest {
	for _, orderReq := range OrderReqDB {
		start := time.Now()

		for _, order := range orderReq.Orders {
			time.Sleep(time.Second * 1)
			order.Status = true
		}

		orderReq.Status = true
		elapsed := time.Since(start)
		log.Printf("order request %v : done with %v", orderReq.OrderRequestID, elapsed.Seconds())
	} 
	
	return OrderReqDB
}