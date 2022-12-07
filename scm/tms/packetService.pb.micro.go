// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: packetService.proto

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

// Api Endpoints for PacketService service

func NewPacketServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PacketService service

type PacketService interface {
	Create(ctx context.Context, in *Packet, opts ...client.CallOption) (*PacketResponse, error)
	Update(ctx context.Context, in *Packet, opts ...client.CallOption) (*PacketResponse, error)
	Get(ctx context.Context, in *PacketRequest, opts ...client.CallOption) (*PacketResponse, error)
	List(ctx context.Context, in *PacketRequest, opts ...client.CallOption) (*PacketResponse, error)
}

type packetService struct {
	c    client.Client
	name string
}

func NewPacketService(name string, c client.Client) PacketService {
	return &packetService{
		c:    c,
		name: name,
	}
}

func (c *packetService) Create(ctx context.Context, in *Packet, opts ...client.CallOption) (*PacketResponse, error) {
	req := c.c.NewRequest(c.name, "PacketService.Create", in)
	out := new(PacketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packetService) Update(ctx context.Context, in *Packet, opts ...client.CallOption) (*PacketResponse, error) {
	req := c.c.NewRequest(c.name, "PacketService.Update", in)
	out := new(PacketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packetService) Get(ctx context.Context, in *PacketRequest, opts ...client.CallOption) (*PacketResponse, error) {
	req := c.c.NewRequest(c.name, "PacketService.Get", in)
	out := new(PacketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packetService) List(ctx context.Context, in *PacketRequest, opts ...client.CallOption) (*PacketResponse, error) {
	req := c.c.NewRequest(c.name, "PacketService.List", in)
	out := new(PacketResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PacketService service

type PacketServiceHandler interface {
	Create(context.Context, *Packet, *PacketResponse) error
	Update(context.Context, *Packet, *PacketResponse) error
	Get(context.Context, *PacketRequest, *PacketResponse) error
	List(context.Context, *PacketRequest, *PacketResponse) error
}

func RegisterPacketServiceHandler(s server.Server, hdlr PacketServiceHandler, opts ...server.HandlerOption) error {
	type packetService interface {
		Create(ctx context.Context, in *Packet, out *PacketResponse) error
		Update(ctx context.Context, in *Packet, out *PacketResponse) error
		Get(ctx context.Context, in *PacketRequest, out *PacketResponse) error
		List(ctx context.Context, in *PacketRequest, out *PacketResponse) error
	}
	type PacketService struct {
		packetService
	}
	h := &packetServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PacketService{h}, opts...))
}

type packetServiceHandler struct {
	PacketServiceHandler
}

func (h *packetServiceHandler) Create(ctx context.Context, in *Packet, out *PacketResponse) error {
	return h.PacketServiceHandler.Create(ctx, in, out)
}

func (h *packetServiceHandler) Update(ctx context.Context, in *Packet, out *PacketResponse) error {
	return h.PacketServiceHandler.Update(ctx, in, out)
}

func (h *packetServiceHandler) Get(ctx context.Context, in *PacketRequest, out *PacketResponse) error {
	return h.PacketServiceHandler.Get(ctx, in, out)
}

func (h *packetServiceHandler) List(ctx context.Context, in *PacketRequest, out *PacketResponse) error {
	return h.PacketServiceHandler.List(ctx, in, out)
}
