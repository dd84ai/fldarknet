package linker

/*
Links data from exported fl-configs
into stuff rendered by fl-darkstat
*/

import (
	"sort"

	"github.com/darklab8/fl-configs/configs/configs_export"
	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-darkstat/darkstat/builder"
	"github.com/darklab8/fl-darkstat/darkstat/front"
	"github.com/darklab8/fl-darkstat/darkstat/front/urls"
	"github.com/darklab8/fl-darkstat/darkstat/settings"
	"github.com/darklab8/fl-darkstat/darkstat/settings/logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type Linker struct {
	configs *configs_export.Exporter
}

type LinkOption func(l *Linker)

func NewLinker(opts ...LinkOption) *Linker {
	l := &Linker{}
	for _, opt := range opts {
		opt(l)
	}

	if l.configs == nil {
		configs := configs_mapped.NewMappedConfigs()
		logus.Log.Debug("scanning freelancer folder", utils_logus.FilePath(settings.FreelancerFolder))
		configs.Read(settings.FreelancerFolder)
		l.configs = configs_export.NewExporter(configs)
	}

	return l
}

func (l *Linker) Link() *builder.Builder {
	data := l.configs.Export()

	sort.Slice(data.Bases, func(i, j int) bool {
		if data.Bases[i].Name != "" && data.Bases[j].Name == "" {
			return true
		}
		return data.Bases[i].Name < data.Bases[j].Name
	})

	for _, base := range data.Bases {
		sort.Slice(base.MarketGoods, func(i, j int) bool {
			if base.MarketGoods[i].Name != "" && base.MarketGoods[j].Name == "" {
				return true
			}
			return base.MarketGoods[i].Name < base.MarketGoods[j].Name
		})
	}

	sort.Slice(data.Factions, func(i, j int) bool {
		if data.Factions[i].Name != "" && data.Factions[j].Name == "" {
			return true
		}
		return data.Factions[i].Name < data.Factions[j].Name
	})

	for _, faction := range data.Factions {
		sort.Slice(faction.Reputations, func(i, j int) bool {
			if faction.Reputations[i].Name != "" && faction.Reputations[j].Name == "" {
				return true
			}
			return faction.Reputations[i].Name < faction.Reputations[j].Name
		})
	}

	sort.Slice(data.Commodities, func(i, j int) bool {
		if data.Commodities[i].Name != "" && data.Commodities[j].Name == "" {
			return true
		}
		return data.Commodities[i].Name < data.Commodities[j].Name
	})

	for _, base_info := range data.Commodities {
		sort.Slice(base_info.Bases, func(i, j int) bool {
			if base_info.Bases[i].BaseName != "" && base_info.Bases[j].BaseName == "" {
				return true
			}
			return base_info.Bases[i].BaseName < base_info.Bases[j].BaseName
		})
	}

	sort.Slice(data.Guns, func(i, j int) bool {
		if data.Guns[i].Name != "" && data.Guns[j].Name == "" {
			return true
		}
		return data.Guns[i].Name < data.Guns[j].Name
	})

	for _, base_info := range data.Guns {
		sort.Slice(base_info.Bases, func(i, j int) bool {
			if base_info.Bases[i].BaseName != "" && base_info.Bases[j].BaseName == "" {
				return true
			}
			return base_info.Bases[i].BaseName < base_info.Bases[j].BaseName
		})
	}

	for _, base_info := range data.Mines {
		sort.Slice(base_info.Bases, func(i, j int) bool {
			if base_info.Bases[i].BaseName != "" && base_info.Bases[j].BaseName == "" {
				return true
			}
			return base_info.Bases[i].BaseName < base_info.Bases[j].BaseName
		})
	}

	sort.Slice(data.Shields, func(i, j int) bool {
		if data.Shields[i].Name != "" && data.Shields[j].Name == "" {
			return true
		}
		return data.Shields[i].Name < data.Shields[j].Name
	})

	for _, base_info := range data.Shields {
		sort.Slice(base_info.Bases, func(i, j int) bool {
			if base_info.Bases[i].BaseName != "" && base_info.Bases[j].BaseName == "" {
				return true
			}
			return base_info.Bases[i].BaseName < base_info.Bases[j].BaseName
		})
	}

	sort.Slice(data.Thrusters, func(i, j int) bool {
		if data.Thrusters[i].Name != "" && data.Thrusters[j].Name == "" {
			return true
		}
		return data.Thrusters[i].Name < data.Thrusters[j].Name
	})

	for _, base_info := range data.Thrusters {
		sort.Slice(base_info.Bases, func(i, j int) bool {
			if base_info.Bases[i].BaseName != "" && base_info.Bases[j].BaseName == "" {
				return true
			}
			return base_info.Bases[i].BaseName < base_info.Bases[j].BaseName
		})
	}

	sort.Slice(data.Ships, func(i, j int) bool {
		if data.Ships[i].Name != "" && data.Ships[j].Name == "" {
			return true
		}
		return data.Ships[i].Name < data.Ships[j].Name
	})

	for _, base_info := range data.Ships {
		sort.Slice(base_info.Bases, func(i, j int) bool {
			if base_info.Bases[i].BaseName != "" && base_info.Bases[j].BaseName == "" {
				return true
			}
			return base_info.Bases[i].BaseName < base_info.Bases[j].BaseName
		})
	}

	build := builder.NewBuilder()
	build.RegComps(
		builder.NewComponent(
			urls.Index,
			front.Index(),
		),
		builder.NewComponent(
			urls.Bases,
			front.BasesT(data.Bases),
		),
		builder.NewComponent(
			urls.Factions,
			front.FactionsT(data.Factions, front.FactionShowBases),
		),
		builder.NewComponent(
			urls.Rephacks,
			front.FactionsT(data.Factions, front.FactionShowRephacks),
		),
		builder.NewComponent(
			urls.Commodities,
			front.CommoditiesT(data.Commodities),
		),
		builder.NewComponent(
			urls.Guns,
			front.GunsT(data.Guns, front.GunsShowBases),
		),
		builder.NewComponent(
			urls.GunModifiers,
			front.GunsT(data.Guns, front.GunsShowDamageBonuses),
		),
		builder.NewComponent(
			urls.Missiles,
			front.GunsT(data.Missiles, front.GunsMissiles),
		),
		builder.NewComponent(
			urls.Mines,
			front.MinesT(data.Mines),
		),
		builder.NewComponent(
			urls.Shields,
			front.ShieldT(data.Shields),
		),
		builder.NewComponent(
			urls.Thrusters,
			front.ThrusterT(data.Thrusters),
		),
		builder.NewComponent(
			urls.Ships,
			front.ShipsT(data.Ships),
		),
	)

	for _, base := range data.Bases {
		// fmt.Println("market_goods, len=", len(base.MarketGoods), " nickname=", base.Nickname, base.Name)
		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.BaseMarketGoodUrl(base)),
				front.BaseMarketGoods(base.MarketGoods),
			),
		)
	}

	for _, faction := range data.Factions {
		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.FactionRepUrl(faction, front.FactionShowBases)),
				front.FactionReps(faction.Reputations),
			),
			builder.NewComponent(
				utils_types.FilePath(front.FactionRepUrl(faction, front.FactionShowRephacks)),
				front.RephackBottom(faction.Rephacks),
			),
		)
	}

	for nickname, infocard := range data.Infocards {
		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.InfocardURL(nickname)),
				front.Infocard(infocard),
			),
		)
	}

	for _, base_info := range data.Commodities {
		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.GoodAtBaseInfoTUrl(base_info)),
				front.GoodAtBaseInfoT(base_info.Bases, front.ShowPricePerVolume(true)),
			),
		)
	}

	for _, gun := range data.Guns {
		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.GunDetailedUrl(gun, front.GunsShowBases)),
				front.GoodAtBaseInfoT(gun.Bases, front.ShowPricePerVolume(false)),
			),
			builder.NewComponent(
				utils_types.FilePath(front.GunDetailedUrl(gun, front.GunsShowDamageBonuses)),
				front.GunShowModifiers(gun),
			),
		)
	}

	for _, missile := range data.Missiles {
		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.GunDetailedUrl(missile, front.GunsMissiles)),
				front.GoodAtBaseInfoT(missile.Bases, front.ShowPricePerVolume(false)),
			),
		)
	}

	for _, mine := range data.Mines {
		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.MineDetailedUrl(mine)),
				front.GoodAtBaseInfoT(mine.Bases, front.ShowPricePerVolume(false)),
			),
		)
	}

	for _, shield := range data.Shields {
		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.ShieldDetailedUrl(shield)),
				front.GoodAtBaseInfoT(shield.Bases, front.ShowPricePerVolume(false)),
			),
		)
	}

	for _, thruster := range data.Thrusters {
		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.ThrusterDetailedUrl(thruster)),
				front.GoodAtBaseInfoT(thruster.Bases, front.ShowPricePerVolume(false)),
			),
		)
	}

	for _, ship := range data.Ships {
		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.ShipDetailedUrl(ship)),
				front.GoodAtBaseInfoT(ship.Bases, front.ShowPricePerVolume(false)),
			),
		)
	}

	return build
}
