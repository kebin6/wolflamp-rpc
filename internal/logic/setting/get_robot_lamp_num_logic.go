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

type GetRobotLampNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRobotLampNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRobotLampNumLogic {
	return &GetRobotLampNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRobotLampNumLogic) GetRobotLampNum(in *wolflamp.Empty) (*wolflamp.RobotLampNumResp, error) {

	if l.svcCtx.Redis.Exists(l.ctx, enum.RobotLampNum.CacheKey()).Val() == 0 {
		setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformSetting.Val()})
		if err != nil {
			return nil, err
		}
		var platformSetting entity.PlatformSetting
		if err := json.Unmarshal([]byte(setting.JsonString), &platformSetting); err != nil {
			return nil, err
		}
		robotLambJson, err := json.Marshal(platformSetting.RobotLampNum)
		if err != nil {
			return nil, err
		}
		_ = l.svcCtx.Redis.Set(l.ctx, enum.RobotLampNum.CacheKey(), robotLambJson, 0)
	}

	cached, err := l.svcCtx.Redis.Get(l.ctx, enum.RobotLampNum.CacheKey()).Result()
	if err != nil {
		return nil, err
	}
	var robotLambNum entity.RobotLampNum
	if err := json.Unmarshal([]byte(cached), &robotLambNum); err != nil {
		return nil, err
	}
	return &wolflamp.RobotLampNumResp{
		Min: robotLambNum.Min,
		Max: robotLambNum.Max,
	}, nil

}
