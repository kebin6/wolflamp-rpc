package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/player"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

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

	resp := &wolflamp.ListPlayerResp{}
	resp.Total = result.PageDetails.Total

	findPlayerLogic := NewFindPlayerLogic(l.ctx, l.svcCtx)
	for _, v := range result.List {
		resp.Data = append(resp.Data, findPlayerLogic.Po2Vo(v))
	}

	return resp, nil

}
