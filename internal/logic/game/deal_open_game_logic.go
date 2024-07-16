package game

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/now"
	"github.com/kebin6/wolflamp-rpc/common/enum/cachekey"
	"github.com/kebin6/wolflamp-rpc/common/enum/poolenum"
	"github.com/kebin6/wolflamp-rpc/common/enum/roundenum"
	"github.com/kebin6/wolflamp-rpc/common/enum/statementenum"
	"github.com/kebin6/wolflamp-rpc/common/util"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
	"github.com/kebin6/wolflamp-rpc/internal/logic/pool"
	"github.com/kebin6/wolflamp-rpc/internal/logic/setting"
	"github.com/kebin6/wolflamp-rpc/internal/logic/statement"
	"github.com/kebin6/wolflamp-rpc/internal/utils/entx"
	"github.com/redis/go-redis/v9"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"math/rand"
	"strconv"
	"time"

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

// GetMaxInvestFoldNo 统计得出投注数最多的羊圈
func (l *DealOpenGameLogic) GetMaxInvestFoldNo(invests []*ent.RoundInvest) uint32 {
	// 统计羊圈的结果
	lambFoldResult := []LambFoldInvestInfo{
		{LambFoldNo: 1, LambNum: 0},
		{LambFoldNo: 2, LambNum: 0},
		{LambFoldNo: 3, LambNum: 0},
		{LambFoldNo: 4, LambNum: 0},
		{LambFoldNo: 5, LambNum: 0},
		{LambFoldNo: 6, LambNum: 0},
		{LambFoldNo: 7, LambNum: 0},
		{LambFoldNo: 8, LambNum: 0},
	}
	// 统计回合结果
	for _, v := range invests {
		lambFoldResult[v.FoldNo-1].LambNum += v.LambNum
	}
	curMaxFoldNo := uint32(0)
	for _, v := range lambFoldResult {
		if curMaxFoldNo == 0 {
			curMaxFoldNo = v.LambFoldNo - 1
			continue
		}
		if lambFoldResult[curMaxFoldNo-1].LambNum < v.LambNum {
			curMaxFoldNo = v.LambFoldNo
		}
	}
	return curMaxFoldNo
}

func (l *DealOpenGameLogic) GetGoldenNum(mode string) (goldenNum *uint32, err error) {

	// 判断时间段
	allowTimeRange, err := setting.NewGetGoldenLambAllowTimeLogic(l.ctx, l.svcCtx).GetGoldenLambAllowTime(&wolflamp.Empty{})
	if err != nil {
		fmt.Printf("ProcessOpen[%s]: check golden allow time, error: %s, exit\n", mode, err.Error())
		return nil, err
	}

	layout := "15:04:05"
	nowTime, _ := time.Parse(layout, fmt.Sprintf("%d:%d:%d", time.Now().Hour(), time.Now().Minute(), time.Now().Second()))
	startTime, err := time.Parse(layout, allowTimeRange.StartTime)
	if err != nil {
		fmt.Printf("ProcessOpen[%s]: check golden allow start time, error: %s, exit\n", mode, err.Error())
		return nil, err
	}
	endTime, err := time.Parse(layout, allowTimeRange.EndTime)
	if err != nil {
		fmt.Printf("ProcessOpen[%s]: check golden allow end time, error: %s, exit\n", mode, err.Error())
		return nil, err
	}
	if nowTime.Before(startTime) || nowTime.After(endTime) {
		fmt.Printf("ProcessOpen[%s]: check golden allow time, no match, exit\n", mode)
		return nil, err
	}

	numRange, err := setting.NewGetGoldenLambNumRangeLogic(l.ctx, l.svcCtx).GetGoldenLambNumRange(&wolflamp.Empty{})
	if err != nil {
		fmt.Printf("ProcessOpen[%s]: check golden num range error: %s, exit\n", mode, err.Error())
		return nil, err
	}

	if numRange.Min < numRange.Max || numRange.Min <= 0 || numRange.Max <= 0 {
		fmt.Printf("ProcessOpen[%s]: check golden num, invalid, exit\n", mode)
		return nil, err
	}

	// 获取当前奖励池剩余量
	sumResp, err := pool.NewGetSumLogic(l.ctx, l.svcCtx).GetSum(&wolflamp.GetSumReq{
		Mode:   mode,
		Type:   poolenum.Reward.Val(),
		Status: 1,
	})
	if err != nil {
		fmt.Printf("ProcessOpen[%s]: check golden, reward pool error: %s, exit\n", mode, err.Error())
		return nil, err
	}
	if sumResp.Amount < float64(numRange.Min) {
		fmt.Printf("ProcessOpen[%s]: check golden, reward pool not enough, exit\n", mode)
		return nil, nil
	}
	if sumResp.Amount < float64(numRange.Max) {
		numRange.Max = uint32(sumResp.Amount)
	}
	rand.Seed(time.Now().UnixNano())
	goldenNum = pointy.GetPointer(uint32(rand.Intn(int(numRange.Max-numRange.Min+1)) + int(numRange.Min)))
	return goldenNum, nil

}

