package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	grpcs "github.com/HenryGunadi/grpcs/services/common/genproto/mcdonald"
	"github.com/HenryGunadi/grpcs/services/common/types"
	"github.com/HenryGunadi/grpcs/services/common/utils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
    grpcClientConn *grpc.ClientConn
}

var (
	orderRequestID int32 = 0;
	orderDBMock = make([]*grpcs.Order, 0)
) 

func NewOrderHTTPHandler(grpcClientConn *grpc.ClientConn) *Handler {
    return &Handler{
        grpcClientConn: grpcClientConn,
    }
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/order", h.createOrder)
}

func (h *Handler) createOrder(w http.ResponseWriter, r *http.Request) {
	var orderPayload types.OrderPayload

	if err := utils.ParseJSON(r, &orderPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON payload"))
		return 
	}

	order := &grpcs.Order{
		OrderID: orderPayload.OrderID,
		ProductID: orderPayload.ProductID,
		Quantity: orderPayload.Quantity,
		Status: orderPayload.Status,
	}

	// create order
	orderDBMock = append(orderDBMock, order)

	c := grpcs.NewOrderServiceClient(h.grpcClientConn)

	ctx, cancel := context.WithTimeout(r.Context(), time.Second * 2)
	defer cancel()

	orderRequestID += 1

	// create order
	orderResponse, err := c.CreateOrder(ctx, &grpcs.CreateOrderRequest{
		OrderRequestID: orderRequestID,
		Orders: orderDBMock,
		Status: false,
		CreatedAt: timestamppb.New(time.Now()),
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error creating an order request : %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, orderResponse)
}