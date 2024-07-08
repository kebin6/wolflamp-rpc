package exchange

import (
	"context"
	"github.com/duke-git/lancet/v2/random"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/internal/utils/entx"
	"github.com/zeromicro/go-zero/core/errorx"
	"time"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExchangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExchangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeLogic {
	return &ExchangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExchangeLogic) Exchange(in *wolflamp.ExchangeReq) (*wolflamp.BaseIDResp, error) {

	// 余额兑换小羊1:10
	// 羊兑币，羊数量必须为10的整数倍
	if in.Type == 2 && in.LampAmount%10 != 0 {
		return nil, errorx.NewInvalidArgumentError("wallet.invalidLambNum")
	}

	info, err := l.svcCtx.DB.Player.Get(l.ctx, in.PlayerId)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.NewInternalError("common.playerNotFound")
		}
		return nil, err
	}

	lambNum := info.CoinLamb
	coinNum := info.Amount
	switch in.Type {
	// 币兑羊
	case 1:
		if info.Amount < float64(in.CoinAmount) {
			return nil, errorx.NewInvalidArgumentError("wallet.insufficientCoin")
		}
		lambNum += float32(in.CoinAmount * 10)
		coinNum -= float64(in.CoinAmount)
	case 2:
		if info.CoinLamb < float32(in.LampAmount) {
			return nil, errorx.NewInvalidArgumentError("wallet.insufficientLamb")
		}
		lambNum -= float32(in.LampAmount)
		coinNum += float64(in.LampAmount / 10)
	}

	transactionId := time.Now().Format("020601021504") + random.RandNumeral(4)
	var exchange *ent.Exchange
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		// 更新用户表数据
		_, err := l.svcCtx.DB.Player.UpdateOneID(in.PlayerId).
			SetCoinLamb(lambNum).
			SetAmount(coinNum).
			Save(l.ctx)
		if err != nil {
			return err
		}
		// 创建兑换单
		exchange, err = l.svcCtx.DB.Exchange.Create().
			SetPlayerID(in.PlayerId).
			SetType(in.Type).
			SetTransactionID(transactionId).
			SetLampNum(in.LampAmount).
			SetCoinNum(in.CoinAmount).
			Save(l.ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &wolflamp.BaseIDResp{
		Id: exchange.ID,
	}, nil

}
