package player

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/common/util"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateGcicsSignLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateGcicsSignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateGcicsSignLogic {
	return &ValidateGcicsSignLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateGcicsSignLogic) ValidateGcicsSign(in *wolflamp.ValidateGcicsSignReq) (*wolflamp.ValidateGcicsSignResp, error) {
	isValid := util.ValidateSignature(in.Timestamp, l.svcCtx.Config.GcicsConf.AppId, l.svcCtx.Config.GcicsConf.AppSecret, in.Sign)
	return &wolflamp.ValidateGcicsSignResp{
		IsValid: isValid,
	}, nil
}
