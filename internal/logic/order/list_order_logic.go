package order

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/order"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrderLogic {
	return &ListOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListOrderLogic) ListOrder(in *wolflamp.ListOrderReq) (*wolflamp.ListOrderResp, error) {

	var predicates []predicate.Order
	if in.Status != nil {
		predicates = append(predicates, order.Status(*pointy.GetStatusPointer(in.Status)))
	}
	if in.Type != nil {
		predicates = append(predicates, order.Type(*in.Type))
	}
	if in.Code != nil {
		predicates = append(predicates, order.Code(*in.Code))
	}
	if in.PlayerId != nil {
		predicates = append(predicates, order.PlayerID(*in.PlayerId))
	}

	result, err := l.svcCtx.DB.Order.Query().Where(predicates...).
		Order(ent.Desc(order.FieldID)).
		Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &wolflamp.ListOrderResp{}
	resp.Total = result.PageDetails.Total

	findLogic := NewFindOrderLogic(l.ctx, l.svcCtx)
	for _, v := range result.List {
		resp.Data = append(resp.Data, findLogic.Po2Vo(v))
	}

	return resp, nil

}
