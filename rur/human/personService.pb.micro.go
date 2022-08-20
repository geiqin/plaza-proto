// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: personService.proto

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

// Api Endpoints for PersonService service

func NewPersonServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PersonService service

type PersonService interface {
	//创建人员
	Create(ctx context.Context, in *Person, opts ...client.CallOption) (*PersonResponse, error)
	// 编辑人员
	Update(ctx context.Context, in *Person, opts ...client.CallOption) (*PersonResponse, error)
	// 删除人员
	Delete(ctx context.Context, in *Person, opts ...client.CallOption) (*PersonResponse, error)
	// 获取人员详情
	Get(ctx context.Context, in *Person, opts ...client.CallOption) (*PersonResponse, error)
	//分页查询人员列表
	Search(ctx context.Context, in *PersonRequest, opts ...client.CallOption) (*PersonResponse, error)
	//获取人口列表
	List(ctx context.Context, in *PersonRequest, opts ...client.CallOption) (*PersonResponse, error)
}

type personService struct {
	c    client.Client
	name string
}

func NewPersonService(name string, c client.Client) PersonService {
	return &personService{
		c:    c,
		name: name,
	}
}

func (c *personService) Create(ctx context.Context, in *Person, opts ...client.CallOption) (*PersonResponse, error) {
	req := c.c.NewRequest(c.name, "PersonService.Create", in)
	out := new(PersonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personService) Update(ctx context.Context, in *Person, opts ...client.CallOption) (*PersonResponse, error) {
	req := c.c.NewRequest(c.name, "PersonService.Update", in)
	out := new(PersonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personService) Delete(ctx context.Context, in *Person, opts ...client.CallOption) (*PersonResponse, error) {
	req := c.c.NewRequest(c.name, "PersonService.Delete", in)
	out := new(PersonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personService) Get(ctx context.Context, in *Person, opts ...client.CallOption) (*PersonResponse, error) {
	req := c.c.NewRequest(c.name, "PersonService.Get", in)
	out := new(PersonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personService) Search(ctx context.Context, in *PersonRequest, opts ...client.CallOption) (*PersonResponse, error) {
	req := c.c.NewRequest(c.name, "PersonService.Search", in)
	out := new(PersonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personService) List(ctx context.Context, in *PersonRequest, opts ...client.CallOption) (*PersonResponse, error) {
	req := c.c.NewRequest(c.name, "PersonService.List", in)
	out := new(PersonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersonService service

type PersonServiceHandler interface {
	//创建人员
	Create(context.Context, *Person, *PersonResponse) error
	// 编辑人员
	Update(context.Context, *Person, *PersonResponse) error
	// 删除人员
	Delete(context.Context, *Person, *PersonResponse) error
	// 获取人员详情
	Get(context.Context, *Person, *PersonResponse) error
	//分页查询人员列表
	Search(context.Context, *PersonRequest, *PersonResponse) error
	//获取人口列表
	List(context.Context, *PersonRequest, *PersonResponse) error
}

func RegisterPersonServiceHandler(s server.Server, hdlr PersonServiceHandler, opts ...server.HandlerOption) error {
	type personService interface {
		Create(ctx context.Context, in *Person, out *PersonResponse) error
		Update(ctx context.Context, in *Person, out *PersonResponse) error
		Delete(ctx context.Context, in *Person, out *PersonResponse) error
		Get(ctx context.Context, in *Person, out *PersonResponse) error
		Search(ctx context.Context, in *PersonRequest, out *PersonResponse) error
		List(ctx context.Context, in *PersonRequest, out *PersonResponse) error
	}
	type PersonService struct {
		personService
	}
	h := &personServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PersonService{h}, opts...))
}

type personServiceHandler struct {
	PersonServiceHandler
}

func (h *personServiceHandler) Create(ctx context.Context, in *Person, out *PersonResponse) error {
	return h.PersonServiceHandler.Create(ctx, in, out)
}

func (h *personServiceHandler) Update(ctx context.Context, in *Person, out *PersonResponse) error {
	return h.PersonServiceHandler.Update(ctx, in, out)
}

func (h *personServiceHandler) Delete(ctx context.Context, in *Person, out *PersonResponse) error {
	return h.PersonServiceHandler.Delete(ctx, in, out)
}

func (h *personServiceHandler) Get(ctx context.Context, in *Person, out *PersonResponse) error {
	return h.PersonServiceHandler.Get(ctx, in, out)
}

func (h *personServiceHandler) Search(ctx context.Context, in *PersonRequest, out *PersonResponse) error {
	return h.PersonServiceHandler.Search(ctx, in, out)
}

func (h *personServiceHandler) List(ctx context.Context, in *PersonRequest, out *PersonResponse) error {
	return h.PersonServiceHandler.List(ctx, in, out)
}
