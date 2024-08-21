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

type GetSliverLambPercentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSliverLambPercentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSliverLambPercentLogic {
	return &GetSliverLambPercentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSliverLambPercentLogic) GetSliverLambPercent(in *wolflamp.Empty) (*wolflamp.PercentResp, error) {
	if l.svcCtx.Redis.Exists(l.ctx, enum.SliverLambPercent.CacheKey()).Val() == 0 {
		setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformSetting.Val()})
		if err != nil {
			return nil, err
		}
		var platformSetting entity.PlatformSetting
		if err := json.Unmarshal([]byte(setting.JsonString), &platformSetting); err != nil {
			return nil, err
		}
		_ = l.svcCtx.Redis.Set(l.ctx, enum.SliverLambPercent.CacheKey(), platformSetting.SliverLambPercent, 0)
	}

	cached, err := l.svcCtx.Redis.Get(l.ctx, enum.SliverLambPercent.CacheKey()).Result()
	if err != nil {
		return nil, err
	}
	sliverLambPercent, err := strconv.ParseFloat(cached, 32)
	if err != nil {
		return nil, err
	}

	return &wolflamp.PercentResp{Percent: float32(sliverLambPercent)}, nil
}
