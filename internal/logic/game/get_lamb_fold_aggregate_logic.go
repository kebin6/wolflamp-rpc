package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLambFoldAggregateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLambFoldAggregateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLambFoldAggregateLogic {
	return &GetLambFoldAggregateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetLambFoldAggregate 计算给定回溯场数，每个羊圈的盈亏数&胜率
func (l *GetLambFoldAggregateLogic) GetLambFoldAggregate(in *wolflamp.GetLambFoldAggregateReq) (*wolflamp.GetLambFoldAggregateResp, error) {

	allInvests, err := l.svcCtx.DB.RoundInvest.Query().
		Where(roundinvest.TotalRoundCountGTE(in.TotalRoundCountMin), roundinvest.TotalRoundCountLTE(in.TotalRoundCountMax)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	// 统计每个玩家的盈亏数&胜率
	// 玩家盈亏数=玩家每场盈亏累加
	lambFoldPlayerSorts := make(map[uint32]map[uint64]map[uint64]*ent.RoundInvest)
	for _, v := range allInvests {
		if _, ok := lambFoldPlayerSorts[v.FoldNo]; !ok {
			lambFoldPlayerSorts[v.FoldNo] = make(map[uint64]map[uint64]*ent.RoundInvest)
		}
		if _, ok := lambFoldPlayerSorts[v.FoldNo][v.PlayerID]; !ok {
			lambFoldPlayerSorts[v.FoldNo][v.PlayerID] = make(map[uint64]*ent.RoundInvest)
		}
		if _, ok := lambFoldPlayerSorts[v.FoldNo][v.PlayerID][v.RoundID]; !ok {
			lambFoldPlayerSorts[v.FoldNo][v.PlayerID][v.RoundID] = v
		} else {
			lambFoldPlayerSorts[v.FoldNo][v.PlayerID][v.RoundID].ProfitAndLoss += v.ProfitAndLoss
		}
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
	// 遍历每个羊圈玩家情况
	for _, playerSorts := range lambFoldPlayerSorts {
		// 遍历每个玩家投注情况
		for _, roundSorts := range playerSorts {
			for _, v := range roundSorts {
				if v.ProfitAndLoss > 0 {
					lambFoldAggregates[v.FoldNo-1].ProfitAndLossCount += 1
				} else {
					lambFoldAggregates[v.FoldNo-1].ProfitAndLossCount -= 1
				}
				lambFoldAggregates[v.FoldNo-1].TotalRoundCount++
				if v.ProfitAndLoss > 0 {
					lambFoldAggregates[v.FoldNo-1].WinCount++
				}
			}
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
