package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/enum/cachekey"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/zeromicro/go-zero/core/errorx"
	"time"

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
	if in.Id == nil {
		currentId, err := l.GetCurrentRoundId()
		if err != nil {
			return nil, err
		}
		in.Id = &currentId
	}
	if in.Id == nil || *in.Id == 0 {
		return nil, errorx.NewInternalError("game.roundNotFound")
	}
	query = l.svcCtx.DB.Round.Query().Where(round.ID(*in.Id))

	result, err := query.WithFold().First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return l.Po2Vo(result), nil

}

func (l *FindRoundLogic) GetCurrentRoundId() (uint64, error) {
	currentIdCached, err := l.svcCtx.Redis.Get(l.ctx, cachekey.CurrentGameRound.Val()).Uint64()
	if err == nil {
		return currentIdCached, err
	}
	listResp, err := l.svcCtx.DB.Round.Query().Order(ent.Desc(round.FieldID)).Page(l.ctx, 1, 2)
	if err != nil {
		return 0, err
	}
	newestOne := listResp.List[0]
	nowTime := time.Now().Unix()
	// 当前时间段没有新的回合，则返回最后一回合数据
	if newestOne.EndAt.Unix() <= nowTime {
		return newestOne.ID, nil
	}
	// 查看最新的回合是否在当前时间段内
	if newestOne.StartAt.Unix() <= nowTime && newestOne.EndAt.Unix() >= nowTime {
		l.svcCtx.Redis.Set(l.ctx, cachekey.CurrentGameRound.Val(), newestOne.ID, 0)
		return newestOne.ID, nil
	}
	newestOne = listResp.List[1]
	if newestOne.StartAt.Unix() <= nowTime && newestOne.EndAt.Unix() >= nowTime {
		l.svcCtx.Redis.Set(l.ctx, cachekey.CurrentGameRound.Val(), newestOne.ID, 0)
		return newestOne.ID, nil
	}
	return 0, errorx.NewInternalError("game.roundNotFound")
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
		Id:              po.ID,
		CreatedAt:       po.CreatedAt.Unix(),
		UpdatedAt:       po.UpdatedAt.Unix(),
		StartAt:         po.StartAt.Unix(),
		OpenAt:          po.OpenAt.Unix(),
		EndAt:           po.EndAt.Unix(),
		Status:          uint32(po.Status),
		SelectedFold:    po.SelectedFold,
		TotalRoundCount: po.TotalRoundCount,
		RoundCount:      po.RoundCount,
		Folds:           folds,
	}

}
