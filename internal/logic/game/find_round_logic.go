package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRoundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRoundLogic {
	return &FindRoundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindRoundLogic) FindRound(in *wolflamp.FindRoundReq) (*wolflamp.RoundInfo, error) {

	var query *ent.RoundQuery
	if in.Id != nil {
		query = l.svcCtx.DB.Round.Query().Where(round.ID(*in.Id))
	} else {
		query = l.svcCtx.DB.Round.Query().Order(ent.Desc(round.FieldID))
	}

	result, err := query.WithFold().First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return l.Po2Vo(result), nil

}

func (l *FindRoundLogic) Po2Vo(po *ent.Round) (vo *wolflamp.RoundInfo) {

	var folds []*wolflamp.FoldInfo
	if po.Edges.Fold != nil {
		for _, fold := range po.Edges.Fold {
			folds = append(folds, &wolflamp.FoldInfo{
				Id:        fold.ID,
				CreatedAt: fold.CreatedAt.Unix(),
				UpdatedAt: fold.UpdatedAt.Unix(),
				RoundId:   fold.RoundID,
				FoldNo:    fold.FoldNo,
				LambNum:   fold.LambNum,
			})
		}
	}

	return &wolflamp.RoundInfo{
		Id:           po.ID,
		CreatedAt:    po.CreatedAt.Unix(),
		UpdatedAt:    po.UpdatedAt.Unix(),
		StartAt:      po.StartAt.Unix(),
		OpenAt:       po.OpenAt.Unix(),
		EndAt:        po.EndAt.Unix(),
		Status:       uint32(po.Status),
		SelectedFold: po.SelectedFold,
		Folds:        folds,
	}

}
