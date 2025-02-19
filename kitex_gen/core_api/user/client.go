// Code generated by Kitex v0.12.2. DO NOT EDIT.

package user

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/ActiManage-IDL-gen/kitex_gen/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Login(ctx context.Context, Req *core_api.LoginReq, callOptions ...callopt.Option) (r *core_api.LoginResp, err error)
	SignUp(ctx context.Context, Req *core_api.SignUpReq, callOptions ...callopt.Option) (r *core_api.SignUpResp, err error)
	GetSetting(ctx context.Context, Req *core_api.GetSettingReq, callOptions ...callopt.Option) (r *core_api.GetSettingResp, err error)
	ListActivities(ctx context.Context, Req *core_api.ListActivitiesReq, callOptions ...callopt.Option) (r *core_api.ListActivitiesResp, err error)
	GetActivity(ctx context.Context, Req *core_api.GetActivityReq, callOptions ...callopt.Option) (r *core_api.GetActivityResp, err error)
	DoFavorite(ctx context.Context, Req *core_api.DoFavoriteReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	CancelFavorite(ctx context.Context, Req *core_api.CancelFavoriteReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	CreateBooking(ctx context.Context, Req *core_api.CreateBookingReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	CancelBookRecord(ctx context.Context, Req *core_api.CancelBookRecordReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	ListActivitiesByBookRecords(ctx context.Context, Req *core_api.ListActivitiesByBookRecordsReq, callOptions ...callopt.Option) (r *core_api.ListActivitiesByBookRecordsResp, err error)
	ListReservers(ctx context.Context, Req *core_api.ListReserversReq, callOptions ...callopt.Option) (r *core_api.ListReserversResp, err error)
	CreateReserver(ctx context.Context, Req *core_api.CreateReserverReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	DeleteReserver(ctx context.Context, Req *core_api.DeleteReserverReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	GetUserInfo(ctx context.Context, Req *core_api.GetUserInfoReq, callOptions ...callopt.Option) (r *core_api.GetUserInfoResp, err error)
	UpdateUserInfo(ctx context.Context, Req *core_api.UpdateUserInfoReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	UpdateNotice(ctx context.Context, Req *core_api.UpdateNoticeReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	GetMerchantInfo(ctx context.Context, Req *core_api.GetMerchantInfoReq, callOptions ...callopt.Option) (r *core_api.GetMerchantInfoResp, err error)
	ListActivitiesByView(ctx context.Context, Req *core_api.ListActivitiesByViewReq, callOptions ...callopt.Option) (r *core_api.ListActivitiesByViewResp, err error)
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
	return &kUserClient{
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

type kUserClient struct {
	*kClient
}

func (p *kUserClient) Login(ctx context.Context, Req *core_api.LoginReq, callOptions ...callopt.Option) (r *core_api.LoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, Req)
}

func (p *kUserClient) SignUp(ctx context.Context, Req *core_api.SignUpReq, callOptions ...callopt.Option) (r *core_api.SignUpResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SignUp(ctx, Req)
}

func (p *kUserClient) GetSetting(ctx context.Context, Req *core_api.GetSettingReq, callOptions ...callopt.Option) (r *core_api.GetSettingResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetSetting(ctx, Req)
}

func (p *kUserClient) ListActivities(ctx context.Context, Req *core_api.ListActivitiesReq, callOptions ...callopt.Option) (r *core_api.ListActivitiesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListActivities(ctx, Req)
}

func (p *kUserClient) GetActivity(ctx context.Context, Req *core_api.GetActivityReq, callOptions ...callopt.Option) (r *core_api.GetActivityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetActivity(ctx, Req)
}

func (p *kUserClient) DoFavorite(ctx context.Context, Req *core_api.DoFavoriteReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DoFavorite(ctx, Req)
}

func (p *kUserClient) CancelFavorite(ctx context.Context, Req *core_api.CancelFavoriteReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelFavorite(ctx, Req)
}

func (p *kUserClient) CreateBooking(ctx context.Context, Req *core_api.CreateBookingReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateBooking(ctx, Req)
}

func (p *kUserClient) CancelBookRecord(ctx context.Context, Req *core_api.CancelBookRecordReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelBookRecord(ctx, Req)
}

func (p *kUserClient) ListActivitiesByBookRecords(ctx context.Context, Req *core_api.ListActivitiesByBookRecordsReq, callOptions ...callopt.Option) (r *core_api.ListActivitiesByBookRecordsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListActivitiesByBookRecords(ctx, Req)
}

func (p *kUserClient) ListReservers(ctx context.Context, Req *core_api.ListReserversReq, callOptions ...callopt.Option) (r *core_api.ListReserversResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListReservers(ctx, Req)
}

func (p *kUserClient) CreateReserver(ctx context.Context, Req *core_api.CreateReserverReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateReserver(ctx, Req)
}

func (p *kUserClient) DeleteReserver(ctx context.Context, Req *core_api.DeleteReserverReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteReserver(ctx, Req)
}

func (p *kUserClient) GetUserInfo(ctx context.Context, Req *core_api.GetUserInfoReq, callOptions ...callopt.Option) (r *core_api.GetUserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}

func (p *kUserClient) UpdateUserInfo(ctx context.Context, Req *core_api.UpdateUserInfoReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserInfo(ctx, Req)
}

func (p *kUserClient) UpdateNotice(ctx context.Context, Req *core_api.UpdateNoticeReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateNotice(ctx, Req)
}

func (p *kUserClient) GetMerchantInfo(ctx context.Context, Req *core_api.GetMerchantInfoReq, callOptions ...callopt.Option) (r *core_api.GetMerchantInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMerchantInfo(ctx, Req)
}

func (p *kUserClient) ListActivitiesByView(ctx context.Context, Req *core_api.ListActivitiesByViewReq, callOptions ...callopt.Option) (r *core_api.ListActivitiesByViewResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListActivitiesByView(ctx, Req)
}
