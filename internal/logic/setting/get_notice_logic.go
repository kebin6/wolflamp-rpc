package setting

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/enum"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeLogic {
	return &GetNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNoticeLogic) GetNotice(in *wolflamp.Empty) (*wolflamp.GetNoticeResp, error) {
	setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformNotice.Val()})
	if err != nil {
		return nil, err
	}
	return &wolflamp.GetNoticeResp{
		Notice: setting.JsonString,
	}, nil
}
