package order

import (
	"context"
	"github.com/duke-git/lancet/v2/random"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"time"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *wolflamp.CreateOrderReq) (*wolflamp.BaseIDResp, error) {

	orderCode := time.Now().Format("06010215") + random.RandString(4)

	result, err := l.svcCtx.DB.Order.Create().
		SetCode(orderCode).
		SetStatus(uint8(in.Status)).
		SetPlayerID(in.PlayerId).
		SetTransactionID("").
		SetHandlingFee(in.HandlingFee).
		SetNum(in.Num).
		SetType(in.Type).
		SetToAddress(in.ToAddress).
		SetFromAddress(in.FromAddress).
		SetNetwork(in.Network).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil

}
