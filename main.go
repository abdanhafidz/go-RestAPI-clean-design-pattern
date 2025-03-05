package main

import (
	"fmt"

	"go-dp.abdanhafidz.com/config"
	"go-dp.abdanhafidz.com/router"
)

func main() {
	fmt.Println("Server started on ", config.TCP_ADDRESS, ", port :", config.HOST_PORT)
	router.StartService()

}
