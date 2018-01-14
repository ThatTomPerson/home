package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/ThatTomPerson/home/internal/coinspot"
)

const key = ""
const secret = ""

func main() {
	csp := coinspot.New(key, secret)

	spot, err := csp.Spot()
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}

	spew.Dump(spot)
}
