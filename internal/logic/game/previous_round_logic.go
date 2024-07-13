package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/enum/cachekey"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"time"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type PreviousRoundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPreviousRoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreviousRoundLogic {
	return &PreviousRoundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PreviousRoundLogic) PreviousRound(in *wolflamp.PreviousRoundReq) (*wolflamp.PreviousRoundResp, error) {

	previousFoldNo, err := l.svcCtx.Redis.Get(l.ctx, cachekey.PreviousSelectedLambFoldNo.ModeVal(in.Mode)).Uint64()
	if err != nil {
		currentRound, err := NewFindRoundLogic(l.ctx, l.svcCtx).FindRound(&wolflamp.FindRoundReq{Mode: in.Mode})
		if err != nil {
			return nil, err
		}
		previousRound, err := l.svcCtx.DB.Round.Query().
			Where(round.Mode(in.Mode)).
			Where(round.IDLT(currentRound.Id)).Order(ent.Desc(round.FieldID)).First(l.ctx)
		if err != nil {
			previousFoldNo = 0
		} else {
			previousFoldNo = uint64(previousRound.SelectedFold)
			l.svcCtx.Redis.Set(l.ctx, cachekey.PreviousSelectedLambFoldNo.ModeVal(in.Mode), previousFoldNo, time.Second*600)
		}
	}
	return &wolflamp.PreviousRoundResp{
		SelectedFoldNo: uint32(previousFoldNo),
	}, nil
}
