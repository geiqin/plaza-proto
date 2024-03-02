// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: dictionayValueService.proto

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

// Api Endpoints for DictionaryValueService service

func NewDictionaryValueServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DictionaryValueService service

type DictionaryValueService interface {
	Create(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error)
	Update(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error)
	Delete(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error)
	Get(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error)
	UpdateSort(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error)
	UpdateEnabled(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error)
}

type dictionaryValueService struct {
	c    client.Client
	name string
}

func NewDictionaryValueService(name string, c client.Client) DictionaryValueService {
	return &dictionaryValueService{
		c:    c,
		name: name,
	}
}

func (c *dictionaryValueService) Create(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error) {
	req := c.c.NewRequest(c.name, "DictionaryValueService.Create", in)
	out := new(DictionaryValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryValueService) Update(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error) {
	req := c.c.NewRequest(c.name, "DictionaryValueService.Update", in)
	out := new(DictionaryValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryValueService) Delete(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error) {
	req := c.c.NewRequest(c.name, "DictionaryValueService.Delete", in)
	out := new(DictionaryValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryValueService) Get(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error) {
	req := c.c.NewRequest(c.name, "DictionaryValueService.Get", in)
	out := new(DictionaryValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryValueService) UpdateSort(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error) {
	req := c.c.NewRequest(c.name, "DictionaryValueService.UpdateSort", in)
	out := new(DictionaryValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryValueService) UpdateEnabled(ctx context.Context, in *DictionaryValue, opts ...client.CallOption) (*DictionaryValueResponse, error) {
	req := c.c.NewRequest(c.name, "DictionaryValueService.UpdateEnabled", in)
	out := new(DictionaryValueResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DictionaryValueService service

type DictionaryValueServiceHandler interface {
	Create(context.Context, *DictionaryValue, *DictionaryValueResponse) error
	Update(context.Context, *DictionaryValue, *DictionaryValueResponse) error
	Delete(context.Context, *DictionaryValue, *DictionaryValueResponse) error
	Get(context.Context, *DictionaryValue, *DictionaryValueResponse) error
	UpdateSort(context.Context, *DictionaryValue, *DictionaryValueResponse) error
	UpdateEnabled(context.Context, *DictionaryValue, *DictionaryValueResponse) error
}

func RegisterDictionaryValueServiceHandler(s server.Server, hdlr DictionaryValueServiceHandler, opts ...server.HandlerOption) error {
	type dictionaryValueService interface {
		Create(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error
		Update(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error
		Delete(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error
		Get(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error
		UpdateSort(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error
		UpdateEnabled(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error
	}
	type DictionaryValueService struct {
		dictionaryValueService
	}
	h := &dictionaryValueServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DictionaryValueService{h}, opts...))
}

type dictionaryValueServiceHandler struct {
	DictionaryValueServiceHandler
}

func (h *dictionaryValueServiceHandler) Create(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error {
	return h.DictionaryValueServiceHandler.Create(ctx, in, out)
}

func (h *dictionaryValueServiceHandler) Update(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error {
	return h.DictionaryValueServiceHandler.Update(ctx, in, out)
}

func (h *dictionaryValueServiceHandler) Delete(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error {
	return h.DictionaryValueServiceHandler.Delete(ctx, in, out)
}

func (h *dictionaryValueServiceHandler) Get(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error {
	return h.DictionaryValueServiceHandler.Get(ctx, in, out)
}

func (h *dictionaryValueServiceHandler) UpdateSort(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error {
	return h.DictionaryValueServiceHandler.UpdateSort(ctx, in, out)
}

func (h *dictionaryValueServiceHandler) UpdateEnabled(ctx context.Context, in *DictionaryValue, out *DictionaryValueResponse) error {
	return h.DictionaryValueServiceHandler.UpdateEnabled(ctx, in, out)
}