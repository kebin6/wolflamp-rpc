package game

import (
	"context"
	"fmt"
	"github.com/kebin6/wolflamp-rpc/common/enum/rewardenum"
	"github.com/kebin6/wolflamp-rpc/common/enum/roundenum"
	"github.com/kebin6/wolflamp-rpc/common/enum/statementenum"
	"github.com/kebin6/wolflamp-rpc/common/util"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
	"github.com/kebin6/wolflamp-rpc/internal/logic/player"
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

	round, err := NewFindRoundLogic(l.ctx, l.svcCtx).FindRound(&wolflamp.FindRoundReq{})
	if err != nil {
		return nil, err
	}
	if round.Status != roundenum.Investing.Val() {
		return nil, errorx.NewAlreadyExistsError("game was already open")
	}
	// 被攻击的小羊按照其他羊圈用户投放的小羊数量占比进行分配
	allInvests, err := l.svcCtx.DB.RoundInvest.Query().Where(roundinvest.RoundID(round.Id)).All(l.ctx)
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

	// 平台预留比例
	commissionResp, err := setting.NewGetGameCommissionLogic(l.ctx, l.svcCtx).GetGameCommission(&wolflamp.Empty{})
	if err != nil {
		return nil, err
	}
	systemCommission := commissionResp.GameCommission
	ReserveNum := totalLoserInvestNum * float64(systemCommission) / 100
	totalLoserInvestNum = totalLoserInvestNum - ReserveNum

	for _, v := range allInvests {
		if v.FoldNo == in.LambFoldNo {
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
	// 统计上级奖励及平台收益
	rewards, err := l.GetRewardList(round.Id, totalLoserInvestNum, investResult)
	if err != nil {
		return nil, err
	}

	statementCreateLogic := statement.NewCreateStatementLogic(l.ctx, l.svcCtx)
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
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
					err := l.svcCtx.DB.Player.UpdateOneID(investInfo.PlayerId).AddCoinLamb(float32(investInfo.LambNum) + investInfo.ProfitAndLoss).Exec(l.ctx)
					if err != nil {
						return err
					}
					// 赢亏为正的，则记录流水账单
					if investInfo.ProfitAndLoss > 0 {
						_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
							PlayerId: investInfo.PlayerId, Status: statementenum.Completed.Val(), Module: statementenum.Invest.Val(),
							Amount: float64(investInfo.ProfitAndLoss), InoutType: statementenum.Income.Val(),
							ReferId: strconv.FormatUint(investInfo.InvestId, 10), Prefix: pointy.GetPointer("IV"),
							Remark: pointy.GetPointer(fmt.Sprintf("%f*%f=%f", float32(totalLoserInvestNum), investInfo.Proportion, investInfo.ProfitAndLoss)),
						})
					}
				} else if investInfo.ProfitAndLoss < 0 {
					_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
						PlayerId: investInfo.PlayerId, Status: statementenum.Completed.Val(), Module: statementenum.Invest.Val(),
						Amount: float64(investInfo.ProfitAndLoss), InoutType: statementenum.Expense.Val(),
						ReferId: strconv.FormatUint(investInfo.InvestId, 10), Prefix: pointy.GetPointer("IV"),
					})
				}
				if err != nil {
					return err
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
		err := l.svcCtx.DB.Round.UpdateOneID(round.Id).SetStatus(uint8(roundenum.Opening.Val())).
			SetSelectedFold(in.LambFoldNo).Exec(l.ctx)
		if err != nil {
			return err
		}

		// 处理奖励
		err = l.dealReward(rewards)
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

// 保存奖励
func (l *DealOpenGameLogic) dealReward(rewards []InvitorRewardInfo) error {

	statementCreateLogic := statement.NewCreateStatementLogic(l.ctx, l.svcCtx)
	for _, reward := range rewards {
		// 记录奖励，平台收益不记录奖励列表
		if reward.InvitorId != 0 {
			err := l.svcCtx.DB.Reward.Create().
				SetStatus(uint8(rewardenum.Completed.Val())).
				SetNum(float32(reward.Reward)).
				SetToID(reward.InvitorId).
				SetContributorID(reward.PlayerId).
				SetContributorEmail(reward.PlayerEmail).
				SetContributorLevel(reward.PlayerLevel).
				SetFormula(reward.Remark).Exec(l.ctx)
			if err != nil {
				return err
			}
		}

		// 记录账单
		prefix := "RW"
		if reward.Module == statementenum.System.Val() {
			prefix = "SY"
		}
		_, err := statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
			PlayerId: 0, Status: statementenum.Completed.Val(), Module: reward.Module,
			Amount: reward.Reward, InoutType: statementenum.Income.Val(),
			ReferId: strconv.FormatUint(reward.PlayerId, 10), Prefix: pointy.GetPointer(prefix),
		})
		if err != nil {
			return err
		}
	}
	return nil

}

