package statement

import (
	"context"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStatementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStatementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatementLogic {
	return &UpdateStatementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStatementLogic) UpdateStatement(in *wolflamp.UpdateStatementReq) (*wolflamp.BaseIDResp, error) {
	// todo: add your logic here and delete this line

	return &wolflamp.BaseIDResp{}, nil
}
