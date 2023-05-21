// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: moduleService.proto

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

// Api Endpoints for ModuleService service

func NewModuleServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ModuleService service

type ModuleService interface {
	Create(ctx context.Context, in *Module, opts ...client.CallOption) (*ModuleResponse, error)
	Update(ctx context.Context, in *Module, opts ...client.CallOption) (*ModuleResponse, error)
	Delete(ctx context.Context, in *Module, opts ...client.CallOption) (*ModuleResponse, error)
	Get(ctx context.Context, in *Module, opts ...client.CallOption) (*ModuleResponse, error)
	Search(ctx context.Context, in *ModuleRequest, opts ...client.CallOption) (*ModuleResponse, error)
	List(ctx context.Context, in *ModuleRequest, opts ...client.CallOption) (*ModuleResponse, error)
}

type moduleService struct {
	c    client.Client
	name string
}

func NewModuleService(name string, c client.Client) ModuleService {
	return &moduleService{
		c:    c,
		name: name,
	}
}

func (c *moduleService) Create(ctx context.Context, in *Module, opts ...client.CallOption) (*ModuleResponse, error) {
	req := c.c.NewRequest(c.name, "ModuleService.Create", in)
	out := new(ModuleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleService) Update(ctx context.Context, in *Module, opts ...client.CallOption) (*ModuleResponse, error) {
	req := c.c.NewRequest(c.name, "ModuleService.Update", in)
	out := new(ModuleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleService) Delete(ctx context.Context, in *Module, opts ...client.CallOption) (*ModuleResponse, error) {
	req := c.c.NewRequest(c.name, "ModuleService.Delete", in)
	out := new(ModuleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleService) Get(ctx context.Context, in *Module, opts ...client.CallOption) (*ModuleResponse, error) {
	req := c.c.NewRequest(c.name, "ModuleService.Get", in)
	out := new(ModuleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleService) Search(ctx context.Context, in *ModuleRequest, opts ...client.CallOption) (*ModuleResponse, error) {
	req := c.c.NewRequest(c.name, "ModuleService.Search", in)
	out := new(ModuleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moduleService) List(ctx context.Context, in *ModuleRequest, opts ...client.CallOption) (*ModuleResponse, error) {
	req := c.c.NewRequest(c.name, "ModuleService.List", in)
	out := new(ModuleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ModuleService service

type ModuleServiceHandler interface {
	Create(context.Context, *Module, *ModuleResponse) error
	Update(context.Context, *Module, *ModuleResponse) error
	Delete(context.Context, *Module, *ModuleResponse) error
	Get(context.Context, *Module, *ModuleResponse) error
	Search(context.Context, *ModuleRequest, *ModuleResponse) error
	List(context.Context, *ModuleRequest, *ModuleResponse) error
}

func RegisterModuleServiceHandler(s server.Server, hdlr ModuleServiceHandler, opts ...server.HandlerOption) error {
	type moduleService interface {
		Create(ctx context.Context, in *Module, out *ModuleResponse) error
		Update(ctx context.Context, in *Module, out *ModuleResponse) error
		Delete(ctx context.Context, in *Module, out *ModuleResponse) error
		Get(ctx context.Context, in *Module, out *ModuleResponse) error
		Search(ctx context.Context, in *ModuleRequest, out *ModuleResponse) error
		List(ctx context.Context, in *ModuleRequest, out *ModuleResponse) error
	}
	type ModuleService struct {
		moduleService
	}
	h := &moduleServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ModuleService{h}, opts...))
}

type moduleServiceHandler struct {
	ModuleServiceHandler
}

func (h *moduleServiceHandler) Create(ctx context.Context, in *Module, out *ModuleResponse) error {
	return h.ModuleServiceHandler.Create(ctx, in, out)
}

func (h *moduleServiceHandler) Update(ctx context.Context, in *Module, out *ModuleResponse) error {
	return h.ModuleServiceHandler.Update(ctx, in, out)
}

func (h *moduleServiceHandler) Delete(ctx context.Context, in *Module, out *ModuleResponse) error {
	return h.ModuleServiceHandler.Delete(ctx, in, out)
}

func (h *moduleServiceHandler) Get(ctx context.Context, in *Module, out *ModuleResponse) error {
	return h.ModuleServiceHandler.Get(ctx, in, out)
}

func (h *moduleServiceHandler) Search(ctx context.Context, in *ModuleRequest, out *ModuleResponse) error {
	return h.ModuleServiceHandler.Search(ctx, in, out)
}

func (h *moduleServiceHandler) List(ctx context.Context, in *ModuleRequest, out *ModuleResponse) error {
	return h.ModuleServiceHandler.List(ctx, in, out)
}
