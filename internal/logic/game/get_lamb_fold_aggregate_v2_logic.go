package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"math/rand"
	"time"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLambFoldAggregateV2Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLambFoldAggregateV2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLambFoldAggregateV2Logic {
	return &GetLambFoldAggregateV2Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLambFoldAggregateV2Logic) GetLambFoldAggregateV2(in *wolflamp.Empty) (*wolflamp.GetLambFoldAggregateResp, error) {

	roundInfo, err := NewFindRoundLogic(l.ctx, l.svcCtx).FindRound(&wolflamp.FindRoundReq{})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	// 羊圈的总盈亏数
	lambFoldAggregates := []*wolflamp.LambFoldAggregateInfo{
		{LambFoldNo: 1, ProfitAndLossCount: 0, WinCount: 0, AvgWinRate: 0, TotalRoundCount: 0},
		{LambFoldNo: 2, ProfitAndLossCount: 0, WinCount: 0, AvgWinRate: 0, TotalRoundCount: 0},
		{LambFoldNo: 3, ProfitAndLossCount: 0, WinCount: 0, AvgWinRate: 0, TotalRoundCount: 0},
		{LambFoldNo: 4, ProfitAndLossCount: 0, WinCount: 0, AvgWinRate: 0, TotalRoundCount: 0},
		{LambFoldNo: 5, ProfitAndLossCount: 0, WinCount: 0, AvgWinRate: 0, TotalRoundCount: 0},
		{LambFoldNo: 6, ProfitAndLossCount: 0, WinCount: 0, AvgWinRate: 0, TotalRoundCount: 0},
		{LambFoldNo: 7, ProfitAndLossCount: 0, WinCount: 0, AvgWinRate: 0, TotalRoundCount: 0},
		{LambFoldNo: 8, ProfitAndLossCount: 0, WinCount: 0, AvgWinRate: 0, TotalRoundCount: 0},
	}

	// 获取当前回合所有玩家列表
	participants, err := l.svcCtx.DB.RoundInvest.Query().Where(roundinvest.RoundID(roundInfo.Id)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	// 记录玩家:羊圈数据，以便后续快速检索
	playerLambFoldMap := make(map[uint64]uint32)
	for _, participant := range participants {
		playerLambFoldMap[participant.PlayerID] = participant.FoldNo
	}

	rand.Seed(time.Now().UnixNano())
	// 随机选择10~30随机数，作为查询值N
	randNum := uint64(rand.Intn(20) + 10)
	for _, participant := range participants {
		page, err := l.svcCtx.DB.RoundInvest.Query().Where(roundinvest.PlayerID(participant.PlayerID), roundinvest.RoundIDNEQ(roundInfo.Id)).
			Order(ent.Desc(roundinvest.FieldID)).Page(l.ctx, 1, randNum)
		if err != nil {
			return nil, err
		}
		for _, invest := range page.List {
			lambFoldNo := playerLambFoldMap[invest.PlayerID]
			lambFoldAggregates[lambFoldNo-1].LambFoldNo = lambFoldNo
			lambFoldAggregates[lambFoldNo-1].ProfitAndLossCount += invest.ProfitAndLoss
			if invest.ProfitAndLoss > 0 {
				lambFoldAggregates[lambFoldNo-1].WinCount += 1
			}
			lambFoldAggregates[lambFoldNo-1].TotalRoundCount += 1
		}
	}
	// 计算胜率
	for i := 0; i < 8; i++ {
		if lambFoldAggregates[i].TotalRoundCount > 0 {
			lambFoldAggregates[i].AvgWinRate = float32(lambFoldAggregates[i].WinCount / lambFoldAggregates[i].TotalRoundCount)
		}
	}
	return &wolflamp.GetLambFoldAggregateResp{
		Data: lambFoldAggregates,
	}, nil

}
