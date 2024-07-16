package roundenum

// OpenType 开奖类型
type OpenType uint32

// 0-未开奖；1-单狼猎杀；2-金羊奖励；3-银羊奖励；4-多狼猎杀
const (
	NoOpen     OpenType = 0 // 未开奖
	SingleWolf OpenType = 1 // 单狼
	GoldenLamb OpenType = 2 // 金羊
	SliverLamb OpenType = 3 // 银羊
	MultiLamb  OpenType = 4 // 多狼
)

func NewOpenType(val uint32) OpenType {
	switch val {
	case 1:
		return SingleWolf
	case 2:
		return GoldenLamb
	case 3:
		return SliverLamb
	case 4:
		return MultiLamb
	default:
		return NoOpen
	}
}

func (p OpenType) Val() uint32 {
	return uint32(p)
}

func (p OpenType) Desc() string {
	switch p {
	case SingleWolf:
		return "单狼"
	case GoldenLamb:
		return "金羊"
	case SliverLamb:
		return "银羊"
	case MultiLamb:
		return "多狼"
	default:
		return "未开奖"
	}
}

func (p OpenType) DescMap() map[OpenType]string {
	return map[OpenType]string{
		SingleWolf: "单狼",
		GoldenLamb: "金羊",
		SliverLamb: "银羊",
		MultiLamb:  "多狼",
		NoOpen:     "未开奖",
	}
}
