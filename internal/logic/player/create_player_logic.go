package player

import (
	"context"
	"github.com/duke-git/lancet/v2/random"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/origininvitecode"
	"github.com/kebin6/wolflamp-rpc/ent/player"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"
	"time"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePlayerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePlayerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePlayerLogic {
	return &CreatePlayerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePlayerLogic) CreatePlayer(in *wolflamp.CreatePlayerReq) (*wolflamp.BaseIDResp, error) {

	// 邀请人信息
	var inviterId uint64 = 0
	if in.InvitedCode == "" {
		return nil, errorx.NewInvalidArgumentError("inviter 's code required.")
	}

	inviter, err := l.svcCtx.DB.Player.Query().Where(player.InviteCode(in.InvitedCode)).First(l.ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	// 校验是否来自原始邀请码
	if ent.IsNotFound(err) {
		_, err := l.svcCtx.DB.OriginInviteCode.Query().Where(origininvitecode.Code(in.InvitedCode)).First(l.ctx)
		if err != nil {
			return nil, errorx.NewInvalidArgumentError("common.inviterNotFound")
		}
	} else {
		inviterId = inviter.ID
	}

	// 邀请码
	// 随机选择10~30随机数，作为查询值N
	inviteCode := time.Now().Format("15") + time.Now().Format("05") + random.RandUpper(3)
	result, err := l.svcCtx.DB.Player.Create().
		SetStatus(uint8(in.Status)).
		SetEmail(in.Email).
		SetPassword(in.Password).
		SetName(in.Name).
		SetCoinLamb(in.CoinLamb).
		SetTokenLamb(in.TokenLamb).
		SetRank(in.Rank).
		SetAmount(in.Amount).
		SetInvitedNum(0).
		SetTotalIncome(0).
		SetProfitAndLoss(0).
		SetInviteCode(inviteCode).
		SetInviterID(inviterId).
		SetInvitedCode(in.InvitedCode).
		SetDepositAddress(in.DepositAddress).
		SetSystemCommission(in.SystemCommission).
		SetGcicsUserID(in.GcicsUserId).
		SetGcicsToken(in.GcicsToken).
		SetGcicsReturnURL(in.ReturnUrl).
		SetGcicsUserName(in.GcicsUserName).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil

}
