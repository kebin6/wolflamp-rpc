package statementenum

// StatementModule 账单所属模块
type StatementModule uint32

const (
	Invest        StatementModule = 1 // 游戏
	Reward        StatementModule = 2 // 佣金奖励
	System        StatementModule = 3 // 系统奖励
	Pool          StatementModule = 4 // 资金池收益
	UnknownModule StatementModule = 0
)

func NewStatementModule(val uint32) StatementModule {
	switch val {
	case Invest.Val():
		return Invest
	case Reward.Val():
		return Reward
	case System.Val():
		return System
	case Pool.Val():
		return Pool
	default:
		return UnknownModule
	}
}

func (p StatementModule) Val() uint32 {
	return uint32(p)
}

func (p StatementModule) Desc() string {
	switch p {
	case Invest:
		return "Invest"
	case Reward:
		return "Reward"
	case System:
		return "System"
	case Pool:
		return "Pool"
	default:
		return "Unknown"
	}
}

func (p StatementModule) DescMap() map[StatementModule]string {
	return map[StatementModule]string{
		Invest:        Invest.Desc(),
		Reward:        Reward.Desc(),
		System:        System.Desc(),
		Pool:          Pool.Desc(),
		UnknownModule: UnknownModule.Desc(),
	}
}
