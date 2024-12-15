package front

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_export"
	"github.com/darklab8/fl-darkstat/darkstat/front/types"
	"github.com/darklab8/go-utils/utils"
)

func TestTechCompat(t *testing.T) {
	var guns []configs_export.Gun
	var gun configs_export.Gun

	var item Item = gun
	_ = item

	GetDiscoCacheMap(utils.CompL(guns, func(x configs_export.Gun) Item { return x }), types.DiscoveryIDs{})
}
