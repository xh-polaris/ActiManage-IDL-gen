// Code generated by Kitex v0.12.3. DO NOT EDIT.

package systemservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	system "github.com/xh-polaris/ActiManage-IDL-gen/kitex_gen/system"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MerchantLogin(ctx context.Context, Req *system.MerchantLoginReq, callOptions ...callopt.Option) (r *system.MerchantLoginResp, err error)
	MerchantSetPassword(ctx context.Context, Req *system.MerchantSetPasswordReq, callOptions ...callopt.Option) (r *system.Response, err error)
	UpdateMerchantSetting(ctx context.Context, Req *system.UpdateSettingReq, callOptions ...callopt.Option) (r *system.Response, err error)
	GetMerchantSetting(ctx context.Context, Req *system.GetMerchantSettingReq, callOptions ...callopt.Option) (r *system.GetMerchantSettingResp, err error)
	UpdateMerchantInfo(ctx context.Context, Req *system.UpdateMerchantInfoReq, callOptions ...callopt.Option) (r *system.Response, err error)
	GetMerchantInfoByUri(ctx context.Context, Req *system.GetMerchantInfoByUriReq, callOptions ...callopt.Option) (r *system.GetMerchantInfoByUriResp, err error)
	GetAd(ctx context.Context, Req *system.GetAdReq, callOptions ...callopt.Option) (r *system.GetAdResp, err error)
	SetAd(ctx context.Context, Req *system.SetAdReq, callOptions ...callopt.Option) (r *system.Response, err error)
	CreateActivity(ctx context.Context, Req *system.CreateActivityReq, callOptions ...callopt.Option) (r *system.Response, err error)
	TopActivity(ctx context.Context, Req *system.TopActivityReq, callOptions ...callopt.Option) (r *system.Response, err error)
	DeleteActivity(ctx context.Context, Req *system.DeleteActivityReq, callOptions ...callopt.Option) (r *system.Response, err error)
	UpdateActivity(ctx context.Context, Req *system.UpdateActivityReq, callOptions ...callopt.Option) (r *system.Response, err error)
	ListActivities(ctx context.Context, Req *system.ListActivitiesReq, callOptions ...callopt.Option) (r *system.ListActivitiesResp, err error)
	GetActivity(ctx context.Context, Req *system.GetActivityReq, callOptions ...callopt.Option) (r *system.GetActivityResp, err error)
	AdminLogin(ctx context.Context, Req *system.AdminLoginReq, callOptions ...callopt.Option) (r *system.AdminLoginResp, err error)
	AdminSetPassword(ctx context.Context, Req *system.AdminSetPasswordReq, callOptions ...callopt.Option) (r *system.Response, err error)
	CreateMerchant(ctx context.Context, Req *system.CreateMerchantReq, callOptions ...callopt.Option) (r *system.Response, err error)
	UpdateMerchant(ctx context.Context, Req *system.UpdateMerchantReq, callOptions ...callopt.Option) (r *system.Response, err error)
	DeleteMerchant(ctx context.Context, Req *system.DeleteMerchantReq, callOptions ...callopt.Option) (r *system.Response, err error)
	GetMerchantInfo(ctx context.Context, Req *system.GetMerchantInfoReq, callOptions ...callopt.Option) (r *system.GetMerchantInfoResp, err error)
	ListMerchants(ctx context.Context, Req *system.ListMerchantsReq, callOptions ...callopt.Option) (r *system.ListMerchantsResp, err error)
	GetMerchantMoreInfo(ctx context.Context, Req *system.GetMerchantMoreInfoReq, callOptions ...callopt.Option) (r *system.GetMerchantMoreInfoResp, err error)
	UpdateMerchantMoreInfo(ctx context.Context, Req *system.UpdateMerchantMoreInfoReq, callOptions ...callopt.Option) (r *system.Response, err error)
	GetMerchantTotalData(ctx context.Context, Req *system.GetMerchantTotalDataReq, callOptions ...callopt.Option) (r *system.GetMerchantTotalDataResp, err error)
	ListMerchantByMerchantId(ctx context.Context, Req *system.ListMerchantsByMerchantIdReq, callOptions ...callopt.Option) (r *system.ListMerchantsByMerchantIdResp, err error)
	ListMerchantByActivityNumber(ctx context.Context, Req *system.ListMerchantsByActivityNumberReq, callOptions ...callopt.Option) (r *system.ListMerchantsByActivityNumberResp, err error)
	ListActivityByActivityId(ctx context.Context, Req *system.ListActivitiesByActivityIdReq, callOptions ...callopt.Option) (r *system.ListActivitiesByActivityIdResp, err error)
	MerchantGetActivityNumber(ctx context.Context, Req *system.MerchantGetActivityNumberReq, callOptions ...callopt.Option) (r *system.MerchantGetActivityNumberResp, err error)
	StsSendVerifyCode(ctx context.Context, Req *system.StsSendVerifyCodeReq, callOptions ...callopt.Option) (r *system.Response, err error)
	StsCheckVerifyCode(ctx context.Context, Req *system.StsCheckVerifyCodeReq, callOptions ...callopt.Option) (r *system.Response, err error)
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
	return &kSystemServiceClient{
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

type kSystemServiceClient struct {
	*kClient
}

func (p *kSystemServiceClient) MerchantLogin(ctx context.Context, Req *system.MerchantLoginReq, callOptions ...callopt.Option) (r *system.MerchantLoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantLogin(ctx, Req)
}

func (p *kSystemServiceClient) MerchantSetPassword(ctx context.Context, Req *system.MerchantSetPasswordReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantSetPassword(ctx, Req)
}

func (p *kSystemServiceClient) UpdateMerchantSetting(ctx context.Context, Req *system.UpdateSettingReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateMerchantSetting(ctx, Req)
}

func (p *kSystemServiceClient) GetMerchantSetting(ctx context.Context, Req *system.GetMerchantSettingReq, callOptions ...callopt.Option) (r *system.GetMerchantSettingResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMerchantSetting(ctx, Req)
}

func (p *kSystemServiceClient) UpdateMerchantInfo(ctx context.Context, Req *system.UpdateMerchantInfoReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateMerchantInfo(ctx, Req)
}

func (p *kSystemServiceClient) GetMerchantInfoByUri(ctx context.Context, Req *system.GetMerchantInfoByUriReq, callOptions ...callopt.Option) (r *system.GetMerchantInfoByUriResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMerchantInfoByUri(ctx, Req)
}

func (p *kSystemServiceClient) GetAd(ctx context.Context, Req *system.GetAdReq, callOptions ...callopt.Option) (r *system.GetAdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAd(ctx, Req)
}

func (p *kSystemServiceClient) SetAd(ctx context.Context, Req *system.SetAdReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetAd(ctx, Req)
}

func (p *kSystemServiceClient) CreateActivity(ctx context.Context, Req *system.CreateActivityReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateActivity(ctx, Req)
}

func (p *kSystemServiceClient) TopActivity(ctx context.Context, Req *system.TopActivityReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TopActivity(ctx, Req)
}

func (p *kSystemServiceClient) DeleteActivity(ctx context.Context, Req *system.DeleteActivityReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteActivity(ctx, Req)
}

func (p *kSystemServiceClient) UpdateActivity(ctx context.Context, Req *system.UpdateActivityReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateActivity(ctx, Req)
}

func (p *kSystemServiceClient) ListActivities(ctx context.Context, Req *system.ListActivitiesReq, callOptions ...callopt.Option) (r *system.ListActivitiesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListActivities(ctx, Req)
}

func (p *kSystemServiceClient) GetActivity(ctx context.Context, Req *system.GetActivityReq, callOptions ...callopt.Option) (r *system.GetActivityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetActivity(ctx, Req)
}

func (p *kSystemServiceClient) AdminLogin(ctx context.Context, Req *system.AdminLoginReq, callOptions ...callopt.Option) (r *system.AdminLoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AdminLogin(ctx, Req)
}

func (p *kSystemServiceClient) AdminSetPassword(ctx context.Context, Req *system.AdminSetPasswordReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AdminSetPassword(ctx, Req)
}

func (p *kSystemServiceClient) CreateMerchant(ctx context.Context, Req *system.CreateMerchantReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateMerchant(ctx, Req)
}

func (p *kSystemServiceClient) UpdateMerchant(ctx context.Context, Req *system.UpdateMerchantReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateMerchant(ctx, Req)
}

func (p *kSystemServiceClient) DeleteMerchant(ctx context.Context, Req *system.DeleteMerchantReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteMerchant(ctx, Req)
}

func (p *kSystemServiceClient) GetMerchantInfo(ctx context.Context, Req *system.GetMerchantInfoReq, callOptions ...callopt.Option) (r *system.GetMerchantInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMerchantInfo(ctx, Req)
}

func (p *kSystemServiceClient) ListMerchants(ctx context.Context, Req *system.ListMerchantsReq, callOptions ...callopt.Option) (r *system.ListMerchantsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListMerchants(ctx, Req)
}

func (p *kSystemServiceClient) GetMerchantMoreInfo(ctx context.Context, Req *system.GetMerchantMoreInfoReq, callOptions ...callopt.Option) (r *system.GetMerchantMoreInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMerchantMoreInfo(ctx, Req)
}

func (p *kSystemServiceClient) UpdateMerchantMoreInfo(ctx context.Context, Req *system.UpdateMerchantMoreInfoReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateMerchantMoreInfo(ctx, Req)
}

func (p *kSystemServiceClient) GetMerchantTotalData(ctx context.Context, Req *system.GetMerchantTotalDataReq, callOptions ...callopt.Option) (r *system.GetMerchantTotalDataResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMerchantTotalData(ctx, Req)
}

func (p *kSystemServiceClient) ListMerchantByMerchantId(ctx context.Context, Req *system.ListMerchantsByMerchantIdReq, callOptions ...callopt.Option) (r *system.ListMerchantsByMerchantIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListMerchantByMerchantId(ctx, Req)
}

func (p *kSystemServiceClient) ListMerchantByActivityNumber(ctx context.Context, Req *system.ListMerchantsByActivityNumberReq, callOptions ...callopt.Option) (r *system.ListMerchantsByActivityNumberResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListMerchantByActivityNumber(ctx, Req)
}

func (p *kSystemServiceClient) ListActivityByActivityId(ctx context.Context, Req *system.ListActivitiesByActivityIdReq, callOptions ...callopt.Option) (r *system.ListActivitiesByActivityIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListActivityByActivityId(ctx, Req)
}

func (p *kSystemServiceClient) MerchantGetActivityNumber(ctx context.Context, Req *system.MerchantGetActivityNumberReq, callOptions ...callopt.Option) (r *system.MerchantGetActivityNumberResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MerchantGetActivityNumber(ctx, Req)
}

func (p *kSystemServiceClient) StsSendVerifyCode(ctx context.Context, Req *system.StsSendVerifyCodeReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StsSendVerifyCode(ctx, Req)
}

func (p *kSystemServiceClient) StsCheckVerifyCode(ctx context.Context, Req *system.StsCheckVerifyCodeReq, callOptions ...callopt.Option) (r *system.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StsCheckVerifyCode(ctx, Req)
}
