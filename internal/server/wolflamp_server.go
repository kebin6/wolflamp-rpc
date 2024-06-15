// Code generated by goctl. DO NOT EDIT.
// Source: wolflamp.proto

package server

import (
	"context"

	"github.com/kebin6/wolflamp-rpc/internal/logic/banner"
	"github.com/kebin6/wolflamp-rpc/internal/logic/base"
	"github.com/kebin6/wolflamp-rpc/internal/logic/exchange"
	"github.com/kebin6/wolflamp-rpc/internal/logic/game"
	"github.com/kebin6/wolflamp-rpc/internal/logic/order"
	"github.com/kebin6/wolflamp-rpc/internal/logic/player"
	"github.com/kebin6/wolflamp-rpc/internal/logic/reward"
	"github.com/kebin6/wolflamp-rpc/internal/logic/setting"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
)

type WolflampServer struct {
	svcCtx *svc.ServiceContext
	wolflamp.UnimplementedWolflampServer
}

func NewWolflampServer(svcCtx *svc.ServiceContext) *WolflampServer {
	return &WolflampServer{
		svcCtx: svcCtx,
	}
}

func (s *WolflampServer) CreateBanner(ctx context.Context, in *wolflamp.CreateBannerReq) (*wolflamp.BaseIDResp, error) {
	l := banner.NewCreateBannerLogic(ctx, s.svcCtx)
	return l.CreateBanner(in)
}

func (s *WolflampServer) UpdateBanner(ctx context.Context, in *wolflamp.UpdateBannerReq) (*wolflamp.BaseIDResp, error) {
	l := banner.NewUpdateBannerLogic(ctx, s.svcCtx)
	return l.UpdateBanner(in)
}

func (s *WolflampServer) DeleteBanner(ctx context.Context, in *wolflamp.DeleteBannerReq) (*wolflamp.BaseIDResp, error) {
	l := banner.NewDeleteBannerLogic(ctx, s.svcCtx)
	return l.DeleteBanner(in)
}

func (s *WolflampServer) FindBanner(ctx context.Context, in *wolflamp.FindBannerReq) (*wolflamp.BannerInfo, error) {
	l := banner.NewFindBannerLogic(ctx, s.svcCtx)
	return l.FindBanner(in)
}

func (s *WolflampServer) ListBanner(ctx context.Context, in *wolflamp.ListBannerReq) (*wolflamp.ListBannerResp, error) {
	l := banner.NewListBannerLogic(ctx, s.svcCtx)
	return l.ListBanner(in)
}

