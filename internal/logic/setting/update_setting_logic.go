package setting

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/enum"
	"github.com/kebin6/wolflamp-rpc/ent/setting"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSettingLogic {
	return &UpdateSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSettingLogic) UpdateSetting(in *wolflamp.UpdateSettingReq) (*wolflamp.BaseIDResp, error) {

	err := l.svcCtx.DB.Setting.Update().
		Where(setting.Module(in.Module)).
		SetValue(in.JsonString).
		Exec(l.ctx)
	l.svcCtx.Redis.Del(l.ctx, enum.GameRule.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.WithdrawCommission.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.MinWithdrawNum.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.RobotNum.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.RobotLampNum.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.IdleTime.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.WithdrawThreshold.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.GameCommission.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.PoolCommission.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.RobPoolCommission.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.RewardPoolCommission.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.GoldenLambAllowTime.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.GoldenLambNum.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.PoolMinNumThenSilver.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.SliverOccurPercent.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.SliverLambNum.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.SliverLambPercent.CacheKey())
	l.svcCtx.Redis.Del(l.ctx, enum.GoldenLambPercent.CacheKey())
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Msg: i18n.UpdateSuccess}, nil

}
