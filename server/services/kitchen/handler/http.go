package handler

import (
	"fmt"
	"log"
	"net/http"

	grpcs "github.com/HenryGunadi/grpcs/services/common/genproto/mcdonald"
	"github.com/HenryGunadi/grpcs/services/common/utils"
	"google.golang.org/grpc"
)

type Handler struct {
	grpcClientConn *grpc.ClientConn
}

func NewKitchenHandler(grpcClientConn *grpc.ClientConn) *Handler {
	return &Handler{
		grpcClientConn: grpcClientConn,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/kitchen", h.handleOrder)
}

func (h *Handler) handleOrder(w http.ResponseWriter, r *http.Request) {
	c := grpcs.NewOrderServiceClient(h.grpcClientConn)

	// get order requets from grpc server
	ordersRequestDone, err := c.GetOrder(r.Context(), &grpcs.GetOrderRequest{})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("doesnt received finished order request yet"))
	}
	log.Println("kitchen handling order request : ", ordersRequestDone)

	utils.WriteJSON(w, http.StatusOK, ordersRequestDone)
}