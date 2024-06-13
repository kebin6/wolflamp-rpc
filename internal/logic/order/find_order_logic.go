package order

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/order"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOrderLogic {
	return &FindOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOrderLogic) FindOrder(in *wolflamp.FindOrderReq) (*wolflamp.OrderInfo, error) {

	query := l.svcCtx.DB.Order.Query().Where(order.ID(in.Id))
	if in.PlayerId != nil {
		query.Where(order.PlayerID(*in.PlayerId))
	}
	result, err := query.First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return l.Po2Vo(result), nil

}

func (l *FindOrderLogic) Po2Vo(po *ent.Order) (vo *wolflamp.OrderInfo) {

	return &wolflamp.OrderInfo{
		Id:            po.ID,
		CreatedAt:     po.CreatedAt.Unix(),
		UpdatedAt:     po.UpdatedAt.Unix(),
		Status:        uint32(po.Status),
		Code:          po.Code,
		Num:           po.Num,
		Type:          po.Type,
		ToAddress:     po.ToAddress,
		FromAddress:   po.FromAddress,
		TransactionId: po.TransactionID,
		HandlingFee:   po.HandlingFee,
		Network:       po.Network,
	}

}
