// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: taxonomyService.proto

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

// Api Endpoints for TaxonomyService service

func NewTaxonomyServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TaxonomyService service

type TaxonomyService interface {
	Create(ctx context.Context, in *Taxonomy, opts ...client.CallOption) (*TaxonomyResponse, error)
	Update(ctx context.Context, in *Taxonomy, opts ...client.CallOption) (*TaxonomyResponse, error)
	Get(ctx context.Context, in *Taxonomy, opts ...client.CallOption) (*TaxonomyResponse, error)
	Delete(ctx context.Context, in *Taxonomy, opts ...client.CallOption) (*TaxonomyResponse, error)
	Search(ctx context.Context, in *TaxonomyRequest, opts ...client.CallOption) (*TaxonomyResponse, error)
	List(ctx context.Context, in *TaxonomyRequest, opts ...client.CallOption) (*TaxonomyResponse, error)
	Tree(ctx context.Context, in *TaxonomyRequest, opts ...client.CallOption) (*TaxonomyResponse, error)
}

type taxonomyService struct {
	c    client.Client
	name string
}

func NewTaxonomyService(name string, c client.Client) TaxonomyService {
	return &taxonomyService{
		c:    c,
		name: name,
	}
}

func (c *taxonomyService) Create(ctx context.Context, in *Taxonomy, opts ...client.CallOption) (*TaxonomyResponse, error) {
	req := c.c.NewRequest(c.name, "TaxonomyService.Create", in)
	out := new(TaxonomyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxonomyService) Update(ctx context.Context, in *Taxonomy, opts ...client.CallOption) (*TaxonomyResponse, error) {
	req := c.c.NewRequest(c.name, "TaxonomyService.Update", in)
	out := new(TaxonomyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxonomyService) Get(ctx context.Context, in *Taxonomy, opts ...client.CallOption) (*TaxonomyResponse, error) {
	req := c.c.NewRequest(c.name, "TaxonomyService.Get", in)
	out := new(TaxonomyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxonomyService) Delete(ctx context.Context, in *Taxonomy, opts ...client.CallOption) (*TaxonomyResponse, error) {
	req := c.c.NewRequest(c.name, "TaxonomyService.Delete", in)
	out := new(TaxonomyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxonomyService) Search(ctx context.Context, in *TaxonomyRequest, opts ...client.CallOption) (*TaxonomyResponse, error) {
	req := c.c.NewRequest(c.name, "TaxonomyService.Search", in)
	out := new(TaxonomyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxonomyService) List(ctx context.Context, in *TaxonomyRequest, opts ...client.CallOption) (*TaxonomyResponse, error) {
	req := c.c.NewRequest(c.name, "TaxonomyService.List", in)
	out := new(TaxonomyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxonomyService) Tree(ctx context.Context, in *TaxonomyRequest, opts ...client.CallOption) (*TaxonomyResponse, error) {
	req := c.c.NewRequest(c.name, "TaxonomyService.Tree", in)
	out := new(TaxonomyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaxonomyService service

type TaxonomyServiceHandler interface {
	Create(context.Context, *Taxonomy, *TaxonomyResponse) error
	Update(context.Context, *Taxonomy, *TaxonomyResponse) error
	Get(context.Context, *Taxonomy, *TaxonomyResponse) error
	Delete(context.Context, *Taxonomy, *TaxonomyResponse) error
	Search(context.Context, *TaxonomyRequest, *TaxonomyResponse) error
	List(context.Context, *TaxonomyRequest, *TaxonomyResponse) error
	Tree(context.Context, *TaxonomyRequest, *TaxonomyResponse) error
}

func RegisterTaxonomyServiceHandler(s server.Server, hdlr TaxonomyServiceHandler, opts ...server.HandlerOption) error {
	type taxonomyService interface {
		Create(ctx context.Context, in *Taxonomy, out *TaxonomyResponse) error
		Update(ctx context.Context, in *Taxonomy, out *TaxonomyResponse) error
		Get(ctx context.Context, in *Taxonomy, out *TaxonomyResponse) error
		Delete(ctx context.Context, in *Taxonomy, out *TaxonomyResponse) error
		Search(ctx context.Context, in *TaxonomyRequest, out *TaxonomyResponse) error
		List(ctx context.Context, in *TaxonomyRequest, out *TaxonomyResponse) error
		Tree(ctx context.Context, in *TaxonomyRequest, out *TaxonomyResponse) error
	}
	type TaxonomyService struct {
		taxonomyService
	}
	h := &taxonomyServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TaxonomyService{h}, opts...))
}

type taxonomyServiceHandler struct {
	TaxonomyServiceHandler
}

func (h *taxonomyServiceHandler) Create(ctx context.Context, in *Taxonomy, out *TaxonomyResponse) error {
	return h.TaxonomyServiceHandler.Create(ctx, in, out)
}

func (h *taxonomyServiceHandler) Update(ctx context.Context, in *Taxonomy, out *TaxonomyResponse) error {
	return h.TaxonomyServiceHandler.Update(ctx, in, out)
}

func (h *taxonomyServiceHandler) Get(ctx context.Context, in *Taxonomy, out *TaxonomyResponse) error {
	return h.TaxonomyServiceHandler.Get(ctx, in, out)
}

func (h *taxonomyServiceHandler) Delete(ctx context.Context, in *Taxonomy, out *TaxonomyResponse) error {
	return h.TaxonomyServiceHandler.Delete(ctx, in, out)
}

func (h *taxonomyServiceHandler) Search(ctx context.Context, in *TaxonomyRequest, out *TaxonomyResponse) error {
	return h.TaxonomyServiceHandler.Search(ctx, in, out)
}

func (h *taxonomyServiceHandler) List(ctx context.Context, in *TaxonomyRequest, out *TaxonomyResponse) error {
	return h.TaxonomyServiceHandler.List(ctx, in, out)
}

func (h *taxonomyServiceHandler) Tree(ctx context.Context, in *TaxonomyRequest, out *TaxonomyResponse) error {
	return h.TaxonomyServiceHandler.Tree(ctx, in, out)
}
