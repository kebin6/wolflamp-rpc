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
			GameCommission: 15, // 默认15个点
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
