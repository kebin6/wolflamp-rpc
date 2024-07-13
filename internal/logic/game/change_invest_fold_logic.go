package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/internal/utils/entx"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeInvestFoldLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeInvestFoldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeInvestFoldLogic {
	return &ChangeInvestFoldLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeInvestFoldLogic) ChangeInvestFold(in *wolflamp.ChangeInvestFoldReq) (*wolflamp.BaseIDResp, error) {

	// 获取当前轮次信息
	roundInfo, err := NewFindRoundLogic(l.ctx, l.svcCtx).FindRound(&wolflamp.FindRoundReq{Mode: in.Mode})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	// 获取用户当前回合的投注信息
	investResp, err := NewGetInvestInfoByPlayerIdLogic(l.ctx, l.svcCtx).GetInvestInfoByPlayerId(&wolflamp.GetInvestInfoByPlayerIdReq{
		PlayerId: in.PlayerId,
		RoundId:  &roundInfo.Id,
		Mode:     in.Mode,
	})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	if len(investResp.Data) > 1 {
		return nil, errorx.NewInternalError("player has twice invest in a round")
	}
	invest := investResp.Data[0]
	if invest.FoldNo == in.FoldNo {
		return &wolflamp.BaseIDResp{}, nil
	}

	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		err := l.svcCtx.DB.RoundInvest.UpdateOneID(invest.Id).SetFoldNo(in.FoldNo).Exec(l.ctx)
		if err != nil {
			return err
		}
		err = l.svcCtx.DB.RoundLambFold.Update().Where(roundlambfold.RoundID(roundInfo.Id), roundlambfold.FoldNo(invest.FoldNo)).
			AddLambNum(-int32(invest.LambNum)).Exec(l.ctx)
		if err != nil {
			return err
		}
		err = l.svcCtx.DB.RoundLambFold.Update().Where(roundlambfold.RoundID(roundInfo.Id), roundlambfold.FoldNo(in.FoldNo)).
			AddLambNum(int32(invest.LambNum)).Exec(l.ctx)
		if err != nil {
			return err
		}
		return nil
	})

	return &wolflamp.BaseIDResp{}, nil
}
