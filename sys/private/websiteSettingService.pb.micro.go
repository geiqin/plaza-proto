// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: websiteSettingService.proto

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

// Api Endpoints for WebsiteSettingService service

func NewWebsiteSettingServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for WebsiteSettingService service

type WebsiteSettingService interface {
	Set(ctx context.Context, in *WebsiteSetting, opts ...client.CallOption) (*WebsiteSettingResponse, error)
	Get(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*WebsiteSettingResponse, error)
}

type websiteSettingService struct {
	c    client.Client
	name string
}

func NewWebsiteSettingService(name string, c client.Client) WebsiteSettingService {
	return &websiteSettingService{
		c:    c,
		name: name,
	}
}

func (c *websiteSettingService) Set(ctx context.Context, in *WebsiteSetting, opts ...client.CallOption) (*WebsiteSettingResponse, error) {
	req := c.c.NewRequest(c.name, "WebsiteSettingService.Set", in)
	out := new(WebsiteSettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *websiteSettingService) Get(ctx context.Context, in *common.Empty, opts ...client.CallOption) (*WebsiteSettingResponse, error) {
	req := c.c.NewRequest(c.name, "WebsiteSettingService.Get", in)
	out := new(WebsiteSettingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WebsiteSettingService service

type WebsiteSettingServiceHandler interface {
	Set(context.Context, *WebsiteSetting, *WebsiteSettingResponse) error
	Get(context.Context, *common.Empty, *WebsiteSettingResponse) error
}

func RegisterWebsiteSettingServiceHandler(s server.Server, hdlr WebsiteSettingServiceHandler, opts ...server.HandlerOption) error {
	type websiteSettingService interface {
		Set(ctx context.Context, in *WebsiteSetting, out *WebsiteSettingResponse) error
		Get(ctx context.Context, in *common.Empty, out *WebsiteSettingResponse) error
	}
	type WebsiteSettingService struct {
		websiteSettingService
	}
	h := &websiteSettingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WebsiteSettingService{h}, opts...))
}

type websiteSettingServiceHandler struct {
	WebsiteSettingServiceHandler
}

func (h *websiteSettingServiceHandler) Set(ctx context.Context, in *WebsiteSetting, out *WebsiteSettingResponse) error {
	return h.WebsiteSettingServiceHandler.Set(ctx, in, out)
}

func (h *websiteSettingServiceHandler) Get(ctx context.Context, in *common.Empty, out *WebsiteSettingResponse) error {
	return h.WebsiteSettingServiceHandler.Get(ctx, in, out)
}
