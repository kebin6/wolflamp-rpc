package banner

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/banner"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBannerLogic {
	return &ListBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListBannerLogic) ListBanner(in *wolflamp.ListBannerReq) (*wolflamp.ListBannerResp, error) {

	var predicates []predicate.Banner
	if in.Status != nil {
		predicates = append(predicates, banner.Status(*pointy.GetStatusPointer(in.Status)))
	}
	if in.Endpoint != nil {
		predicates = append(predicates, banner.Endpoint(*in.Endpoint))
	}
	if in.Module != nil {
		predicates = append(predicates, banner.Module(*in.Module))
	}

	result, err := l.svcCtx.DB.Banner.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &wolflamp.ListBannerResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		var file wolflamp.FileInfo
		if v.Edges.File != nil && v.Edges.File.Status == 1 {
			file = wolflamp.FileInfo{
				Id:       v.Edges.File.ID.String(),
				Name:     v.Edges.File.Name,
				Path:     v.Edges.File.Path,
				FileType: uint32(v.Edges.File.FileType),
			}
		}
		resp.Data = append(resp.Data, &wolflamp.BannerInfo{
			Id:        v.ID,
			Endpoint:  v.Endpoint,
			Module:    v.Module,
			Path:      v.Path,
			JumpUrl:   v.JumpURL,
			FileType:  uint32(v.FileType),
			Status:    uint32(v.Status),
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.UpdatedAt.Unix(),
			File:      &file,
		})
	}

	return resp, nil

}
