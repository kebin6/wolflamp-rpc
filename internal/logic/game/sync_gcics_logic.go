package game

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/api"
	"github.com/kebin6/wolflamp-rpc/common/enum/roundenum"
	"github.com/kebin6/wolflamp-rpc/common/util"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/round"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncGcicsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncGcicsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncGcicsLogic {
	return &SyncGcicsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SyncGcicsLogic) SyncGcics(in *wolflamp.SyncGcicsReq) (*wolflamp.BaseIDResp, error) {

	syncList, err := l.svcCtx.DB.Round.Query().
		WithInvest().
		Where(round.SyncStatus(roundenum.NotYet.Val())).
		Limit(int(in.ChunkSize)).
		Select(round.FieldID, round.FieldSyncStatus, round.FieldComputeAmount).
		All(l.ctx)
	if err != nil {
		return nil, err
	}

	if len(syncList) == 0 {
		return &wolflamp.BaseIDResp{
			Msg: "No Data, Exist",
		}, nil
	}

	for _, syncItem := range syncList {
		err := l.ProcessSync(syncItem)
		if err != nil {
			return nil, err
		}
	}
	return &wolflamp.BaseIDResp{
		Msg: "Success",
	}, nil

}

func (l *SyncGcicsLogic) ProcessSync(syncItem *ent.Round) error {
	commissonInfoList := make([]api.CommissionInfo, 0)
	investList := syncItem.Edges.Invest
	for _, invest := range investList {
		if util.IsRealPlayer(invest.PlayerID) {
			player, err := l.svcCtx.DB.Player.Get(l.ctx, invest.PlayerID)
			if err != nil {
				return err
			}
			commissonInfoList = append(commissonInfoList, api.CommissionInfo{
				LambAmount: float64(invest.LambNum),
				UserId:     player.GcicsUserID,
			})
		}
	}
	if len(commissonInfoList) == 0 {
		err := l.svcCtx.DB.Round.UpdateOneID(syncItem.ID).SetSyncStatus(roundenum.NoNeed.Val()).Exec(l.ctx)
		if err != nil {
			return err
		}
		return nil
	}
	gcicsApi := api.GcicsApi{
		Ctx:    l.ctx,
		SvcCtx: l.svcCtx,
	}
	err := gcicsApi.Commission(syncItem.Mode, syncItem.ComputeAmount, commissonInfoList)
	if err != nil {
		return err
	}
	return nil
}
