package setting

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/ent/setting"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSettingLogic {
	return &FindSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindSettingLogic) FindSetting(in *wolflamp.FindSettingReq) (*wolflamp.FindSettingResp, error) {

	result, err := l.svcCtx.DB.Setting.Query().Where(setting.Module(in.Module)).First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.FindSettingResp{
		JsonString: *result.Value,
	}, nil

}
