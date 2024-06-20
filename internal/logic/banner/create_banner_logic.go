package banner

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBannerLogic {
	return &CreateBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBannerLogic) CreateBanner(in *wolflamp.CreateBannerReq) (*wolflamp.BaseIDResp, error) {

	result, err := l.svcCtx.DB.Banner.Create().
		SetNotNilStatus(pointy.GetStatusPointer(&in.Status)).
		SetNotNilModule(pointy.GetPointer(in.Module)).
		SetNotNilPath(pointy.GetPointer(in.Path)).
		SetNotNilEndpoint(pointy.GetPointer(in.Endpoint)).
		SetNotNilJumpURL(pointy.GetPointer(in.JumpUrl)).
		SetNotNilFileType(pointy.GetPointer(uint8(in.FileType))).
		SetFileID(uuidx.ParseUUIDString(in.FileId)).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil

}
