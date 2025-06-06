// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: storeOperatorService.proto

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

// Api Endpoints for StoreOperatorService service

func NewStoreOperatorServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for StoreOperatorService service

type StoreOperatorService interface {
	// 运营人员保存
	Save(ctx context.Context, in *StoreOperator, opts ...client.CallOption) (*StoreOperatorResponse, error)
	// 运营人员删除
	Delete(ctx context.Context, in *StoreOperator, opts ...client.CallOption) (*StoreOperatorResponse, error)
	// 运营人员获取
	Get(ctx context.Context, in *StoreOperator, opts ...client.CallOption) (*StoreOperatorResponse, error)
	// 运营人员查询
	Search(ctx context.Context, in *StoreOperatorRequest, opts ...client.CallOption) (*StoreOperatorResponse, error)
	// 运营人员列表
	List(ctx context.Context, in *StoreOperatorRequest, opts ...client.CallOption) (*StoreOperatorResponse, error)
}

type storeOperatorService struct {
	c    client.Client
	name string
}

func NewStoreOperatorService(name string, c client.Client) StoreOperatorService {
	return &storeOperatorService{
		c:    c,
		name: name,
	}
}

func (c *storeOperatorService) Save(ctx context.Context, in *StoreOperator, opts ...client.CallOption) (*StoreOperatorResponse, error) {
	req := c.c.NewRequest(c.name, "StoreOperatorService.Save", in)
	out := new(StoreOperatorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeOperatorService) Delete(ctx context.Context, in *StoreOperator, opts ...client.CallOption) (*StoreOperatorResponse, error) {
	req := c.c.NewRequest(c.name, "StoreOperatorService.Delete", in)
	out := new(StoreOperatorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeOperatorService) Get(ctx context.Context, in *StoreOperator, opts ...client.CallOption) (*StoreOperatorResponse, error) {
	req := c.c.NewRequest(c.name, "StoreOperatorService.Get", in)
	out := new(StoreOperatorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeOperatorService) Search(ctx context.Context, in *StoreOperatorRequest, opts ...client.CallOption) (*StoreOperatorResponse, error) {
	req := c.c.NewRequest(c.name, "StoreOperatorService.Search", in)
	out := new(StoreOperatorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeOperatorService) List(ctx context.Context, in *StoreOperatorRequest, opts ...client.CallOption) (*StoreOperatorResponse, error) {
	req := c.c.NewRequest(c.name, "StoreOperatorService.List", in)
	out := new(StoreOperatorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StoreOperatorService service

type StoreOperatorServiceHandler interface {
	// 运营人员保存
	Save(context.Context, *StoreOperator, *StoreOperatorResponse) error
	// 运营人员删除
	Delete(context.Context, *StoreOperator, *StoreOperatorResponse) error
	// 运营人员获取
	Get(context.Context, *StoreOperator, *StoreOperatorResponse) error
	// 运营人员查询
	Search(context.Context, *StoreOperatorRequest, *StoreOperatorResponse) error
	// 运营人员列表
	List(context.Context, *StoreOperatorRequest, *StoreOperatorResponse) error
}

func RegisterStoreOperatorServiceHandler(s server.Server, hdlr StoreOperatorServiceHandler, opts ...server.HandlerOption) error {
	type storeOperatorService interface {
		Save(ctx context.Context, in *StoreOperator, out *StoreOperatorResponse) error
		Delete(ctx context.Context, in *StoreOperator, out *StoreOperatorResponse) error
		Get(ctx context.Context, in *StoreOperator, out *StoreOperatorResponse) error
		Search(ctx context.Context, in *StoreOperatorRequest, out *StoreOperatorResponse) error
		List(ctx context.Context, in *StoreOperatorRequest, out *StoreOperatorResponse) error
	}
	type StoreOperatorService struct {
		storeOperatorService
	}
	h := &storeOperatorServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&StoreOperatorService{h}, opts...))
}

type storeOperatorServiceHandler struct {
	StoreOperatorServiceHandler
}

func (h *storeOperatorServiceHandler) Save(ctx context.Context, in *StoreOperator, out *StoreOperatorResponse) error {
	return h.StoreOperatorServiceHandler.Save(ctx, in, out)
}

func (h *storeOperatorServiceHandler) Delete(ctx context.Context, in *StoreOperator, out *StoreOperatorResponse) error {
	return h.StoreOperatorServiceHandler.Delete(ctx, in, out)
}

func (h *storeOperatorServiceHandler) Get(ctx context.Context, in *StoreOperator, out *StoreOperatorResponse) error {
	return h.StoreOperatorServiceHandler.Get(ctx, in, out)
}

func (h *storeOperatorServiceHandler) Search(ctx context.Context, in *StoreOperatorRequest, out *StoreOperatorResponse) error {
	return h.StoreOperatorServiceHandler.Search(ctx, in, out)
}

func (h *storeOperatorServiceHandler) List(ctx context.Context, in *StoreOperatorRequest, out *StoreOperatorResponse) error {
	return h.StoreOperatorServiceHandler.List(ctx, in, out)
}
