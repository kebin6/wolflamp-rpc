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

type GetSliverOccurPercentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSliverOccurPercentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSliverOccurPercentLogic {
	return &GetSliverOccurPercentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSliverOccurPercentLogic) GetSliverOccurPercent(in *wolflamp.Empty) (*wolflamp.PercentResp, error) {

	if l.svcCtx.Redis.Exists(l.ctx, enum.SliverOccurPercent.CacheKey()).Val() == 0 {
		setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformSetting.Val()})
		if err != nil {
			return nil, err
		}
		var platformSetting entity.PlatformSetting
		if err := json.Unmarshal([]byte(setting.JsonString), &platformSetting); err != nil {
			return nil, err
		}
		_ = l.svcCtx.Redis.Set(l.ctx, enum.SliverOccurPercent.CacheKey(), platformSetting.SliverOccurPercent, 0)
	}

	cached, err := l.svcCtx.Redis.Get(l.ctx, enum.SliverOccurPercent.CacheKey()).Result()
	if err != nil {
		return nil, err
	}
	SliverOccurPercent, err := strconv.ParseFloat(cached, 32)
	if err != nil {
		return nil, err
	}

	return &wolflamp.PercentResp{Percent: float32(SliverOccurPercent)}, nil

}
