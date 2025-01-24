// Code generated by Kitex v0.12.1. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "github.com/xh-polaris/ActiManage-IDL-gen/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserSignUp(ctx context.Context, Req *user.UserSignUpReq, callOptions ...callopt.Option) (r *user.UserSignUpResp, err error)
	UserLogin(ctx context.Context, Req *user.UserLoginReq, callOptions ...callopt.Option) (r *user.UserLoginResp, err error)
	GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq, callOptions ...callopt.Option) (r *user.GetUserInfoResp, err error)
	UpdateUserInfo(ctx context.Context, Req *user.UpdateUserInfoReq, callOptions ...callopt.Option) (r *user.Response, err error)
	CreateReserver(ctx context.Context, Req *user.CreateReserverReq, callOptions ...callopt.Option) (r *user.Response, err error)
	DeleteReserver(ctx context.Context, Req *user.DeleteReserverReq, callOptions ...callopt.Option) (r *user.Response, err error)
	UpdateReserver(ctx context.Context, Req *user.UpdateReserverReq, callOptions ...callopt.Option) (r *user.Response, err error)
	ListReservers(ctx context.Context, Req *user.ListReserversReq, callOptions ...callopt.Option) (r *user.ListReserversResp, err error)
	CreateBooking(ctx context.Context, Req *user.CreateBookRecordReq, callOptions ...callopt.Option) (r *user.Response, err error)
	DeleteBooking(ctx context.Context, Req *user.DeleteBookRecordReq, callOptions ...callopt.Option) (r *user.Response, err error)
	GetBookingDetail(ctx context.Context, Req *user.GetBookRecordReq, callOptions ...callopt.Option) (r *user.GetBookRecordResp, err error)
	ListBookRecordsByUser(ctx context.Context, Req *user.ListBookRecordsByUserReq, callOptions ...callopt.Option) (r *user.ListBookRecordsByUserResp, err error)
	CreateReceipt(ctx context.Context, Req *user.CreateReceiptReq, callOptions ...callopt.Option) (r *user.Response, err error)
	DeleteReceipt(ctx context.Context, Req *user.DeleteReceiptReq, callOptions ...callopt.Option) (r *user.Response, err error)
	ListReceipts(ctx context.Context, Req *user.ListReceiptsReq, callOptions ...callopt.Option) (r *user.ListReceiptsResp, err error)
	MarkReceiptRead(ctx context.Context, Req *user.MarkReceiptReadReq, callOptions ...callopt.Option) (r *user.Response, err error)
	CreateFavorite(ctx context.Context, Req *user.CreateFavoriteReq, callOptions ...callopt.Option) (r *user.Response, err error)
	CancelFavorite(ctx context.Context, Req *user.CancelFavoriteReq, callOptions ...callopt.Option) (r *user.Response, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) UserSignUp(ctx context.Context, Req *user.UserSignUpReq, callOptions ...callopt.Option) (r *user.UserSignUpResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserSignUp(ctx, Req)
}

func (p *kUserServiceClient) UserLogin(ctx context.Context, Req *user.UserLoginReq, callOptions ...callopt.Option) (r *user.UserLoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserLogin(ctx, Req)
}

func (p *kUserServiceClient) GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq, callOptions ...callopt.Option) (r *user.GetUserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}

func (p *kUserServiceClient) UpdateUserInfo(ctx context.Context, Req *user.UpdateUserInfoReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserInfo(ctx, Req)
}

func (p *kUserServiceClient) CreateReserver(ctx context.Context, Req *user.CreateReserverReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateReserver(ctx, Req)
}

func (p *kUserServiceClient) DeleteReserver(ctx context.Context, Req *user.DeleteReserverReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteReserver(ctx, Req)
}

func (p *kUserServiceClient) UpdateReserver(ctx context.Context, Req *user.UpdateReserverReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateReserver(ctx, Req)
}

func (p *kUserServiceClient) ListReservers(ctx context.Context, Req *user.ListReserversReq, callOptions ...callopt.Option) (r *user.ListReserversResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListReservers(ctx, Req)
}

func (p *kUserServiceClient) CreateBooking(ctx context.Context, Req *user.CreateBookRecordReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateBooking(ctx, Req)
}

func (p *kUserServiceClient) DeleteBooking(ctx context.Context, Req *user.DeleteBookRecordReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteBooking(ctx, Req)
}

func (p *kUserServiceClient) GetBookingDetail(ctx context.Context, Req *user.GetBookRecordReq, callOptions ...callopt.Option) (r *user.GetBookRecordResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetBookingDetail(ctx, Req)
}

func (p *kUserServiceClient) ListBookRecordsByUser(ctx context.Context, Req *user.ListBookRecordsByUserReq, callOptions ...callopt.Option) (r *user.ListBookRecordsByUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListBookRecordsByUser(ctx, Req)
}

func (p *kUserServiceClient) CreateReceipt(ctx context.Context, Req *user.CreateReceiptReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateReceipt(ctx, Req)
}

func (p *kUserServiceClient) DeleteReceipt(ctx context.Context, Req *user.DeleteReceiptReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteReceipt(ctx, Req)
}

func (p *kUserServiceClient) ListReceipts(ctx context.Context, Req *user.ListReceiptsReq, callOptions ...callopt.Option) (r *user.ListReceiptsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListReceipts(ctx, Req)
}

func (p *kUserServiceClient) MarkReceiptRead(ctx context.Context, Req *user.MarkReceiptReadReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MarkReceiptRead(ctx, Req)
}

func (p *kUserServiceClient) CreateFavorite(ctx context.Context, Req *user.CreateFavoriteReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateFavorite(ctx, Req)
}

func (p *kUserServiceClient) CancelFavorite(ctx context.Context, Req *user.CancelFavoriteReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelFavorite(ctx, Req)
}
