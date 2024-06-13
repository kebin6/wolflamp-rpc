package exchange

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

type CreateExchangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateExchangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExchangeLogic {
	return &CreateExchangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateExchangeLogic) CreateExchange(in *wolflamp.CreateExchangeReq) (*wolflamp.BaseIDResp, error) {

	transactionId := time.Now().Format("020601021504") + random.RandNumeral(4)
	result, err := l.svcCtx.DB.Exchange.Create().
		SetStatus(uint8(1)).
		SetPlayerID(in.PlayerId).
		SetTransactionID(transactionId).
		SetCoinNum(in.CoinNum).
		SetLampNum(in.LampNum).
		SetType(in.Type).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil

}
