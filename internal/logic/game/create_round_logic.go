package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/enum/roundenum"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/internal/utils/entx"
	"github.com/suyuan32/simple-admin-common/i18n"
	"strconv"
	"time"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoundLogic {
	return &CreateRoundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoundLogic) CreateRound(in *wolflamp.CreateRoundReq) (*wolflamp.BaseIDResp, error) {

	// 获取当前轮次信息
	roundInfo, err := l.svcCtx.DB.Round.Query().Order(ent.Desc(round.FieldID)).First(l.ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	idStr := time.Now().Format("060102150405")
	roundCount := uint32(1)
	if roundInfo != nil {
		roundCount = uint32(roundInfo.RoundCount) + 1
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, err
	}

	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		result, err := l.svcCtx.DB.Round.Create().
			SetID(id).
			SetStatus(uint8(roundenum.Investing.Val())).
			SetStartAt(time.Unix(in.StartAt, 0)).
			SetOpenAt(time.Unix(in.OpenAt, 0)).
			SetEndAt(time.Unix(in.EndAt, 0)).
			SetRoundCount(roundCount).
			Save(l.ctx)
		if err != nil {
			return err
		}
		id = result.ID

		var folds []*ent.RoundLambFoldCreate
		for i := 1; i <= 8; i++ {
			folds = append(folds, l.svcCtx.DB.RoundLambFold.Create().
				SetRoundID(id).
				SetRoundCount(roundCount).
				SetFoldNo(uint32(i)).
				SetLambNum(uint32(0)).
				SetProfitAndLoss(float32(0)))
		}
		err = l.svcCtx.DB.RoundLambFold.CreateBulk(folds...).Exec(l.ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Id: id, Msg: i18n.CreateSuccess}, nil

}
