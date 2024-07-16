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

type GetPoolMinNumThenSilverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPoolMinNumThenSilverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoolMinNumThenSilverLogic {
	return &GetPoolMinNumThenSilverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPoolMinNumThenSilverLogic) GetPoolMinNumThenSilver(in *wolflamp.Empty) (*wolflamp.PoolMinNumThenSilverResp, error) {

	if l.svcCtx.Redis.Exists(l.ctx, enum.PoolMinNumThenSilver.CacheKey()).Val() == 0 {
		setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformSetting.Val()})
		if err != nil {
			return nil, err
		}
		var platformSetting entity.PlatformSetting
		if err := json.Unmarshal([]byte(setting.JsonString), &platformSetting); err != nil {
			return nil, err
		}
		_ = l.svcCtx.Redis.Set(l.ctx, enum.PoolMinNumThenSilver.CacheKey(), platformSetting.PoolMinNumThenSilver, 0)
	}

	cached, err := l.svcCtx.Redis.Get(l.ctx, enum.PoolMinNumThenSilver.CacheKey()).Result()
	if err != nil {
		return nil, err
	}
	min, err := strconv.ParseUint(cached, 10, 32)
	if err != nil {
		return nil, err
	}

	return &wolflamp.PoolMinNumThenSilverResp{Min: uint32(min)}, nil

}
