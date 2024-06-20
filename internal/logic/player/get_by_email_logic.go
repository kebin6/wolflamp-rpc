package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/player"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByEmailLogic {
	return &GetByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByEmailLogic) GetByEmail(in *wolflamp.GetByEmailReq) (*wolflamp.PlayerInfo, error) {

	result, err := l.svcCtx.DB.Player.Query().Where(player.Email(in.Email)).First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	playerInfo := NewFindPlayerLogic(l.ctx, l.svcCtx).Po2Vo(result)
	playerInfo.Password = result.Password
	return playerInfo, nil

}
