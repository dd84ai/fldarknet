package darkhttp

import (
	"net/http"

	"github.com/darklab8/fl-darkstat/darkapis/darkhttp/apiutils"
	"github.com/darklab8/fl-darkstat/darkcore/web"
	"github.com/darklab8/fl-darkstat/darkcore/web/registry"
	"github.com/darklab8/fl-darkstat/darkstat/configs_export"
)

type Engine struct {
	*configs_export.Engine
	MarketGoods []*configs_export.MarketGood        `json:"market_goods"`
	TechCompat  *configs_export.DiscoveryTechCompat `json:"tech_compat"`
}

// ShowAccount godoc
// @Summary      Getting list of Engines
// @Tags         equipment
// @Accept       json
// @Produce      json
// @Success      200  {array}  	darkhttp.Engine
// @Router       /api/engines [get]
// @Param        filter_to_useful    query     string  false  "insert 'true' if wish to filter items only to useful, usually they are sold, or have goods, or craftable or findable in loot, or bases that are flight reachable from manhattan"
// @Param        include_market_goods    query     string  false  "insert 'true' if wish to include market goods under 'market goods' key or not. Such data can add a lot of extra weight"
// @Param        include_tech_compat    query     string  false  "insert 'true' if wish to include tech compat info too for the item. Such data can add a lot of extra weight"
func GetEngines(webapp *web.Web, api *Api) *registry.Endpoint {
	return &registry.Endpoint{
		Url: "GET " + ApiRoute + "/engines",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if webapp.AppDataMutex != nil {
				webapp.AppDataMutex.Lock()
				defer webapp.AppDataMutex.Unlock()
			}

			filter_to_useful := r.URL.Query().Get("filter_to_useful") == "true"
			include_market_goods := r.URL.Query().Get("include_market_goods") == "true"
			include_tech_compat := r.URL.Query().Get("include_tech_compat") == "true"

			var result []configs_export.Engine
			if filter_to_useful {
				result = api.app_data.Configs.FilterToUsefulEngines(api.app_data.Configs.Engines)
			} else {
				result = api.app_data.Configs.Engines
			}

			var output []*Engine
			for _, item := range result {
				answer := &Engine{
					Engine: &item,
				}
				if include_market_goods {
					for _, good := range item.Bases {
						answer.MarketGoods = append(answer.MarketGoods, good)
					}
				}
				if include_tech_compat {
					answer.TechCompat = item.DiscoveryTechCompat
				}
				output = append(output, answer)
			}

			apiutils.ReturnJson(&w, output)
		},
	}
}
