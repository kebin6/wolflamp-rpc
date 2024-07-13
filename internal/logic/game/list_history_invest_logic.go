package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListHistoryInvestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListHistoryInvestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHistoryInvestLogic {
	return &ListHistoryInvestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListHistoryInvestLogic) ListHistoryInvest(in *wolflamp.ListHistoryInvestReq) (*wolflamp.ListHistoryInvestResp, error) {

	// 获取当前回合数据
	currentRound, err := NewFindRoundLogic(l.ctx, l.svcCtx).FindRound(&wolflamp.FindRoundReq{Mode: in.Mode})
	if err != nil {
		return nil, err
	}
	var predicates []predicate.RoundInvest
	if in.PlayerId != nil {
		predicates = append(predicates, roundinvest.PlayerID(*in.PlayerId))
	}
	predicates = append(predicates, roundinvest.Mode(in.Mode))
	predicates = append(predicates, roundinvest.TotalRoundCountLT(currentRound.TotalRoundCount))
	result, err := l.svcCtx.DB.RoundInvest.Query().Where(predicates...).
		Order(ent.Desc(roundinvest.FieldID)).
		Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &wolflamp.ListHistoryInvestResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, l.Po2Vo(v))
	}

	return resp, nil

}

func (l *ListHistoryInvestLogic) Po2Vo(po *ent.RoundInvest) (vo *wolflamp.InvestInfo) {

	return &wolflamp.InvestInfo{
		Id:            po.ID,
		CreatedAt:     po.CreatedAt.Unix(),
		UpdatedAt:     po.UpdatedAt.Unix(),
		RoundId:       po.RoundID,
		PlayerId:      po.PlayerID,
		PlayerEmail:   po.PlayerEmail,
		FoldNo:        po.FoldNo,
		LambNum:       po.LambNum,
		ProfitAndLoss: po.ProfitAndLoss,
		Mode:          po.Mode,
	}

}
