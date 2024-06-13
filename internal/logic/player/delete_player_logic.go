package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/player"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePlayerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePlayerLogic {
	return &DeletePlayerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeletePlayerLogic) DeletePlayer(in *wolflamp.DeletePlayerReq) (*wolflamp.BaseIDResp, error) {

	_, err := l.svcCtx.DB.Player.Delete().Where(player.ID(in.Id)).Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Msg: i18n.DeleteSuccess}, nil

}