// GetRewardList 处理玩家奖励
func (l *DealOpenGameLogic) GetRewardList(roundId uint64, amount float64, investInfos []PlayerInvestInfo) ([]InvitorRewardInfo, error) {

	// 记录上级奖励
	var invitorRewards []InvitorRewardInfo
	// 记录扣减掉上级奖励后平台的收益
	platformAmount := amount
	// 获取玩家ID以便查询上级列表
	var playerIds []uint64
	var winners []PlayerInvestInfo
	for _, v := range investInfos {
		if util.IsRealPlayer(v.PlayerId) {
			playerIds = append(playerIds, v.PlayerId)
			winners = append(winners, v)
		}
	}
	// 没有真实用户则全部算入平台收益
	if len(playerIds) == 0 {
		invitorRewards = append(invitorRewards, InvitorRewardInfo{
			Module: statementenum.System.Val(),
			Reward: platformAmount,
		})
		return invitorRewards, nil
	}

	// 计算每个胜利的玩家上级的奖励
	result, err := player.NewGetInvitorListByIdsLogic(l.ctx, l.svcCtx).GetInvitorListByIds(&wolflamp.GetInvitorListByIdsReq{PlayerIds: playerIds})
	if err != nil {
		return nil, err
	}
	// 玩家没有上级则全部算入平台收益
	if len(result.Data) == 0 {
		invitorRewards = append(invitorRewards, InvitorRewardInfo{
			Module: statementenum.System.Val(),
			Reward: platformAmount,
		})
		return invitorRewards, nil
	}

	invitors := result.Data
	invitorGroup := make(map[uint64][]*wolflamp.InvitorInfo)
	// 对上级按照获胜玩家ID进行分组
	for _, v := range invitors {
		if _, ok := invitorGroup[v.PlayerId]; ok {
			invitorGroup[v.PlayerId] = append(invitorGroup[v.PlayerId], v)
		} else {
			invitorGroup[v.PlayerId] = []*wolflamp.InvitorInfo{v}
		}
	}
	// 遍历每个获胜玩家ID，计算上级奖励
	for _, winner := range winners {
		// 计算上级奖励比例的基数
		base := amount * float64(winner.Proportion)

		for _, invitor := range invitorGroup[winner.PlayerId] {
			remark := ""
			// 根据上级等级分配奖励
			if invitor.Rank > 0 {
				reward := float64(0)
				switch invitor.Rank {
				case 1:
					reward = base * 0.01
					remark = fmt.Sprintf("一级奖励%f*%f=%f", base, 0.01, reward)
				case 2:
					reward = base * 0.02
					remark = fmt.Sprintf("二级奖励%f*%f=%f", base, 0.01, reward)
				case 3:
					reward = base * 0.03
					remark = fmt.Sprintf("三级奖励%f*%f=%f", base, 0.01, reward)
				}
				platformAmount -= reward
				invitorRewards = append(invitorRewards, InvitorRewardInfo{
					Module:      statementenum.Reward.Val(),
					PlayerId:    invitor.PlayerId,
					PlayerEmail: winner.PlayerEmail,
					PlayerLevel: invitor.PlayerLevel,
					InvitorId:   invitor.InvitorId,
					Reward:      reward,
					Remark:      remark,
				})
			}

			// 记录平台奖励
			if invitor.SystemCommission > 0 {
				systemCommission := float64(invitor.SystemCommission) / 100
				reward := base * systemCommission
				remark = fmt.Sprintf("平台奖励%f*%f=%f", base, systemCommission, reward)
				platformAmount -= reward
				invitorRewards = append(invitorRewards, InvitorRewardInfo{
					Module:      statementenum.System.Val(),
					PlayerId:    invitor.PlayerId,
					PlayerEmail: winner.PlayerEmail,
					PlayerLevel: invitor.PlayerLevel,
					InvitorId:   invitor.InvitorId,
					Reward:      reward,
					Remark:      remark,
				})
			}
		}
	}
	// 记录平台收益
	if platformAmount > 0 {
		invitorRewards = append(invitorRewards, InvitorRewardInfo{
			Module: statementenum.System.Val(),
			Reward: platformAmount,
		})
	}
	return invitorRewards, nil

}
