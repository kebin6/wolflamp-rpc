package orderenum

// OrderStatus 订单状态
type OrderStatus uint32

const (
	Applying   OrderStatus = 1 // 申请中
	Pending    OrderStatus = 2 // 待处理
	Processing OrderStatus = 3 // 处理中
	Success    OrderStatus = 4 // 成功
	Fail       OrderStatus = 5 // 失败
	Unknown    OrderStatus = 0 // 未知
)

func (p OrderStatus) Val() uint32 {
	return uint32(p)
}

func NewOrderStatus(v uint32) OrderStatus {
	switch v {
	case Applying.Val():
		return Applying
	case Pending.Val():
		return Pending
	case Processing.Val():
		return Processing
	case Success.Val():
		return Success
	case Fail.Val():
		return Fail
	default:
		return Unknown

	}
}

func (p OrderStatus) Desc() string {
	switch p {
	case Applying:
		return "Waiting for Withdrawal"
	case Pending:
		return "Pending"
	case Processing:
		return "Processing"
	case Success:
		return "Completed"
	case Fail:
		return "Failed"
	default:
		return "Unknown Status"
	}
}

func (p OrderStatus) DescMap() map[OrderStatus]string {
	return map[OrderStatus]string{
		Applying:   Applying.Desc(),
		Pending:    Pending.Desc(),
		Processing: Processing.Desc(),
		Success:    Success.Desc(),
		Fail:       Fail.Desc(),
		Unknown:    Unknown.Desc(),
	}
}
