package darkapi

import (
	"github.com/darklab8/fl-darkstat/darkcore/web"
	"github.com/darklab8/fl-darkstat/darkcore/web/registry"
	_ "github.com/darklab8/fl-darkstat/darkstat/configs_export"
)

// ShowAccount godoc
// @Summary      Getting list of Player Owned Bases
// @Description  in difference to Disco API, it is enriched with Nicknames/Infocard Names,Region names
// @Description  Sector coordinates, and extra information written in Infocard (totally reflecting Darkstat itself)
// @Tags         player_owned_bases
// @Accept       json
// @Produce      json
// @Success      200  {array}  	configs_export.PoB
// @Router       /api/pobs [get]
func GetPoBs(webapp *web.Web, api *Api) *registry.Endpoint {
	return &registry.Endpoint{
		Url:     "GET " + ApiRoute + "/pobs",
		Handler: GetItemsT(webapp, api.app_data.Configs.PoBs),
	}
}

// ShowAccount godoc
// @Summary      PoB Goods
// @Tags         player_owned_bases
// @Accept       json
// @Produce      json
// @Success      200  {array}  	configs_export.PoBGood
// @Router       /api/pob_goods [get]
func GetPobGoods(webapp *web.Web, api *Api) *registry.Endpoint {
	return &registry.Endpoint{
		Url:     "GET " + ApiRoute + "/pob_goods",
		Handler: GetItemsT(webapp, api.app_data.Configs.PoBGoods),
	}
}

// ShowAccount godoc
// @Summary      Getting list of Player Owned Bases in Bases format. Lists only pobs that have known position coordinate
// @Tags         player_owned_bases
// @Accept       json
// @Produce      json
// @Success      200  {array}  	configs_export.Base
// @Router       /api/pobs/bases [get]
func GetPoBBases(webapp *web.Web, api *Api) *registry.Endpoint {
	return &registry.Endpoint{
		Url:     "GET " + ApiRoute + "/pobs/bases",
		Handler: GetItemsT(webapp, api.app_data.Configs.PoBsToBases(api.app_data.Configs.PoBs)),
	}
}
