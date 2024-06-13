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

type GetRobotNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRobotNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRobotNumLogic {
	return &GetRobotNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRobotNumLogic) GetRobotNum(in *wolflamp.Empty) (*wolflamp.RobotNumResp, error) {

	if l.svcCtx.Redis.Exists(l.ctx, enum.RobotNum.CacheKey()).Val() == 0 {
		setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformSetting.Val()})
		if err != nil {
			return nil, err
		}
		var platformSetting entity.PlatformSetting
		if err := json.Unmarshal([]byte(setting.JsonString), &platformSetting); err != nil {
			return nil, err
		}
		robotNumJson, err := json.Marshal(platformSetting.RobotNum)
		if err != nil {
			return nil, err
		}
		_ = l.svcCtx.Redis.Set(l.ctx, enum.RobotNum.CacheKey(), robotNumJson, 0)
	}

	cached, err := l.svcCtx.Redis.Get(l.ctx, enum.RobotNum.CacheKey()).Result()
	if err != nil {
		return nil, err
	}
	var robotNum entity.RobotNum
	if err := json.Unmarshal([]byte(cached), &robotNum); err != nil {
		return nil, err
	}
	return &wolflamp.RobotNumResp{
		Min: robotNum.Min,
		Max: robotNum.Max,
	}, nil

}
