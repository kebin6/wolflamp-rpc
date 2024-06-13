package game

import (
	"context"
	"fmt"
	"github.com/kebin6/wolflamp-rpc/common/enum/roundenum"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"
	"time"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type InvestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvestLogic {
	return &InvestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvestLogic) Invest(in *wolflamp.CreateInvestReq) (*wolflamp.BaseIDResp, error) {

	// 获取当前轮次信息
	roundInfo, err := l.svcCtx.DB.Round.Query().Order(ent.Desc(round.FieldID)).First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	if in.RoundId != roundInfo.ID {
		return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("round id %d has expired", in.RoundId))
	}
	// 不在投注阶段
	if roundInfo.Status != uint8(roundenum.Investing.Val()) || roundInfo.StartAt.Unix() > time.Now().Unix() ||
		roundInfo.OpenAt.Unix() < time.Now().Unix() {
		return nil, errorx.NewInvalidArgumentError(fmt.Sprint("investing time over"))
	}
	result, err := l.svcCtx.DB.RoundInvest.Create().
		SetPlayerID(in.PlayerId).
		SetRoundID(in.RoundId).
		SetFoldNo(in.FoldNo).
		SetLambNum(in.LambNum).
		SetProfitAndLoss(0.0).
		SetRoundCount(roundInfo.RoundCount).
		Save(l.ctx)

	if err != nil {
		return nil, err
	}
	return &wolflamp.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
