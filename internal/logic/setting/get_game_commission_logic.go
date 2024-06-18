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

type GetGameCommissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGameCommissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGameCommissionLogic {
	return &GetGameCommissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGameCommissionLogic) GetGameCommission(in *wolflamp.Empty) (*wolflamp.GameCommissionResp, error) {

	if l.svcCtx.Redis.Exists(l.ctx, enum.GameCommission.CacheKey()).Val() == 0 {
		setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformSetting.Val()})
		if err != nil {
			return nil, err
		}
		var platformSetting entity.PlatformSetting
		if err := json.Unmarshal([]byte(setting.JsonString), &platformSetting); err != nil {
			return nil, err
		}
		_ = l.svcCtx.Redis.Set(l.ctx, enum.GameCommission.CacheKey(), platformSetting.GameCommission, 0)
	}

	cached, err := l.svcCtx.Redis.Get(l.ctx, enum.GameCommission.CacheKey()).Result()
	if err != nil {
		return nil, err
	}
	gameCommission, err := strconv.ParseFloat(cached, 32)
	if err != nil {
		return nil, err
	}
	return &wolflamp.GameCommissionResp{
		GameCommission: float32(gameCommission),
	}, nil

}
