package exchange

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/exchange"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteExchangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteExchangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteExchangeLogic {
	return &DeleteExchangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteExchangeLogic) DeleteExchange(in *wolflamp.DeleteExchangeReq) (*wolflamp.BaseIDResp, error) {

	_, err := l.svcCtx.DB.Exchange.Delete().Where(exchange.ID(in.Id)).Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Msg: i18n.DeleteSuccess}, nil

}
