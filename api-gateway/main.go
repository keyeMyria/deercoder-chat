package main

import (
	"deercoder-chat/api-gateway/routers"
	"github.com/dreamlu/deercoder-gin"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-web"
	"log"
)

//var (
//	UserClient user.UserService
//)

func main() {

	// registry
	registry := consul.NewRegistry(consul.Config(
		&api.Config{
			Address: deercoder.GetDevModeConfig("consul.address"),
			Scheme:  deercoder.GetDevModeConfig("consul.scheme"),
		}))

	// Create service
	service := web.NewService(
		web.Name("deercoder-chat.api"),
		web.Registry(registry),
		web.Address(":"+deercoder.GetDevModeConfig("http_port")),
	)

	_ = service.Init()

	// Create RESTful handler (using Gin)
	// Register Handler
	gin.SetMode(gin.DebugMode)
	service.Handle("/", routers.SetRouter())

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
