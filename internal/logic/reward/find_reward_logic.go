package reward

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/reward"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRewardLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRewardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRewardLogic {
	return &FindRewardLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindRewardLogic) FindReward(in *wolflamp.FindRewardReq) (*wolflamp.RewardInfo, error) {

	query := l.svcCtx.DB.Reward.Query().Where(reward.ID(in.Id))
	if in.ToId != nil {
		query.Where(reward.ToID(*in.ToId))
	}
	result, err := query.First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return l.Po2Vo(result), nil

}

func (l *FindRewardLogic) Po2Vo(po *ent.Reward) (vo *wolflamp.RewardInfo) {

	return &wolflamp.RewardInfo{
		Id:               po.ID,
		CreatedAt:        po.CreatedAt.Unix(),
		UpdatedAt:        po.UpdatedAt.Unix(),
		Status:           uint32(po.Status),
		ContributorId:    po.ContributorID,
		ContributorEmail: po.ContributorEmail,
		ContributorLevel: po.ContributorLevel,
		Num:              po.Num,
		Remark:           po.Remark,
		ToId:             po.ToID,
	}

}
