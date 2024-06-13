package rewardenum

// RewardStatus 奖励状态
type RewardStatus uint32

const (
	Completed RewardStatus = 1 // 处理完成
	Unknown   RewardStatus = 0
)

func New(val uint32) RewardStatus {
	switch val {
	case Completed.Val():
		return Completed
	default:
		return Unknown
	}
}

func (p RewardStatus) Val() uint32 {
	return uint32(p)
}

func (p RewardStatus) Desc() string {
	switch p {
	case Completed:
		return "Completed"
	default:
		return "Unknown Status"
	}
}

func (p RewardStatus) DescMap() map[RewardStatus]string {
	return map[RewardStatus]string{
		Completed: Completed.Desc(),
		Unknown:   Unknown.Desc(),
	}
}
