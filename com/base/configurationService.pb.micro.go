// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: configurationService.proto

package services

import (
	fmt "fmt"
	_ "github.com/geiqin/micro-kit/protobuf/common"
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

// Api Endpoints for ConfigurationService service

func NewConfigurationServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ConfigurationService service

type ConfigurationService interface {
	Save(ctx context.Context, in *Configuration, opts ...client.CallOption) (*ConfigurationResponse, error)
	Get(ctx context.Context, in *ConfigurationRequest, opts ...client.CallOption) (*ConfigurationResponse, error)
	Search(ctx context.Context, in *ConfigurationRequest, opts ...client.CallOption) (*ConfigurationResponse, error)
	List(ctx context.Context, in *ConfigurationRequest, opts ...client.CallOption) (*ConfigurationResponse, error)
}

type configurationService struct {
	c    client.Client
	name string
}

func NewConfigurationService(name string, c client.Client) ConfigurationService {
	return &configurationService{
		c:    c,
		name: name,
	}
}

func (c *configurationService) Save(ctx context.Context, in *Configuration, opts ...client.CallOption) (*ConfigurationResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigurationService.Save", in)
	out := new(ConfigurationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationService) Get(ctx context.Context, in *ConfigurationRequest, opts ...client.CallOption) (*ConfigurationResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigurationService.Get", in)
	out := new(ConfigurationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationService) Search(ctx context.Context, in *ConfigurationRequest, opts ...client.CallOption) (*ConfigurationResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigurationService.Search", in)
	out := new(ConfigurationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationService) List(ctx context.Context, in *ConfigurationRequest, opts ...client.CallOption) (*ConfigurationResponse, error) {
	req := c.c.NewRequest(c.name, "ConfigurationService.List", in)
	out := new(ConfigurationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConfigurationService service

type ConfigurationServiceHandler interface {
	Save(context.Context, *Configuration, *ConfigurationResponse) error
	Get(context.Context, *ConfigurationRequest, *ConfigurationResponse) error
	Search(context.Context, *ConfigurationRequest, *ConfigurationResponse) error
	List(context.Context, *ConfigurationRequest, *ConfigurationResponse) error
}

func RegisterConfigurationServiceHandler(s server.Server, hdlr ConfigurationServiceHandler, opts ...server.HandlerOption) error {
	type configurationService interface {
		Save(ctx context.Context, in *Configuration, out *ConfigurationResponse) error
		Get(ctx context.Context, in *ConfigurationRequest, out *ConfigurationResponse) error
		Search(ctx context.Context, in *ConfigurationRequest, out *ConfigurationResponse) error
		List(ctx context.Context, in *ConfigurationRequest, out *ConfigurationResponse) error
	}
	type ConfigurationService struct {
		configurationService
	}
	h := &configurationServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ConfigurationService{h}, opts...))
}

type configurationServiceHandler struct {
	ConfigurationServiceHandler
}

func (h *configurationServiceHandler) Save(ctx context.Context, in *Configuration, out *ConfigurationResponse) error {
	return h.ConfigurationServiceHandler.Save(ctx, in, out)
}

func (h *configurationServiceHandler) Get(ctx context.Context, in *ConfigurationRequest, out *ConfigurationResponse) error {
	return h.ConfigurationServiceHandler.Get(ctx, in, out)
}

func (h *configurationServiceHandler) Search(ctx context.Context, in *ConfigurationRequest, out *ConfigurationResponse) error {
	return h.ConfigurationServiceHandler.Search(ctx, in, out)
}

func (h *configurationServiceHandler) List(ctx context.Context, in *ConfigurationRequest, out *ConfigurationResponse) error {
	return h.ConfigurationServiceHandler.List(ctx, in, out)
}