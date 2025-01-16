package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/darklab8/fl-configs/configs/configs_export"
	"github.com/darklab8/fl-darkcore/darkcore/web"
	"github.com/darklab8/fl-darkcore/darkcore/web/registry"
	"github.com/darklab8/fl-darkstat/darkstat/settings/logus"
)

// ShowAccount godoc
// @Summary      Getting list of Player Owned Bases
// @Tags         pobs
// @Accept       json
// @Produce      json
// @Success      200  {array}  	configs_export.PoB
// @Router       /api/pobs [get]
func NewEndpointPoBs(webapp *web.Web, api *Api) *registry.Endpoint {
	return &registry.Endpoint{
		Url: "GET " + ApiRoute + "/pobs",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if webapp.AppDataMutex != nil {
				webapp.AppDataMutex.Lock()
				defer webapp.AppDataMutex.Unlock()
			}

			data, err := json.Marshal(api.app_data.Configs.PoBs)
			logus.Log.CheckPanic(err, "should be marshable")
			fmt.Fprint(w, string(data))
		},
	}
}

// ShowAccount godoc
// @Summary      PoB Goods
// @Tags         pobs
// @Accept       json
// @Produce      json
// @Success      200  {array}  	configs_export.PoBGood
// @Router       /api/pob_goods [get]
func NewEndpointPoBGoods(webapp *web.Web, api *Api) *registry.Endpoint {
	return &registry.Endpoint{
		Url: "GET " + ApiRoute + "/pob_goods",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if webapp.AppDataMutex != nil {
				webapp.AppDataMutex.Lock()
				defer webapp.AppDataMutex.Unlock()
			}

			data, err := json.Marshal(api.app_data.Configs.PoBGoods)
			logus.Log.CheckPanic(err, "should be marshable")
			fmt.Fprint(w, string(data))
		},
	}
}
