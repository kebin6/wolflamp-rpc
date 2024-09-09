package pool

import (
	"context"
	"fmt"
	"github.com/jinzhu/now"
	"github.com/kebin6/wolflamp-rpc/ent"
	"github.com/kebin6/wolflamp-rpc/ent/pool"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"strings"
	"time"
)

type GetSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSumLogic {
	return &GetSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSumLogic) GetSum(in *wolflamp.GetSumReq) (*wolflamp.GetSumResp, error) {
	amountCached, err := l.GetLambSumCache(in)
	if err != nil {
		return nil, err
	}
	amount, err := l.svcCtx.DB.Pool.Query().
		Where(pool.Mode(in.Mode)).
		Where(pool.Status(uint8(in.Status))).
		Where(pool.Type(in.Type)).
		Where(pool.CreatedAtGTE(now.BeginningOfDay())).
		Aggregate(ent.Sum(pool.FieldLambNum)).
		Float64(l.ctx)

	if err != nil {
		if !strings.Contains(err.Error(), "converting NULL to float64 is unsupported") {
			return nil, err
		}
	}
	return &wolflamp.GetSumResp{
		Amount: amount + amountCached,
	}, nil

}

// GetLambSumCache 获取mode模式下，截止至昨天的总机器人池剩余数量 /*
func (l *GetSumLogic) GetLambSumCache(in *wolflamp.GetSumReq) (float64, error) {
	nowTime := now.BeginningOfDay().In(time.UTC)

	//yesterday := nowTime.AddDate(0, 0, -1)
	format := "2006-01-02"
	formattedToday := nowTime.Format(format)
	todayCacheKey := fmt.Sprintf("%s_%d_%d_lamb_pool_sum:%s", in.Mode, in.Type, in.Status, formattedToday)

	// 如果有昨天的统计缓存，则直接返回
	if l.svcCtx.Redis.Exists(l.ctx, todayCacheKey).Val() != 0 {
		cached, err := l.svcCtx.Redis.Get(l.ctx, todayCacheKey).Result()
		if err == nil {
			f, err := strconv.ParseFloat(cached, 64)
			if err == nil {
				return f, nil
			}
		}
	}

	// 获取前天的统计缓存，有则统计昨天的数据并进行累加，缓存
	yesterday := nowTime.AddDate(0, 0, -1)
	formattedYesterday := yesterday.Format(format)
	cacheKey := fmt.Sprintf("%s_%d_%d_lamb_pool_sum:%s", in.Mode, in.Type, in.Status, formattedYesterday)
	if l.svcCtx.Redis.Exists(l.ctx, cacheKey).Val() != 0 {
		cached, err := l.svcCtx.Redis.Get(l.ctx, cacheKey).Result()
		if err == nil {
			f, err := strconv.ParseFloat(cached, 64)
			if err == nil {
				// 截止到前天有缓存，则查询昨天的数据进行累加
				amount, err := l.svcCtx.DB.Pool.Query().
					Where(pool.Mode(in.Mode)).
					Where(pool.Status(uint8(in.Status))).
					Where(pool.Type(in.Type)).
					Where(pool.CreatedAtLT(nowTime)).
					Where(pool.CreatedAtGTE(yesterday)).
					Aggregate(ent.Sum(pool.FieldLambNum)).
					Float64(l.ctx)
				if err == nil {
					amount += f
					_ = l.svcCtx.Redis.Set(l.ctx, todayCacheKey, amount, time.Hour*24*2)
					return amount, err
				}
			}
		}
	}
	amount, err := l.svcCtx.DB.Pool.Query().
		Where(pool.Mode(in.Mode)).
		Where(pool.Status(uint8(in.Status))).
		Where(pool.Type(in.Type)).
		Where(pool.CreatedAtLT(nowTime)).
		Aggregate(ent.Sum(pool.FieldLambNum)).
		Float64(l.ctx)
	if err != nil {
		if !strings.Contains(err.Error(), "converting NULL to float64 is unsupported") {
			return 0, err
		}
		return 0, nil
	}
	// 保存3天
	_ = l.svcCtx.Redis.Set(l.ctx, todayCacheKey, amount, time.Hour*24*2)
	return amount, nil
}
