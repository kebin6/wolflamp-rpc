package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInvestInfoByPlayerIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInvestInfoByPlayerIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInvestInfoByPlayerIdLogic {
	return &GetInvestInfoByPlayerIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInvestInfoByPlayerIdLogic) GetInvestInfoByPlayerId(in *wolflamp.GetInvestInfoByPlayerIdReq) (*wolflamp.GetInvestInfoByPlayerIdResp, error) {

	var predicates []predicate.RoundInvest
	predicates = append(predicates, roundinvest.PlayerID(in.PlayerId))
	predicates = append(predicates, roundinvest.RoundID(*in.RoundId))

	result, err := l.svcCtx.DB.RoundInvest.Query().Where(predicates...).
		Order(ent.Desc("id")).
		All(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &wolflamp.GetInvestInfoByPlayerIdResp{}

	for _, v := range result {
		resp.Data = append(resp.Data, &wolflamp.InvestInfo{
			Id:            v.ID,
			CreatedAt:     v.CreatedAt.Unix(),
			UpdatedAt:     v.UpdatedAt.Unix(),
			RoundId:       v.RoundID,
			PlayerId:      v.PlayerID,
			PlayerEmail:   v.PlayerEmail,
			FoldNo:        v.FoldNo,
			LambNum:       v.LambNum,
			ProfitAndLoss: v.ProfitAndLoss,
		})
	}

	return resp, nil

}
