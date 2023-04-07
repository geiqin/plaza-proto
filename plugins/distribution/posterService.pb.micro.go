// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: posterService.proto

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

// Api Endpoints for PosterService service

func NewPosterServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PosterService service

type PosterService interface {
	//保存推广海报
	SavePosterPromotion(ctx context.Context, in *PosterPromotion, opts ...client.CallOption) (*PosterResponse, error)
	//保存商品海报
	SavePosterGoods(ctx context.Context, in *PosterGoods, opts ...client.CallOption) (*PosterResponse, error)
	//获取推广海报
	GetPosterPromotion(ctx context.Context, in *PosterRequest, opts ...client.CallOption) (*PosterResponse, error)
	//获取推广海报
	GetPosterGoods(ctx context.Context, in *PosterRequest, opts ...client.CallOption) (*PosterResponse, error)
}

type posterService struct {
	c    client.Client
	name string
}

func NewPosterService(name string, c client.Client) PosterService {
	return &posterService{
		c:    c,
		name: name,
	}
}

func (c *posterService) SavePosterPromotion(ctx context.Context, in *PosterPromotion, opts ...client.CallOption) (*PosterResponse, error) {
	req := c.c.NewRequest(c.name, "PosterService.SavePosterPromotion", in)
	out := new(PosterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posterService) SavePosterGoods(ctx context.Context, in *PosterGoods, opts ...client.CallOption) (*PosterResponse, error) {
	req := c.c.NewRequest(c.name, "PosterService.SavePosterGoods", in)
	out := new(PosterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posterService) GetPosterPromotion(ctx context.Context, in *PosterRequest, opts ...client.CallOption) (*PosterResponse, error) {
	req := c.c.NewRequest(c.name, "PosterService.GetPosterPromotion", in)
	out := new(PosterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posterService) GetPosterGoods(ctx context.Context, in *PosterRequest, opts ...client.CallOption) (*PosterResponse, error) {
	req := c.c.NewRequest(c.name, "PosterService.GetPosterGoods", in)
	out := new(PosterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PosterService service

type PosterServiceHandler interface {
	//保存推广海报
	SavePosterPromotion(context.Context, *PosterPromotion, *PosterResponse) error
	//保存商品海报
	SavePosterGoods(context.Context, *PosterGoods, *PosterResponse) error
	//获取推广海报
	GetPosterPromotion(context.Context, *PosterRequest, *PosterResponse) error
	//获取推广海报
	GetPosterGoods(context.Context, *PosterRequest, *PosterResponse) error
}

func RegisterPosterServiceHandler(s server.Server, hdlr PosterServiceHandler, opts ...server.HandlerOption) error {
	type posterService interface {
		SavePosterPromotion(ctx context.Context, in *PosterPromotion, out *PosterResponse) error
		SavePosterGoods(ctx context.Context, in *PosterGoods, out *PosterResponse) error
		GetPosterPromotion(ctx context.Context, in *PosterRequest, out *PosterResponse) error
		GetPosterGoods(ctx context.Context, in *PosterRequest, out *PosterResponse) error
	}
	type PosterService struct {
		posterService
	}
	h := &posterServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PosterService{h}, opts...))
}

type posterServiceHandler struct {
	PosterServiceHandler
}

func (h *posterServiceHandler) SavePosterPromotion(ctx context.Context, in *PosterPromotion, out *PosterResponse) error {
	return h.PosterServiceHandler.SavePosterPromotion(ctx, in, out)
}

func (h *posterServiceHandler) SavePosterGoods(ctx context.Context, in *PosterGoods, out *PosterResponse) error {
	return h.PosterServiceHandler.SavePosterGoods(ctx, in, out)
}

func (h *posterServiceHandler) GetPosterPromotion(ctx context.Context, in *PosterRequest, out *PosterResponse) error {
	return h.PosterServiceHandler.GetPosterPromotion(ctx, in, out)
}

func (h *posterServiceHandler) GetPosterGoods(ctx context.Context, in *PosterRequest, out *PosterResponse) error {
	return h.PosterServiceHandler.GetPosterGoods(ctx, in, out)
}