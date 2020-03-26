package main

import (
	"fmt"
	"log"

	apiclient "github.com/julianpitt/go-swagger-test/sdk/client"
	helloClient "github.com/julianpitt/go-swagger-test/sdk/client/hello"
)

func main() {
	params := helloClient.NewSayHelloParams()
	resp, err := apiclient.Default.Hello.SayHello(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}
