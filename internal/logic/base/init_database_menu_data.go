package base

import (
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

// insert initial menu data
func (l *InitDatabaseLogic) insertMenuData() error {
	menuListResp, err := l.svcCtx.CoreRpc.GetMenuList(l.ctx, &core.PageInfoReq{})
	if err != nil {
		return err
	}

	var menuList map[string]uint64
	if menuListResp.Total > 0 {
		menuList = make(map[string]uint64, menuListResp.Total)
		for _, menu := range menuListResp.Data {
			menuList[*menu.Path] = *menu.Id
		}
	}

	// 目录
	var parentId uint64
	var path string

	// PLATFORM MANAGEMENT
	path = "/platform_management"
	if id, ok := menuList[path]; ok {
		parentId = id
	} else {
		menuData, err := l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
			Level:       pointy.GetPointer(uint32(1)),
			ParentId:    pointy.GetPointer(common.DefaultParentId),
			Name:        pointy.GetPointer("PlatformManagementDirectory"),
			Component:   pointy.GetPointer("LAYOUT"),
			Path:        pointy.GetPointer(path),
			Sort:        pointy.GetPointer(uint32(1)),
			Disabled:    pointy.GetPointer(false),
			ServiceName: pointy.GetPointer("WolfLamp"),
			Meta: &core.Meta{
				Title:              pointy.GetPointer("route.platformManagement"),
				Icon:               pointy.GetPointer("uil:setting"),
				HideMenu:           pointy.GetPointer(false),
				HideBreadcrumb:     pointy.GetPointer(false),
				IgnoreKeepAlive:    pointy.GetPointer(false),
				HideTab:            pointy.GetPointer(false),
				CarryParam:         pointy.GetPointer(false),
				HideChildrenInMenu: pointy.GetPointer(false),
				Affix:              pointy.GetPointer(false),
			},
			MenuType: pointy.GetPointer(uint32(0)),
		})

		if err != nil {
			return err
		}

		parentId = menuData.Id
	}

	// 轮播图管理
	path = "/platform_management/banner"
	if _, ok := menuList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
			Level:       pointy.GetPointer(uint32(2)),
			ParentId:    pointy.GetPointer(parentId),
			Path:        pointy.GetPointer(path),
			Name:        pointy.GetPointer("BannerManagement"),
			Component:   pointy.GetPointer("/platform_management/banner/index"),
			Sort:        pointy.GetPointer(uint32(1)),
			Disabled:    pointy.GetPointer(false),
			ServiceName: pointy.GetPointer("WolfLamp"),
			Meta: &core.Meta{
				Title:              pointy.GetPointer("route.bannerManagement"),
				Icon:               pointy.GetPointer("material-symbols:planner-banner-ad-pt-outline"),
				HideMenu:           pointy.GetPointer(false),
				HideBreadcrumb:     pointy.GetPointer(false),
				IgnoreKeepAlive:    pointy.GetPointer(false),
				HideTab:            pointy.GetPointer(false),
				CarryParam:         pointy.GetPointer(false),
				HideChildrenInMenu: pointy.GetPointer(false),
				Affix:              pointy.GetPointer(false),
			},
			MenuType: pointy.GetPointer(uint32(1)),
		})

		if err != nil {
			return err
		}
	}

	// 平台配置
	path = "/platform_management/setting"
	if _, ok := menuList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
			Level:       pointy.GetPointer(uint32(2)),
			ParentId:    pointy.GetPointer(parentId),
			Path:        pointy.GetPointer(path),
			Name:        pointy.GetPointer("SettingManagement"),
			Component:   pointy.GetPointer("/platform_management/setting/index"),
			Sort:        pointy.GetPointer(uint32(1)),
			Disabled:    pointy.GetPointer(false),
			ServiceName: pointy.GetPointer("WolfLamp"),
			Meta: &core.Meta{
				Title:              pointy.GetPointer("route.settingManagement"),
				Icon:               pointy.GetPointer("uil:setting"),
				HideMenu:           pointy.GetPointer(false),
				HideBreadcrumb:     pointy.GetPointer(false),
				IgnoreKeepAlive:    pointy.GetPointer(false),
				HideTab:            pointy.GetPointer(false),
				CarryParam:         pointy.GetPointer(false),
				HideChildrenInMenu: pointy.GetPointer(false),
				Affix:              pointy.GetPointer(false),
			},
			MenuType: pointy.GetPointer(uint32(1)),
		})

		if err != nil {
			return err
		}
	}

	// 玩家管理
	path = "/platform_management/player"
	if _, ok := menuList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
			Level:       pointy.GetPointer(uint32(2)),
			ParentId:    pointy.GetPointer(parentId),
			Path:        pointy.GetPointer(path),
			Name:        pointy.GetPointer("PlayerManagement"),
			Component:   pointy.GetPointer("/platform_management/player/index"),
			Sort:        pointy.GetPointer(uint32(1)),
			Disabled:    pointy.GetPointer(false),
			ServiceName: pointy.GetPointer("WolfLamp"),
			Meta: &core.Meta{
				Title:              pointy.GetPointer("route.playerManagement"),
				Icon:               pointy.GetPointer("ph:user"),
				HideMenu:           pointy.GetPointer(false),
				HideBreadcrumb:     pointy.GetPointer(false),
				IgnoreKeepAlive:    pointy.GetPointer(false),
				HideTab:            pointy.GetPointer(false),
				CarryParam:         pointy.GetPointer(false),
				HideChildrenInMenu: pointy.GetPointer(false),
				Affix:              pointy.GetPointer(false),
			},
			MenuType: pointy.GetPointer(uint32(1)),
		})

		if err != nil {
			return err
		}
	}

	// 订单管理
	path = "/platform_management/order"
	if _, ok := menuList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
			Level:       pointy.GetPointer(uint32(2)),
			ParentId:    pointy.GetPointer(parentId),
			Path:        pointy.GetPointer(path),
			Name:        pointy.GetPointer("OrderManagement"),
			Component:   pointy.GetPointer("/platform_management/order/index"),
			Sort:        pointy.GetPointer(uint32(1)),
			Disabled:    pointy.GetPointer(false),
			ServiceName: pointy.GetPointer("WolfLamp"),
			Meta: &core.Meta{
				Title:              pointy.GetPointer("route.orderManagement"),
				Icon:               pointy.GetPointer("lets-icons:order"),
				HideMenu:           pointy.GetPointer(false),
				HideBreadcrumb:     pointy.GetPointer(false),
				IgnoreKeepAlive:    pointy.GetPointer(false),
				HideTab:            pointy.GetPointer(false),
				CarryParam:         pointy.GetPointer(false),
				HideChildrenInMenu: pointy.GetPointer(false),
				Affix:              pointy.GetPointer(false),
			},
			MenuType: pointy.GetPointer(uint32(1)),
		})

		if err != nil {
			return err
		}
	}

	return nil
}
