package setting

import (
	"context"
	"encoding/json"
	"github.com/kebin6/wolflamp-rpc/common/entity"
	"github.com/kebin6/wolflamp-rpc/common/enum"
	"strconv"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWithdrawThresholdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWithdrawThresholdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWithdrawThresholdLogic {
	return &GetWithdrawThresholdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWithdrawThresholdLogic) GetWithdrawThreshold(in *wolflamp.Empty) (*wolflamp.WithdrawThresholdResp, error) {

	if l.svcCtx.Redis.Exists(l.ctx, enum.WithdrawThreshold.CacheKey()).Val() == 0 {
		setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformSetting.Val()})
		if err != nil {
			return nil, err
		}
		var platformSetting entity.PlatformSetting
		if err := json.Unmarshal([]byte(setting.JsonString), &platformSetting); err != nil {
			return nil, err
		}
		_ = l.svcCtx.Redis.Set(l.ctx, enum.WithdrawThreshold.CacheKey(), platformSetting.WithdrawThreshold, 0)
	}

	cached, err := l.svcCtx.Redis.Get(l.ctx, enum.WithdrawThreshold.CacheKey()).Result()
	if err != nil {
		return nil, err
	}
	withdrawThreshold, err := strconv.ParseFloat(cached, 32)
	if err != nil {
		return nil, err
	}

	return &wolflamp.WithdrawThresholdResp{WithdrawThreshold: float32(withdrawThreshold)}, nil

}
