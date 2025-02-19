package darkhttp

import (
	"net/http"

	pb "github.com/darklab8/fl-darkstat/darkapis/darkgrpc/statproto"
	"github.com/darklab8/fl-darkstat/darkapis/darkhttp/apiutils"
	"github.com/darklab8/fl-darkstat/darkapis/services"
	"github.com/darklab8/fl-darkstat/darkcore/web"
	"github.com/darklab8/fl-darkstat/darkcore/web/registry"
	"github.com/darklab8/fl-darkstat/darkstat/configs_export"
)

type Base struct {
	*configs_export.Base
	MarketGoods []*configs_export.MarketGood `json:"market_goods"`
}

// ShowAccount godoc
// @Summary      Getting list of NPC Bases
// @Tags         bases
// @Accept       json
// @Produce      json
// @Success      200  {array}  	darkhttp.Base
// @Router       /api/npc_bases [post]
// @Param request body pb.GetBasesInput true "input variables, description in Models of api 2.0"
func GetBases(webapp *web.Web, api *Api) *registry.Endpoint {
	return &registry.Endpoint{
		Url: "" + ApiRoute + "/npc_bases",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if webapp.AppDataMutex != nil {
				webapp.AppDataMutex.Lock()
				defer webapp.AppDataMutex.Unlock()
			}

			var in pb.GetBasesInput
			if err := ReadJsonInput(w, r, &in); err != nil && r.Method == "POST" {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			filter_to_useful := in.FilterToUseful
			include_market_goods := in.IncludeMarketGoods
			if r.URL.Query().Get("filter_to_useful") == "true" {
				filter_to_useful = true
			}
			if r.URL.Query().Get("include_market_goods") == "true" {
				include_market_goods = true
			}

			var result []*configs_export.Base
			if filter_to_useful {
				result = configs_export.FilterToUserfulBases(api.app_data.Configs.Bases)
			} else {
				result = api.app_data.Configs.Bases
			}
			result = services.FilterNicknames(in.FilterNicknames, result)

			var output []*Base
			for _, item := range result {
				answer := &Base{
					Base: item,
				}
				if include_market_goods {
					for _, good := range services.FilterMarketGoodCategory(in.FilterMarketGoodCategory, item.MarketGoodsPerNick) {
						answer.MarketGoods = append(answer.MarketGoods, good)
					}
				}
				output = append(output, answer)
			}
			apiutils.ReturnJson(&w, output)
		},
	}

}

// ShowAccount godoc
// @Summary      Getting list of Mining Operations
// @Tags         bases
// @Accept       json
// @Produce      json
// @Success      200  {array}  	darkhttp.Base
// @Router       /api/mining_operations [post]
// @Param request body pb.GetBasesInput true "input variables, description in Models of api 2.0"
func GetOreFields(webapp *web.Web, api *Api) *registry.Endpoint {
	return &registry.Endpoint{
		Url: "" + ApiRoute + "/mining_operations",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if webapp.AppDataMutex != nil {
				webapp.AppDataMutex.Lock()
				defer webapp.AppDataMutex.Unlock()
			}
			var in pb.GetBasesInput
			if err := ReadJsonInput(w, r, &in); err != nil && r.Method == "POST" {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			filter_to_useful := in.FilterToUseful
			include_market_goods := in.IncludeMarketGoods
			if r.URL.Query().Get("filter_to_useful") == "true" {
				filter_to_useful = true
			}
			if r.URL.Query().Get("include_market_goods") == "true" {
				include_market_goods = true
			}

			var result []*configs_export.Base
			if filter_to_useful {
				result = configs_export.FitlerToUsefulOres(api.app_data.Configs.MiningOperations)
			} else {
				result = api.app_data.Configs.MiningOperations
			}
			result = services.FilterNicknames(in.FilterNicknames, result)

			var output []*Base
			for _, item := range result {
				answer := &Base{
					Base: item,
				}
				if include_market_goods {
					for _, good := range services.FilterMarketGoodCategory(in.FilterMarketGoodCategory, item.MarketGoodsPerNick) {
						answer.MarketGoods = append(answer.MarketGoods, good)
					}
				}
				output = append(output, answer)
			}
			apiutils.ReturnJson(&w, output)
		},
	}
}

// ShowAccount godoc
// @Summary      Getting list of Player Owned Bases in Bases format. Lists only pobs that have known position coordinates
// @Tags         bases
// @Accept       json
// @Produce      json
// @Success      200  {array}  	darkhttp.Base
// @Router       /api/pobs/bases [post]
// @Param request body pb.GetBasesInput true "input variables, description in Models of api 2.0"
func GetPoBBases(webapp *web.Web, api *Api) *registry.Endpoint {
	return &registry.Endpoint{
		Url: "" + ApiRoute + "/pobs/bases",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if webapp.AppDataMutex != nil {
				webapp.AppDataMutex.Lock()
				defer webapp.AppDataMutex.Unlock()
			}
			var in pb.GetBasesInput
			if err := ReadJsonInput(w, r, &in); err != nil && r.Method == "POST" {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			include_market_goods := in.IncludeMarketGoods
			if r.URL.Query().Get("include_market_goods") == "true" {
				include_market_goods = true
			}
			var result []*configs_export.Base = api.app_data.Configs.PoBsToBases(api.app_data.Configs.PoBs)
			result = services.FilterNicknames(in.FilterNicknames, result)

			var output []*Base
			for _, item := range result {
				answer := &Base{
					Base: item,
				}
				if include_market_goods {
					for _, good := range services.FilterMarketGoodCategory(in.FilterMarketGoodCategory, item.MarketGoodsPerNick) {
						answer.MarketGoods = append(answer.MarketGoods, good)
					}
				}
				output = append(output, answer)
			}

			apiutils.ReturnJson(&w, output)
		},
	}
}
