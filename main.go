package main

import (
	"fmt"

	"godp.abdanhafidz.com/config"
	"godp.abdanhafidz.com/router"
)

func main() {
	fmt.Println("Server started on ", config.TCP_ADDRESS, ", port :", config.HOST_PORT)
	router.StartService()

}
