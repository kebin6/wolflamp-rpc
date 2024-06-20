package reward

import (
	"context"
	"github.com/jinzhu/now"
	"github.com/kebin6/wolflamp-rpc/common/enum/rewardenum"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/reward"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/zeromicro/go-zero/core/logx"
)

type RewardAggregateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRewardAggregateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RewardAggregateLogic {
	return &RewardAggregateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RewardAggregateLogic) RewardAggregate(in *wolflamp.RewardAggregateReq) (resp *wolflamp.RewardAggregateResp, err error) {

	var total float64
	if in.GetTotal {
		total, _ = l.svcCtx.DB.Reward.Query().
			Where(reward.ToID(in.ToId)).
			Where(reward.Status(uint8(rewardenum.Completed.Val()))).
			Aggregate(ent.Sum(reward.FieldNum)).Float64(l.ctx)
	}
	var today float64
	if in.GetToday {
		today, _ = l.svcCtx.DB.Reward.Query().
			Where(reward.ToID(in.ToId)).
			Where(reward.Status(uint8(rewardenum.Completed.Val()))).
			Where(reward.CreatedAtGTE(now.BeginningOfDay())).
			//Where(reward.CreatedAtLTE(now.EndOfDay())).
			Aggregate(ent.Sum(reward.FieldNum)).
			Float64(l.ctx)
	}
	return &wolflamp.RewardAggregateResp{
		Total: &total,
		Today: &today,
	}, nil

}
