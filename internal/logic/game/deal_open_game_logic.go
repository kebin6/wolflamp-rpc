package game

import (
	"context"
	"fmt"
	"github.com/kebin6/wolflamp-rpc/common/enum/poolenum"
	"github.com/kebin6/wolflamp-rpc/common/enum/roundenum"
	"github.com/kebin6/wolflamp-rpc/common/enum/statementenum"
	"github.com/kebin6/wolflamp-rpc/common/util"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
	"github.com/kebin6/wolflamp-rpc/internal/logic/setting"
	"github.com/kebin6/wolflamp-rpc/internal/logic/statement"
	"github.com/kebin6/wolflamp-rpc/internal/utils/entx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"strconv"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type DealOpenGameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type PlayerInvestInfo struct {
	InvestId      uint64
	PlayerId      uint64
	PlayerEmail   string
	RoundId       uint64
	LambFoldNo    uint32
	LambNum       uint32
	Proportion    float32
	ProfitAndLoss float32
}

type LambFoldInvestInfo struct {
	RoundId       uint64
	LambFoldNo    uint32
	LambNum       uint32
	ProfitAndLoss float32
}

type InvitorRewardInfo struct {
	Module      uint32
	PlayerId    uint64
	PlayerEmail string
	PlayerLevel uint32
	InvitorId   uint64
	Reward      float64
	Remark      string
}

func NewDealOpenGameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DealOpenGameLogic {
	return &DealOpenGameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DealOpenGameLogic) DealOpenGame(in *wolflamp.DealOpenGameReq) (*wolflamp.BaseIDResp, error) {

	round, err := NewFindRoundLogic(l.ctx, l.svcCtx).FindRound(&wolflamp.FindRoundReq{Mode: in.Mode})
	if err != nil {
		return nil, err
	}
	if round.Status != roundenum.Investing.Val() {
		return nil, errorx.NewAlreadyExistsError("game was already open")
	}
	// 被攻击的小羊按照其他羊圈用户投放的小羊数量占比进行分配
	allInvests, err := l.svcCtx.DB.RoundInvest.Query().Where(roundinvest.Mode(in.Mode)).
		Where(roundinvest.RoundID(round.Id)).All(l.ctx)
	if err != nil {
		return nil, err
	}

	// 按照羊圈+用户纬度统计投注数量
	//lambFoldPlayerSorts := make(map[uint32]map[uint64]uint32)
	// 统计盈利方总投注数量
	totalWinnerInvestNum := uint32(0)
	totalLoserInvestNum := float64(0)
	playerNum := 0
	investResult := make([]PlayerInvestInfo, playerNum)
	// 以下逻辑前置条件是：玩家每回合只能投一个羊圈
	// 统计每个参与玩家的结果
	for _, v := range allInvests {
		playerNum++
		if v.FoldNo != in.LambFoldNo {
			totalWinnerInvestNum += v.LambNum
		} else {
			totalLoserInvestNum += float64(v.LambNum)
		}
	}

	// 平台佣金比例
	commissionResp, err := setting.NewGetGameCommissionLogic(l.ctx, l.svcCtx).GetGameCommission(&wolflamp.Empty{})
	if err != nil {
		return nil, err
	}
	systemCommission := commissionResp.GameCommission
	systemCommissionNum := totalLoserInvestNum * float64(systemCommission) / 100

	// 资金池总预留比例
	poolCommissionResp, err := setting.NewGetPoolCommissionLogic(l.ctx, l.svcCtx).GetPoolCommission(&wolflamp.Empty{})
	if err != nil {
		return nil, err
	}
	poolCommission := poolCommissionResp.Commission
	poolCommissionNum := totalLoserInvestNum * float64(poolCommission) / 100
	// 机器人池预留占比
	robPoolCommissionResp, err := setting.NewGetRobPoolCommissionLogic(l.ctx, l.svcCtx).GetRobPoolCommission(&wolflamp.Empty{})
	if err != nil {
		return nil, err
	}
	// 奖金池预留占比
	rewardPoolCommissionResp, err := setting.NewGetRewardPoolCommissionLogic(l.ctx, l.svcCtx).GetRewardPoolCommission(&wolflamp.Empty{})
	if err != nil {
		return nil, err
	}
	// 机器人池预留数
	robPoolCommissionNum := poolCommissionNum * float64(robPoolCommissionResp.Commission) / 100
	// 奖金池预留数
	rewardPoolCommissionNum := poolCommissionNum * float64(rewardPoolCommissionResp.Commission) / 100
	// 其他剩余池预留数
	restPoolCommissionNum := poolCommissionNum - robPoolCommissionNum - rewardPoolCommissionNum

	// 传给第三方的被杀总数S=被杀总羊数-平台佣金-资金池预留-N
	// 传给第三方的被杀总数S*0.09*A/(A+B+C)=A上级总共所得分佣
	// 传给第三方的被杀总数S*0.09*B/(A+B+C)=B上级总共所得分佣
	// 传给第三方的被杀总数S*0.09*C/(A+B+C)=C上级总共所得分佣
	// 传给第三方的被杀总数S = 真实用户上级总分佣 / 0.09

	// 要分给用户上级的预留数
	invitorCommission := 0.09
	invitorCommissionNum := totalLoserInvestNum * invitorCommission

	// 剩余部分是赢的玩家按照投注占比进行瓜分的
	totalLoserInvestNum = totalLoserInvestNum - systemCommissionNum - poolCommissionNum - invitorCommissionNum

	for _, v := range allInvests {
		if v.FoldNo == in.LambFoldNo {
			// 记录输家亏掉的羊数量
			investResult = append(investResult, PlayerInvestInfo{
				InvestId:      v.ID,
				PlayerId:      v.PlayerID,
				PlayerEmail:   v.PlayerEmail,
				RoundId:       v.RoundID,
				LambFoldNo:    v.FoldNo,
				LambNum:       v.LambNum,
				ProfitAndLoss: -float32(v.LambNum),
			})
		} else {
			// 记录赢家赢到的羊数量
			proportion := float32(v.LambNum) / float32(totalWinnerInvestNum)
			investResult = append(investResult, PlayerInvestInfo{
				InvestId:      v.ID,
				PlayerId:      v.PlayerID,
				PlayerEmail:   v.PlayerEmail,
				RoundId:       v.RoundID,
				LambFoldNo:    v.FoldNo,
				LambNum:       v.LambNum,
				Proportion:    proportion,
				ProfitAndLoss: float32(totalLoserInvestNum) * proportion,
			})
		}
	}

	// 统计羊圈的结果
	lambFoldResult := []LambFoldInvestInfo{
		{RoundId: round.Id, LambFoldNo: 1, LambNum: 0, ProfitAndLoss: 0},
		{RoundId: round.Id, LambFoldNo: 2, LambNum: 0, ProfitAndLoss: 0},
		{RoundId: round.Id, LambFoldNo: 3, LambNum: 0, ProfitAndLoss: 0},
		{RoundId: round.Id, LambFoldNo: 4, LambNum: 0, ProfitAndLoss: 0},
		{RoundId: round.Id, LambFoldNo: 5, LambNum: 0, ProfitAndLoss: 0},
		{RoundId: round.Id, LambFoldNo: 6, LambNum: 0, ProfitAndLoss: 0},
		{RoundId: round.Id, LambFoldNo: 7, LambNum: 0, ProfitAndLoss: 0},
		{RoundId: round.Id, LambFoldNo: 8, LambNum: 0, ProfitAndLoss: 0},
	}
	// 统计回合结果
	for _, v := range investResult {
		lambFoldResult[v.LambFoldNo-1].LambNum += v.LambNum
		lambFoldResult[v.LambFoldNo-1].ProfitAndLoss += v.ProfitAndLoss
	}

	statementCreateLogic := statement.NewCreateStatementLogic(l.ctx, l.svcCtx)
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		// 记录资金池记录
		if robPoolCommissionNum > 0 {
			err := l.svcCtx.DB.Pool.Create().SetMode(in.Mode).
				SetStatus(1).SetRoundID(round.Id).SetType(poolenum.Robot.Val()).
				SetLambNum(robPoolCommissionNum).SetRemark("机器人池预留").Exec(l.ctx)
			if err != nil {
				return err
			}
		}
		if rewardPoolCommissionNum > 0 {
			err := l.svcCtx.DB.Pool.Create().SetMode(in.Mode).
				SetStatus(1).SetRoundID(round.Id).SetType(poolenum.Reward.Val()).
				SetLambNum(rewardPoolCommissionNum).Exec(l.ctx)
			if err != nil {
				return err
			}
		}
		if restPoolCommissionNum > 0 {
			err := l.svcCtx.DB.Pool.Create().SetMode(in.Mode).
				SetStatus(1).SetRoundID(round.Id).SetType(poolenum.Other.Val()).
				SetLambNum(restPoolCommissionNum).Exec(l.ctx)
			if err != nil {
				return err
			}
		}
		// 记录资金池预留收益账单
		_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
			Mode:     in.Mode,
			PlayerId: 0, Status: statementenum.Completed.Val(), Module: statementenum.Pool.Val(),
			Amount: poolCommissionNum, InoutType: statementenum.Income.Val(),
			ReferId: strconv.FormatUint(round.Id, 10), Prefix: pointy.GetPointer("IV"),
			Remark: pointy.GetPointer(fmt.Sprintf("%f*%f=%f", float32(totalLoserInvestNum), poolCommission/100, poolCommissionNum)),
		})
		if err != nil {
			return err
		}

		// 记录平台收益账单
		_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
			Mode:     in.Mode,
			PlayerId: 0, Status: statementenum.Completed.Val(), Module: statementenum.System.Val(),
			Amount: systemCommissionNum, InoutType: statementenum.Income.Val(),
			ReferId: strconv.FormatUint(round.Id, 10), Prefix: pointy.GetPointer("IV"),
			Remark: pointy.GetPointer(fmt.Sprintf("%f*%f=%f", float32(totalLoserInvestNum), systemCommission/100, systemCommissionNum)),
		})
		if err != nil {
			return err
		}

		// 记录预留上级佣金账单
		_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
			Mode:     in.Mode,
			PlayerId: 0, Status: statementenum.Completed.Val(), Module: statementenum.Reward.Val(),
			Amount: invitorCommissionNum, InoutType: statementenum.Income.Val(),
			ReferId: strconv.FormatUint(round.Id, 10), Prefix: pointy.GetPointer("IV"),
			Remark: pointy.GetPointer(fmt.Sprintf("%f*%f=%f", float32(totalLoserInvestNum), invitorCommission, invitorCommissionNum)),
		})
		if err != nil {
			return err
		}

		// 回传第三方的总数量=真实用户上级总分佣
		computedAmount := invitorCommissionNum
		// 更新投注记录的盈亏数据
		for _, investInfo := range investResult {
			err := l.svcCtx.DB.RoundInvest.Update().Where(roundinvest.ID(investInfo.InvestId)).
				SetProfitAndLoss(investInfo.ProfitAndLoss).Exec(l.ctx)
			if err != nil {
				return err
			}
			if util.IsRealPlayer(investInfo.PlayerId) {
				// 真实用户将收益及投注本金加回账户
				if investInfo.ProfitAndLoss >= 0 {
					if in.Mode == "coin" {
						err = l.svcCtx.DB.Player.UpdateOneID(investInfo.PlayerId).AddCoinLamb(float32(investInfo.LambNum) + investInfo.ProfitAndLoss).Exec(l.ctx)
					} else {
						err = l.svcCtx.DB.Player.UpdateOneID(investInfo.PlayerId).AddTokenLamb(float32(investInfo.LambNum) + investInfo.ProfitAndLoss).Exec(l.ctx)
					}

					if err != nil {
						return err
					}
					// 赢亏为正的，则记录流水账单
					if investInfo.ProfitAndLoss > 0 {
						_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
							Mode:     in.Mode,
							PlayerId: investInfo.PlayerId, Status: statementenum.Completed.Val(), Module: statementenum.Invest.Val(),
							Amount: float64(investInfo.ProfitAndLoss), InoutType: statementenum.Income.Val(),
							ReferId: strconv.FormatUint(investInfo.InvestId, 10), Prefix: pointy.GetPointer("IV"),
							Remark: pointy.GetPointer(fmt.Sprintf("%f*%f=%f", float32(totalLoserInvestNum), investInfo.Proportion, investInfo.ProfitAndLoss)),
						})
					}
				} else if investInfo.ProfitAndLoss < 0 {
					_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
						Mode:     in.Mode,
						PlayerId: investInfo.PlayerId, Status: statementenum.Completed.Val(), Module: statementenum.Invest.Val(),
						Amount: float64(investInfo.ProfitAndLoss), InoutType: statementenum.Expense.Val(),
						ReferId: strconv.FormatUint(investInfo.InvestId, 10), Prefix: pointy.GetPointer("IV"),
					})
				}
				if err != nil {
					return err
				}
			} else {
				if investInfo.ProfitAndLoss > 0 {
					playerInvitorCommission := invitorCommissionNum * float64(investInfo.Proportion)
					computedAmount -= playerInvitorCommission
					// 机器人赢的钱回归到机器人池
					err := l.svcCtx.DB.Pool.Create().SetMode(in.Mode).
						SetStatus(1).SetRoundID(round.Id).SetType(poolenum.Robot.Val()).
						SetLambNum(float64(investInfo.ProfitAndLoss)).SetRemark("机器人获胜").Exec(l.ctx)
					if err != nil {
						return err
					}

					_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
						Mode:     in.Mode,
						PlayerId: investInfo.PlayerId, Status: statementenum.Completed.Val(), Module: statementenum.Invest.Val(),
						Amount: float64(investInfo.ProfitAndLoss) + playerInvitorCommission, InoutType: statementenum.Income.Val(),
						ReferId: strconv.FormatUint(investInfo.InvestId, 10), Prefix: pointy.GetPointer("IV"),
						Remark: pointy.GetPointer(fmt.Sprintf("机器人获胜：总待分羊数量（%f）*玩家投注占比（%f）+玩家上级分佣金（%f）=%f",
							float32(totalLoserInvestNum), investInfo.Proportion, playerInvitorCommission, float64(investInfo.ProfitAndLoss)+playerInvitorCommission)),
					})
				} else if investInfo.ProfitAndLoss < 0 {
					_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
						Mode:     in.Mode,
						PlayerId: investInfo.PlayerId, Status: statementenum.Completed.Val(), Module: statementenum.Invest.Val(),
						Amount: float64(investInfo.ProfitAndLoss), InoutType: statementenum.Expense.Val(),
						ReferId: strconv.FormatUint(investInfo.InvestId, 10), Prefix: pointy.GetPointer("IV"),
					})
				}
			}
		}

		// 更新羊圈的盈亏数据
		for _, foldInfo := range lambFoldResult {
			// 更新羊圈投注记录
			err := l.svcCtx.DB.RoundLambFold.Update().Where(roundlambfold.RoundID(round.Id), roundlambfold.FoldNo(foldInfo.LambFoldNo)).
				SetLambNum(foldInfo.LambNum).SetProfitAndLoss(foldInfo.ProfitAndLoss).Exec(l.ctx)
			if err != nil {
				return err
			}
		}
		// 更新回合状态
		err := l.svcCtx.DB.Round.UpdateOneID(round.Id).
			SetStatus(uint8(roundenum.Opening.Val())).
			SetComputeAmount(computedAmount).
			SetSyncStatus(uint32(roundenum.NotYet)).
			SetSelectedFold(in.LambFoldNo).Exec(l.ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &wolflamp.BaseIDResp{}, nil

}
