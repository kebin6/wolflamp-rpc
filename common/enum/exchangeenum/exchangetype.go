package exchangeenum

// ExchangeType 兑换类型
type ExchangeType uint32

const (
	CoinToLamb ExchangeType = 1 // 币兑羊
	LambToCoin ExchangeType = 2 // 羊兑币
	Unknown    ExchangeType = 0
)

func (p ExchangeType) Val() uint32 {
	return uint32(p)
}

func NewExchangeType(v uint32) ExchangeType {
	switch v {
	case CoinToLamb.Val():
		return CoinToLamb
	case LambToCoin.Val():
		return LambToCoin
	default:
		return Unknown
	}
}

func (p ExchangeType) Desc() string {
	switch p {
	case CoinToLamb:
		return "BNB/LAMB"
	case LambToCoin:
		return "LAMB/BNB"
	default:
		return "Unknown Status"
	}
}

func (p ExchangeType) DescMap() map[ExchangeType]string {
	return map[ExchangeType]string{
		CoinToLamb: CoinToLamb.Desc(),
		LambToCoin: LambToCoin.Desc(),
	}
}
