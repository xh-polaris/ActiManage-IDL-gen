// Code generated by Kitex v0.12.3. DO NOT EDIT.

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
	SetPassword(ctx context.Context, Req *user.SetPasswordReq, callOptions ...callopt.Option) (r *user.Response, err error)
	SetNotice(ctx context.Context, Req *user.SetNoticeReq, callOptions ...callopt.Option) (r *user.Response, err error)
	CreateReserver(ctx context.Context, Req *user.CreateReserverReq, callOptions ...callopt.Option) (r *user.Response, err error)
	DeleteReserver(ctx context.Context, Req *user.DeleteReserverReq, callOptions ...callopt.Option) (r *user.Response, err error)
	UpdateReserver(ctx context.Context, Req *user.UpdateReserverReq, callOptions ...callopt.Option) (r *user.Response, err error)
	ListReservers(ctx context.Context, Req *user.ListReserversReq, callOptions ...callopt.Option) (r *user.ListReserversResp, err error)
	CreateBookRecord(ctx context.Context, Req *user.CreateBookRecordReq, callOptions ...callopt.Option) (r *user.Response, err error)
	CancelBookRecord(ctx context.Context, Req *user.CancelBookRecordReq, callOptions ...callopt.Option) (r *user.Response, err error)
	GetBookRecordDetail(ctx context.Context, Req *user.GetBookRecordReq, callOptions ...callopt.Option) (r *user.GetBookRecordResp, err error)
	ListBookRecordsByUser(ctx context.Context, Req *user.ListBookRecordsByUserReq, callOptions ...callopt.Option) (r *user.ListBookRecordsByUserResp, err error)
	ListBookRecordsByActivity(ctx context.Context, Req *user.ListBookRecordsByActivityReq, callOptions ...callopt.Option) (r *user.ListBookRecordsByActivityResp, err error)
	CheckBookRecordByUserIdAndActivityId(ctx context.Context, Req *user.CheckBookRecordByUserIdAndActivityIdReq, callOptions ...callopt.Option) (r *user.CheckBookRecordByUserIdAndActivityIdResp, err error)
	CreateReceipt(ctx context.Context, Req *user.CreateReceiptReq, callOptions ...callopt.Option) (r *user.Response, err error)
	DeleteReceipt(ctx context.Context, Req *user.DeleteReceiptReq, callOptions ...callopt.Option) (r *user.Response, err error)
	ListReceipts(ctx context.Context, Req *user.ListReceiptsReq, callOptions ...callopt.Option) (r *user.ListReceiptsResp, err error)
	MarkReceiptRead(ctx context.Context, Req *user.MarkReceiptReadReq, callOptions ...callopt.Option) (r *user.Response, err error)
	DoFavorite(ctx context.Context, Req *user.DoFavoriteReq, callOptions ...callopt.Option) (r *user.Response, err error)
	CancelFavorite(ctx context.Context, Req *user.CancelFavoriteReq, callOptions ...callopt.Option) (r *user.Response, err error)
	CheckFavorite(ctx context.Context, Req *user.CheckFavoriteReq, callOptions ...callopt.Option) (r *user.Response, err error)
	CreateView(ctx context.Context, Req *user.CreateViewReq, callOptions ...callopt.Option) (r *user.Response, err error)
	GetFavoriteAndViewOfActivity(ctx context.Context, Req *user.GetFavoriteAndViewOfActivityReq, callOptions ...callopt.Option) (r *user.GetFavoriteAndViewOfActivityResp, err error)
	GetViewOfMerchant(ctx context.Context, Req *user.GetViewOfMerchantReq, callOptions ...callopt.Option) (r *user.GetViewOfMerchantResp, err error)
	ListActivityIdsByView(ctx context.Context, Req *user.ListActivityIdsByViewReq, callOptions ...callopt.Option) (r *user.ListActivityIdsByViewResp, err error)
	ListMerchantIdByViewRank(ctx context.Context, Req *user.ListMerchantIdsByViewRankReq, callOptions ...callopt.Option) (r *user.ListMerchantIdsByViewRankResp, err error)
	ListMerchantIdByBookRecordRank(ctx context.Context, Req *user.ListMerchantIdsByBookRecordRankReq, callOptions ...callopt.Option) (r *user.ListMerchantIdsByBookRecordRankResp, err error)
	ListActivityIdByBookRecordRank(ctx context.Context, Req *user.ListActivityIdsByBookRecordRankReq, callOptions ...callopt.Option) (r *user.ListActivityIdsByBookRecordRankResp, err error)
	GetViewDataByMerchant(ctx context.Context, Req *user.GetViewDataByMerchantReq, callOptions ...callopt.Option) (r *user.GetViewDataByMerchantResp, err error)
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

