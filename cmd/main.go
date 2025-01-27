package main

import (
	"fmt"

	"github.com/abhinandpn/GoCassandra/pkg/Router"
)

func main() {
	fmt.Println("Starting Server....")
	Router.StartServer()
}
