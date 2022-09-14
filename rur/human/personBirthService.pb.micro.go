// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: personBirthService.proto

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

// Api Endpoints for PersonBirthService service

func NewPersonBirthServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PersonBirthService service

type PersonBirthService interface {
	//创建出生状况
	Create(ctx context.Context, in *PersonBirth, opts ...client.CallOption) (*PersonBirthResponse, error)
	// 编辑出生状况
	Update(ctx context.Context, in *PersonBirth, opts ...client.CallOption) (*PersonBirthResponse, error)
	// 删除出生状况
	Delete(ctx context.Context, in *PersonBirth, opts ...client.CallOption) (*PersonBirthResponse, error)
	// 获取出生状况详情
	Get(ctx context.Context, in *PersonBirth, opts ...client.CallOption) (*PersonBirthResponse, error)
	//分页查询家庭列表
	Search(ctx context.Context, in *PersonBirthRequest, opts ...client.CallOption) (*PersonBirthResponse, error)
}

type personBirthService struct {
	c    client.Client
	name string
}

func NewPersonBirthService(name string, c client.Client) PersonBirthService {
	return &personBirthService{
		c:    c,
		name: name,
	}
}

func (c *personBirthService) Create(ctx context.Context, in *PersonBirth, opts ...client.CallOption) (*PersonBirthResponse, error) {
	req := c.c.NewRequest(c.name, "PersonBirthService.Create", in)
	out := new(PersonBirthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personBirthService) Update(ctx context.Context, in *PersonBirth, opts ...client.CallOption) (*PersonBirthResponse, error) {
	req := c.c.NewRequest(c.name, "PersonBirthService.Update", in)
	out := new(PersonBirthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personBirthService) Delete(ctx context.Context, in *PersonBirth, opts ...client.CallOption) (*PersonBirthResponse, error) {
	req := c.c.NewRequest(c.name, "PersonBirthService.Delete", in)
	out := new(PersonBirthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personBirthService) Get(ctx context.Context, in *PersonBirth, opts ...client.CallOption) (*PersonBirthResponse, error) {
	req := c.c.NewRequest(c.name, "PersonBirthService.Get", in)
	out := new(PersonBirthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personBirthService) Search(ctx context.Context, in *PersonBirthRequest, opts ...client.CallOption) (*PersonBirthResponse, error) {
	req := c.c.NewRequest(c.name, "PersonBirthService.Search", in)
	out := new(PersonBirthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersonBirthService service

type PersonBirthServiceHandler interface {
	//创建出生状况
	Create(context.Context, *PersonBirth, *PersonBirthResponse) error
	// 编辑出生状况
	Update(context.Context, *PersonBirth, *PersonBirthResponse) error
	// 删除出生状况
	Delete(context.Context, *PersonBirth, *PersonBirthResponse) error
	// 获取出生状况详情
	Get(context.Context, *PersonBirth, *PersonBirthResponse) error
	//分页查询家庭列表
	Search(context.Context, *PersonBirthRequest, *PersonBirthResponse) error
}

func RegisterPersonBirthServiceHandler(s server.Server, hdlr PersonBirthServiceHandler, opts ...server.HandlerOption) error {
	type personBirthService interface {
		Create(ctx context.Context, in *PersonBirth, out *PersonBirthResponse) error
		Update(ctx context.Context, in *PersonBirth, out *PersonBirthResponse) error
		Delete(ctx context.Context, in *PersonBirth, out *PersonBirthResponse) error
		Get(ctx context.Context, in *PersonBirth, out *PersonBirthResponse) error
		Search(ctx context.Context, in *PersonBirthRequest, out *PersonBirthResponse) error
	}
	type PersonBirthService struct {
		personBirthService
	}
	h := &personBirthServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PersonBirthService{h}, opts...))
}

type personBirthServiceHandler struct {
	PersonBirthServiceHandler
}

func (h *personBirthServiceHandler) Create(ctx context.Context, in *PersonBirth, out *PersonBirthResponse) error {
	return h.PersonBirthServiceHandler.Create(ctx, in, out)
}

func (h *personBirthServiceHandler) Update(ctx context.Context, in *PersonBirth, out *PersonBirthResponse) error {
	return h.PersonBirthServiceHandler.Update(ctx, in, out)
}

func (h *personBirthServiceHandler) Delete(ctx context.Context, in *PersonBirth, out *PersonBirthResponse) error {
	return h.PersonBirthServiceHandler.Delete(ctx, in, out)
}

func (h *personBirthServiceHandler) Get(ctx context.Context, in *PersonBirth, out *PersonBirthResponse) error {
	return h.PersonBirthServiceHandler.Get(ctx, in, out)
}

func (h *personBirthServiceHandler) Search(ctx context.Context, in *PersonBirthRequest, out *PersonBirthResponse) error {
	return h.PersonBirthServiceHandler.Search(ctx, in, out)
}