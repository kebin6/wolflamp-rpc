package order

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderLogic) UpdateOrder(in *wolflamp.UpdateOrderReq) (*wolflamp.BaseIDResp, error) {

	err := l.svcCtx.DB.Order.UpdateOneID(in.Id).
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilType(in.Type).
		SetNotNilToAddress(in.ToAddress).
		SetNotNilFromAddress(in.FromAddress).
		SetNotNilTransactionID(in.TransactionId).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Msg: i18n.UpdateSuccess}, nil

}