func (l *DealOpenGameLogic) DealGoldenCase(mode string, invests []*ent.RoundInvest) (choiceFoldNo uint32, totalRewardNum float64, investResult []PlayerInvestInfo, err error) {

	goldenLock := cachekey.TodayGoldenLambLock.ModeVal(mode)
	occurTime, err := l.svcCtx.Redis.Get(l.ctx, goldenLock).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		fmt.Printf("ProcessOpen[%s]: checking golden, golden lock error : %s, exit\n", mode, err.Error())
		return 0, 0, nil, err
	}
	if occurTime != "" {
		fmt.Printf("ProcessOpen[%s]: today golden lamb has occur, exit\n", mode)
		return 0, 0, nil, err
	}

	num, err := l.GetGoldenNum(mode)
	if err != nil || num == nil {
		return 0, 0, nil, err
	}
	// 符合触发条件，计算出中奖羊圈
	maxInvestFoldNo := l.GetMaxInvestFoldNo(invests)
	totalWinnerInvestNum := uint32(0)
	playerNum := 0
	// 以下逻辑前置条件是：玩家每回合只能投一个羊圈
	// 统计每个参与玩家的结果
	for _, v := range invests {
		playerNum++
		if v.FoldNo == maxInvestFoldNo {
			totalWinnerInvestNum += v.LambNum
		}
	}

	// 计算每个投注的盈亏情况
	for _, v := range invests {
		if v.FoldNo != maxInvestFoldNo {
			// 没有中奖的玩家
			investResult = append(investResult, PlayerInvestInfo{
				InvestId:      v.ID,
				PlayerId:      v.PlayerID,
				PlayerEmail:   v.PlayerEmail,
				RoundId:       v.RoundID,
				LambFoldNo:    v.FoldNo,
				LambNum:       v.LambNum,
				ProfitAndLoss: float32(0),
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
				ProfitAndLoss: float32(*num) * proportion,
			})
		}
	}

	expiredTime := time.Duration(now.EndOfDay().Unix()-time.Now().Unix()) * time.Second
	_, err = l.svcCtx.Redis.SetNX(l.ctx, goldenLock,
		time.Now().Unix(), expiredTime).Result()
	if err != nil {
		fmt.Printf("ProcessOpen[%s]: today golden lamb has occur, exit\n", mode)
		return 0, 0, nil, err
	}

	fmt.Printf("ProcessOpen[%s]: Golden Lamb Choose LambFold: %d\n", mode, maxInvestFoldNo)
	return maxInvestFoldNo, float64(*num), investResult, nil
}

