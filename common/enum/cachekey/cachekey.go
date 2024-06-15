package cachekey

// CacheKey 缓存键
type CacheKey string

const (
	CurrentGameRound           CacheKey = "current_game:round_id"          // 游戏当前轮次ID
	CurrentGameRobotNum        CacheKey = "current_game:robot_num"         // 游戏当前机器人数量
	CurrentGameLastRobotTime   CacheKey = "current_game:last_robot_time"   // 最近机器人投入时间
	CurrentGameRobotLambNum    CacheKey = "current_game:robot_lamb_num"    // 机器人投入羊数量
	PreviousSelectedLambFoldNo CacheKey = "current_game:previous_selected" // 上一局被选中的羊圈
)

func (p CacheKey) Val() string {
	return string(p)
}
