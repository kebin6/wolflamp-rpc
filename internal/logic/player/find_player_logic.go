package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPlayerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPlayerLogic {
	return &FindPlayerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindPlayerLogic) FindPlayer(in *wolflamp.FindPlayerReq) (*wolflamp.PlayerInfo, error) {

	result, err := l.svcCtx.DB.Player.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return l.Po2Vo(result), nil

}

func (l *FindPlayerLogic) Po2Vo(po *ent.Player) (vo *wolflamp.PlayerInfo) {

	return &wolflamp.PlayerInfo{
		Id:                   po.ID,
		CreatedAt:            po.CreatedAt.Unix(),
		UpdatedAt:            po.UpdatedAt.Unix(),
		Status:               uint32(po.Status),
		Name:                 po.Name,
		Email:                po.Email,
		Lamp:                 po.Lamp,
		Rank:                 po.Rank,
		Amount:               po.Amount,
		InvitedNum:           po.InvitedNum,
		TotalIncome:          po.TotalIncome,
		ProfitAndLoss:        po.ProfitAndLoss,
		Recent_100WinPercent: po.Recent100WinPercent,
		InviteCode:           po.InviteCode,
		InviterId:            po.InviterID,
		InvitedCode:          po.InvitedCode,
		TransactionPassword:  po.TransactionPassword,
		DepositAddress:       po.DepositAddress,
	}

}
