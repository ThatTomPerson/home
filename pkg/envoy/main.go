// Envoy
package main

import (
	"log"

	"github.com/micro/go-micro/registry/mdns"
)

func main() {
	s := mdns.NewRegistry()

	w, err := s.Watch()
	if err != nil {
		log.Fatal(err)
	}
	for {
		r, err := w.Next()
		if err != nil {
			log.Print(err)
			w.Stop()
		}

		log.Printf("%s %s %s", r.Action, r.Service.Name, r.Service.Version)
	}
}