func (l *DealOpenGameLogic) GetSliverNum(mode string) (sliverNum *uint32, err error) {
	minResp, err := setting.NewGetPoolMinNumThenSilverLogic(l.ctx, l.svcCtx).GetPoolMinNumThenSilver(&wolflamp.Empty{})
	if err != nil {
		fmt.Printf("ProcessOpen[%s]: checking sliver, no min num pool setting, exit\n", mode)
		return nil, err
	}
	// 获取当前奖励池剩余量
	sumResp, err := pool.NewGetSumLogic(l.ctx, l.svcCtx).GetSum(&wolflamp.GetSumReq{
		Mode:   mode,
		Type:   poolenum.Reward.Val(),
		Status: 1,
	})
	if err != nil {
		fmt.Printf("ProcessOpen[%s]: checking sliver, reward sum no setting, exit\n", mode)
		return nil, err
	}
	if sumResp.Amount < float64(minResp.Min) {
		fmt.Printf("ProcessOpen[%s]: checking sliver, reward sum is not enough, exit\n", mode)
		return nil, nil
	}
	percentResp, err := setting.NewGetSliverOccurPercentLogic(l.ctx, l.svcCtx).GetSliverOccurPercent(&wolflamp.Empty{})
	if err != nil {
		fmt.Printf("ProcessOpen[%s]: checking sliver, no percent setting, exit\n", mode)
		return nil, err
	}
	if percentResp.Percent <= 0 {
		fmt.Printf("ProcessOpen[%s]: checking sliver, occur percent is zero, exit\n", mode)
		return nil, nil
	}
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(100)+1 > int(percentResp.Percent) {
		fmt.Printf("ProcessOpen[%s]: checking sliver, no matching occur percent, exit\n", mode)
		return nil, nil
	}

	numRange, err := setting.NewGetSliverLambNumRangeLogic(l.ctx, l.svcCtx).GetSliverLambNumRange(&wolflamp.Empty{})
	if err != nil {
		fmt.Printf("ProcessOpen[%s]: checking sliver, no sliver num range setting, exit\n", mode)
		return nil, err
	}

	if numRange.Min < numRange.Max || numRange.Min <= 0 || numRange.Max <= 0 {
		fmt.Printf("ProcessOpen[%s]: checking sliver, num range setting is invalid, exit\n", mode)
		return nil, err
	}

	if sumResp.Amount < float64(numRange.Min) {
		fmt.Printf("ProcessOpen[%s]: checking sliver, sliver num is not enough, exit\n", mode)
		return nil, nil
	}
	if sumResp.Amount < float64(numRange.Max) {
		numRange.Max = uint32(sumResp.Amount)
	}
	rand.Seed(time.Now().UnixNano())
	sliverNum = pointy.GetPointer(uint32(rand.Intn(int(numRange.Max-numRange.Min+1)) + int(numRange.Min)))
	return sliverNum, nil
}

func (l *DealOpenGameLogic) DealSliverCase(mode string, invests []*ent.RoundInvest) (choiceFoldNo uint32, totalRewardNum float64, investResult []PlayerInvestInfo, err error) {
	sliverNum, err := l.GetSliverNum(mode)
	if err != nil || sliverNum == nil {
		return 0, 0, nil, err
	}

	// 符合触发条件，计算出中奖羊圈
	maxInvestFoldNo := l.GetMaxInvestFoldNo(invests)
	totalWinnerInvestNum := uint32(0)
	playerNum := 0
	// 以下逻辑前置条件是：玩家每回合只能投一个羊圈
	// 统计每个参与玩家的结果
	for _, v := range invests {
		playerNum++
		if v.FoldNo == maxInvestFoldNo {
			totalWinnerInvestNum += v.LambNum
		}
	}

	// 计算每个投注的盈亏情况
	for _, v := range invests {
		if v.FoldNo != maxInvestFoldNo {
			// 没有中奖的玩家
			investResult = append(investResult, PlayerInvestInfo{
				InvestId:      v.ID,
				PlayerId:      v.PlayerID,
				PlayerEmail:   v.PlayerEmail,
				RoundId:       v.RoundID,
				LambFoldNo:    v.FoldNo,
				LambNum:       v.LambNum,
				ProfitAndLoss: float32(0),
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
				ProfitAndLoss: float32(*sliverNum) * proportion,
			})
		}
	}
	fmt.Printf("ProcessOpen[%s]: Silver Lamb Choose LambFold: %d\n", mode, maxInvestFoldNo)
	return maxInvestFoldNo, float64(*sliverNum), investResult, nil
}