func (p *kUserServiceClient) SetPassword(ctx context.Context, Req *user.SetPasswordReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetPassword(ctx, Req)
}

func (p *kUserServiceClient) SetNotice(ctx context.Context, Req *user.SetNoticeReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetNotice(ctx, Req)
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

func (p *kUserServiceClient) CreateBookRecord(ctx context.Context, Req *user.CreateBookRecordReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateBookRecord(ctx, Req)
}

func (p *kUserServiceClient) CancelBookRecord(ctx context.Context, Req *user.CancelBookRecordReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelBookRecord(ctx, Req)
}

func (p *kUserServiceClient) GetBookRecordDetail(ctx context.Context, Req *user.GetBookRecordReq, callOptions ...callopt.Option) (r *user.GetBookRecordResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetBookRecordDetail(ctx, Req)
}

func (p *kUserServiceClient) ListBookRecordsByUser(ctx context.Context, Req *user.ListBookRecordsByUserReq, callOptions ...callopt.Option) (r *user.ListBookRecordsByUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListBookRecordsByUser(ctx, Req)
}

func (p *kUserServiceClient) ListBookRecordsByActivity(ctx context.Context, Req *user.ListBookRecordsByActivityReq, callOptions ...callopt.Option) (r *user.ListBookRecordsByActivityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListBookRecordsByActivity(ctx, Req)
}

func (p *kUserServiceClient) CheckBookRecordByUserIdAndActivityId(ctx context.Context, Req *user.CheckBookRecordByUserIdAndActivityIdReq, callOptions ...callopt.Option) (r *user.CheckBookRecordByUserIdAndActivityIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckBookRecordByUserIdAndActivityId(ctx, Req)
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

func (p *kUserServiceClient) DoFavorite(ctx context.Context, Req *user.DoFavoriteReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DoFavorite(ctx, Req)
}

func (p *kUserServiceClient) CancelFavorite(ctx context.Context, Req *user.CancelFavoriteReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelFavorite(ctx, Req)
}

func (p *kUserServiceClient) CheckFavorite(ctx context.Context, Req *user.CheckFavoriteReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckFavorite(ctx, Req)
}

func (p *kUserServiceClient) CreateView(ctx context.Context, Req *user.CreateViewReq, callOptions ...callopt.Option) (r *user.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateView(ctx, Req)
}

func (p *kUserServiceClient) GetFavoriteAndViewOfActivity(ctx context.Context, Req *user.GetFavoriteAndViewOfActivityReq, callOptions ...callopt.Option) (r *user.GetFavoriteAndViewOfActivityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFavoriteAndViewOfActivity(ctx, Req)
}

func (p *kUserServiceClient) GetViewOfMerchant(ctx context.Context, Req *user.GetViewOfMerchantReq, callOptions ...callopt.Option) (r *user.GetViewOfMerchantResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetViewOfMerchant(ctx, Req)
}

func (p *kUserServiceClient) ListActivityIdsByView(ctx context.Context, Req *user.ListActivityIdsByViewReq, callOptions ...callopt.Option) (r *user.ListActivityIdsByViewResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListActivityIdsByView(ctx, Req)
}

func (p *kUserServiceClient) ListMerchantIdByViewRank(ctx context.Context, Req *user.ListMerchantIdsByViewRankReq, callOptions ...callopt.Option) (r *user.ListMerchantIdsByViewRankResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListMerchantIdByViewRank(ctx, Req)
}

func (p *kUserServiceClient) ListMerchantIdByBookRecordRank(ctx context.Context, Req *user.ListMerchantIdsByBookRecordRankReq, callOptions ...callopt.Option) (r *user.ListMerchantIdsByBookRecordRankResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListMerchantIdByBookRecordRank(ctx, Req)
}

func (p *kUserServiceClient) ListActivityIdByBookRecordRank(ctx context.Context, Req *user.ListActivityIdsByBookRecordRankReq, callOptions ...callopt.Option) (r *user.ListActivityIdsByBookRecordRankResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListActivityIdByBookRecordRank(ctx, Req)
}

func (p *kUserServiceClient) GetViewDataByMerchant(ctx context.Context, Req *user.GetViewDataByMerchantReq, callOptions ...callopt.Option) (r *user.GetViewDataByMerchantResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetViewDataByMerchant(ctx, Req)
}
