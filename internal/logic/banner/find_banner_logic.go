package banner

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindBannerLogic {
	return &FindBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindBannerLogic) FindBanner(in *wolflamp.FindBannerReq) (*wolflamp.BannerInfo, error) {

	result, err := l.svcCtx.DB.Banner.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	var file wolflamp.FileInfo
	if result.Edges.File != nil && result.Edges.File.Status == 1 {
		file = wolflamp.FileInfo{
			Id:       result.Edges.File.ID.String(),
			Name:     result.Edges.File.Name,
			Path:     result.Edges.File.Path,
			FileType: uint32(result.Edges.File.FileType),
		}
	}
	return &wolflamp.BannerInfo{
		Id:        result.ID,
		CreatedAt: result.CreatedAt.Unix(),
		UpdatedAt: result.UpdatedAt.Unix(),
		Status:    uint32(result.Status),
		Endpoint:  result.Endpoint,
		Module:    result.Module,
		FileType:  uint32(result.FileType),
		Path:      result.Path,
		JumpUrl:   result.JumpURL,
		File:      &file,
	}, nil

}
