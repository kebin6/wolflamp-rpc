package player

import (
	"context"
	"fmt"
	"github.com/kebin6/wolflamp-rpc/ent/player"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"strconv"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPlayerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPlayerLogic {
	return &ListPlayerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPlayerLogic) ListPlayer(in *wolflamp.ListPlayerReq) (*wolflamp.ListPlayerResp, error) {

	var predicates []predicate.Player
	if in.Status != nil {
		predicates = append(predicates, player.Status(*pointy.GetStatusPointer(in.Status)))
	}
	if in.InviteCode != nil {
		predicates = append(predicates, player.InviteCode(*in.InviteCode))
	}
	if in.Name != nil {
		predicates = append(predicates, player.Name(*in.Name))
	}
	if in.Rank != nil {
		predicates = append(predicates, player.Rank(*in.Rank))
	}
	if in.InvitedCode != nil {
		predicates = append(predicates, player.InvitedCode(*in.InvitedCode))
	}
	if in.Email != nil {
		predicates = append(predicates, player.Email(*in.Email))
	}

	result, err := l.svcCtx.DB.Player.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	playerIds := make([]uint64, result.PageDetails.Total)
	for i, v := range result.List {
		playerIds[i] = v.ID
	}
	// 获取玩家投注信息
	invests, _ := l.svcCtx.DB.RoundInvest.Query().
		Where(roundinvest.PlayerIDIn(playerIds...)).
		Select(roundinvest.FieldPlayerID, roundinvest.FieldMode, roundinvest.FieldLambNum, roundinvest.FieldProfitAndLoss).
		All(l.ctx)
	playerInvests := map[uint64]PlayerInvestInfo{}
	for _, v := range invests {
		var playerInvestInfo PlayerInvestInfo
		if _, ok := playerInvests[v.PlayerID]; !ok {
			playerInvestInfo = PlayerInvestInfo{}
		} else {
			playerInvestInfo = playerInvests[v.PlayerID]
		}

		if v.Mode == "coin" {
			playerInvestInfo.CoinTotalInvestCount++
			if v.ProfitAndLoss > 0 {
				playerInvestInfo.CoinTotalWinCount++
				playerInvestInfo.CoinTotalWinAmount += float64(v.ProfitAndLoss)
			} else {
				playerInvestInfo.CoinTotalLoseCount++
				playerInvestInfo.CoinTotalLoseAmount += float64(v.ProfitAndLoss)
			}
		} else {
			playerInvestInfo.TokenTotalInvestCount++
			if v.ProfitAndLoss > 0 {
				playerInvestInfo.TokenTotalWinCount++
				playerInvestInfo.TokenTotalWinAmount += float64(v.ProfitAndLoss)
			} else {
				playerInvestInfo.TokenTotalLoseCount++
				playerInvestInfo.TokenTotalLoseAmount += float64(v.ProfitAndLoss)
			}
		}
		playerInvests[v.PlayerID] = playerInvestInfo
	}

	resp := &wolflamp.ListPlayerResp{}
	resp.Total = result.PageDetails.Total

	findPlayerLogic := NewFindPlayerLogic(l.ctx, l.svcCtx)
	for _, v := range result.List {
		info := findPlayerLogic.Po2Vo(v)
		if v, ok := playerInvests[v.ID]; ok {
			info.CoinTotalWinCount = v.CoinTotalWinCount
			info.CoinTotalLoseCount = v.CoinTotalLoseCount
			info.CoinTotalWinAmount = v.CoinTotalWinAmount
			info.CoinTotalLoseAmount = v.CoinTotalLoseAmount
			if v.CoinTotalInvestCount > 0 {
				percent := fmt.Sprintf("%.2f", float32(v.CoinTotalWinCount)/float32(v.CoinTotalInvestCount)*100)
				float64Val, err := strconv.ParseFloat(percent, 32)
				if err == nil {
					float32Val := float32(float64Val)
					info.CoinWinPercent = float32Val
				}
			}
			info.TokenTotalWinCount = v.TokenTotalWinCount
			info.TokenTotalLoseCount = v.TokenTotalLoseCount
			info.TokenTotalWinAmount = v.TokenTotalWinAmount
			info.TokenTotalLoseAmount = v.TokenTotalLoseAmount
			if v.TokenTotalInvestCount > 0 {
				percent := fmt.Sprintf("%.2f", float32(v.TokenTotalWinCount)/float32(v.TokenTotalInvestCount)*100)
				float64Val, err := strconv.ParseFloat(percent, 32)
				if err == nil {
					float32Val := float32(float64Val)
					info.TokenWinPercent = float32Val
				}
			}
		}
		resp.Data = append(resp.Data, info)
	}

	return resp, nil

}

type PlayerInvestInfo struct {
	CoinTotalInvestCount  uint32
	CoinTotalWinCount     uint32
	CoinTotalLoseCount    uint32
	CoinTotalWinAmount    float64
	CoinTotalLoseAmount   float64
	TokenTotalInvestCount uint32
	TokenTotalWinCount    uint32
	TokenTotalLoseCount   uint32
	TokenTotalWinAmount   float64
	TokenTotalLoseAmount  float64
}