func (l *DealOpenGameLogic) DealSingleWolfCase(mode string, round *wolflamp.RoundInfo, invests []*ent.RoundInvest) (choiceFoldNo uint32, totalRewardNum float64, investResult []PlayerInvestInfo, err error) {

	// 抽选羊圈
	foldChoice, err := l.ChooseLambFold(mode, round)
	if err != nil {
		return
	}
	// 按照羊圈+用户纬度统计投注数量
	//lambFoldPlayerSorts := make(map[uint32]map[uint64]uint32)
	// 统计盈利方总投注数量
	totalWinnerInvestNum := uint32(0)
	totalLoserInvestNum := float64(0)
	playerNum := 0
	// 以下逻辑前置条件是：玩家每回合只能投一个羊圈
	// 统计每个参与玩家的结果
	for _, v := range invests {
		playerNum++
		if v.FoldNo != *foldChoice {
			totalWinnerInvestNum += v.LambNum
		} else {
			totalLoserInvestNum += float64(v.LambNum)
		}
	}

	// investResult = make([]PlayerInvestInfo, playerNum)
	for _, v := range invests {
		if v.FoldNo == *foldChoice {
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
	fmt.Printf("ProcessOpen[%s]: Single Wolf Choose LambFold is : %d\n", mode, *foldChoice)
	return *foldChoice, totalLoserInvestNum, investResult, nil

}

//func (l *DealOpenGameLogic) DealMultiWolfCase(in *wolflamp.DealOpenGameReq) (*wolflamp.BaseIDResp, error) {
//}

func (l *DealOpenGameLogic) DealOpenGame(in *wolflamp.DealOpenGameReq) (*wolflamp.BaseIDResp, error) {

	openLock := fmt.Sprintf("current_game:%s_opening_lock", in.Mode)

	round, err := NewFindRoundLogic(l.ctx, l.svcCtx).FindRound(&wolflamp.FindRoundReq{Mode: in.Mode})
	if err != nil {
		return nil, err
	}
	if round.Status != roundenum.Investing.Val() {
		l.svcCtx.Redis.Del(l.ctx, openLock)
		fmt.Printf("ProcessOpen[%s]: round status is not investing, exist", in.Mode)
		return nil, errorx.NewAlreadyExistsError("game was already open")
	}

	// 获取所有投注信息
	allInvests, err := l.svcCtx.DB.RoundInvest.Query().Where(roundinvest.Mode(in.Mode)).
		Where(roundinvest.RoundID(round.Id)).All(l.ctx)
	if err != nil {
		fmt.Printf("ProcessOpen[%s]: get all invest list faild: %s", in.Mode, err.Error())
		l.svcCtx.Redis.Del(l.ctx, openLock)
		return nil, err
	}

	openType := uint32(0)
	// 先触发金羊逻辑
	choiceFoldNo, totalRewardNum, investResult, err := l.DealGoldenCase(in.Mode, allInvests)

	// 金羊没有触发再触发银羊逻辑
	if investResult == nil {
		choiceFoldNo, totalRewardNum, investResult, err = l.DealSliverCase(in.Mode, allInvests)
	} else {
		openType = roundenum.GoldenLamb.Val()
	}

	if investResult == nil {
		choiceFoldNo, totalRewardNum, investResult, err = l.DealSingleWolfCase(in.Mode, round, allInvests)
		if err != nil {
			fmt.Printf("ProcessOpen[%s]: deal single wolf case error : %s", in.Mode, err.Error())
			l.svcCtx.Redis.Del(l.ctx, openLock)
			return nil, err
		}
	} else if openType == roundenum.NoOpen.Val() {
		openType = roundenum.SliverLamb.Val()
	}

	if openType == roundenum.NoOpen.Val() {
		openType = roundenum.SingleWolf.Val()
	}

	// 平台佣金比例
	commissionResp, err := setting.NewGetGameCommissionLogic(l.ctx, l.svcCtx).GetGameCommission(&wolflamp.Empty{})
	if err != nil {
		l.svcCtx.Redis.Del(l.ctx, openLock)
		fmt.Printf("ProcessOpen[%s]: get system commission setting error : %s", in.Mode, err.Error())
		return nil, err
	}
	systemCommission := commissionResp.GameCommission
	systemCommissionNum := totalRewardNum * float64(systemCommission) / 100

	// 资金池总预留比例
	poolCommissionResp, err := setting.NewGetPoolCommissionLogic(l.ctx, l.svcCtx).GetPoolCommission(&wolflamp.Empty{})
	if err != nil {
		l.svcCtx.Redis.Del(l.ctx, openLock)
		fmt.Printf("ProcessOpen[%s]: get pool commission setting error : %s", in.Mode, err.Error())
		return nil, err
	}
	poolCommission := poolCommissionResp.Commission
	poolCommissionNum := totalRewardNum * float64(poolCommission) / 100
	// 机器人池预留占比
	robPoolCommissionResp, err := setting.NewGetRobPoolCommissionLogic(l.ctx, l.svcCtx).GetRobPoolCommission(&wolflamp.Empty{})
	if err != nil {
		l.svcCtx.Redis.Del(l.ctx, openLock)
		fmt.Printf("ProcessOpen[%s]: get robot pool commission setting error : %s", in.Mode, err.Error())
		return nil, err
	}
	// 奖金池预留占比
	rewardPoolCommissionResp, err := setting.NewGetRewardPoolCommissionLogic(l.ctx, l.svcCtx).GetRewardPoolCommission(&wolflamp.Empty{})
	if err != nil {
		l.svcCtx.Redis.Del(l.ctx, openLock)
		fmt.Printf("ProcessOpen[%s]: get reward pool commission setting error : %s", in.Mode, err.Error())
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
	invitorCommissionNum := totalRewardNum * invitorCommission

	// 剩余部分是赢的玩家按照投注占比进行瓜分的
	totalRewardNum = totalRewardNum - systemCommissionNum - poolCommissionNum - invitorCommissionNum

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
				SetLambNum(robPoolCommissionNum).SetRemark("机器人池预留").
				Exec(l.ctx)
			if err != nil {
				fmt.Printf("ProcessOpen[%s]: create robot pool commission error : %s", in.Mode, err.Error())
				return err
			}
		}
		if rewardPoolCommissionNum > 0 {
			err := l.svcCtx.DB.Pool.Create().
				SetMode(in.Mode).
				SetStatus(1).SetRoundID(round.Id).SetType(poolenum.Reward.Val()).
				SetLambNum(rewardPoolCommissionNum).
				SetRemark("奖金池预留").
				Exec(l.ctx)
			if err != nil {
				fmt.Printf("ProcessOpen[%s]: create reward pool commission error : %s", in.Mode, err.Error())
				return err
			}
		}
		if restPoolCommissionNum > 0 {
			err := l.svcCtx.DB.Pool.Create().SetMode(in.Mode).
				SetStatus(1).SetRoundID(round.Id).SetType(poolenum.Other.Val()).
				SetLambNum(restPoolCommissionNum).Exec(l.ctx)
			if err != nil {
				fmt.Printf("ProcessOpen[%s]: create rest pool commission error : %s", in.Mode, err.Error())
				return err
			}
		}
		// 记录资金池预留收益账单
		_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
			Mode:     in.Mode,
			PlayerId: 0, Status: statementenum.Completed.Val(), Module: statementenum.Pool.Val(),
			Amount: poolCommissionNum, InoutType: statementenum.Income.Val(),
			ReferId: strconv.FormatUint(round.Id, 10), Prefix: pointy.GetPointer("IV"),
			Remark: pointy.GetPointer(fmt.Sprintf("资金池预留：%f*%f=%f", float32(totalRewardNum), poolCommission/100, poolCommissionNum)),
		})
		if err != nil {
			fmt.Printf("ProcessOpen[%s]: create pool statement error : %s", in.Mode, err.Error())
			return err
		}

		// 记录平台收益账单
		_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
			Mode:     in.Mode,
			PlayerId: 0, Status: statementenum.Completed.Val(), Module: statementenum.System.Val(),
			Amount: systemCommissionNum, InoutType: statementenum.Income.Val(),
			ReferId: strconv.FormatUint(round.Id, 10), Prefix: pointy.GetPointer("IV"),
			Remark: pointy.GetPointer(fmt.Sprintf("平台收益：%f*%f=%f", float32(totalRewardNum), systemCommission/100, systemCommissionNum)),
		})
		if err != nil {
			fmt.Printf("ProcessOpen[%s]: create system commission statement error : %s", in.Mode, err.Error())
			return err
		}

		// 记录预留给上级佣金账单
		_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
			Mode:     in.Mode,
			PlayerId: 0, Status: statementenum.Completed.Val(), Module: statementenum.Reward.Val(),
			Amount: invitorCommissionNum, InoutType: statementenum.Income.Val(),
			ReferId: strconv.FormatUint(round.Id, 10), Prefix: pointy.GetPointer("IV"),
			Remark: pointy.GetPointer(fmt.Sprintf("上级总分佣：%f*%f=%f", float32(totalRewardNum), invitorCommission, invitorCommissionNum)),
		})
		if err != nil {
			fmt.Printf("ProcessOpen[%s]: create invitor reward commission statement error : %s", in.Mode, err.Error())
			return err
		}

		// 回传第三方的总数量=真实用户上级总分佣
		computedAmount := invitorCommissionNum
		// 更新投注记录的盈亏数据
		for _, investInfo := range investResult {
			err := l.svcCtx.DB.RoundInvest.Update().Where(roundinvest.ID(investInfo.InvestId)).
				SetProfitAndLoss(investInfo.ProfitAndLoss).Exec(l.ctx)
			if err != nil {
				fmt.Printf("ProcessOpen[%s]: update invest error : %s", in.Mode, err.Error())
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
						fmt.Printf("ProcessOpen[%s]: update player error : %s", in.Mode, err.Error())
						return err
					}
					// 赢亏为正的，则记录流水账单
					if investInfo.ProfitAndLoss > 0 {
						_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
							Mode:     in.Mode,
							PlayerId: investInfo.PlayerId, Status: statementenum.Completed.Val(), Module: statementenum.Invest.Val(),
							Amount: float64(investInfo.ProfitAndLoss), InoutType: statementenum.Income.Val(),
							ReferId: strconv.FormatUint(investInfo.InvestId, 10), Prefix: pointy.GetPointer("IV"),
							Remark: pointy.GetPointer(fmt.Sprintf("玩家获胜：%f*%f=%f", float32(totalRewardNum), investInfo.Proportion, investInfo.ProfitAndLoss)),
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
					fmt.Printf("ProcessOpen[%s]: create player invest statement error : %s", in.Mode, err.Error())
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
						fmt.Printf("ProcessOpen[%s]: create robot pool invest error : %s", in.Mode, err.Error())
						return err
					}

					_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
						Mode:     in.Mode,
						PlayerId: investInfo.PlayerId, Status: statementenum.Completed.Val(), Module: statementenum.Invest.Val(),
						Amount: float64(investInfo.ProfitAndLoss), InoutType: statementenum.Income.Val(),
						ReferId: strconv.FormatUint(investInfo.InvestId, 10), Prefix: pointy.GetPointer("IV"),
						Remark: pointy.GetPointer(fmt.Sprintf("机器人获胜：%f*%f=%f",
							float32(totalRewardNum), investInfo.Proportion, float64(investInfo.ProfitAndLoss))),
					})
				} else if investInfo.ProfitAndLoss < 0 {
					_, err = statementCreateLogic.CreateStatement(&wolflamp.CreateStatementReq{
						Mode:     in.Mode,
						PlayerId: investInfo.PlayerId, Status: statementenum.Completed.Val(), Module: statementenum.Invest.Val(),
						Amount: float64(investInfo.ProfitAndLoss), InoutType: statementenum.Expense.Val(),
						ReferId: strconv.FormatUint(investInfo.InvestId, 10), Prefix: pointy.GetPointer("IV"),
					})
				}

				if err != nil {
					fmt.Printf("ProcessOpen[%s]: create robot invest statement error : %s", in.Mode, err.Error())
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
				fmt.Printf("ProcessOpen[%s]: update fold error : %s", in.Mode, err.Error())
				return err
			}
		}
		// 更新回合状态
		err := l.svcCtx.DB.Round.UpdateOneID(round.Id).
			SetStatus(uint8(roundenum.Opening.Val())).
			SetComputeAmount(computedAmount / invitorCommission).
			SetOpenType(openType).
			SetSyncStatus(uint32(roundenum.NotYet)).
			SetSelectedFold(choiceFoldNo).Exec(l.ctx)
		if err != nil {
			fmt.Printf("ProcessOpen[%s]: update round error : %s", in.Mode, err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		l.svcCtx.Redis.Del(l.ctx, openLock)
		return nil, err
	}

	return &wolflamp.BaseIDResp{}, nil

}

