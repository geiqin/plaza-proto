// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: exampleService.proto

package services

import (
	_ "github.com/geiqin/micro-kit/protobuf/common"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ExampleService service

func NewExampleServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ExampleService service

type ExampleService interface {
	Create(ctx context.Context, in *Example, opts ...client.CallOption) (*ExampleResponse, error)
	Update(ctx context.Context, in *Example, opts ...client.CallOption) (*ExampleResponse, error)
	Delete(ctx context.Context, in *Example, opts ...client.CallOption) (*ExampleResponse, error)
	Get(ctx context.Context, in *Example, opts ...client.CallOption) (*ExampleResponse, error)
	Search(ctx context.Context, in *ExampleRequest, opts ...client.CallOption) (*ExampleResponse, error)
}

type exampleService struct {
	c    client.Client
	name string
}

func NewExampleService(name string, c client.Client) ExampleService {
	return &exampleService{
		c:    c,
		name: name,
	}
}

func (c *exampleService) Create(ctx context.Context, in *Example, opts ...client.CallOption) (*ExampleResponse, error) {
	req := c.c.NewRequest(c.name, "ExampleService.Create", in)
	out := new(ExampleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) Update(ctx context.Context, in *Example, opts ...client.CallOption) (*ExampleResponse, error) {
	req := c.c.NewRequest(c.name, "ExampleService.Update", in)
	out := new(ExampleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) Delete(ctx context.Context, in *Example, opts ...client.CallOption) (*ExampleResponse, error) {
	req := c.c.NewRequest(c.name, "ExampleService.Delete", in)
	out := new(ExampleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) Get(ctx context.Context, in *Example, opts ...client.CallOption) (*ExampleResponse, error) {
	req := c.c.NewRequest(c.name, "ExampleService.Get", in)
	out := new(ExampleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) Search(ctx context.Context, in *ExampleRequest, opts ...client.CallOption) (*ExampleResponse, error) {
	req := c.c.NewRequest(c.name, "ExampleService.Search", in)
	out := new(ExampleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExampleService service

type ExampleServiceHandler interface {
	Create(context.Context, *Example, *ExampleResponse) error
	Update(context.Context, *Example, *ExampleResponse) error
	Delete(context.Context, *Example, *ExampleResponse) error
	Get(context.Context, *Example, *ExampleResponse) error
	Search(context.Context, *ExampleRequest, *ExampleResponse) error
}

func RegisterExampleServiceHandler(s server.Server, hdlr ExampleServiceHandler, opts ...server.HandlerOption) error {
	type exampleService interface {
		Create(ctx context.Context, in *Example, out *ExampleResponse) error
		Update(ctx context.Context, in *Example, out *ExampleResponse) error
		Delete(ctx context.Context, in *Example, out *ExampleResponse) error
		Get(ctx context.Context, in *Example, out *ExampleResponse) error
		Search(ctx context.Context, in *ExampleRequest, out *ExampleResponse) error
	}
	type ExampleService struct {
		exampleService
	}
	h := &exampleServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ExampleService{h}, opts...))
}

type exampleServiceHandler struct {
	ExampleServiceHandler
}

func (h *exampleServiceHandler) Create(ctx context.Context, in *Example, out *ExampleResponse) error {
	return h.ExampleServiceHandler.Create(ctx, in, out)
}

func (h *exampleServiceHandler) Update(ctx context.Context, in *Example, out *ExampleResponse) error {
	return h.ExampleServiceHandler.Update(ctx, in, out)
}

func (h *exampleServiceHandler) Delete(ctx context.Context, in *Example, out *ExampleResponse) error {
	return h.ExampleServiceHandler.Delete(ctx, in, out)
}

func (h *exampleServiceHandler) Get(ctx context.Context, in *Example, out *ExampleResponse) error {
	return h.ExampleServiceHandler.Get(ctx, in, out)
}

func (h *exampleServiceHandler) Search(ctx context.Context, in *ExampleRequest, out *ExampleResponse) error {
	return h.ExampleServiceHandler.Search(ctx, in, out)
}
