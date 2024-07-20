package exchange

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/enum/exchangeenum"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/internal/utils/entx"
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
	if exchangeLog.Status != uint8(exchangeenum.Created) {
		return nil, errorx.NewCodeInvalidArgumentError("Forbidden")
	}

	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		updateQuery := exchangeLog.Update()
		updatePlayer := l.svcCtx.DB.Player.UpdateOneID(exchangeLog.PlayerID)
		if in.IsPaid {
			updateQuery.SetStatus(uint8(exchangeenum.Completed))
			if exchangeLog.Mode == "coin" {
				updatePlayer.AddCoinLamb(float32(in.Amount))
			} else {
				updatePlayer.AddTokenLamb(float32(in.Amount))
			}
			err := updateQuery.Exec(l.ctx)
			if err != nil {
				return err
			}
		} else {
			updateQuery.SetStatus(uint8(exchangeenum.Failed))
		}

		err = updateQuery.Exec(l.ctx)
		if err != nil {
			return errorx.NewCodeInternalError(err.Error())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return &wolflamp.BaseIDResp{Id: in.Id}, nil

}
