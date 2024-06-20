package overview

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/jinzhu/now"
	"github.com/kebin6/wolflamp-rpc/common/enum/statementenum"
	"github.com/kebin6/wolflamp-rpc/common/util"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/player"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/ent/statement"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOverviewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOverviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOverviewLogic {
	return &GetOverviewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOverviewLogic) GetOverview(in *wolflamp.GetOverviewReq) (*wolflamp.GetOverviewResp, error) {

	// 统计今日参与玩家数
	players, err := l.svcCtx.DB.RoundInvest.Query().
		Select(roundinvest.FieldPlayerID).
		Where(func(s *sql.Selector) {
			roundTable := sql.Table(round.Table)
			s.Join(roundTable).On(s.C(roundinvest.FieldRoundID), roundTable.C(round.FieldID))
			s.Where(sql.GTE(roundTable.C(round.FieldStartAt), now.BeginningOfDay()))
		}).Where(roundinvest.PlayerIDLT(util.PlayerMaxId)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	// 去重
	todayParticipates := make(map[uint64]bool)
	for _, v := range players {
		todayParticipates[v.PlayerID] = true
	}
	todayParticipateCount := len(todayParticipates)

	// 今日新增游戏玩家数
	todayNewPlayerCount, err := l.svcCtx.DB.Player.Query().Where(player.CreatedAtGTE(now.BeginningOfDay())).Count(l.ctx)
	if err != nil {
		return nil, err
	}

	// 今日游戏开局次数
	todayRoundCount, err := l.svcCtx.DB.Round.Query().Where(round.StartAtGTE(now.BeginningOfDay())).Count(l.ctx)

	// 今日累积吃羊数
	todayEatCount, err := l.svcCtx.DB.RoundInvest.Query().Where(func(s *sql.Selector) {
		roundTable := sql.Table(round.Table)
		s.Join(roundTable).On(s.C(roundinvest.FieldRoundID), roundTable.C(round.FieldID))
		s.Where(sql.GTE(roundTable.C(round.FieldStartAt), now.BeginningOfDay()))
	}).Where(roundinvest.ProfitAndLossLT(0)).Aggregate(ent.Sum(roundinvest.FieldLambNum)).Int(l.ctx)
	if err != nil {
		return nil, err
	}

	// 今日平台收益
	todayPlatformProfitAmount, err := l.svcCtx.DB.Statement.Query().Where(statement.CreatedAtGTE(now.BeginningOfDay()),
		statement.PlayerID(0), statement.InoutType(statementenum.Income.Val()), statement.Module(statementenum.System.Val())).
		Aggregate(ent.Sum(statement.FieldAmount)).Float64(l.ctx)
	if err != nil {
		return nil, err
	}

	// 平台累积收益（可按月查询明细）
	todayPlatformProfitAmountQuery := l.svcCtx.DB.Statement.Query().
		Where(statement.PlayerID(0), statement.InoutType(statementenum.Income.Val()), statement.Module(statementenum.System.Val()))
	if in.StartTime != nil {
		todayPlatformProfitAmountQuery.Where(statement.CreatedAtGTE(time.Unix(*in.StartTime, 0)))
	}
	if in.EndTime != nil {
		todayPlatformProfitAmountQuery.Where(statement.CreatedAtLTE(time.Unix(*in.EndTime, 0)))
	}
	totalPlatformProfitAmount, err := todayPlatformProfitAmountQuery.Aggregate(ent.Sum(statement.FieldAmount)).Float64(l.ctx)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	// 平台累积人数
	playerCount, err := l.svcCtx.DB.Player.Query().Count(l.ctx)
	if err != nil {
		return nil, err
	}

	return &wolflamp.GetOverviewResp{
		Data: &wolflamp.OverviewInfo{
			TodayParticipateCount: uint64(todayParticipateCount),
			TodayNewPlayerCount:   uint64(todayNewPlayerCount),
			TodayRoundCount:       uint64(todayRoundCount),
			TodayEatCount:         uint64(todayEatCount),
			TodayPlatformProfit:   uint64(todayPlatformProfitAmount),
			TotalPlatformProfit:   uint64(totalPlatformProfitAmount),
			TotalPlayerCount:      uint64(playerCount),
		},
	}, nil

}
