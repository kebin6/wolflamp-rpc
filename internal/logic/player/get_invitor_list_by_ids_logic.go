package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/player"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInvitorListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInvitorListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInvitorListByIdsLogic {
	return &GetInvitorListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInvitorListByIdsLogic) GetInvitorListByIds(in *wolflamp.GetInvitorListByIdsReq) (*wolflamp.GetInvitorListByIdsResp, error) {

	var invitors []*wolflamp.InvitorInfo
	// 获取直接上级列表
	players, err := l.svcCtx.DB.Player.Query().Where(player.IDIn(in.PlayerIds...)).
		WithInviter(func(q *ent.PlayerQuery) {
			q.WithInviter(func(q *ent.PlayerQuery) {
				q.WithInviter(func(q *ent.PlayerQuery) {
					q.Select(player.FieldID, player.FieldRank, player.FieldSystemCommission)
				}).Select(player.FieldID, player.FieldRank, player.FieldSystemCommission)
			}).Select(player.FieldID, player.FieldRank, player.FieldSystemCommission)
		}).Select(player.FieldID, player.FieldRank, player.FieldSystemCommission).All(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	if len(players) == 0 {
		return &wolflamp.GetInvitorListByIdsResp{}, nil
	}
	for _, playerOne := range players {
		invitors = append(invitors, l.AppendInvitorInfo(0, playerOne, 1)...)
	}
	return &wolflamp.GetInvitorListByIdsResp{Data: invitors}, nil

}

// AppendInvitorInfo 递归获取上级信息列表，level代表与当前玩家之间的代理层级关系
func (l *GetInvitorListByIdsLogic) AppendInvitorInfo(playerId uint64, player *ent.Player, level uint32) []*wolflamp.InvitorInfo {

	var invitors []*wolflamp.InvitorInfo
	if player == nil {
		return invitors
	}
	if player.Edges.Inviter == nil {
		return invitors
	}
	invitor := player.Edges.Inviter
	if playerId == 0 {
		playerId = player.ID
	}
	invitors = append(invitors, &wolflamp.InvitorInfo{
		InvitorId:        invitor.ID,
		PlayerId:         playerId,
		Rank:             invitor.Rank,
		SystemCommission: invitor.SystemCommission,
		PlayerEmail:      player.Email,
	})
	return append(invitors, l.AppendInvitorInfo(playerId, invitor, level+1)...)

}