func (s *WolflampServer) InitDatabase(ctx context.Context, in *wolflamp.Empty) (*wolflamp.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

func (s *WolflampServer) CreateExchange(ctx context.Context, in *wolflamp.CreateExchangeReq) (*wolflamp.BaseIDResp, error) {
	l := exchange.NewCreateExchangeLogic(ctx, s.svcCtx)
	return l.CreateExchange(in)
}

func (s *WolflampServer) UpdateExchange(ctx context.Context, in *wolflamp.UpdateExchangeReq) (*wolflamp.BaseIDResp, error) {
	l := exchange.NewUpdateExchangeLogic(ctx, s.svcCtx)
	return l.UpdateExchange(in)
}

func (s *WolflampServer) DeleteExchange(ctx context.Context, in *wolflamp.DeleteExchangeReq) (*wolflamp.BaseIDResp, error) {
	l := exchange.NewDeleteExchangeLogic(ctx, s.svcCtx)
	return l.DeleteExchange(in)
}

func (s *WolflampServer) FindExchange(ctx context.Context, in *wolflamp.FindExchangeReq) (*wolflamp.ExchangeInfo, error) {
	l := exchange.NewFindExchangeLogic(ctx, s.svcCtx)
	return l.FindExchange(in)
}

func (s *WolflampServer) ListExchange(ctx context.Context, in *wolflamp.ListExchangeReq) (*wolflamp.ListExchangeResp, error) {
	l := exchange.NewListExchangeLogic(ctx, s.svcCtx)
	return l.ListExchange(in)
}

func (s *WolflampServer) Exchange(ctx context.Context, in *wolflamp.ExchangeReq) (*wolflamp.BaseIDResp, error) {
	l := exchange.NewExchangeLogic(ctx, s.svcCtx)
	return l.Exchange(in)
}

func (s *WolflampServer) CreateRound(ctx context.Context, in *wolflamp.CreateRoundReq) (*wolflamp.BaseIDResp, error) {
	l := game.NewCreateRoundLogic(ctx, s.svcCtx)
	return l.CreateRound(in)
}

func (s *WolflampServer) FindRound(ctx context.Context, in *wolflamp.FindRoundReq) (*wolflamp.RoundInfo, error) {
	l := game.NewFindRoundLogic(ctx, s.svcCtx)
	return l.FindRound(in)
}

func (s *WolflampServer) PreviousRound(ctx context.Context, in *wolflamp.Empty) (*wolflamp.PreviousRoundResp, error) {
	l := game.NewPreviousRoundLogic(ctx, s.svcCtx)
	return l.PreviousRound(in)
}

func (s *WolflampServer) ListFold(ctx context.Context, in *wolflamp.ListFoldReq) (*wolflamp.ListFoldResp, error) {
	l := game.NewListFoldLogic(ctx, s.svcCtx)
	return l.ListFold(in)
}

func (s *WolflampServer) Invest(ctx context.Context, in *wolflamp.CreateInvestReq) (*wolflamp.BaseIDResp, error) {
	l := game.NewInvestLogic(ctx, s.svcCtx)
	return l.Invest(in)
}

func (s *WolflampServer) GetInvestInfoByPlayerId(ctx context.Context, in *wolflamp.GetInvestInfoByPlayerIdReq) (*wolflamp.GetInvestInfoByPlayerIdResp, error) {
	l := game.NewGetInvestInfoByPlayerIdLogic(ctx, s.svcCtx)
	return l.GetInvestInfoByPlayerId(in)
}

func (s *WolflampServer) GetLambFoldAggregate(ctx context.Context, in *wolflamp.GetLambFoldAggregateReq) (*wolflamp.GetLambFoldAggregateResp, error) {
	l := game.NewGetLambFoldAggregateLogic(ctx, s.svcCtx)
	return l.GetLambFoldAggregate(in)
}

func (s *WolflampServer) DealOpenGame(ctx context.Context, in *wolflamp.DealOpenGameReq) (*wolflamp.BaseIDResp, error) {
	l := game.NewDealOpenGameLogic(ctx, s.svcCtx)
	return l.DealOpenGame(in)
}

func (s *WolflampServer) CreateOrder(ctx context.Context, in *wolflamp.CreateOrderReq) (*wolflamp.BaseIDResp, error) {
	l := order.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}

func (s *WolflampServer) UpdateOrder(ctx context.Context, in *wolflamp.UpdateOrderReq) (*wolflamp.BaseIDResp, error) {
	l := order.NewUpdateOrderLogic(ctx, s.svcCtx)
	return l.UpdateOrder(in)
}

func (s *WolflampServer) DeleteOrder(ctx context.Context, in *wolflamp.DeleteOrderReq) (*wolflamp.BaseIDResp, error) {
	l := order.NewDeleteOrderLogic(ctx, s.svcCtx)
	return l.DeleteOrder(in)
}

func (s *WolflampServer) FindOrder(ctx context.Context, in *wolflamp.FindOrderReq) (*wolflamp.OrderInfo, error) {
	l := order.NewFindOrderLogic(ctx, s.svcCtx)
	return l.FindOrder(in)
}

func (s *WolflampServer) ListOrder(ctx context.Context, in *wolflamp.ListOrderReq) (*wolflamp.ListOrderResp, error) {
	l := order.NewListOrderLogic(ctx, s.svcCtx)
	return l.ListOrder(in)
}

func (s *WolflampServer) CalculateWithdrawFee(ctx context.Context, in *wolflamp.CalculateWithdrawFeeReq) (*wolflamp.CalculateWithdrawFeeResp, error) {
	l := order.NewCalculateWithdrawFeeLogic(ctx, s.svcCtx)
	return l.CalculateWithdrawFee(in)
}

func (s *WolflampServer) CreatePlayer(ctx context.Context, in *wolflamp.CreatePlayerReq) (*wolflamp.BaseIDResp, error) {
	l := player.NewCreatePlayerLogic(ctx, s.svcCtx)
	return l.CreatePlayer(in)
}

func (s *WolflampServer) UpdatePlayer(ctx context.Context, in *wolflamp.UpdatePlayerReq) (*wolflamp.BaseIDResp, error) {
	l := player.NewUpdatePlayerLogic(ctx, s.svcCtx)
	return l.UpdatePlayer(in)
}

func (s *WolflampServer) DeletePlayer(ctx context.Context, in *wolflamp.DeletePlayerReq) (*wolflamp.BaseIDResp, error) {
	l := player.NewDeletePlayerLogic(ctx, s.svcCtx)
	return l.DeletePlayer(in)
}

func (s *WolflampServer) FindPlayer(ctx context.Context, in *wolflamp.FindPlayerReq) (*wolflamp.PlayerInfo, error) {
	l := player.NewFindPlayerLogic(ctx, s.svcCtx)
	return l.FindPlayer(in)
}

func (s *WolflampServer) ListPlayer(ctx context.Context, in *wolflamp.ListPlayerReq) (*wolflamp.ListPlayerResp, error) {
	l := player.NewListPlayerLogic(ctx, s.svcCtx)
	return l.ListPlayer(in)
}

func (s *WolflampServer) GetByEmail(ctx context.Context, in *wolflamp.GetByEmailReq) (*wolflamp.PlayerInfo, error) {
	l := player.NewGetByEmailLogic(ctx, s.svcCtx)
	return l.GetByEmail(in)
}

func (s *WolflampServer) GetByInviteCode(ctx context.Context, in *wolflamp.GetByInviteCodeReq) (*wolflamp.PlayerInfo, error) {
	l := player.NewGetByInviteCodeLogic(ctx, s.svcCtx)
	return l.GetByInviteCode(in)
}

func (s *WolflampServer) CreateReward(ctx context.Context, in *wolflamp.CreateRewardReq) (*wolflamp.BaseIDResp, error) {
	l := reward.NewCreateRewardLogic(ctx, s.svcCtx)
	return l.CreateReward(in)
}

func (s *WolflampServer) UpdateReward(ctx context.Context, in *wolflamp.UpdateRewardReq) (*wolflamp.BaseIDResp, error) {
	l := reward.NewUpdateRewardLogic(ctx, s.svcCtx)
	return l.UpdateReward(in)
}

func (s *WolflampServer) DeleteReward(ctx context.Context, in *wolflamp.DeleteRewardReq) (*wolflamp.BaseIDResp, error) {
	l := reward.NewDeleteRewardLogic(ctx, s.svcCtx)
	return l.DeleteReward(in)
}

func (s *WolflampServer) FindReward(ctx context.Context, in *wolflamp.FindRewardReq) (*wolflamp.RewardInfo, error) {
	l := reward.NewFindRewardLogic(ctx, s.svcCtx)
	return l.FindReward(in)
}

func (s *WolflampServer) ListReward(ctx context.Context, in *wolflamp.ListRewardReq) (*wolflamp.ListRewardResp, error) {
	l := reward.NewListRewardLogic(ctx, s.svcCtx)
	return l.ListReward(in)
}

func (s *WolflampServer) RewardAggregate(ctx context.Context, in *wolflamp.RewardAggregateReq) (*wolflamp.RewardAggregateResp, error) {
	l := reward.NewRewardAggregateLogic(ctx, s.svcCtx)
	return l.RewardAggregate(in)
}

func (s *WolflampServer) UpdateSetting(ctx context.Context, in *wolflamp.UpdateSettingReq) (*wolflamp.BaseIDResp, error) {
	l := setting.NewUpdateSettingLogic(ctx, s.svcCtx)
	return l.UpdateSetting(in)
}

func (s *WolflampServer) FindSetting(ctx context.Context, in *wolflamp.FindSettingReq) (*wolflamp.FindSettingResp, error) {
	l := setting.NewFindSettingLogic(ctx, s.svcCtx)
	return l.FindSetting(in)
}

func (s *WolflampServer) GetWithdrawThreshold(ctx context.Context, in *wolflamp.Empty) (*wolflamp.WithdrawThresholdResp, error) {
	l := setting.NewGetWithdrawThresholdLogic(ctx, s.svcCtx)
	return l.GetWithdrawThreshold(in)
}

func (s *WolflampServer) GetMinWithdrawNum(ctx context.Context, in *wolflamp.Empty) (*wolflamp.MinWithdrawNumResp, error) {
	l := setting.NewGetMinWithdrawNumLogic(ctx, s.svcCtx)
	return l.GetMinWithdrawNum(in)
}

func (s *WolflampServer) GetWithdrawCommission(ctx context.Context, in *wolflamp.Empty) (*wolflamp.WithdrawCommissionResp, error) {
	l := setting.NewGetWithdrawCommissionLogic(ctx, s.svcCtx)
	return l.GetWithdrawCommission(in)
}

func (s *WolflampServer) GetIdleTime(ctx context.Context, in *wolflamp.Empty) (*wolflamp.IdleTimeResp, error) {
	l := setting.NewGetIdleTimeLogic(ctx, s.svcCtx)
	return l.GetIdleTime(in)
}

func (s *WolflampServer) GetRobotLampNum(ctx context.Context, in *wolflamp.Empty) (*wolflamp.RobotLampNumResp, error) {
	l := setting.NewGetRobotLampNumLogic(ctx, s.svcCtx)
	return l.GetRobotLampNum(in)
}

func (s *WolflampServer) GetGameRule(ctx context.Context, in *wolflamp.Empty) (*wolflamp.GameRuleResp, error) {
	l := setting.NewGetGameRuleLogic(ctx, s.svcCtx)
	return l.GetGameRule(in)
}

func (s *WolflampServer) GetRobotNum(ctx context.Context, in *wolflamp.Empty) (*wolflamp.RobotNumResp, error) {
	l := setting.NewGetRobotNumLogic(ctx, s.svcCtx)
	return l.GetRobotNum(in)
}
