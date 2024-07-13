package exchange

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/enum/exchangeenum"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyLogic {
	return &NotifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NotifyLogic) Notify(in *wolflamp.NotifyExchangeReq) (*wolflamp.BaseIDResp, error) {

	exchangeLog, err := l.svcCtx.DB.Exchange.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	if float64(exchangeLog.LampNum) != in.Amount {
		return nil, errorx.NewCodeInvalidArgumentError("lamb amount not match")
	}

	if in.IsPaid {
		exchangeLog.Status = uint8(exchangeenum.Completed)
	} else {
		exchangeLog.Status = uint8(exchangeenum.Failed)
	}

	err = exchangeLog.Update().Exec(l.ctx)
	if err != nil {
		return nil, errorx.NewCodeInternalError(err.Error())
	}
	return &wolflamp.BaseIDResp{Id: in.Id}, nil

}
