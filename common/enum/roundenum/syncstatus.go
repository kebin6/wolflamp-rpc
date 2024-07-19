package roundenum

// SyncStatus 回传状态
type SyncStatus uint32

// 状态
const (
	Created SyncStatus = 0
	NotYet  SyncStatus = 1 // 未同步
	Success SyncStatus = 2 // 已同步
	Failed  SyncStatus = 3 // 同步失败
	NoNeed  SyncStatus = 4 // 无需同步
)

func NewSyncStatus(val uint32) SyncStatus {
	switch val {
	case 0:
		return Created
	case 1:
		return NotYet
	case 2:
		return Success
	case 3:
		return Failed
	case 4:
		return NoNeed
	default:
		return Created
	}
}

func (p SyncStatus) Val() uint32 {
	return uint32(p)
}

func (p SyncStatus) Desc() string {
	switch p {
	case Created:
		return "创建"
	case NotYet:
		return "未同步"
	case Success:
		return "已同步"
	case Failed:
		return "同步失败"
	case NoNeed:
		return "无需同步"
	default:
		return "未知"
	}
}

func (p SyncStatus) DescMap() map[SyncStatus]string {
	return map[SyncStatus]string{
		Created: "创建",
		NotYet:  "未同步",
		Success: "已同步",
		Failed:  "同步失败",
		NoNeed:  "无需同步",
	}
}
