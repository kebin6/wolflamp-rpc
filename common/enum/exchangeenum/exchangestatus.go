package exchangeenum

// ExchangeStatus 兑换状态
type ExchangeStatus uint32

const (
	Created   ExchangeStatus = 0 // 创建
	Completed ExchangeStatus = 1 // 完成
	Failed    ExchangeStatus = 2 // 失败
)

func (p ExchangeStatus) Val() uint32 {
	return uint32(p)
}

func NewExchangeStatus(v uint32) ExchangeStatus {
	switch v {
	case Completed.Val():
		return Completed
	case Created.Val():
		return Created
	case Failed.Val():
		return Failed
	default:
		return 0
	}
}

func (p ExchangeStatus) Desc() string {
	switch p {
	case Completed:
		return "Completed"
	case Created:
		return "Created"
	case Failed:
		return "Failed"
	default:
		return "Unknown Status"
	}
}

func (p ExchangeStatus) DescMap() map[ExchangeStatus]string {
	return map[ExchangeStatus]string{
		Completed: Completed.Desc(),
	}
}
