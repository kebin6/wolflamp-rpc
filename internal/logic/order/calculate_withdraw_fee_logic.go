package order

import (
	"context"
	"github.com/kebin6/wolflamp-rpc/internal/logic/setting"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/kebin6/wolflamp-rpc/types/wolflamp"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type CalculateWithdrawFeeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCalculateWithdrawFeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalculateWithdrawFeeLogic {
	return &CalculateWithdrawFeeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CalculateWithdrawFeeLogic) CalculateWithdrawFee(in *wolflamp.CalculateWithdrawFeeReq) (*wolflamp.CalculateWithdrawFeeResp, error) {

	if in.Balance < in.Amount {
		return nil, errorx.NewInvalidArgumentError("wallet.insufficientBalance")
	}
	// 计算手续费
	commissionConfig, err := setting.NewGetWithdrawCommissionLogic(l.ctx, l.svcCtx).GetWithdrawCommission(&wolflamp.Empty{})
	if err != nil {
		return nil, err
	}
	handingFee := in.Amount * float64(commissionConfig.WithdrawCommission) / 100

	// 计算总金额
	receivedAmount := in.Amount
	totalAmount := in.Amount + handingFee
	if in.Balance < totalAmount {
		receivedAmount = in.Balance - handingFee
	}
	totalAmount = receivedAmount + handingFee

	return &wolflamp.CalculateWithdrawFeeResp{
		ReceivedAmount: receivedAmount,
		HandlingFee:    handingFee,
		TotalAmount:    totalAmount,
	}, nil
}
