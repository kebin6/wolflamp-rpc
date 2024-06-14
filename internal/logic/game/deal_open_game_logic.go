package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/enum/roundenum"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
	"github.com/kebin6/wolflamp-rpc/internal/utils/entx"
	"github.com/zeromicro/go-zero/core/errorx"

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
	totalLoserInvestNum := uint32(0)
	playerNum := 0
	investResult := make([]PlayerInvestInfo, playerNum)
	// 以下逻辑前置条件是：玩家每回合只能投一个羊圈
	// 统计每个参与玩家的结果
	for _, v := range allInvests {
		playerNum++
		if v.FoldNo != in.LambFoldNo {
			totalWinnerInvestNum += v.LambNum
		} else {
			totalLoserInvestNum += v.LambNum
		}
	}
	for _, v := range allInvests {
		if v.FoldNo == in.LambFoldNo {
			investResult = append(investResult, PlayerInvestInfo{
				PlayerId:      v.PlayerID,
				RoundId:       v.RoundID,
				LambFoldNo:    v.FoldNo,
				LambNum:       v.LambNum,
				ProfitAndLoss: -v.ProfitAndLoss,
			})
		} else {
			proportion := float32(v.LambNum) / float32(totalWinnerInvestNum)
			investResult = append(investResult, PlayerInvestInfo{
				PlayerId:      v.PlayerID,
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

	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		for _, investInfo := range investResult {
			// 更新投注记录
			err := l.svcCtx.DB.RoundInvest.Update().Where(roundinvest.ID(investInfo.InvestId)).
				SetProfitAndLoss(investInfo.ProfitAndLoss).Exec(l.ctx)
			if err != nil {
				return err
			}
			if investInfo.ProfitAndLoss > 0 {
				l.svcCtx.DB.Player.UpdateOneID(investInfo.PlayerId).AddLamp(float32(investInfo.LambNum) + investInfo.ProfitAndLoss)
			}
		}
		for _, foldInfo := range lambFoldResult {
			// 更新羊圈投注记录
			err := l.svcCtx.DB.RoundLambFold.Update().Where(roundlambfold.RoundID(round.Id), roundlambfold.FoldNo(foldInfo.LambFoldNo)).
				SetLambNum(foldInfo.LambNum).SetProfitAndLoss(foldInfo.ProfitAndLoss).Exec(l.ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &wolflamp.BaseIDResp{}, nil

}
