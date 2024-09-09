package overview

import (
	"context"
	"fmt"
	"github.com/jinzhu/now"
	"github.com/kebin6/wolflamp-rpc/common/enum/poolenum"
	"github.com/kebin6/wolflamp-rpc/common/enum/statementenum"
	"github.com/kebin6/wolflamp-rpc/common/util"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/player"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/ent/statement"
	poollogic "github.com/kebin6/wolflamp-rpc/internal/logic/pool"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
	"time"
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
	todayRounds, err := l.svcCtx.DB.Round.Query().
		Where(round.CreatedAtGTE(now.BeginningOfDay())).
		Select(round.FieldID, round.FieldCreatedAt).All(l.ctx)
	if err != nil {
		return nil, err
	}
	todayRoundIds := make([]uint64, len(todayRounds))
	for i, v := range todayRounds {
		todayRoundIds[i] = v.ID
	}
	todayInvests, err := l.svcCtx.DB.RoundInvest.Query().
		//Where(roundinvest.PlayerIDLT(util.PlayerMaxId)).
		Where(roundinvest.RoundIDIn(todayRoundIds...)).
		Select(roundinvest.FieldPlayerID, roundinvest.FieldMode, roundinvest.FieldProfitAndLoss,
			roundinvest.FieldLambNum).
		All(l.ctx)

	// 去重
	coinTodayParticipates := make(map[uint64]bool)
	tokenTodayParticipates := make(map[uint64]bool)
	for _, v := range todayInvests {
		if v.PlayerID >= util.PlayerMaxId {
			continue
		}
		if v.Mode == "coin" {
			coinTodayParticipates[v.PlayerID] = true
		} else {
			tokenTodayParticipates[v.PlayerID] = true
		}
	}
	// coin场今日参与者数
	coinTodayParticipateCount := len(coinTodayParticipates)
	// token场今日参与者数
	tokenTodayParticipateCount := len(tokenTodayParticipates)

	// 今日新增游戏玩家数
	todayNewPlayerCount, err := l.svcCtx.DB.Player.Query().Where(player.CreatedAtGTE(now.BeginningOfDay())).Count(l.ctx)
	if err != nil {
		return nil, err
	}

	// 今日游戏开局次数(coin场)
	coinTodayRoundCount := 0
	// 今日游戏开局次数(token场)
	tokenTodayRoundCount := 0
	for _, v := range todayRounds {
		if v.Mode == "coin" {
			coinTodayRoundCount++
		} else {
			tokenTodayRoundCount++
		}
	}
	// 今日吃羊数羊(coin场)
	coinTodayEatCount := 0
	// 今日吃羊数羊(token场)
	tokenTodayEatCount := 0
	for _, v := range todayInvests {
		if v.ProfitAndLoss < 0 {
			if v.Mode == "coin" {
				coinTodayEatCount += int(v.LambNum)
			} else {
				tokenTodayEatCount += int(v.LambNum)
			}
		}
	}

	// 今日平台收益(coin场)
	coinTodayPlatformProfitAmount, err := l.svcCtx.DB.Statement.Query().Where(statement.CreatedAtGTE(now.BeginningOfDay()),
		statement.PlayerID(0), statement.InoutType(statementenum.Income.Val()), statement.Module(statementenum.System.Val()),
		statement.Mode("coin")).
		Aggregate(ent.Sum(statement.FieldAmount)).Float64(l.ctx)
	if err != nil {
		coinTodayPlatformProfitAmount = 0
	}

	// 今日平台收益(token场)
	tokenTodayPlatformProfitAmount, err := l.svcCtx.DB.Statement.Query().Where(statement.CreatedAtGTE(now.BeginningOfDay()),
		statement.PlayerID(0), statement.InoutType(statementenum.Income.Val()), statement.Module(statementenum.System.Val()),
		statement.Mode("token")).
		Aggregate(ent.Sum(statement.FieldAmount)).Float64(l.ctx)
	if err != nil {
		tokenTodayPlatformProfitAmount = 0
	}

	// 平台累积收益(coin场)
	coinTotalPlatformProfitAmount, err := l.GetTotalPlatformProfitCache("coin")
	if err != nil {
		fmt.Printf("[coin] get platform profit amount error: %v\n", err)
		coinTotalPlatformProfitAmount = 0
	} else {
		coinTotalPlatformProfitAmount += coinTodayPlatformProfitAmount
	}

	// 平台累积收益(token场)
	tokenTotalPlatformProfitAmount, err := l.GetTotalPlatformProfitCache("token")
	if err != nil {
		fmt.Printf("[token] get platform profit amount error: %v\n", err)
		tokenTotalPlatformProfitAmount = 0
	} else {
		tokenTotalPlatformProfitAmount += tokenTodayPlatformProfitAmount
	}

	// 平台累积人数
	playerCount, err := l.svcCtx.DB.Player.Query().Count(l.ctx)
	if err != nil {
		fmt.Printf("get player count error: %v\n", err)
		return nil, err
	}

	// coin机器人池剩余量
	coinRobResult, err := poollogic.NewGetSumLogic(l.ctx, l.svcCtx).GetSum(&wolflamp.GetSumReq{Mode: "coin", Status: 1, Type: poolenum.Robot.Val()})
	coinRobotPoolRestNum := float64(0)
	if err != nil {
		fmt.Printf("get coin robot pool rest num error: %v\n", err)
		coinRobotPoolRestNum = 0
	} else {
		coinRobotPoolRestNum = coinRobResult.GetAmount()
	}

	// token机器人池剩余量
	tokenRobResult, err := poollogic.NewGetSumLogic(l.ctx, l.svcCtx).GetSum(&wolflamp.GetSumReq{Mode: "token", Status: 1, Type: poolenum.Robot.Val()})
	tokenRobotPoolRestNum := float64(0)
	if err != nil {
		fmt.Printf("get token robot pool rest num error: %v\n", err)
		tokenRobotPoolRestNum = 0
	} else {
		tokenRobotPoolRestNum = tokenRobResult.GetAmount()
	}

	// coin奖金池剩余量
	coinRewardResult, err := poollogic.NewGetSumLogic(l.ctx, l.svcCtx).GetSum(&wolflamp.GetSumReq{Mode: "coin", Status: 1, Type: poolenum.Reward.Val()})
	coinRewardPoolRestNum := float64(0)
	if err != nil {
		fmt.Printf("get coin reward pool rest num error: %v\n", err)
		coinRewardPoolRestNum = 0
	} else {
		coinRewardPoolRestNum = coinRewardResult.GetAmount()
	}

	// coin奖金池剩余量
	tokenRewardResult, err := poollogic.NewGetSumLogic(l.ctx, l.svcCtx).GetSum(&wolflamp.GetSumReq{Mode: "token", Status: 1, Type: poolenum.Reward.Val()})
	tokenRewardPoolRestNum := float64(0)
	if err != nil {
		fmt.Printf("get token reward pool rest num error: %v\n", err)
		tokenRewardPoolRestNum = 0
	} else {
		tokenRewardPoolRestNum = tokenRewardResult.GetAmount()
	}

	return &wolflamp.GetOverviewResp{
		Data: &wolflamp.OverviewInfo{
			TodayNewPlayerCount:        uint64(todayNewPlayerCount),
			TotalPlayerCount:           uint64(playerCount),
			CoinTodayParticipateCount:  uint64(coinTodayParticipateCount),
			TokenTodayParticipateCount: uint64(tokenTodayParticipateCount),
			CoinTodayRoundCount:        uint64(coinTodayRoundCount),
			TokenTodayRoundCount:       uint64(tokenTodayRoundCount),
			CoinTodayEatCount:          uint64(coinTodayEatCount),
			TokenTodayEatCount:         uint64(tokenTodayEatCount),
			CoinTodayPlatformProfit:    uint64(coinTodayPlatformProfitAmount),
			TokenTodayPlatformProfit:   uint64(tokenTodayPlatformProfitAmount),
			CoinTotalPlatformProfit:    uint64(coinTotalPlatformProfitAmount),
			CoinRobotPoolRestNum:       coinRobotPoolRestNum,
			TokenRobotPoolRestNum:      tokenRobotPoolRestNum,
			CoinRewardPoolRestNum:      coinRewardPoolRestNum,
			TokenRewardPoolRestNum:     tokenRewardPoolRestNum,
		},
	}, nil

}

