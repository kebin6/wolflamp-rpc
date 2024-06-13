package orderenum

// OrderType 订单类型
type OrderType string

const (
	Deposit  OrderType = "deposit"  // 充值订单
	Withdraw OrderType = "withdraw" // 提现订单
)

func (p OrderType) Val() string {
	return string(p)
}

func (p OrderType) NewOrderType(v string) OrderType {
	switch v {
	case Deposit.Val():
		return Deposit
	case Withdraw.Val():
		return Withdraw
	default:
		return ""
	}
}

func (p OrderType) Desc() string {
	switch p {
	case Deposit:
		return "Deposit"
	case Withdraw:
		return "Withdraw"
	default:
		return "Unknown Status"
	}
}

func (p OrderType) DescMap() map[OrderType]string {
	return map[OrderType]string{
		Deposit:  Deposit.Desc(),
		Withdraw: Withdraw.Desc(),
	}
}
