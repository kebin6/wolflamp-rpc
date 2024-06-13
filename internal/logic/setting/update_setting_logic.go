package setting

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/setting"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSettingLogic {
	return &UpdateSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSettingLogic) UpdateSetting(in *wolflamp.UpdateSettingReq) (*wolflamp.BaseIDResp, error) {

	err := l.svcCtx.DB.Setting.Update().
		Where(setting.Module(in.Module)).
		SetValue(in.JsonString).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Msg: i18n.UpdateSuccess}, nil

}
