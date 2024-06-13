package exchangeenum

// ExchangeStatus 兑换状态
type ExchangeStatus uint32

const (
	Completed ExchangeStatus = 1 // 完成
)

func (p ExchangeStatus) Val() uint32 {
	return uint32(p)
}

func NewExchangeStatus(v uint32) ExchangeStatus {
	switch v {
	case Completed.Val():
		return Completed
	default:
		return 0
	}
}

func (p ExchangeStatus) Desc() string {
	switch p {
	case Completed:
		return "Completed"
	default:
		return "Unknown Status"
	}
}

func (p ExchangeStatus) DescMap() map[ExchangeStatus]string {
	return map[ExchangeStatus]string{
		Completed: Completed.Desc(),
	}
}
