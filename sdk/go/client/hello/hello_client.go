// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new hello API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for hello API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SayHello returns a hello world message
*/
func (a *Client) SayHello(params *SayHelloParams) (*SayHelloOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSayHelloParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SayHello",
		Method:             "GET",
		PathPattern:        "/hello",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SayHelloReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SayHelloOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SayHello: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
