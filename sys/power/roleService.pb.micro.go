// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: roleService.proto

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

// Api Endpoints for RoleService service

func NewRoleServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RoleService service

type RoleService interface {
	Create(ctx context.Context, in *Role, opts ...client.CallOption) (*RoleResponse, error)
	Update(ctx context.Context, in *Role, opts ...client.CallOption) (*RoleResponse, error)
	Delete(ctx context.Context, in *Role, opts ...client.CallOption) (*RoleResponse, error)
	Get(ctx context.Context, in *Role, opts ...client.CallOption) (*RoleResponse, error)
	List(ctx context.Context, in *RoleRequest, opts ...client.CallOption) (*RoleResponse, error)
	Search(ctx context.Context, in *RoleRequest, opts ...client.CallOption) (*RoleResponse, error)
}

type roleService struct {
	c    client.Client
	name string
}

func NewRoleService(name string, c client.Client) RoleService {
	return &roleService{
		c:    c,
		name: name,
	}
}

func (c *roleService) Create(ctx context.Context, in *Role, opts ...client.CallOption) (*RoleResponse, error) {
	req := c.c.NewRequest(c.name, "RoleService.Create", in)
	out := new(RoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) Update(ctx context.Context, in *Role, opts ...client.CallOption) (*RoleResponse, error) {
	req := c.c.NewRequest(c.name, "RoleService.Update", in)
	out := new(RoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) Delete(ctx context.Context, in *Role, opts ...client.CallOption) (*RoleResponse, error) {
	req := c.c.NewRequest(c.name, "RoleService.Delete", in)
	out := new(RoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) Get(ctx context.Context, in *Role, opts ...client.CallOption) (*RoleResponse, error) {
	req := c.c.NewRequest(c.name, "RoleService.Get", in)
	out := new(RoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) List(ctx context.Context, in *RoleRequest, opts ...client.CallOption) (*RoleResponse, error) {
	req := c.c.NewRequest(c.name, "RoleService.List", in)
	out := new(RoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleService) Search(ctx context.Context, in *RoleRequest, opts ...client.CallOption) (*RoleResponse, error) {
	req := c.c.NewRequest(c.name, "RoleService.Search", in)
	out := new(RoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoleService service

type RoleServiceHandler interface {
	Create(context.Context, *Role, *RoleResponse) error
	Update(context.Context, *Role, *RoleResponse) error
	Delete(context.Context, *Role, *RoleResponse) error
	Get(context.Context, *Role, *RoleResponse) error
	List(context.Context, *RoleRequest, *RoleResponse) error
	Search(context.Context, *RoleRequest, *RoleResponse) error
}

func RegisterRoleServiceHandler(s server.Server, hdlr RoleServiceHandler, opts ...server.HandlerOption) error {
	type roleService interface {
		Create(ctx context.Context, in *Role, out *RoleResponse) error
		Update(ctx context.Context, in *Role, out *RoleResponse) error
		Delete(ctx context.Context, in *Role, out *RoleResponse) error
		Get(ctx context.Context, in *Role, out *RoleResponse) error
		List(ctx context.Context, in *RoleRequest, out *RoleResponse) error
		Search(ctx context.Context, in *RoleRequest, out *RoleResponse) error
	}
	type RoleService struct {
		roleService
	}
	h := &roleServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RoleService{h}, opts...))
}

type roleServiceHandler struct {
	RoleServiceHandler
}

func (h *roleServiceHandler) Create(ctx context.Context, in *Role, out *RoleResponse) error {
	return h.RoleServiceHandler.Create(ctx, in, out)
}

func (h *roleServiceHandler) Update(ctx context.Context, in *Role, out *RoleResponse) error {
	return h.RoleServiceHandler.Update(ctx, in, out)
}

func (h *roleServiceHandler) Delete(ctx context.Context, in *Role, out *RoleResponse) error {
	return h.RoleServiceHandler.Delete(ctx, in, out)
}

func (h *roleServiceHandler) Get(ctx context.Context, in *Role, out *RoleResponse) error {
	return h.RoleServiceHandler.Get(ctx, in, out)
}

func (h *roleServiceHandler) List(ctx context.Context, in *RoleRequest, out *RoleResponse) error {
	return h.RoleServiceHandler.List(ctx, in, out)
}

func (h *roleServiceHandler) Search(ctx context.Context, in *RoleRequest, out *RoleResponse) error {
	return h.RoleServiceHandler.Search(ctx, in, out)
}
