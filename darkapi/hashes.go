package darkapi

import (
	"net/http"

	"github.com/darklab8/fl-darkstat/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld/flhash"
	"github.com/darklab8/fl-darkstat/darkcore/web"
	"github.com/darklab8/fl-darkstat/darkcore/web/registry"
)

type Hashes struct {
	NicknameToHash map[string]flhash.HashCode `json:"nickname_to_hash"`
}

// ShowAccount godoc
// @Summary      Hashes
// @Tags         hashes
// @Accept       json
// @Produce      json
// @Success      200  {object}  	Hashes
// @Router       /api/hashes [get]
func GetHashes(webapp *web.Web, api *Api) *registry.Endpoint {
	return &registry.Endpoint{
		Url: "GET " + ApiRoute + "/hashes",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if webapp.AppDataMutex != nil {
				webapp.AppDataMutex.Lock()
				defer webapp.AppDataMutex.Unlock()
			}

			ReturnJson(&w, Hashes{NicknameToHash: api.app_data.Configs.Hashes})
		},
	}
}
