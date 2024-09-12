package main

import (
	"log"

	"github.com/HenryGunadi/grpcs/services/kitchen/api"
)

func main() {
	// start kitchen server
	kitchenAPIServer := api.NewKitchenAPIServer(":3000")
	if err := kitchenAPIServer.Run(); err != nil {
		log.Fatal(err)
	}
}