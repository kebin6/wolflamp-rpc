package roundenum

// RoundStatus 奖励状态
type RoundStatus uint32

// 状态: 1-投注中 2-开奖中 3-已结束
const (
	Investing RoundStatus = 1 // 投注中
	Opening   RoundStatus = 2 // 开奖中
	Completed RoundStatus = 3 // 已结束
	Unknown   RoundStatus = 0
)

func New(val uint32) RoundStatus {
	switch val {
	case Investing.Val():
		return Investing
	case Opening.Val():
		return Opening
	case Completed.Val():
		return Completed
	default:
		return Unknown
	}
}

func (p RoundStatus) Val() uint32 {
	return uint32(p)
}

func (p RoundStatus) Desc() string {
	switch p {
	case Investing:
		return "Investing"
	case Opening:
		return "Opening"
	case Completed:
		return "Completed"
	default:
		return "Unknown Status"
	}
}

func (p RoundStatus) DescMap() map[RoundStatus]string {
	return map[RoundStatus]string{
		Investing: Investing.Desc(),
		Opening:   Opening.Desc(),
		Completed: Completed.Desc(),
		Unknown:   Unknown.Desc(),
	}
}
