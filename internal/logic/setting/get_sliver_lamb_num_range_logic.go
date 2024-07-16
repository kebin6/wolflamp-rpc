package setting

import (
	"context"
	"encoding/json"
	"github.com/kebin6/wolflamp-rpc/common/entity"
	"github.com/kebin6/wolflamp-rpc/common/enum"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSliverLambNumRangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSliverLambNumRangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSliverLambNumRangeLogic {
	return &GetSliverLambNumRangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSliverLambNumRangeLogic) GetSliverLambNumRange(in *wolflamp.Empty) (*wolflamp.NumRangeResp, error) {
	if l.svcCtx.Redis.Exists(l.ctx, enum.SliverLambNum.CacheKey()).Val() == 0 {
		setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformSetting.Val()})
		if err != nil {
			return nil, err
		}
		var platformSetting entity.PlatformSetting
		if err := json.Unmarshal([]byte(setting.JsonString), &platformSetting); err != nil {
			return nil, err
		}
		settingJson, err := json.Marshal(platformSetting.SliverLambNum)
		if err != nil {
			return nil, err
		}
		_ = l.svcCtx.Redis.Set(l.ctx, enum.SliverLambNum.CacheKey(), settingJson, 0)
	}

	cached, err := l.svcCtx.Redis.Get(l.ctx, enum.SliverLambNum.CacheKey()).Result()
	if err != nil {
		return nil, err
	}
	var numRange entity.GoldenLambNum
	if err := json.Unmarshal([]byte(cached), &numRange); err != nil {
		return nil, err
	}
	return &wolflamp.NumRangeResp{
		Min: numRange.Min,
		Max: numRange.Max,
	}, nil
}
