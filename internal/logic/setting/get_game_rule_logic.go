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

type GetGameRuleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGameRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGameRuleLogic {
	return &GetGameRuleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGameRuleLogic) GetGameRule(in *wolflamp.Empty) (*wolflamp.GameRuleResp, error) {

	if l.svcCtx.Redis.Exists(l.ctx, enum.GameRule.CacheKey()).Val() == 0 {
		setting, err := NewFindSettingLogic(l.ctx, l.svcCtx).FindSetting(&wolflamp.FindSettingReq{Module: enum.PlatformSetting.Val()})
		if err != nil {
			return nil, err
		}
		var platformSetting entity.PlatformSetting
		if err := json.Unmarshal([]byte(setting.JsonString), &platformSetting); err != nil {
			return nil, err
		}
		gameRuleJson, err := json.Marshal(platformSetting.GameRule)
		if err != nil {
			return nil, err
		}
		_ = l.svcCtx.Redis.Set(l.ctx, enum.GameRule.CacheKey(), gameRuleJson, 0)
	}

	cached, err := l.svcCtx.Redis.Get(l.ctx, enum.GameRule.CacheKey()).Result()
	if err != nil {
		return nil, err
	}
	var gameRule entity.GameRule
	if err := json.Unmarshal([]byte(cached), &gameRule); err != nil {
		return nil, err
	}
	return &wolflamp.GameRuleResp{
		Title:   gameRule.Title,
		Content: gameRule.Content,
	}, nil

}
