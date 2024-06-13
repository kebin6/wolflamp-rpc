package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/player"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByInviteCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByInviteCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByInviteCodeLogic {
	return &GetByInviteCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByInviteCodeLogic) GetByInviteCode(in *wolflamp.GetByInviteCodeReq) (*wolflamp.PlayerInfo, error) {

	result, err := l.svcCtx.DB.Player.Query().Where(player.InviteCode(in.InviteCode)).First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return NewFindPlayerLogic(l.ctx, l.svcCtx).Po2Vo(result), nil

}
