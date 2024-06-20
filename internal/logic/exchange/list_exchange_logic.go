package exchange

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/exchange"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListExchangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListExchangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListExchangeLogic {
	return &ListExchangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListExchangeLogic) ListExchange(in *wolflamp.ListExchangeReq) (*wolflamp.ListExchangeResp, error) {

	var predicates []predicate.Exchange
	if in.Type != nil {
		predicates = append(predicates, exchange.Type(*in.Type))
	}
	if in.PlayerId != nil {
		predicates = append(predicates, exchange.PlayerID(*in.PlayerId))
	}
	if in.PlayerId != nil {
		predicates = append(predicates, exchange.PlayerID(*in.PlayerId))
	}

	result, err := l.svcCtx.DB.Exchange.Query().Where(predicates...).
		Order(ent.Desc("id")).
		Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &wolflamp.ListExchangeResp{}
	resp.Total = result.PageDetails.Total

	findLogic := NewFindExchangeLogic(l.ctx, l.svcCtx)
	for _, v := range result.List {
		resp.Data = append(resp.Data, findLogic.Po2Vo(v))
	}

	return resp, nil

}
