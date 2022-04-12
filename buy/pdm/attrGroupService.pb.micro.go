// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: attrGroupService.proto

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

// Api Endpoints for AttrGroupService service

func NewAttrGroupServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AttrGroupService service

type AttrGroupService interface {
	Create(ctx context.Context, in *AttrGroup, opts ...client.CallOption) (*AttrGroupResponse, error)
	Update(ctx context.Context, in *AttrGroup, opts ...client.CallOption) (*AttrGroupResponse, error)
	Delete(ctx context.Context, in *AttrGroup, opts ...client.CallOption) (*AttrGroupResponse, error)
	Get(ctx context.Context, in *AttrGroup, opts ...client.CallOption) (*AttrGroupResponse, error)
	Search(ctx context.Context, in *AttrGroupRequest, opts ...client.CallOption) (*AttrGroupResponse, error)
}

type attrGroupService struct {
	c    client.Client
	name string
}

func NewAttrGroupService(name string, c client.Client) AttrGroupService {
	return &attrGroupService{
		c:    c,
		name: name,
	}
}

func (c *attrGroupService) Create(ctx context.Context, in *AttrGroup, opts ...client.CallOption) (*AttrGroupResponse, error) {
	req := c.c.NewRequest(c.name, "AttrGroupService.Create", in)
	out := new(AttrGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrGroupService) Update(ctx context.Context, in *AttrGroup, opts ...client.CallOption) (*AttrGroupResponse, error) {
	req := c.c.NewRequest(c.name, "AttrGroupService.Update", in)
	out := new(AttrGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrGroupService) Delete(ctx context.Context, in *AttrGroup, opts ...client.CallOption) (*AttrGroupResponse, error) {
	req := c.c.NewRequest(c.name, "AttrGroupService.Delete", in)
	out := new(AttrGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrGroupService) Get(ctx context.Context, in *AttrGroup, opts ...client.CallOption) (*AttrGroupResponse, error) {
	req := c.c.NewRequest(c.name, "AttrGroupService.Get", in)
	out := new(AttrGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attrGroupService) Search(ctx context.Context, in *AttrGroupRequest, opts ...client.CallOption) (*AttrGroupResponse, error) {
	req := c.c.NewRequest(c.name, "AttrGroupService.Search", in)
	out := new(AttrGroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AttrGroupService service

type AttrGroupServiceHandler interface {
	Create(context.Context, *AttrGroup, *AttrGroupResponse) error
	Update(context.Context, *AttrGroup, *AttrGroupResponse) error
	Delete(context.Context, *AttrGroup, *AttrGroupResponse) error
	Get(context.Context, *AttrGroup, *AttrGroupResponse) error
	Search(context.Context, *AttrGroupRequest, *AttrGroupResponse) error
}

func RegisterAttrGroupServiceHandler(s server.Server, hdlr AttrGroupServiceHandler, opts ...server.HandlerOption) error {
	type attrGroupService interface {
		Create(ctx context.Context, in *AttrGroup, out *AttrGroupResponse) error
		Update(ctx context.Context, in *AttrGroup, out *AttrGroupResponse) error
		Delete(ctx context.Context, in *AttrGroup, out *AttrGroupResponse) error
		Get(ctx context.Context, in *AttrGroup, out *AttrGroupResponse) error
		Search(ctx context.Context, in *AttrGroupRequest, out *AttrGroupResponse) error
	}
	type AttrGroupService struct {
		attrGroupService
	}
	h := &attrGroupServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AttrGroupService{h}, opts...))
}

type attrGroupServiceHandler struct {
	AttrGroupServiceHandler
}

func (h *attrGroupServiceHandler) Create(ctx context.Context, in *AttrGroup, out *AttrGroupResponse) error {
	return h.AttrGroupServiceHandler.Create(ctx, in, out)
}

func (h *attrGroupServiceHandler) Update(ctx context.Context, in *AttrGroup, out *AttrGroupResponse) error {
	return h.AttrGroupServiceHandler.Update(ctx, in, out)
}

func (h *attrGroupServiceHandler) Delete(ctx context.Context, in *AttrGroup, out *AttrGroupResponse) error {
	return h.AttrGroupServiceHandler.Delete(ctx, in, out)
}

func (h *attrGroupServiceHandler) Get(ctx context.Context, in *AttrGroup, out *AttrGroupResponse) error {
	return h.AttrGroupServiceHandler.Get(ctx, in, out)
}

func (h *attrGroupServiceHandler) Search(ctx context.Context, in *AttrGroupRequest, out *AttrGroupResponse) error {
	return h.AttrGroupServiceHandler.Search(ctx, in, out)
}