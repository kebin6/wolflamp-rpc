package base

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/msg/logmsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *wolflamp.Empty) (*wolflamp.BaseResp, error) {
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewInternalError(err.Error())
	}

	if l.svcCtx.Config.CoreRpc.Enabled {
		err := l.insertApiData()
		if err != nil {
			return nil, err
		}

		err = l.insertMenuData()
		if err != nil {
			return nil, err
		}

		err = l.insertSettingData()
		if err != nil {
			return nil, err
		}
	}

	return &wolflamp.BaseResp{Msg: errormsg.Success}, nil
}
