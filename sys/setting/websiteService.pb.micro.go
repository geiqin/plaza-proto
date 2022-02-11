// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: websiteService.proto

package services

import (
	fmt "fmt"
	common "github.com/geiqin/micro-kit/protobuf/common"
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

// Api Endpoints for WebsiteService service

func NewWebsiteServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for WebsiteService service

type WebsiteService interface {
	Set(ctx context.Context, in *Website, opts ...client.CallOption) (*WebsiteResponse, error)
	Get(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*WebsiteResponse, error)
}

type websiteService struct {
	c    client.Client
	name string
}

func NewWebsiteService(name string, c client.Client) WebsiteService {
	return &websiteService{
		c:    c,
		name: name,
	}
}

func (c *websiteService) Set(ctx context.Context, in *Website, opts ...client.CallOption) (*WebsiteResponse, error) {
	req := c.c.NewRequest(c.name, "WebsiteService.Set", in)
	out := new(WebsiteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteService) Get(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*WebsiteResponse, error) {
	req := c.c.NewRequest(c.name, "WebsiteService.Get", in)
	out := new(WebsiteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WebsiteService service

type WebsiteServiceHandler interface {
	Set(context.Context, *Website, *WebsiteResponse) error
	Get(context.Context, *common.Empty, *WebsiteResponse) error
}

func RegisterWebsiteServiceHandler(s server.Server, hdlr WebsiteServiceHandler, opts ...server.HandlerOption) error {
	type websiteService interface {
		Set(ctx context.Context, in *Website, out *WebsiteResponse) error
		Get(ctx context.Context, in *common.Empty, out *WebsiteResponse) error
	}
	type WebsiteService struct {
		websiteService
	}
	h := &websiteServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WebsiteService{h}, opts...))
}

type websiteServiceHandler struct {
	WebsiteServiceHandler
}

func (h *websiteServiceHandler) Set(ctx context.Context, in *Website, out *WebsiteResponse) error {
	return h.WebsiteServiceHandler.Set(ctx, in, out)
}

func (h *websiteServiceHandler) Get(ctx context.Context, in *common.Empty, out *WebsiteResponse) error {
	return h.WebsiteServiceHandler.Get(ctx, in, out)
}
