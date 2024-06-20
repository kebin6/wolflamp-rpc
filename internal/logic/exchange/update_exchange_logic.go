package exchange

import (
	"context"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExchangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateExchangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExchangeLogic {
	return &UpdateExchangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateExchangeLogic) UpdateExchange(in *wolflamp.UpdateExchangeReq) (*wolflamp.BaseIDResp, error) {
	// todo: add your logic here and delete this line

	return &wolflamp.BaseIDResp{}, nil
}
