package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/ent/roundlambfold"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFoldLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListFoldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFoldLogic {
	return &ListFoldLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListFoldLogic) ListFold(in *wolflamp.ListFoldReq) (*wolflamp.ListFoldResp, error) {

	var predicates []predicate.RoundLambFold
	if in.RoundId != nil {
		predicates = append(predicates, roundlambfold.RoundID(*in.RoundId))
	}
	if in.TotalRoundCount != nil {
		predicates = append(predicates, roundlambfold.TotalRoundCount(*in.TotalRoundCount))
	}
	if in.TotalRoundCoundMin != nil {
		predicates = append(predicates, roundlambfold.TotalRoundCountGTE(*in.TotalRoundCoundMin))
	}
	if in.TotalRoundCoundMax != nil {
		predicates = append(predicates, roundlambfold.TotalRoundCountLTE(*in.TotalRoundCoundMax))
	}

	result, err := l.svcCtx.DB.RoundLambFold.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &wolflamp.ListFoldResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &wolflamp.FoldInfo{
			Id:              v.ID,
			RoundId:         v.RoundID,
			LambNum:         v.LambNum,
			FoldNo:          v.FoldNo,
			RoundCount:      v.RoundCount,
			TotalRoundCount: v.TotalRoundCount,
			CreatedAt:       v.CreatedAt.Unix(),
		})
	}

	return resp, nil

}