// ChooseLambFold 抽选羊圈
func (l *DealOpenGameLogic) ChooseLambFold(mode string, round *wolflamp.RoundInfo) (*uint32, error) {

	fmt.Printf("ChooseLambFold[%s]\n", mode)

	// 统计当前回合没有投注的羊圈
	lambFoldInvestInfo := [8]bool{false, false, false, false, false, false, false, false}
	lambFoldInvested := make([]uint32, 0)
	for _, fold := range round.Folds {
		if fold.LambNum > 0 {
			lambFoldInvestInfo[fold.FoldNo-1] = true
			lambFoldInvested = append(lambFoldInvested, fold.FoldNo)
		}
	}

	// 根据查询值获取羊圈统计数据
	aggregateResult, err := NewGetLambFoldAggregateV2Logic(l.ctx, l.svcCtx).
		GetLambFoldAggregateV2(&wolflamp.GetLambFoldAggregateV2Req{Mode: mode})
	if err != nil {
		return nil, err
	}
	// 先排除掉没有投注的羊圈
	aggregateExcludeResult := make(map[uint32]*wolflamp.LambFoldAggregateInfo)
	for _, v := range aggregateResult.Data {
		if lambFoldInvestInfo[v.LambFoldNo-1] {
			aggregateExcludeResult[v.LambFoldNo] = v
		}
	}
	if len(aggregateExcludeResult) < 1 {
		//return nil, errorx.NewNotFoundError("no lamb fold to choose")
		return pointy.GetPointer(uint32(1)), nil
	}
	// 只有一个有投注则直接返回
	if len(aggregateExcludeResult) == 1 {
		for _, v := range aggregateExcludeResult {
			return pointy.GetPointer(v.LambFoldNo), nil
		}
	}

	// 盈亏最小的羊圈编号
	excludeFirstOne := uint32(0)
	// 盈亏第二小的羊圈编号
	excludeSecondOne := uint32(0)
	// 盈亏数最小的2个羊圈优先排除
	// 双变量遍历一次即可得出两个盈亏最小的羊圈索引值
	for foldNo, v := range aggregateExcludeResult {
		if excludeFirstOne == 0 {
			excludeFirstOne = foldNo
			excludeSecondOne = foldNo
			continue
		}
		if v.ProfitAndLossCount < aggregateExcludeResult[excludeFirstOne].ProfitAndLossCount {
			excludeSecondOne = excludeFirstOne
			excludeFirstOne = foldNo
		} else if v.ProfitAndLossCount < aggregateExcludeResult[excludeSecondOne].ProfitAndLossCount {
			excludeSecondOne = foldNo
		}
	}
	// 如果只有两个羊圈有投注，则排除盈亏最小的那个羊圈
	if len(aggregateExcludeResult) == 2 {
		return &excludeSecondOne, nil
	}
	// 如果只有三个羊圈有投注，则排除盈亏最小的两个羊圈
	if len(aggregateExcludeResult) == 3 {
		for foldNo := range aggregateExcludeResult {
			if foldNo == excludeFirstOne || foldNo == excludeSecondOne {
				continue
			}
			return &foldNo, nil
		}
	}
	// 排除盈亏最小的两个羊圈
	delete(aggregateExcludeResult, excludeFirstOne)
	delete(aggregateExcludeResult, excludeSecondOne)

	// 排除胜率最低的羊圈
	excludeThirdOne := uint32(0)
	for foldNo, v := range aggregateExcludeResult {
		if excludeThirdOne == 0 {
			excludeThirdOne = foldNo
			continue
		}
		if v.AvgWinRate < aggregateExcludeResult[excludeThirdOne].AvgWinRate {
			excludeThirdOne = foldNo
		}
	}
	// 排除胜率最低的羊圈
	delete(aggregateExcludeResult, excludeThirdOne)

	// 在剩下的羊圈（至少1个）中随机抽选
	rand.Seed(time.Now().UnixNano())
	foldRand := rand.Intn(len(aggregateExcludeResult))
	count := 0
	for foldNo := range aggregateExcludeResult {
		if count == foldRand {
			return pointy.GetPointer(foldNo), nil
		}
		count++
	}
	return nil, errorx.NewNotFoundError("no lamb fold to choose")

}
