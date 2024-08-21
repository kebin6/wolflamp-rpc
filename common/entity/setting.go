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
	// 平台抽佣比例（%）
	GameCommission float32 `json:"gameCommission"`
	// 资金池预留比例（%）
	PoolCommission float32 `json:"poolCommission"`
	// 机器人池预留占资金池比例（%）
	RobPoolCommission float32 `json:"robPoolCommission"`
	// 奖金池预留占资金池比例（%）
	RewardPoolCommission float32 `json:"rewardPoolCommission"`
	// 金羊出现时间段
	GoldenLambAllowTime GoldenLambAllowTime `json:"goldenLambAllowTime"`
	// 金羊数量区间设定
	GoldenLambNum GoldenLambNum `json:"goldenLambNum"`
	// 银羊触发条件-奖金池最低阈值
	PoolMinNumThenSilver uint32 `json:"poolMinNumThenSilver"`
	// 银羊触发概率（%）
	SliverOccurPercent float32 `json:"sliverOccurPercent"`
	// 银羊数量区间设定
	SliverLambNum GoldenLambNum `json:"sliverLambNum"`
	// 银羊触发数量比例
	SliverLambPercent float32 `json:"sliverLambPercent"`
	// 金羊触发数量比例
	GoldenLambPercent float32 `json:"goldenLambPercent"`
}

type GoldenLambNum struct {
	Min uint32 `json:"min"`
	Max uint32 `json:"max"`
}

type GoldenLambAllowTime struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
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
