package main

import (
	"fmt"
	"log"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/julianpitt/go-swagger-test/sdk/go/client"
	helloClient "github.com/julianpitt/go-swagger-test/sdk/go/client/hello"
)

func main() {
	transport := httptransport.New(os.Getenv("HELLO_HOST"), "", nil)
	client := apiclient.New(transport, strfmt.Default)

	resp, err := client.Hello.SayHello(helloClient.NewSayHelloParams())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp.Payload)
}
