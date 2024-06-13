package exchange

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/exchange"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindExchangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindExchangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindExchangeLogic {
	return &FindExchangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindExchangeLogic) FindExchange(in *wolflamp.FindExchangeReq) (*wolflamp.ExchangeInfo, error) {

	query := l.svcCtx.DB.Exchange.Query().Where(exchange.ID(in.Id))
	if in.PlayerId != nil {
		query.Where(exchange.PlayerID(*in.PlayerId))
	}
	result, err := query.First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return l.Po2Vo(result), nil

}

func (l *FindExchangeLogic) Po2Vo(po *ent.Exchange) (vo *wolflamp.ExchangeInfo) {

	return &wolflamp.ExchangeInfo{
		Id:            po.ID,
		CreatedAt:     po.CreatedAt.Unix(),
		UpdatedAt:     po.UpdatedAt.Unix(),
		Status:        uint32(po.Status),
		Type:          po.Type,
		TransactionId: po.TransactionID,
		LampNum:       po.LampNum,
		PlayerId:      po.PlayerID,
		CoinNum:       po.CoinNum,
	}

}
