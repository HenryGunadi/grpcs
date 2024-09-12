package main

import (
	"log"

	"github.com/HenryGunadi/grpcs/services/orders/api"
)

func main() {
	// start http server
	orderAPIServer := api.NewOrderAPIServer(":8000")
	if err := orderAPIServer.Run(); err != nil {
		log.Fatal("error running order api server : ", err)
	}
}