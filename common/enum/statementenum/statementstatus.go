package statementenum

// StatementStatus 账单状态
type StatementStatus uint32

// 状态: 1-已完成
const (
	Completed StatementStatus = 1 // 已完成
	Unknown   StatementStatus = 0
)

func New(val uint32) StatementStatus {
	switch val {
	case Completed.Val():
		return Completed
	default:
		return Unknown
	}
}

func (p StatementStatus) Val() uint32 {
	return uint32(p)
}

func (p StatementStatus) Desc() string {
	switch p {
	case Completed:
		return "Completed"
	default:
		return "Unknown Status"
	}
}

func (p StatementStatus) DescMap() map[StatementStatus]string {
	return map[StatementStatus]string{
		Completed: Completed.Desc(),
		Unknown:   Unknown.Desc(),
	}
}
