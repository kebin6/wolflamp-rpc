package reward

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/ent/reward"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListRewardLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListRewardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRewardLogic {
	return &ListRewardLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListRewardLogic) ListReward(in *wolflamp.ListRewardReq) (*wolflamp.ListRewardResp, error) {

	var predicates []predicate.Reward
	if in.Status != nil {
		predicates = append(predicates, reward.Status(*pointy.GetStatusPointer(in.Status)))
	}
	if in.ToId != nil {
		predicates = append(predicates, reward.ToID(*in.ToId))
	}
	if in.ContributorLevel != nil {
		predicates = append(predicates, reward.ContributorLevel(*in.ContributorLevel))
	}

	result, err := l.svcCtx.DB.Reward.Query().Where(predicates...).
		Order(ent.Desc("id")).
		Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &wolflamp.ListRewardResp{}
	resp.Total = result.PageDetails.Total

	findLogic := NewFindRewardLogic(l.ctx, l.svcCtx)
	for _, v := range result.List {
		resp.Data = append(resp.Data, findLogic.Po2Vo(v))
	}

	return resp, nil

}
