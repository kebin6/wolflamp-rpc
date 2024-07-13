package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/api"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGcicsBalanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGcicsBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGcicsBalanceLogic {
	return &GetGcicsBalanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGcicsBalanceLogic) GetGcicsBalance(in *wolflamp.GetGcicsBalanceReq) (*wolflamp.GetGcicsBalanceResp, error) {
	player, err := l.svcCtx.DB.Player.Get(l.ctx, in.PlayerId)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	// 获取GCICS平台用户的钱包余额信息
	gcicsApi := api.GcicsApi{
		Ctx: l.ctx, SvcCtx: l.svcCtx,
		UserId: player.ID,
	}
	balance, err := gcicsApi.GetBalance(player.GcicsToken)
	if err != nil {
		return nil, err
	}
	return &wolflamp.GetGcicsBalanceResp{
		Coin:  *balance.Coin,
		Token: *balance.Token,
	}, nil
}
