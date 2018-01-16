package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	chattyPb "github.com/ThatTomPerson/home/internal/api/chatty"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/mdns"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("chatty"),
		micro.Version("latest"),
		micro.Registry(mdns.NewRegistry()),
	)

	// Init will parse the command line flags. Any flags set will
	// override the above settings. Options defined here will
	// override anything set on the command line.
	service.Init()

	http.Handle("/", handle(service))

	http.ListenAndServe(":8082", nil)
}

func handle(service micro.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		name := r.Form.Get("name")
		if name == "" {
			name = "Tom"
		}

		// create the greeter client using the service name and client
		chatty := chattyPb.NewChattyClient("chatty", service.Client())

		// request the Hello method on the Chatty handler
		rsp, err := chatty.Hello(context.TODO(), &chattyPb.HelloRequest{
			Name: name,
		})

		log.Printf("Rx: %s\n", rsp.Greeting)

		if err != nil {
			fmt.Println(err)
			return
		}

		w.Write([]byte(rsp.Greeting))
	})
}
