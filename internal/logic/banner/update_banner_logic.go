package banner

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBannerLogic {
	return &UpdateBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBannerLogic) UpdateBanner(in *wolflamp.UpdateBannerReq) (*wolflamp.BaseIDResp, error) {

	err := l.svcCtx.DB.Banner.UpdateOneID(in.Id).
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilPath(in.Path).
		SetNotNilFileType(pointy.GetStatusPointer(in.FileType)).
		SetNotNilModule(in.Module).
		SetNotNilEndpoint(in.Endpoint).
		SetNotNilJumpURL(in.JumpUrl).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Msg: i18n.UpdateSuccess}, nil

}
