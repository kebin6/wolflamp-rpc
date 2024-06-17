package statement

import (
	"context"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindStatementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindStatementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindStatementLogic {
	return &FindStatementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindStatementLogic) FindStatement(in *wolflamp.FindStatementReq) (*wolflamp.StatementInfo, error) {
	// todo: add your logic here and delete this line

	return &wolflamp.StatementInfo{}, nil
}
