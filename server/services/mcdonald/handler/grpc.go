package handler

import (
	"context"

	grpcs "github.com/HenryGunadi/grpcs/services/common/genproto/mcdonald"
	"github.com/HenryGunadi/grpcs/services/mcdonald/types"
	"google.golang.org/grpc"
)

type GRPCHandler struct {
	orderService types.OrderService
	grpcs.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server, orderService types.OrderService) {
	grpcHandler := &GRPCHandler{
		orderService: orderService,
	}

	grpcs.RegisterOrderServiceServer(grpcServer, grpcHandler)
}

func (h *GRPCHandler) CreateOrder(ctx context.Context, req *grpcs.CreateOrderRequest) (*grpcs.OrderRequestResponse, error) {
	// create order
	err := h.orderService.CreateOrder(ctx, req)
	if err != nil {
		return nil, err
	}

	response := &grpcs.OrderRequestResponse{
		Message: "order added",
	}

	return response, nil
}

func (h *GRPCHandler) GetOrder(ctx context.Context, orderReq *grpcs.GetOrderRequest) (*grpcs.OrderResponse, error) {
	orderRequests := h.orderService.GetOrder(ctx)
	
	orderResponse := &grpcs.OrderResponse{
		Message: "order received",
		OrderRequests: orderRequests,
	}

	return orderResponse, nil
}