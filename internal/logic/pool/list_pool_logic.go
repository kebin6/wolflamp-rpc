package pool

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/pool"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPoolLogic {
	return &ListPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPoolLogic) ListPool(in *wolflamp.ListPoolReq) (*wolflamp.ListPoolResp, error) {

	var predicates []predicate.Pool
	if in.Status != nil {
		predicates = append(predicates, pool.Status(*pointy.GetStatusPointer(in.Status)))
	}
	if in.Mode != nil {
		predicates = append(predicates, pool.Mode(*in.Mode))
	}
	if in.Type != nil {
		predicates = append(predicates, pool.Type(*in.Type))
	}
	if in.Remark != nil {
		predicates = append(predicates, pool.RemarkContains(*in.Remark))
	}

	result, err := l.svcCtx.DB.Pool.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	playerIds := make([]uint64, result.PageDetails.Total)
	for i, v := range result.List {
		playerIds[i] = v.ID
	}
	resp := &wolflamp.ListPoolResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &wolflamp.PoolInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.UpdatedAt.Unix(),
			Mode:      v.Mode,
			Status:    uint32(v.Status),
			Type:      v.Type,
			LambNum:   v.LambNum,
			Remark:    v.Remark,
		})
	}
	return resp, nil

}