// GetTotalPlatformProfitCache 统计平台今日以前累计收益
func (l *GetOverviewLogic) GetTotalPlatformProfitCache(mode string) (float64, error) {
	nowTime := now.BeginningOfDay().In(time.UTC)

	//yesterday := nowTime.AddDate(0, 0, -1)
	format := "2006-01-02"
	formattedToday := nowTime.Format(format)
	todayCacheKey := fmt.Sprintf("%s_platform_profit_sum:%s", mode, formattedToday)

	// 如果有昨天的统计缓存，则直接返回
	if l.svcCtx.Redis.Exists(l.ctx, todayCacheKey).Val() != 0 {
		cached, err := l.svcCtx.Redis.Get(l.ctx, todayCacheKey).Result()
		if err == nil {
			f, err := strconv.ParseFloat(cached, 64)
			if err == nil {
				return f, nil
			}
		}
	}

	// 获取前天的统计缓存，有则统计昨天的数据并进行累加，缓存
	yesterday := nowTime.AddDate(0, 0, -1)
	formattedYesterday := yesterday.Format(format)
	cacheKey := fmt.Sprintf("%s_platform_profit_sum:%s", mode, formattedYesterday)
	if l.svcCtx.Redis.Exists(l.ctx, cacheKey).Val() != 0 {
		cached, err := l.svcCtx.Redis.Get(l.ctx, cacheKey).Result()
		if err == nil {
			f, err := strconv.ParseFloat(cached, 64)
			if err == nil {
				// 截止到前天有缓存，则查询昨天的数据进行累加
				amount, err := l.svcCtx.DB.Statement.Query().
					Where(statement.CreatedAtLT(nowTime)).
					Where(statement.CreatedAtGTE(yesterday)).
					Where(statement.PlayerID(0), statement.InoutType(statementenum.Income.Val()),
						statement.Module(statementenum.System.Val()), statement.Mode(mode)).
					Aggregate(ent.Sum(statement.FieldAmount)).Float64(l.ctx)
				if err == nil {
					amount += f
					_ = l.svcCtx.Redis.Set(l.ctx, todayCacheKey, amount, time.Hour*24*2)
					return amount, err
				}
			}
		}
	}
	amount, err := l.svcCtx.DB.Statement.Query().
		Where(statement.CreatedAtLT(nowTime)).
		Where(statement.PlayerID(0), statement.InoutType(statementenum.Income.Val()),
			statement.Module(statementenum.System.Val()), statement.Mode(mode)).
		Aggregate(ent.Sum(statement.FieldAmount)).Float64(l.ctx)
	if err != nil {
		if !strings.Contains(err.Error(), "converting NULL to float64 is unsupported") {
			return 0, err
		}
		return 0, nil
	}
	// 保存3天
	_ = l.svcCtx.Redis.Set(l.ctx, todayCacheKey, amount, time.Hour*24*2)
	return amount, nil
}
