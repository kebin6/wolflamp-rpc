package pool

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePoolLogic {
	return &CreatePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePoolLogic) CreatePool(in *wolflamp.CreatePoolReq) (*wolflamp.BaseIDResp, error) {

	result, err := l.svcCtx.DB.Pool.Create().
		SetStatus(uint8(in.Status)).
		SetMode(in.Mode).
		SetType(in.Type).
		SetRemark(in.Remark).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &wolflamp.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil

}
