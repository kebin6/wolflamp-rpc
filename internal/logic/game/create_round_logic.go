package game

import (
	"context"
	"github.com/jinzhu/now"
	"github.com/kebin6/wolflamp-rpc/common/enum/roundenum"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/internal/utils/entx"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
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
	roundInfo, err := NewFindRoundLogic(l.ctx, l.svcCtx).FindRound(&wolflamp.FindRoundReq{Mode: in.Mode})
	if err != nil && err.Error() != errormsg.TargetNotFound {
		return nil, err
	}
	idStr := time.Now().Format("060102150405")
	// 游戏回合记录当天累计和历史累计
	roundCount := uint32(1)
	totalRoundCount := uint64(1)
	if roundInfo != nil {
		// 如果当天有记录，则继续累加
		// 当前回合的开始时间在今天，新增的记录开始时间在下一天，则新增记录重置
		if roundInfo.StartAt <= now.EndOfDay().Unix() && in.StartAt > now.EndOfDay().Unix() {
			roundCount = uint32(1)
		} else {
			roundCount = uint32(roundInfo.RoundCount) + 1
		}
		totalRoundCount = uint64(roundInfo.TotalRoundCount) + 1
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, err
	}

	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		if roundInfo != nil {
			err := tx.Round.UpdateOneID(roundInfo.Id).SetStatus(uint8(roundenum.Completed.Val())).Exec(l.ctx)
			if err != nil {
				return err
			}
		}
		result, err := tx.Round.Create().
			SetID(id).
			SetStatus(uint8(roundenum.Investing.Val())).
			SetStartAt(time.Unix(in.StartAt, 0)).
			SetOpenAt(time.Unix(in.OpenAt, 0)).
			SetEndAt(time.Unix(in.EndAt, 0)).
			SetRoundCount(roundCount).
			SetTotalRoundCount(totalRoundCount).
			SetMode(in.Mode).
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
				SetTotalRoundCount(totalRoundCount).
				SetFoldNo(uint32(i)).
				SetLambNum(uint32(0)).
				SetProfitAndLoss(float32(0)).SetMode(in.Mode))
		}
		err = tx.RoundLambFold.CreateBulk(folds...).Exec(l.ctx)
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
