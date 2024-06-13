package reward

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRewardLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRewardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRewardLogic {
	return &CreateRewardLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRewardLogic) CreateReward(in *wolflamp.CreateRewardReq) (*wolflamp.BaseIDResp, error) {

	result, err := l.svcCtx.DB.Reward.Create().
		SetStatus(uint8(1)).
		SetToID(in.ToId).
		SetContributorID(in.ContributorId).
		SetContributorEmail(in.ContributorEmail).
		SetContributorLevel(in.ContributorLevel).
		SetNum(in.Num).
		SetRemark(in.Remark).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil

}
