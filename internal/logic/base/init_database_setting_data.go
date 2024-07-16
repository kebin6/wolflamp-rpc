package base

import (
	"encoding/json"
	"github.com/kebin6/wolflamp-rpc/common/entity"
	"github.com/kebin6/wolflamp-rpc/common/enum"
)

// insert initial api data
func (l *InitDatabaseLogic) insertSettingData() error {
	settingListResp, err := l.svcCtx.DB.Setting.Query().All(l.ctx)
	if err != nil {
		return err
	}

	var settingList map[string]uint64
	if len(settingListResp) > 0 {
		settingList = make(map[string]uint64, len(settingListResp))
		for _, setting := range settingListResp {
			settingList[setting.Module] = setting.ID
		}
	}

	// 平台配置
	if _, ok := settingList[enum.PlatformSetting.Val()]; !ok {
		setting := &entity.PlatformSetting{
			WithdrawCommission: 0,
			MinWithdrawNum:     0,
			WithdrawThreshold:  0,
			RobotLampNum: entity.RobotLampNum{
				Max: 0,
				Min: 0,
			},
			RobotNum: entity.RobotNum{
				Max: 0,
				Min: 0,
			},
			IdleTime: 0,
			GameRule: entity.GameRule{
				Title:   "",
				Content: "",
			},
			GameCommission:       3,  // 默认3个点
			PoolCommission:       3,  // 默认3个点
			RobPoolCommission:    50, // 默认占资金池50%
			RewardPoolCommission: 50, // 默认占资金池50%
			GoldenLambAllowTime: entity.GoldenLambAllowTime{
				StartTime: "",
				EndTime:   "",
			},
			GoldenLambNum: entity.GoldenLambNum{
				Max: 0,
				Min: 0,
			},
			PoolMinNumThenSilver: 0,
			SliverOccurPercent:   0,
			SliverLambNum: entity.GoldenLambNum{
				Max: 0,
				Min: 0,
			},
		}

		gameRuleJsonByte, err := json.Marshal(setting)
		if err != nil {
			return err
		}
		_, err = l.svcCtx.DB.Setting.
			Create().
			SetName("setting.platformSetting").
			SetModule(enum.PlatformSetting.Val()).
			SetValue(string(gameRuleJsonByte)).Save(l.ctx)
		if err != nil {
			return err
		}
	}

	// 平台公告
	if _, ok := settingList[enum.PlatformNotice.Val()]; !ok {
		_, err = l.svcCtx.DB.Setting.
			Create().
			SetName("setting.platformNotice").
			SetModule(enum.PlatformNotice.Val()).
			SetValue("").Save(l.ctx)
		if err != nil {
			return err
		}
	}

	return err
}
