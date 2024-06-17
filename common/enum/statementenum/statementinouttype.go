package statementenum

// InOutType 收支类型
type InOutType uint32

const (
	Income      InOutType = 1 // 入账
	Expense     InOutType = 2 // 出账
	UnknownType InOutType = 0
)

func NewInOutType(val uint32) InOutType {
	switch val {
	case Income.Val():
		return Income
	case Expense.Val():
		return Expense
	default:
		return UnknownType
	}
}

func (p InOutType) Val() uint32 {
	return uint32(p)
}

func (p InOutType) Desc() string {
	switch p {
	case Income:
		return "Income"
	case Expense:
		return "Expense"
	default:
		return "Unknown"
	}
}

func (p InOutType) DescMap() map[InOutType]string {
	return map[InOutType]string{
		Income:      Income.Desc(),
		Expense:     Expense.Desc(),
		UnknownType: UnknownType.Desc(),
	}
}
