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

type GetWithdrawCommissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWithdrawCommissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWithdrawCommissionLogic {
	return &GetWithdrawCommissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWithdrawCommissionLogic) GetWithdrawCommission(in *wolflamp.Empty) (*wolflamp.WithdrawCommissionResp, error) {

	if l.svcCtx.Redis.Exists(l.ctx, enum.WithdrawCommission.CacheKey()).Val() == 0 {
		setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformSetting.Val()})
		if err != nil {
			return nil, err
		}
		var platformSetting entity.PlatformSetting
		if err := json.Unmarshal([]byte(setting.JsonString), &platformSetting); err != nil {
			return nil, err
		}
		_ = l.svcCtx.Redis.Set(l.ctx, enum.WithdrawCommission.CacheKey(), platformSetting.WithdrawCommission, 0)
	}

	cached, err := l.svcCtx.Redis.Get(l.ctx, enum.WithdrawCommission.CacheKey()).Result()
	if err != nil {
		return nil, err
	}
	withdrawCommission, err := strconv.ParseFloat(cached, 32)
	if err != nil {
		return nil, err
	}

	return &wolflamp.WithdrawCommissionResp{WithdrawCommission: float32(withdrawCommission)}, nil

}
