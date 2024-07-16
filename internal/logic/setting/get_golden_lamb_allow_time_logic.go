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

type GetGoldenLambAllowTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoldenLambAllowTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoldenLambAllowTimeLogic {
	return &GetGoldenLambAllowTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoldenLambAllowTimeLogic) GetGoldenLambAllowTime(in *wolflamp.Empty) (*wolflamp.HourTimeRangeResp, error) {

	if l.svcCtx.Redis.Exists(l.ctx, enum.GoldenLambAllowTime.CacheKey()).Val() == 0 {
		setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).
			FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformSetting.Val()})
		if err != nil {
			return nil, err
		}
		var platformSetting entity.PlatformSetting
		if err := json.Unmarshal([]byte(setting.JsonString), &platformSetting); err != nil {
			return nil, err
		}
		settingJson, err := json.Marshal(platformSetting.GoldenLambAllowTime)
		if err != nil {
			return nil, err
		}
		_ = l.svcCtx.Redis.Set(l.ctx, enum.GoldenLambAllowTime.CacheKey(), settingJson, 0)
	}

	cached, err := l.svcCtx.Redis.Get(l.ctx, enum.GoldenLambAllowTime.CacheKey()).Result()
	if err != nil {
		return nil, err
	}
	var timeRange entity.GoldenLambAllowTime
	if err := json.Unmarshal([]byte(cached), &timeRange); err != nil {
		return nil, err
	}
	return &wolflamp.HourTimeRangeResp{
		StartTime: timeRange.StartTime,
		EndTime:   timeRange.EndTime,
	}, nil

}
