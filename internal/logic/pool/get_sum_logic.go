package pool

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/pool"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSumLogic {
	return &GetSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSumLogic) GetSum(in *wolflamp.GetSumReq) (*wolflamp.GetSumResp, error) {
	// TODO:: 待优化
	amount, _ := l.svcCtx.DB.Pool.Query().
		Where(pool.Mode(in.Mode)).
		Where(pool.Status(uint8(in.Status))).
		Where(pool.Type(in.Type)).
		Aggregate(ent.Sum(pool.FieldLambNum)).
		Float64(l.ctx)
	return &wolflamp.GetSumResp{
		Amount: amount,
	}, nil

}
