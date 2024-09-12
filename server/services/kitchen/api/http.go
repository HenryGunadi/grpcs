package api

import (
	"log"
	"net/http"

	"github.com/HenryGunadi/grpcs/services/common"
	"github.com/HenryGunadi/grpcs/services/kitchen/handler"
	"github.com/go-chi/cors"
)

type APIServer struct {
	addr string
}

func NewKitchenAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	// create router
	router := http.NewServeMux()

	// register new grpc client
	conn := common.NewGRPCClient(":9000")
	defer conn.Close()
	
	// kitchen client service
	kitchenHandler := handler.NewKitchenHandler(conn)
	kitchenHandler.RegisterRoutes(router)

	// cors settings
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		AllowedMethods: []string{"POST", "GET", "DELETE", "PATCH"},
		AllowCredentials: true,
	})

	corsHandler := c.Handler(router)

	log.Println("Listening to kicthen server on port", s.addr)
	
	return http.ListenAndServe(s.addr, corsHandler)
}