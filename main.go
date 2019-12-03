package main

import (
	"github.com/ipan97/go-dig-sample/config"
	"github.com/ipan97/go-dig-sample/database"
	"github.com/ipan97/go-dig-sample/repository"
	"github.com/ipan97/go-dig-sample/server"
	"github.com/ipan97/go-dig-sample/service"
	"go.uber.org/dig"
)

func main() {
	container := BuildContainer()
	err := container.Invoke(func(server *server.Server) {
		server.Run()
	})
	if err != nil {
		panic(err)
	}
}

func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(config.NewConfig)
	_ = container.Provide(database.Connect)
	_ = container.Provide(repository.NewPersonRepository)
	_ = container.Provide(service.NewPersonService)
	_ = container.Provide(server.NewServer)
	return container
}
