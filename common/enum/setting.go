package enum

// SettingModule 平台配置项所属模块
type SettingModule string

const (
	PlatformSetting SettingModule = "platform_setting"
	PlatformNotice  SettingModule = "platform_notice"
)

func (p SettingModule) Val() string {
	return string(p)
}

func (p SettingModule) Desc() string {
	switch p {
	case PlatformSetting:
		return "平台配置"
	case PlatformNotice:
		return "平台公告"
	default:
		return ""
	}
}

func (p SettingModule) DescMap() map[SettingModule]string {
	return map[SettingModule]string{
		PlatformSetting: PlatformSetting.Desc(),
		PlatformNotice:  PlatformNotice.Desc(),
	}
}

// SettingKey 平台配置项代码
type SettingKey string

const (
	GameRule             SettingKey = "game_rule"
	WithdrawCommission   SettingKey = "withdraw_commission"
	MinWithdrawNum       SettingKey = "min_withdraw_num"
	RobotNum             SettingKey = "robot_num"
	RobotLampNum         SettingKey = "robot_lamp_num"
	IdleTime             SettingKey = "idle_time"
	WithdrawThreshold    SettingKey = "withdraw_threshold"
	GameCommission       SettingKey = "game_commission"
	PoolCommission       SettingKey = "pool_commission"
	RobPoolCommission    SettingKey = "rob_pool_commission"
	RewardPoolCommission SettingKey = "reward_pool_commission"
	GoldenLambAllowTime  SettingKey = "golden_lamb_allow_time"
	GoldenLambNum        SettingKey = "golden_lamb_num"
	PoolMinNumThenSilver SettingKey = "pool_min_num_then_silver"
	SliverOccurPercent   SettingKey = "sliver_occur_percent"
	SliverLambNum        SettingKey = "sliver_lamb_num"
)

func (p SettingKey) Val() string {
	return string(p)
}

func (p SettingKey) Desc() string {
	switch p {
	case GameRule:
		return "游戏规则"
	case WithdrawCommission:
		return "提币手续费(%)"
	case MinWithdrawNum:
		return "最小提币金额"
	case RobotNum:
		return "机器人数量"
	case RobotLampNum:
		return "机器人投入羊数量"
	case IdleTime:
		return "机器人空闲秒数"
	case WithdrawThreshold:
		return "提币免审阈值"
	case GameCommission:
		return "平台抽成比例(%)"
	case PoolCommission:
		return "资金池抽成比例(%)"
	case RobPoolCommission:
		return "机器人池抽成比例(%)"
	case RewardPoolCommission:
		return "奖励池抽成比例(%)"
	case GoldenLambAllowTime:
		return "金羊可投时间区间"
	case GoldenLambNum:
		return "金羊触发数量"
	case PoolMinNumThenSilver:
		return "资金池满足N后触发银羊出现"
	case SliverOccurPercent:
		return "银羊出现概率(%)"
	case SliverLambNum:
		return "银羊触发数量"
	default:
		return "未知配置项"
	}
}

func (p SettingKey) DescMap() map[SettingKey]string {
	return map[SettingKey]string{
		GameRule:             GameRule.Desc(),
		WithdrawCommission:   WithdrawCommission.Desc(),
		MinWithdrawNum:       MinWithdrawNum.Desc(),
		RobotNum:             RobotNum.Desc(),
		RobotLampNum:         RobotLampNum.Desc(),
		IdleTime:             IdleTime.Desc(),
		WithdrawThreshold:    WithdrawThreshold.Desc(),
		GameCommission:       GameCommission.Desc(),
		PoolCommission:       PoolCommission.Desc(),
		RobPoolCommission:    RobPoolCommission.Desc(),
		RewardPoolCommission: RewardPoolCommission.Desc(),
		GoldenLambAllowTime:  GoldenLambAllowTime.Desc(),
		GoldenLambNum:        GoldenLambNum.Desc(),
		PoolMinNumThenSilver: PoolMinNumThenSilver.Desc(),
		SliverOccurPercent:   SliverOccurPercent.Desc(),
		SliverLambNum:        SliverLambNum.Desc(),
	}
}

func (p SettingKey) CacheKey() string {
	return "cache_setting:" + p.Val()
}
