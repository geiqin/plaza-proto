// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: complaintService.proto

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

// Api Endpoints for ComplaintService service

func NewComplaintServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ComplaintService service

type ComplaintService interface {
	//售后申请页【user】
	ApplyIndex(ctx context.Context, in *ComplaintRequest, opts ...client.CallOption) (*ComplaintResponse, error)
	// 售后申请创建【user】
	ApplyCreate(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error)
	//售后取消【admin/user】
	Cancel(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error)
	//售后单删除【admin/user】
	Delete(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error)
	//用户退货【user：换货/退货】
	ReturnGoods(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error)
	//*管理员审核【admin】
	//拒绝申请，拒绝原因
	//仅退款：确认退款
	//退货退款：同意并发送退货地址
	//换货：同意并发送退货地址
	ConfirmReview(ctx context.Context, in *ComplaintReviewRequest, opts ...client.CallOption) (*ComplaintResponse, error)
	//确认收货(换货：[确认并重新发货]确认收货并重新发货)【admin：换货/退货】
	ConfirmCollect(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error)
	//确认退款【admin】
	ConfirmRefund(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error)
	//重新退款【admin】：对退款失败重新执行退款命令
	RefundAgain(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error)
	//售后单详情【admin/user】
	Detail(ctx context.Context, in *ComplaintRequest, opts ...client.CallOption) (*ComplaintResponse, error)
	// 售后单获取
	Get(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error)
	// 售后单查询
	Search(ctx context.Context, in *ComplaintRequest, opts ...client.CallOption) (*ComplaintResponse, error)
	// 售后单列表
	List(ctx context.Context, in *ComplaintRequest, opts ...client.CallOption) (*ComplaintResponse, error)
}

type complaintService struct {
	c    client.Client
	name string
}

func NewComplaintService(name string, c client.Client) ComplaintService {
	return &complaintService{
		c:    c,
		name: name,
	}
}

func (c *complaintService) ApplyIndex(ctx context.Context, in *ComplaintRequest, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.ApplyIndex", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) ApplyCreate(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.ApplyCreate", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) Cancel(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.Cancel", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) Delete(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.Delete", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) ReturnGoods(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.ReturnGoods", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) ConfirmReview(ctx context.Context, in *ComplaintReviewRequest, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.ConfirmReview", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) ConfirmCollect(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.ConfirmCollect", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) ConfirmRefund(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.ConfirmRefund", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) RefundAgain(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.RefundAgain", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) Detail(ctx context.Context, in *ComplaintRequest, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.Detail", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) Get(ctx context.Context, in *Complaint, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.Get", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) Search(ctx context.Context, in *ComplaintRequest, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.Search", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complaintService) List(ctx context.Context, in *ComplaintRequest, opts ...client.CallOption) (*ComplaintResponse, error) {
	req := c.c.NewRequest(c.name, "ComplaintService.List", in)
	out := new(ComplaintResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ComplaintService service

type ComplaintServiceHandler interface {
	//售后申请页【user】
	ApplyIndex(context.Context, *ComplaintRequest, *ComplaintResponse) error
	// 售后申请创建【user】
	ApplyCreate(context.Context, *Complaint, *ComplaintResponse) error
	//售后取消【admin/user】
	Cancel(context.Context, *Complaint, *ComplaintResponse) error
	//售后单删除【admin/user】
	Delete(context.Context, *Complaint, *ComplaintResponse) error
	//用户退货【user：换货/退货】
	ReturnGoods(context.Context, *Complaint, *ComplaintResponse) error
	//*管理员审核【admin】
	//拒绝申请，拒绝原因
	//仅退款：确认退款
	//退货退款：同意并发送退货地址
	//换货：同意并发送退货地址
	ConfirmReview(context.Context, *ComplaintReviewRequest, *ComplaintResponse) error
	//确认收货(换货：[确认并重新发货]确认收货并重新发货)【admin：换货/退货】
	ConfirmCollect(context.Context, *Complaint, *ComplaintResponse) error
	//确认退款【admin】
	ConfirmRefund(context.Context, *Complaint, *ComplaintResponse) error
	//重新退款【admin】：对退款失败重新执行退款命令
	RefundAgain(context.Context, *Complaint, *ComplaintResponse) error
	//售后单详情【admin/user】
	Detail(context.Context, *ComplaintRequest, *ComplaintResponse) error
	// 售后单获取
	Get(context.Context, *Complaint, *ComplaintResponse) error
	// 售后单查询
	Search(context.Context, *ComplaintRequest, *ComplaintResponse) error
	// 售后单列表
	List(context.Context, *ComplaintRequest, *ComplaintResponse) error
}

func RegisterComplaintServiceHandler(s server.Server, hdlr ComplaintServiceHandler, opts ...server.HandlerOption) error {
	type complaintService interface {
		ApplyIndex(ctx context.Context, in *ComplaintRequest, out *ComplaintResponse) error
		ApplyCreate(ctx context.Context, in *Complaint, out *ComplaintResponse) error
		Cancel(ctx context.Context, in *Complaint, out *ComplaintResponse) error
		Delete(ctx context.Context, in *Complaint, out *ComplaintResponse) error
		ReturnGoods(ctx context.Context, in *Complaint, out *ComplaintResponse) error
		ConfirmReview(ctx context.Context, in *ComplaintReviewRequest, out *ComplaintResponse) error
		ConfirmCollect(ctx context.Context, in *Complaint, out *ComplaintResponse) error
		ConfirmRefund(ctx context.Context, in *Complaint, out *ComplaintResponse) error
		RefundAgain(ctx context.Context, in *Complaint, out *ComplaintResponse) error
		Detail(ctx context.Context, in *ComplaintRequest, out *ComplaintResponse) error
		Get(ctx context.Context, in *Complaint, out *ComplaintResponse) error
		Search(ctx context.Context, in *ComplaintRequest, out *ComplaintResponse) error
		List(ctx context.Context, in *ComplaintRequest, out *ComplaintResponse) error
	}
	type ComplaintService struct {
		complaintService
	}
	h := &complaintServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ComplaintService{h}, opts...))
}

