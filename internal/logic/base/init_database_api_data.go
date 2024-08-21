package base

import (
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

// insert initial api data
func (l *InitDatabaseLogic) insertApiData() error {
	apiListResp, err := l.svcCtx.CoreRpc.GetApiList(l.ctx, &core.ApiListReq{})
	if err != nil {
		return err
	}

	var apiList map[string]uint64
	if apiListResp.Total > 0 {
		apiList = make(map[string]uint64, apiListResp.Total)
		for _, api := range apiListResp.Data {
			apiList[*api.Path] = *api.Id
		}
	}

	var path string

	// BANNER
	path = "/banner/create"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.createBanner"),
			ApiGroup:    pointy.GetPointer("banner"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/banner/update"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.updateBanner"),
			ApiGroup:    pointy.GetPointer("banner"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/banner/delete"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.deleteBanner"),
			ApiGroup:    pointy.GetPointer("banner"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/banner/list"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.listBanner"),
			ApiGroup:    pointy.GetPointer("banner"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/banner/find"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.findBanner"),
			ApiGroup:    pointy.GetPointer("banner"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	// SETTING
	path = "/setting/create"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.createSetting"),
			ApiGroup:    pointy.GetPointer("setting"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/setting/update"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.updateSetting"),
			ApiGroup:    pointy.GetPointer("setting"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/setting/delete"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.deleteSetting"),
			ApiGroup:    pointy.GetPointer("setting"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/setting/list"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.listSetting"),
			ApiGroup:    pointy.GetPointer("setting"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/setting/find"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.findSetting"),
			ApiGroup:    pointy.GetPointer("setting"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	// PLAYER
	path = "/player/create"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.createPlayer"),
			ApiGroup:    pointy.GetPointer("player"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/player/update"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.updatePlayer"),
			ApiGroup:    pointy.GetPointer("player"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/player/list"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.listPlayer"),
			ApiGroup:    pointy.GetPointer("player"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/player/find"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.findPlayer"),
			ApiGroup:    pointy.GetPointer("player"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/player/delete"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.deletePlayer"),
			ApiGroup:    pointy.GetPointer("player"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	// ORDER
	path = "/order/create"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.createOrder"),
			ApiGroup:    pointy.GetPointer("order"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/order/update"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.updateOrder"),
			ApiGroup:    pointy.GetPointer("order"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/order/list"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.listOrder"),
			ApiGroup:    pointy.GetPointer("order"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/order/find"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.findOrder"),
			ApiGroup:    pointy.GetPointer("order"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/order/delete"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.deleteOrder"),
			ApiGroup:    pointy.GetPointer("order"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	// POOL
	path = "/pool/create_pool"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.createPool"),
			ApiGroup:    pointy.GetPointer("pool"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	path = "/pool/list_pool"
	if _, ok := apiList[path]; !ok {
		_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
			ServiceName: pointy.GetPointer("WolfLamp"),
			Path:        pointy.GetPointer(path),
			Description: pointy.GetPointer("apiDesc.listPool"),
			ApiGroup:    pointy.GetPointer("pool"),
			Method:      pointy.GetPointer("POST"),
		})
		if err != nil {
			return err
		}
	}

	return err
}
