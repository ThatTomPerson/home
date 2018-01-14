package main

import (
	"context"
	"log"

	proto "github.com/ThatTomPerson/home/internal/api/chatty"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/mdns"
)

type Chatty struct{}

func (g *Chatty) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("chatty"),
		micro.Version("latest"),
		micro.Registry(mdns.NewRegistry()),
	)

	service.Init()

	proto.RegisterChattyHandler(service.Server(), new(Chatty))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
