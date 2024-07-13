package exchange

import (
	"context"
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/kebin6/wolflamp-rpc/common/api"
	"github.com/kebin6/wolflamp-rpc/common/enum/exchangeenum"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/internal/logic/player"
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

func (l *ExchangeLogic) Exchange(in *wolflamp.ExchangeReq) (*wolflamp.ExchangeResp, error) {

	// 余额兑换小羊1:10
	// 羊兑币，羊数量必须为10的整数倍
	if in.Type == 2 && in.LampAmount%10 != 0 {
		return nil, errorx.NewInvalidArgumentError("wallet.invalidLambNum")
	}

	// 币兑羊
	if in.Type == 1 {
		return l.DoExchange(in)
	} else {
		return l.DoWithdraw(in)
	}

}

func (l *ExchangeLogic) DoExchange(in *wolflamp.ExchangeReq) (*wolflamp.ExchangeResp, error) {

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

	balance, err := player.NewGetGcicsBalanceLogic(l.ctx, l.svcCtx).
		GetGcicsBalance(&wolflamp.GetGcicsBalanceReq{PlayerId: info.ID})
	if err != nil {
		return nil, err
	}

	var amount float64
	if in.Mode == "coin" {
		amount = balance.Coin
	} else {
		amount = balance.Token
	}

	// 检查余额是否充足
	if amount < float64(in.CoinAmount) {
		return nil, errorx.NewInvalidArgumentError("wallet.insufficientCoin")
	}

	// 创建交易单
	exchageInfo, err := l.CreateExchangeLog(in)
	if err != nil {
		return nil, err
	}

	// 获取GCICS平台用户的钱包余额信息
	gcicsApi := api.GcicsApi{
		Ctx: l.ctx, SvcCtx: l.svcCtx,
		UserId: info.ID,
	}
	link, err := gcicsApi.GeneratePaymentLink(exchageInfo.ID, in.Mode, float64(in.CoinAmount))
	if err != nil {
		return nil, err
	}
	dealLink := fmt.Sprintf(*link+"?q=%s&k=%s&lang=en", "", "")

	return &wolflamp.ExchangeResp{
		Id:   exchageInfo.ID,
		Link: &dealLink,
	}, nil
}

func (l *ExchangeLogic) DoWithdraw(in *wolflamp.ExchangeReq) (*wolflamp.ExchangeResp, error) {
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

	var lambNum float32
	if in.Mode == "coin" {
		lambNum = info.CoinLamb
	} else {
		lambNum = info.TokenLamb
	}

	if lambNum < float32(in.LampAmount) {
		return nil, errorx.NewInvalidArgumentError("wallet.insufficientLamb")
	}

	// 创建交易单
	exchageInfo, err := l.CreateExchangeLog(in)
	if err != nil {
		return nil, err
	}

	// 获取GCICS平台用户的钱包余额信息
	gcicsApi := api.GcicsApi{
		Ctx: l.ctx, SvcCtx: l.svcCtx,
		UserId: info.ID,
	}

	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		err = gcicsApi.Withdraw(in.Mode, float64(in.LampAmount))
		if err != nil {
			exchageInfo.Status = uint8(exchangeenum.Failed)
			exchageInfo.Remark = "gcics occur an error: " + err.Error()
			err := exchageInfo.Update().Exec(l.ctx)
			if err != nil {
				return err
			}
			return nil
		}

		// 更新用户表数据
		if in.Mode == "coin" {
			_, err := l.svcCtx.DB.Player.UpdateOneID(in.PlayerId).
				AddCoinLamb(-float32(in.LampAmount)).
				Save(l.ctx)
			if err != nil {
				return err
			}
		} else {
			_, err := l.svcCtx.DB.Player.UpdateOneID(in.PlayerId).
				AddTokenLamb(-float32(in.LampAmount)).
				Save(l.ctx)
			if err != nil {
				return err
			}
		}

		exchageInfo.Status = uint8(exchangeenum.Completed)
		return exchageInfo.Update().Exec(l.ctx)
	})

	return &wolflamp.ExchangeResp{
		Id: exchageInfo.ID,
	}, nil
}

func (l *ExchangeLogic) CreateExchangeLog(in *wolflamp.ExchangeReq) (*ent.Exchange, error) {
	transactionId := time.Now().Format("020601021504") + random.RandNumeral(4)
	// 创建兑换单
	exchangeInfo, err := l.svcCtx.DB.Exchange.Create().
		SetPlayerID(in.PlayerId).
		SetType(in.Type).
		SetTransactionID(transactionId).
		SetLampNum(in.LampAmount).
		SetCoinNum(in.CoinAmount).
		SetStatus(uint8(exchangeenum.Created.Val())).
		SetMode(in.Mode).
		Save(l.ctx)

	if err != nil {
		return nil, err
	}

	return exchangeInfo, nil
}
