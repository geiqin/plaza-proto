// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: authorizeService.proto

package services

import (
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

// Api Endpoints for AuthorizeService service

func NewAuthorizeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AuthorizeService service

type AuthorizeService interface {
	//获取当前店铺(简单信息)
	Info(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*StoreResponse, error)
	//获取当前店铺(详细信息)
	Detail(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*StoreResponse, error)
}

type authorizeService struct {
	c    client.Client
	name string
}

func NewAuthorizeService(name string, c client.Client) AuthorizeService {
	return &authorizeService{
		c:    c,
		name: name,
	}
}

func (c *authorizeService) Info(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.Info", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizeService) Detail(ctx context.Context, in *AuthorizeRequest, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "AuthorizeService.Detail", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthorizeService service

type AuthorizeServiceHandler interface {
	//获取当前店铺(简单信息)
	Info(context.Context, *AuthorizeRequest, *StoreResponse) error
	//获取当前店铺(详细信息)
	Detail(context.Context, *AuthorizeRequest, *StoreResponse) error
}

func RegisterAuthorizeServiceHandler(s server.Server, hdlr AuthorizeServiceHandler, opts ...server.HandlerOption) error {
	type authorizeService interface {
		Info(ctx context.Context, in *AuthorizeRequest, out *StoreResponse) error
		Detail(ctx context.Context, in *AuthorizeRequest, out *StoreResponse) error
	}
	type AuthorizeService struct {
		authorizeService
	}
	h := &authorizeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AuthorizeService{h}, opts...))
}

type authorizeServiceHandler struct {
	AuthorizeServiceHandler
}

func (h *authorizeServiceHandler) Info(ctx context.Context, in *AuthorizeRequest, out *StoreResponse) error {
	return h.AuthorizeServiceHandler.Info(ctx, in, out)
}

func (h *authorizeServiceHandler) Detail(ctx context.Context, in *AuthorizeRequest, out *StoreResponse) error {
	return h.AuthorizeServiceHandler.Detail(ctx, in, out)
}
