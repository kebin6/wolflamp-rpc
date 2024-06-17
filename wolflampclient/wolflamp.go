// Code generated by goctl. DO NOT EDIT.
// Source: wolflamp.proto

package wolflampclient

import (
	"context"

	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BannerInfo                  = wolflamp.BannerInfo
	BaseIDInt32Resp             = wolflamp.BaseIDInt32Resp
	BaseIDInt64Resp             = wolflamp.BaseIDInt64Resp
	BaseIDResp                  = wolflamp.BaseIDResp
	BaseIDUint32Resp            = wolflamp.BaseIDUint32Resp
	BaseResp                    = wolflamp.BaseResp
	BaseUUIDResp                = wolflamp.BaseUUIDResp
	CalculateWithdrawFeeReq     = wolflamp.CalculateWithdrawFeeReq
	CalculateWithdrawFeeResp    = wolflamp.CalculateWithdrawFeeResp
	ChangeInvestFoldReq         = wolflamp.ChangeInvestFoldReq
	CreateBannerReq             = wolflamp.CreateBannerReq
	CreateExchangeReq           = wolflamp.CreateExchangeReq
	CreateInvestReq             = wolflamp.CreateInvestReq
	CreateOrderReq              = wolflamp.CreateOrderReq
	CreatePlayerReq             = wolflamp.CreatePlayerReq
	CreateRewardReq             = wolflamp.CreateRewardReq
	CreateRoundReq              = wolflamp.CreateRoundReq
	CreateStatementReq          = wolflamp.CreateStatementReq
	DealOpenGameReq             = wolflamp.DealOpenGameReq
	DeleteBannerReq             = wolflamp.DeleteBannerReq
	DeleteExchangeReq           = wolflamp.DeleteExchangeReq
	DeleteOrderReq              = wolflamp.DeleteOrderReq
	DeletePlayerReq             = wolflamp.DeletePlayerReq
	DeleteRewardReq             = wolflamp.DeleteRewardReq
	DeleteStatementReq          = wolflamp.DeleteStatementReq
	Empty                       = wolflamp.Empty
	ExchangeInfo                = wolflamp.ExchangeInfo
	ExchangeReq                 = wolflamp.ExchangeReq
	FileInfo                    = wolflamp.FileInfo
	FindBannerReq               = wolflamp.FindBannerReq
	FindExchangeReq             = wolflamp.FindExchangeReq
	FindOrderReq                = wolflamp.FindOrderReq
	FindPlayerReq               = wolflamp.FindPlayerReq
	FindRewardReq               = wolflamp.FindRewardReq
	FindRoundReq                = wolflamp.FindRoundReq
	FindSettingReq              = wolflamp.FindSettingReq
	FindSettingResp             = wolflamp.FindSettingResp
	FindStatementReq            = wolflamp.FindStatementReq
	FoldInfo                    = wolflamp.FoldInfo
	GameCommissionResp          = wolflamp.GameCommissionResp
	GameRuleResp                = wolflamp.GameRuleResp
	GetByEmailReq               = wolflamp.GetByEmailReq
	GetByInviteCodeReq          = wolflamp.GetByInviteCodeReq
	GetInvestByRoundIdResp      = wolflamp.GetInvestByRoundIdResp
	GetInvestInfoByPlayerIdReq  = wolflamp.GetInvestInfoByPlayerIdReq
	GetInvestInfoByPlayerIdResp = wolflamp.GetInvestInfoByPlayerIdResp
	GetInvestsByRoundIdReq      = wolflamp.GetInvestsByRoundIdReq
	GetInvitorListByIdsReq      = wolflamp.GetInvitorListByIdsReq
	GetInvitorListByIdsResp     = wolflamp.GetInvitorListByIdsResp
	GetLambFoldAggregateReq     = wolflamp.GetLambFoldAggregateReq
	GetLambFoldAggregateResp    = wolflamp.GetLambFoldAggregateResp
	IDInt32Req                  = wolflamp.IDInt32Req
	IDInt64Req                  = wolflamp.IDInt64Req
	IDReq                       = wolflamp.IDReq
	IDUint32Req                 = wolflamp.IDUint32Req
	IDsInt32Req                 = wolflamp.IDsInt32Req
	IDsInt64Req                 = wolflamp.IDsInt64Req
	IDsReq                      = wolflamp.IDsReq
	IDsUint32Req                = wolflamp.IDsUint32Req
	IdleTimeResp                = wolflamp.IdleTimeResp
	InvestInfo                  = wolflamp.InvestInfo
	InvitorInfo                 = wolflamp.InvitorInfo
	LambFoldAggregateInfo       = wolflamp.LambFoldAggregateInfo
	ListBannerReq               = wolflamp.ListBannerReq
	ListBannerResp              = wolflamp.ListBannerResp
	ListExchangeReq             = wolflamp.ListExchangeReq
	ListExchangeResp            = wolflamp.ListExchangeResp
	ListFoldReq                 = wolflamp.ListFoldReq
	ListFoldResp                = wolflamp.ListFoldResp
	ListHistoryInvestReq        = wolflamp.ListHistoryInvestReq
	ListHistoryInvestResp       = wolflamp.ListHistoryInvestResp
	ListOrderReq                = wolflamp.ListOrderReq
	ListOrderResp               = wolflamp.ListOrderResp
	ListPlayerReq               = wolflamp.ListPlayerReq
	ListPlayerResp              = wolflamp.ListPlayerResp
	ListRewardReq               = wolflamp.ListRewardReq
	ListRewardResp              = wolflamp.ListRewardResp
	ListStatementReq            = wolflamp.ListStatementReq
	ListStatementResp           = wolflamp.ListStatementResp
	MinWithdrawNumResp          = wolflamp.MinWithdrawNumResp
	OrderInfo                   = wolflamp.OrderInfo
	PageInfoReq                 = wolflamp.PageInfoReq
	PlayerInfo                  = wolflamp.PlayerInfo
	PreviousRoundResp           = wolflamp.PreviousRoundResp
	RewardAggregateReq          = wolflamp.RewardAggregateReq
	RewardAggregateResp         = wolflamp.RewardAggregateResp
	RewardInfo                  = wolflamp.RewardInfo
	RobotLampNumResp            = wolflamp.RobotLampNumResp
	RobotNumResp                = wolflamp.RobotNumResp
	RoundInfo                   = wolflamp.RoundInfo
	StatementInfo               = wolflamp.StatementInfo
	UUIDReq                     = wolflamp.UUIDReq
	UUIDsReq                    = wolflamp.UUIDsReq
	UpdateBannerReq             = wolflamp.UpdateBannerReq
	UpdateExchangeReq           = wolflamp.UpdateExchangeReq
	UpdateInvestReq             = wolflamp.UpdateInvestReq
	UpdateOrderReq              = wolflamp.UpdateOrderReq
	UpdatePlayerReq             = wolflamp.UpdatePlayerReq
	UpdateRewardReq             = wolflamp.UpdateRewardReq
	UpdateSettingReq            = wolflamp.UpdateSettingReq
	UpdateStatementReq          = wolflamp.UpdateStatementReq
	WithdrawCommissionResp      = wolflamp.WithdrawCommissionResp
	WithdrawThresholdResp       = wolflamp.WithdrawThresholdResp

	Wolflamp interface {
		CreateBanner(ctx context.Context, in *CreateBannerReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateBanner(ctx context.Context, in *UpdateBannerReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		DeleteBanner(ctx context.Context, in *DeleteBannerReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		FindBanner(ctx context.Context, in *FindBannerReq, opts ...grpc.CallOption) (*BannerInfo, error)
		ListBanner(ctx context.Context, in *ListBannerReq, opts ...grpc.CallOption) (*ListBannerResp, error)
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		CreateExchange(ctx context.Context, in *CreateExchangeReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateExchange(ctx context.Context, in *UpdateExchangeReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		DeleteExchange(ctx context.Context, in *DeleteExchangeReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		FindExchange(ctx context.Context, in *FindExchangeReq, opts ...grpc.CallOption) (*ExchangeInfo, error)
		ListExchange(ctx context.Context, in *ListExchangeReq, opts ...grpc.CallOption) (*ListExchangeResp, error)
		Exchange(ctx context.Context, in *ExchangeReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		CreateRound(ctx context.Context, in *CreateRoundReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		FindRound(ctx context.Context, in *FindRoundReq, opts ...grpc.CallOption) (*RoundInfo, error)
		PreviousRound(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PreviousRoundResp, error)
		ListHistoryInvest(ctx context.Context, in *ListHistoryInvestReq, opts ...grpc.CallOption) (*ListHistoryInvestResp, error)
		ListFold(ctx context.Context, in *ListFoldReq, opts ...grpc.CallOption) (*ListFoldResp, error)
		Invest(ctx context.Context, in *CreateInvestReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		ChangeInvestFold(ctx context.Context, in *ChangeInvestFoldReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		GetInvestInfoByPlayerId(ctx context.Context, in *GetInvestInfoByPlayerIdReq, opts ...grpc.CallOption) (*GetInvestInfoByPlayerIdResp, error)
		GetInvestByRoundId(ctx context.Context, in *GetInvestsByRoundIdReq, opts ...grpc.CallOption) (*GetInvestByRoundIdResp, error)
		GetLambFoldAggregate(ctx context.Context, in *GetLambFoldAggregateReq, opts ...grpc.CallOption) (*GetLambFoldAggregateResp, error)
		DealOpenGame(ctx context.Context, in *DealOpenGameReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		FindOrder(ctx context.Context, in *FindOrderReq, opts ...grpc.CallOption) (*OrderInfo, error)
		ListOrder(ctx context.Context, in *ListOrderReq, opts ...grpc.CallOption) (*ListOrderResp, error)
		CalculateWithdrawFee(ctx context.Context, in *CalculateWithdrawFeeReq, opts ...grpc.CallOption) (*CalculateWithdrawFeeResp, error)
		CreatePlayer(ctx context.Context, in *CreatePlayerReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdatePlayer(ctx context.Context, in *UpdatePlayerReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		DeletePlayer(ctx context.Context, in *DeletePlayerReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		FindPlayer(ctx context.Context, in *FindPlayerReq, opts ...grpc.CallOption) (*PlayerInfo, error)
		ListPlayer(ctx context.Context, in *ListPlayerReq, opts ...grpc.CallOption) (*ListPlayerResp, error)
		GetByEmail(ctx context.Context, in *GetByEmailReq, opts ...grpc.CallOption) (*PlayerInfo, error)
		GetByInviteCode(ctx context.Context, in *GetByInviteCodeReq, opts ...grpc.CallOption) (*PlayerInfo, error)
		GetInvitorListByIds(ctx context.Context, in *GetInvitorListByIdsReq, opts ...grpc.CallOption) (*GetInvitorListByIdsResp, error)
		CreateReward(ctx context.Context, in *CreateRewardReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateReward(ctx context.Context, in *UpdateRewardReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		DeleteReward(ctx context.Context, in *DeleteRewardReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		FindReward(ctx context.Context, in *FindRewardReq, opts ...grpc.CallOption) (*RewardInfo, error)
		ListReward(ctx context.Context, in *ListRewardReq, opts ...grpc.CallOption) (*ListRewardResp, error)
		RewardAggregate(ctx context.Context, in *RewardAggregateReq, opts ...grpc.CallOption) (*RewardAggregateResp, error)
		UpdateSetting(ctx context.Context, in *UpdateSettingReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		FindSetting(ctx context.Context, in *FindSettingReq, opts ...grpc.CallOption) (*FindSettingResp, error)
		GetWithdrawThreshold(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WithdrawThresholdResp, error)
		GetMinWithdrawNum(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MinWithdrawNumResp, error)
		GetWithdrawCommission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WithdrawCommissionResp, error)
		GetIdleTime(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IdleTimeResp, error)
		GetRobotLampNum(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RobotLampNumResp, error)
		GetGameRule(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameRuleResp, error)
		GetRobotNum(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RobotNumResp, error)
		GetGameCommission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameCommissionResp, error)
		CreateStatement(ctx context.Context, in *CreateStatementReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateStatement(ctx context.Context, in *UpdateStatementReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		DeleteStatement(ctx context.Context, in *DeleteStatementReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		FindStatement(ctx context.Context, in *FindStatementReq, opts ...grpc.CallOption) (*StatementInfo, error)
		ListStatement(ctx context.Context, in *ListStatementReq, opts ...grpc.CallOption) (*ListStatementResp, error)
	}

	defaultWolflamp struct {
		cli zrpc.Client
	}
)

func NewWolflamp(cli zrpc.Client) Wolflamp {
	return &defaultWolflamp{
		cli: cli,
	}
}

func (m *defaultWolflamp) CreateBanner(ctx context.Context, in *CreateBannerReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.CreateBanner(ctx, in, opts...)
}

func (m *defaultWolflamp) UpdateBanner(ctx context.Context, in *UpdateBannerReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.UpdateBanner(ctx, in, opts...)
}

func (m *defaultWolflamp) DeleteBanner(ctx context.Context, in *DeleteBannerReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.DeleteBanner(ctx, in, opts...)
}

func (m *defaultWolflamp) FindBanner(ctx context.Context, in *FindBannerReq, opts ...grpc.CallOption) (*BannerInfo, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.FindBanner(ctx, in, opts...)
}

func (m *defaultWolflamp) ListBanner(ctx context.Context, in *ListBannerReq, opts ...grpc.CallOption) (*ListBannerResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.ListBanner(ctx, in, opts...)
}

func (m *defaultWolflamp) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

func (m *defaultWolflamp) CreateExchange(ctx context.Context, in *CreateExchangeReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.CreateExchange(ctx, in, opts...)
}

func (m *defaultWolflamp) UpdateExchange(ctx context.Context, in *UpdateExchangeReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.UpdateExchange(ctx, in, opts...)
}

func (m *defaultWolflamp) DeleteExchange(ctx context.Context, in *DeleteExchangeReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.DeleteExchange(ctx, in, opts...)
}

func (m *defaultWolflamp) FindExchange(ctx context.Context, in *FindExchangeReq, opts ...grpc.CallOption) (*ExchangeInfo, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.FindExchange(ctx, in, opts...)
}

func (m *defaultWolflamp) ListExchange(ctx context.Context, in *ListExchangeReq, opts ...grpc.CallOption) (*ListExchangeResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.ListExchange(ctx, in, opts...)
}

func (m *defaultWolflamp) Exchange(ctx context.Context, in *ExchangeReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.Exchange(ctx, in, opts...)
}

func (m *defaultWolflamp) CreateRound(ctx context.Context, in *CreateRoundReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.CreateRound(ctx, in, opts...)
}

func (m *defaultWolflamp) FindRound(ctx context.Context, in *FindRoundReq, opts ...grpc.CallOption) (*RoundInfo, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.FindRound(ctx, in, opts...)
}

func (m *defaultWolflamp) PreviousRound(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PreviousRoundResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.PreviousRound(ctx, in, opts...)
}

func (m *defaultWolflamp) ListHistoryInvest(ctx context.Context, in *ListHistoryInvestReq, opts ...grpc.CallOption) (*ListHistoryInvestResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.ListHistoryInvest(ctx, in, opts...)
}

func (m *defaultWolflamp) ListFold(ctx context.Context, in *ListFoldReq, opts ...grpc.CallOption) (*ListFoldResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.ListFold(ctx, in, opts...)
}

func (m *defaultWolflamp) Invest(ctx context.Context, in *CreateInvestReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.Invest(ctx, in, opts...)
}

func (m *defaultWolflamp) ChangeInvestFold(ctx context.Context, in *ChangeInvestFoldReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.ChangeInvestFold(ctx, in, opts...)
}

func (m *defaultWolflamp) GetInvestInfoByPlayerId(ctx context.Context, in *GetInvestInfoByPlayerIdReq, opts ...grpc.CallOption) (*GetInvestInfoByPlayerIdResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetInvestInfoByPlayerId(ctx, in, opts...)
}

func (m *defaultWolflamp) GetInvestByRoundId(ctx context.Context, in *GetInvestsByRoundIdReq, opts ...grpc.CallOption) (*GetInvestByRoundIdResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetInvestByRoundId(ctx, in, opts...)
}

func (m *defaultWolflamp) GetLambFoldAggregate(ctx context.Context, in *GetLambFoldAggregateReq, opts ...grpc.CallOption) (*GetLambFoldAggregateResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetLambFoldAggregate(ctx, in, opts...)
}

func (m *defaultWolflamp) DealOpenGame(ctx context.Context, in *DealOpenGameReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.DealOpenGame(ctx, in, opts...)
}

func (m *defaultWolflamp) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.CreateOrder(ctx, in, opts...)
}

func (m *defaultWolflamp) UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.UpdateOrder(ctx, in, opts...)
}

func (m *defaultWolflamp) DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.DeleteOrder(ctx, in, opts...)
}

func (m *defaultWolflamp) FindOrder(ctx context.Context, in *FindOrderReq, opts ...grpc.CallOption) (*OrderInfo, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.FindOrder(ctx, in, opts...)
}

func (m *defaultWolflamp) ListOrder(ctx context.Context, in *ListOrderReq, opts ...grpc.CallOption) (*ListOrderResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.ListOrder(ctx, in, opts...)
}

func (m *defaultWolflamp) CalculateWithdrawFee(ctx context.Context, in *CalculateWithdrawFeeReq, opts ...grpc.CallOption) (*CalculateWithdrawFeeResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.CalculateWithdrawFee(ctx, in, opts...)
}

func (m *defaultWolflamp) CreatePlayer(ctx context.Context, in *CreatePlayerReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.CreatePlayer(ctx, in, opts...)
}

func (m *defaultWolflamp) UpdatePlayer(ctx context.Context, in *UpdatePlayerReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.UpdatePlayer(ctx, in, opts...)
}

func (m *defaultWolflamp) DeletePlayer(ctx context.Context, in *DeletePlayerReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.DeletePlayer(ctx, in, opts...)
}

func (m *defaultWolflamp) FindPlayer(ctx context.Context, in *FindPlayerReq, opts ...grpc.CallOption) (*PlayerInfo, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.FindPlayer(ctx, in, opts...)
}

func (m *defaultWolflamp) ListPlayer(ctx context.Context, in *ListPlayerReq, opts ...grpc.CallOption) (*ListPlayerResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.ListPlayer(ctx, in, opts...)
}

func (m *defaultWolflamp) GetByEmail(ctx context.Context, in *GetByEmailReq, opts ...grpc.CallOption) (*PlayerInfo, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetByEmail(ctx, in, opts...)
}

func (m *defaultWolflamp) GetByInviteCode(ctx context.Context, in *GetByInviteCodeReq, opts ...grpc.CallOption) (*PlayerInfo, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetByInviteCode(ctx, in, opts...)
}

func (m *defaultWolflamp) GetInvitorListByIds(ctx context.Context, in *GetInvitorListByIdsReq, opts ...grpc.CallOption) (*GetInvitorListByIdsResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetInvitorListByIds(ctx, in, opts...)
}

func (m *defaultWolflamp) CreateReward(ctx context.Context, in *CreateRewardReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.CreateReward(ctx, in, opts...)
}

func (m *defaultWolflamp) UpdateReward(ctx context.Context, in *UpdateRewardReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.UpdateReward(ctx, in, opts...)
}

func (m *defaultWolflamp) DeleteReward(ctx context.Context, in *DeleteRewardReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.DeleteReward(ctx, in, opts...)
}

func (m *defaultWolflamp) FindReward(ctx context.Context, in *FindRewardReq, opts ...grpc.CallOption) (*RewardInfo, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.FindReward(ctx, in, opts...)
}

func (m *defaultWolflamp) ListReward(ctx context.Context, in *ListRewardReq, opts ...grpc.CallOption) (*ListRewardResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.ListReward(ctx, in, opts...)
}

func (m *defaultWolflamp) RewardAggregate(ctx context.Context, in *RewardAggregateReq, opts ...grpc.CallOption) (*RewardAggregateResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.RewardAggregate(ctx, in, opts...)
}

func (m *defaultWolflamp) UpdateSetting(ctx context.Context, in *UpdateSettingReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.UpdateSetting(ctx, in, opts...)
}

func (m *defaultWolflamp) FindSetting(ctx context.Context, in *FindSettingReq, opts ...grpc.CallOption) (*FindSettingResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.FindSetting(ctx, in, opts...)
}

func (m *defaultWolflamp) GetWithdrawThreshold(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WithdrawThresholdResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetWithdrawThreshold(ctx, in, opts...)
}

func (m *defaultWolflamp) GetMinWithdrawNum(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MinWithdrawNumResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetMinWithdrawNum(ctx, in, opts...)
}

func (m *defaultWolflamp) GetWithdrawCommission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WithdrawCommissionResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetWithdrawCommission(ctx, in, opts...)
}

func (m *defaultWolflamp) GetIdleTime(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IdleTimeResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetIdleTime(ctx, in, opts...)
}

func (m *defaultWolflamp) GetRobotLampNum(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RobotLampNumResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetRobotLampNum(ctx, in, opts...)
}

func (m *defaultWolflamp) GetGameRule(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameRuleResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetGameRule(ctx, in, opts...)
}

func (m *defaultWolflamp) GetRobotNum(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RobotNumResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetRobotNum(ctx, in, opts...)
}

func (m *defaultWolflamp) GetGameCommission(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameCommissionResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.GetGameCommission(ctx, in, opts...)
}

func (m *defaultWolflamp) CreateStatement(ctx context.Context, in *CreateStatementReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.CreateStatement(ctx, in, opts...)
}

func (m *defaultWolflamp) UpdateStatement(ctx context.Context, in *UpdateStatementReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.UpdateStatement(ctx, in, opts...)
}

func (m *defaultWolflamp) DeleteStatement(ctx context.Context, in *DeleteStatementReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.DeleteStatement(ctx, in, opts...)
}

func (m *defaultWolflamp) FindStatement(ctx context.Context, in *FindStatementReq, opts ...grpc.CallOption) (*StatementInfo, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.FindStatement(ctx, in, opts...)
}

func (m *defaultWolflamp) ListStatement(ctx context.Context, in *ListStatementReq, opts ...grpc.CallOption) (*ListStatementResp, error) {
	client := wolflamp.NewWolflampClient(m.cli.Conn())
	return client.ListStatement(ctx, in, opts...)
}
