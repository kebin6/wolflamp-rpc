package game

import (
	"context"
	"fmt"
	"github.com/kebin6/wolflamp-rpc/common/enum/poolenum"
	"github.com/kebin6/wolflamp-rpc/common/enum/roundenum"
	"github.com/kebin6/wolflamp-rpc/common/util"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/internal/utils/entx"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"
	"time"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type InvestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvestLogic {
	return &InvestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvestLogic) Invest(in *wolflamp.CreateInvestReq) (*wolflamp.BaseIDResp, error) {

	var player *ent.Player
	// 判断当前玩家剩余羊数量是否满足要求
	if util.IsRealPlayer(in.PlayerId) {
		playerInfo, err := l.svcCtx.DB.Player.Get(l.ctx, in.PlayerId)
		if err != nil {
			return nil, err
		}
		if playerInfo.CoinLamb < float32(in.LambNum) {
			return nil, errorx.NewInvalidArgumentError("game.lambNotEnough")
		}
		player = playerInfo
	}
	// 获取当前轮次信息
	roundInfo, err := NewFindRoundLogic(l.ctx, l.svcCtx).FindRound(&wolflamp.FindRoundReq{Mode: in.Mode})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	if in.RoundId != roundInfo.Id {
		return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("round id %d has expired", in.RoundId))
	}

	// 不在投注阶段
	if roundInfo.Status != roundenum.Investing.Val() || roundInfo.StartAt > time.Now().Unix() ||
		roundInfo.OpenAt < time.Now().Unix() {
		return nil, errorx.NewInvalidArgumentError(fmt.Sprint("investing time over"))
	}

	// 判断当前真人投注羊圈是否大于等于3
	//l.svcCtx.DB.RoundInvest.Query().Where(roundinvest.PlayerID(in.PlayerId), roundinvest.RoundID(in.RoundId))

	var result *ent.RoundInvest
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		investedOne, err := tx.RoundInvest.Query().
			Where(roundinvest.PlayerID(in.PlayerId), roundinvest.RoundID(in.RoundId)).First(l.ctx)
		if err != nil && !ent.IsNotFound(err) {
			return err
		}
		if investedOne != nil && investedOne.FoldNo != in.FoldNo {
			return errorx.NewInvalidArgumentError("game.foldNoNotMatch")
		}

		playerEmail := ""
		if player != nil {
			playerEmail = player.Email
			if in.Mode == "coin" {
				err = tx.Player.UpdateOneID(in.PlayerId).
					AddCoinLamb(-float32(in.LambNum)).
					Exec(l.ctx)
			} else {
				err = tx.Player.UpdateOneID(in.PlayerId).
					AddTokenLamb(-float32(in.LambNum)).
					Exec(l.ctx)
			}
			if err != nil {
				return err
			}
		} else {
			// 机器人投注需要扣除机器人池数量
			err := tx.Pool.Create().SetMode(in.Mode).
				SetStatus(1).SetRoundID(roundInfo.Id).
				SetType(poolenum.Robot.Val()).
				SetLambNum(-float64(in.LambNum)).
				SetRemark("机器人投注").
				Exec(l.ctx)
			if err != nil {
				return err
			}
		}
		if investedOne == nil {
			result, err = tx.RoundInvest.Create().
				SetMode(in.Mode).
				SetPlayerID(in.PlayerId).
				SetPlayerEmail(playerEmail).
				SetRoundID(in.RoundId).
				SetFoldNo(in.FoldNo).
				SetLambNum(in.LambNum).
				SetProfitAndLoss(0.0).
				SetRoundCount(roundInfo.RoundCount).
				SetTotalRoundCount(roundInfo.TotalRoundCount).
				Save(l.ctx)
		} else {
			result, err = tx.RoundInvest.UpdateOneID(investedOne.ID).
				AddLambNum(int32(in.LambNum)).Save(l.ctx)
		}
		if err != nil {
			return err
		}
		err = tx.RoundLambFold.Update().
			Where(roundlambfold.RoundID(in.RoundId), roundlambfold.FoldNo(in.FoldNo)).
			AddLambNum(int32(in.LambNum)).Exec(l.ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &wolflamp.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
