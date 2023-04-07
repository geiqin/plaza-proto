// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: configureService.proto

package services

import (
	_ "../common"
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

// Api Endpoints for ConfigureService service

func NewConfigureServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ConfigureService service

type ConfigureService interface {
	//获取分销配置信息
	Get(ctx context.Context, in *BaseConfig, opts ...client.CallOption) (*ConfigureResponse, error)
	//设置销分销配置信息
	Set(ctx context.Context, in *BaseConfig, opts ...client.CallOption) (*ConfigureResponse, error)
}

type configureService struct {
	c    client.Client
	name string
}

func NewConfigureService(name string, c client.Client) ConfigureService {
	return &configureService{
		c:    c,
		name: name,
	}
}

func (c *configureService) Get(ctx context.Context, in *BaseConfig, opts ...client.CallOption) (*ConfigureResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigureService.Get", in)
	out := new(ConfigureResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configureService) Set(ctx context.Context, in *BaseConfig, opts ...client.CallOption) (*ConfigureResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigureService.Set", in)
	out := new(ConfigureResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConfigureService service

type ConfigureServiceHandler interface {
	//获取分销配置信息
	Get(context.Context, *BaseConfig, *ConfigureResponse) error
	//设置销分销配置信息
	Set(context.Context, *BaseConfig, *ConfigureResponse) error
}

func RegisterConfigureServiceHandler(s server.Server, hdlr ConfigureServiceHandler, opts ...server.HandlerOption) error {
	type configureService interface {
		Get(ctx context.Context, in *BaseConfig, out *ConfigureResponse) error
		Set(ctx context.Context, in *BaseConfig, out *ConfigureResponse) error
	}
	type ConfigureService struct {
		configureService
	}
	h := &configureServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ConfigureService{h}, opts...))
}

type configureServiceHandler struct {
	ConfigureServiceHandler
}

func (h *configureServiceHandler) Get(ctx context.Context, in *BaseConfig, out *ConfigureResponse) error {
	return h.ConfigureServiceHandler.Get(ctx, in, out)
}

func (h *configureServiceHandler) Set(ctx context.Context, in *BaseConfig, out *ConfigureResponse) error {
	return h.ConfigureServiceHandler.Set(ctx, in, out)
}