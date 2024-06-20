package reward

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/reward"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRewardLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRewardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRewardLogic {
	return &DeleteRewardLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRewardLogic) DeleteReward(in *wolflamp.DeleteRewardReq) (*wolflamp.BaseIDResp, error) {

	_, err := l.svcCtx.DB.Reward.Delete().Where(reward.ID(in.Id)).Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Msg: i18n.DeleteSuccess}, nil

}
