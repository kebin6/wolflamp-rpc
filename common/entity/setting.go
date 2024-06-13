package entity

type PlatformSetting struct {
	// 游戏规则
	GameRule GameRule `json:"gameRule"`
	// 机器人数量
	RobotNum RobotNum `json:"robotNum"`
	// 机器人投入羊数量
	RobotLampNum RobotLampNum `json:"robotLampNum"`
	// 提现手续费
	WithdrawCommission float32 `json:"withdrawCommission"`
	// 最小提现数量
	MinWithdrawNum float32 `json:"minWithdrawNum"`
	// 提现免审阈值
	WithdrawThreshold float32 `json:"withdrawThreshold"`
	// 机器人空闲秒数
	IdleTime uint32 `json:"idleTime"`
}

// GameRule 游戏规则
type GameRule struct {
	// 标题
	Title string `json:"title"`
	// 规则内容
	Content string `json:"content"`
}

// RobotNum 机器人数量
type RobotNum struct {
	Min uint32 `json:"min"`
	Max uint32 `json:"max"`
}

// RobotLampNum 机器人投入羊数量
type RobotLampNum struct {
	Min uint32 `json:"min"`
	Max uint32 `json:"max"`
}
