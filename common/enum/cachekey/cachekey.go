package cachekey

import "fmt"

// CacheKey 缓存键
type CacheKey string

const (
	GameAuthToken              CacheKey = "game_auth_token:%d"             // 游戏登陆认证token
	CurrentGameRound           CacheKey = "current_game:round_id"          // 游戏当前轮次ID
	CurrentGameRobotNum        CacheKey = "current_game:robot_num"         // 游戏当前机器人数量
	CurrentGameLastRobotTime   CacheKey = "current_game:last_robot_time"   // 最近机器人投入时间
	CurrentGameRobotLambNum    CacheKey = "current_game:robot_lamb_num"    // 机器人投入羊数量
	PreviousSelectedLambFoldNo CacheKey = "current_game:previous_selected" // 上一局被选中的羊圈
	TodayGoldenLambLock        CacheKey = "today_golden_lamb_lock"         // 记录今日是否出现过黄金羊
)

func (p CacheKey) Val() string {
	return string(p)
}

func (p CacheKey) ModeVal(mode string) string {
	return fmt.Sprintf("%s_%s", mode, string(p))
}
