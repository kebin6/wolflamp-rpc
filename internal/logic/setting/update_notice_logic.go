package setting

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/setting"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeLogic {
	return &UpdateNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateNoticeLogic) UpdateNotice(in *wolflamp.UpdateNoticeReq) (*wolflamp.BaseIDResp, error) {
	err := l.svcCtx.DB.Setting.Update().
		Where(setting.Module(in.Notice)).
		SetValue(in.Notice).
		Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &wolflamp.BaseIDResp{}, nil
}
