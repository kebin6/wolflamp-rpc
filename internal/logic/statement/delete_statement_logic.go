package statement

import (
	"context"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStatementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStatementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStatementLogic {
	return &DeleteStatementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStatementLogic) DeleteStatement(in *wolflamp.DeleteStatementReq) (*wolflamp.BaseIDResp, error) {
	// todo: add your logic here and delete this line

	return &wolflamp.BaseIDResp{}, nil
}
