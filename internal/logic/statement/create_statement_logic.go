package statement

import (
	"context"
	"github.com/duke-git/lancet/v2/random"
	"github.com/kebin6/wolflamp-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-common/i18n"
	"strconv"
	"time"

	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStatementLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStatementLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStatementLogic {
	return &CreateStatementLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStatementLogic) CreateStatement(in *wolflamp.CreateStatementReq) (*wolflamp.BaseIDResp, error) {

	prefix := ""
	if in.Prefix != nil {
		prefix = *in.Prefix
	}
	// 账单号
	code := prefix + strconv.FormatInt(time.Now().UnixMilli()*100+int64(random.RandInt(100, 999)), 10)
	result, err := l.svcCtx.DB.Statement.Create().
		SetStatus(uint8(in.Status)).
		SetCode(code).
		SetAmount(in.Amount).
		SetPlayerID(in.PlayerId).
		SetModule(in.Module).
		SetInoutType(in.InoutType).
		SetReferID(in.ReferId).
		SetNillableRemark(in.Remark).
		SetMode(in.Mode).
		Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &wolflamp.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil

}
