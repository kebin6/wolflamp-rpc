package reward

import (
	"context"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRewardLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRewardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRewardLogic {
	return &UpdateRewardLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRewardLogic) UpdateReward(in *wolflamp.UpdateRewardReq) (*wolflamp.BaseIDResp, error) {
	// todo: add your logic here and delete this line

	return &wolflamp.BaseIDResp{}, nil
}
