// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: transferService.proto

package services

import (
	fmt "fmt"
	_ "github.com/geiqin/microkit/protobuf/common"
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

// Api Endpoints for TransferService service

func NewTransferServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TransferService service

type TransferService interface {
	//创建转账
	Create(ctx context.Context, in *Transfer, opts ...client.CallOption) (*TransferResponse, error)
	//获得转账信息
	Get(ctx context.Context, in *Transfer, opts ...client.CallOption) (*TransferResponse, error)
	//查询转账
	Search(ctx context.Context, in *TransferWhere, opts ...client.CallOption) (*TransferResponse, error)
	// 转账到企业银行账户
	ToCorpBank(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error)
	//转账到微信个人钱包
	ToWxWallet(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error)
	//转账到支付宝账户
	ToAliAccount(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error)
}

type transferService struct {
	c    client.Client
	name string
}

func NewTransferService(name string, c client.Client) TransferService {
	return &transferService{
		c:    c,
		name: name,
	}
}

func (c *transferService) Create(ctx context.Context, in *Transfer, opts ...client.CallOption) (*TransferResponse, error) {
	req := c.c.NewRequest(c.name, "TransferService.Create", in)
	out := new(TransferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferService) Get(ctx context.Context, in *Transfer, opts ...client.CallOption) (*TransferResponse, error) {
	req := c.c.NewRequest(c.name, "TransferService.Get", in)
	out := new(TransferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferService) Search(ctx context.Context, in *TransferWhere, opts ...client.CallOption) (*TransferResponse, error) {
	req := c.c.NewRequest(c.name, "TransferService.Search", in)
	out := new(TransferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferService) ToCorpBank(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error) {
	req := c.c.NewRequest(c.name, "TransferService.ToCorpBank", in)
	out := new(TransferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferService) ToWxWallet(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error) {
	req := c.c.NewRequest(c.name, "TransferService.ToWxWallet", in)
	out := new(TransferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferService) ToAliAccount(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error) {
	req := c.c.NewRequest(c.name, "TransferService.ToAliAccount", in)
	out := new(TransferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TransferService service

type TransferServiceHandler interface {
	//创建转账
	Create(context.Context, *Transfer, *TransferResponse) error
	//获得转账信息
	Get(context.Context, *Transfer, *TransferResponse) error
	//查询转账
	Search(context.Context, *TransferWhere, *TransferResponse) error
	// 转账到企业银行账户
	ToCorpBank(context.Context, *TransferRequest, *TransferResponse) error
	//转账到微信个人钱包
	ToWxWallet(context.Context, *TransferRequest, *TransferResponse) error
	//转账到支付宝账户
	ToAliAccount(context.Context, *TransferRequest, *TransferResponse) error
}

func RegisterTransferServiceHandler(s server.Server, hdlr TransferServiceHandler, opts ...server.HandlerOption) error {
	type transferService interface {
		Create(ctx context.Context, in *Transfer, out *TransferResponse) error
		Get(ctx context.Context, in *Transfer, out *TransferResponse) error
		Search(ctx context.Context, in *TransferWhere, out *TransferResponse) error
		ToCorpBank(ctx context.Context, in *TransferRequest, out *TransferResponse) error
		ToWxWallet(ctx context.Context, in *TransferRequest, out *TransferResponse) error
		ToAliAccount(ctx context.Context, in *TransferRequest, out *TransferResponse) error
	}
	type TransferService struct {
		transferService
	}
	h := &transferServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TransferService{h}, opts...))
}

type transferServiceHandler struct {
	TransferServiceHandler
}

func (h *transferServiceHandler) Create(ctx context.Context, in *Transfer, out *TransferResponse) error {
	return h.TransferServiceHandler.Create(ctx, in, out)
}

func (h *transferServiceHandler) Get(ctx context.Context, in *Transfer, out *TransferResponse) error {
	return h.TransferServiceHandler.Get(ctx, in, out)
}

func (h *transferServiceHandler) Search(ctx context.Context, in *TransferWhere, out *TransferResponse) error {
	return h.TransferServiceHandler.Search(ctx, in, out)
}

func (h *transferServiceHandler) ToCorpBank(ctx context.Context, in *TransferRequest, out *TransferResponse) error {
	return h.TransferServiceHandler.ToCorpBank(ctx, in, out)
}

func (h *transferServiceHandler) ToWxWallet(ctx context.Context, in *TransferRequest, out *TransferResponse) error {
	return h.TransferServiceHandler.ToWxWallet(ctx, in, out)
}

func (h *transferServiceHandler) ToAliAccount(ctx context.Context, in *TransferRequest, out *TransferResponse) error {
	return h.TransferServiceHandler.ToAliAccount(ctx, in, out)
}