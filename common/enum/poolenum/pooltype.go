package poolenum

// PoolType 资金池记录类型
type PoolType uint32

const (
	Robot   PoolType = 1 // 机器人池
	Reward  PoolType = 2 // 奖励池
	Other   PoolType = 3 // 其他
	Unknown PoolType = 0
)

func New(val uint32) PoolType {
	switch val {
	case Robot.Val():
		return Robot
	case Reward.Val():
		return Reward
	case Other.Val():
		return Other
	default:
		return Unknown
	}
}

func (p PoolType) Val() uint32 {
	return uint32(p)
}

func (p PoolType) Desc() string {
	switch p {
	case Robot:
		return "robot"
	case Reward:
		return "reward"
	case Other:
		return "other"
	default:
		return "Unknown Status"
	}
}

func (p PoolType) DescMap() map[PoolType]string {
	return map[PoolType]string{
		Robot:   Robot.Desc(),
		Reward:  Reward.Desc(),
		Other:   Other.Desc(),
		Unknown: Unknown.Desc(),
	}
}