type complaintServiceHandler struct {
	ComplaintServiceHandler
}

func (h *complaintServiceHandler) ApplyIndex(ctx context.Context, in *ComplaintRequest, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.ApplyIndex(ctx, in, out)
}

func (h *complaintServiceHandler) ApplyCreate(ctx context.Context, in *Complaint, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.ApplyCreate(ctx, in, out)
}

func (h *complaintServiceHandler) Cancel(ctx context.Context, in *Complaint, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.Cancel(ctx, in, out)
}

func (h *complaintServiceHandler) Delete(ctx context.Context, in *Complaint, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.Delete(ctx, in, out)
}

func (h *complaintServiceHandler) ReturnGoods(ctx context.Context, in *Complaint, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.ReturnGoods(ctx, in, out)
}

func (h *complaintServiceHandler) ConfirmReview(ctx context.Context, in *ComplaintReviewRequest, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.ConfirmReview(ctx, in, out)
}

func (h *complaintServiceHandler) ConfirmCollect(ctx context.Context, in *Complaint, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.ConfirmCollect(ctx, in, out)
}

func (h *complaintServiceHandler) ConfirmRefund(ctx context.Context, in *Complaint, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.ConfirmRefund(ctx, in, out)
}

func (h *complaintServiceHandler) RefundAgain(ctx context.Context, in *Complaint, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.RefundAgain(ctx, in, out)
}

func (h *complaintServiceHandler) Detail(ctx context.Context, in *ComplaintRequest, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.Detail(ctx, in, out)
}

func (h *complaintServiceHandler) Get(ctx context.Context, in *Complaint, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.Get(ctx, in, out)
}

func (h *complaintServiceHandler) Search(ctx context.Context, in *ComplaintRequest, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.Search(ctx, in, out)
}

func (h *complaintServiceHandler) List(ctx context.Context, in *ComplaintRequest, out *ComplaintResponse) error {
	return h.ComplaintServiceHandler.List(ctx, in, out)
}
