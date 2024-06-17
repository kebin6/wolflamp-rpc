package statement

import (
	"context"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListStatementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListStatementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStatementLogic {
	return &ListStatementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListStatementLogic) ListStatement(in *wolflamp.ListStatementReq) (*wolflamp.ListStatementResp, error) {
	// todo: add your logic here and delete this line

	return &wolflamp.ListStatementResp{}, nil
}
