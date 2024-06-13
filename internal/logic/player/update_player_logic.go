package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePlayerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePlayerLogic {
	return &UpdatePlayerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePlayerLogic) UpdatePlayer(in *wolflamp.UpdatePlayerReq) (*wolflamp.BaseIDResp, error) {

	err := l.svcCtx.DB.Player.UpdateOneID(in.Id).
		SetNotNilEmail(in.Email).
		SetNotNilPassword(in.Password).
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilName(in.Name).
		SetNotNilLamp(in.Lamp).
		SetNotNilRank(in.Rank).
		SetNotNilAmount(in.Amount).
		SetNotNilInvitedNum(in.InvitedNum).
		SetNotNilTotalIncome(in.TotalIncome).
		SetNotNilProfitAndLoss(in.ProfitAndLoss).
		SetNotNilRecent100WinPercent(in.Recent_100WinPercent).
		SetNotNilTransactionPassword(in.TransactionPassword).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Msg: i18n.UpdateSuccess}, nil

}
