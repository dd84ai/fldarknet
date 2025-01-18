package router

import (
	"sort"

	"github.com/darklab8/fl-darkcore/darkcore/builder"
	"github.com/darklab8/fl-darkstat/darkstat/configs_export"
	"github.com/darklab8/fl-darkstat/darkstat/front"
	"github.com/darklab8/fl-darkstat/darkstat/front/tab"
	"github.com/darklab8/fl-darkstat/darkstat/front/types"
	"github.com/darklab8/fl-darkstat/darkstat/front/urls"
	"github.com/darklab8/go-utils/utils/utils_types"
)

func (l *Router) LinkBases(
	build *builder.Builder,
	data *configs_export.Exporter,
	shared *types.SharedData,
) {

	sort.Slice(data.Bases, func(i, j int) bool {
		return data.Bases[i].Name < data.Bases[j].Name
	})

	for _, base := range data.Bases {
		sort.Slice(base.TradeRoutes, func(i, j int) bool {
			return base.TradeRoutes[i].Transport.GetProffitPerTime() > base.TradeRoutes[j].Transport.GetProffitPerTime()
		})
	}

	build.RegComps(
		builder.NewComponent(
			urls.Bases,
			front.BasesT(configs_export.FilterToUserfulBases(data.Bases), front.BaseShowShops, tab.ShowEmpty(false), shared),
		),
		builder.NewComponent(
			tab.AllItemsUrl(urls.Bases),
			front.BasesT(data.Bases, front.BaseShowShops, tab.ShowEmpty(true), shared),
		),
		builder.NewComponent(
			urls.Missions,
			front.BasesT(configs_export.FilterToUserfulBases(data.Bases), front.BaseShowMissions, tab.ShowEmpty(false), shared),
		),
		builder.NewComponent(
			tab.AllItemsUrl(urls.Missions),
			front.BasesT(data.Bases, front.BaseShowMissions, tab.ShowEmpty(true), shared),
		),
	)

	for _, base := range data.Bases {
		if base.Missions != nil {
			build.RegComps(
				builder.NewComponent(
					utils_types.FilePath(front.BaseDetailedUrl(base, front.BaseShowMissions)),
					front.BaseMissions(base.Name, *base.Missions, front.BaseShowMissions),
				),
			)
		}

		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.BaseDetailedUrl(base, front.BaseShowShops)),
				front.BaseMarketGoods(base.Name, base.MarketGoodsPerNick, front.BaseShowShops),
			),

			builder.NewComponent(
				utils_types.FilePath(front.BaseDetailedUrl(base, front.BaseTabTrades)),
				front.BaseTrades(base.Name, base.TradeRoutes, front.BaseTabTrades, shared),
			),
			builder.NewComponent(
				utils_types.FilePath(front.BaseDetailedUrl(base, front.BaseAllRoutes)),
				front.BaseRoutes(base.Name, base.AllRoutes, front.BaseAllRoutes, shared),
			),
		)

		for _, combo_route := range base.AllRoutes {
			build.RegComps(
				builder.NewComponent(
					utils_types.FilePath(front.RouteUrl(combo_route.Transport.Route)),
					front.TradeRouteInfo(combo_route.Transport.Route, combo_route.Frigate.Route, combo_route.Freighter.Route, shared),
				),
			)
		}
	}

	sort.Slice(data.Bases, func(i, j int) bool {
		if data.Bases[j].BestTransportRoute == nil {
			return true
		}
		if data.Bases[i].BestTransportRoute == nil {
			return false
		}
		return data.Bases[i].BestTransportRoute.GetProffitPerTime() > data.Bases[j].BestTransportRoute.GetProffitPerTime()
	})

	build.RegComps(
		builder.NewComponent(
			urls.Trades,
			front.BasesT(configs_export.FilterToUserfulBases(data.Bases), front.BaseTabTrades, tab.ShowEmpty(false), shared),
		),
		builder.NewComponent(
			tab.AllItemsUrl(urls.Trades),
			front.BasesT(data.Bases, front.BaseTabTrades, tab.ShowEmpty(true), shared),
		),
		builder.NewComponent(
			urls.Asteroids,
			front.BasesT(configs_export.FitlerToUsefulOres(data.MiningOperations), front.BaseTabOres, tab.ShowEmpty(false), shared),
		),
		builder.NewComponent(
			tab.AllItemsUrl(urls.Asteroids),
			front.BasesT(data.MiningOperations, front.BaseTabOres, tab.ShowEmpty(true), shared),
		),
		builder.NewComponent(
			urls.TravelRoutes,
			front.BasesT(configs_export.FilterToUserfulBases(data.Bases), front.BaseAllRoutes, tab.ShowEmpty(false), shared),
		),
		builder.NewComponent(
			tab.AllItemsUrl(urls.TravelRoutes),
			front.BasesT(data.Bases, front.BaseAllRoutes, tab.ShowEmpty(true), shared),
		),
	)

	for _, base := range data.MiningOperations {
		build.RegComps(
			builder.NewComponent(
				utils_types.FilePath(front.BaseDetailedUrl(base, front.BaseTabOres)),
				front.BaseTrades(base.Name, base.TradeRoutes, front.BaseTabOres, shared),
			),
		)

		for _, combo_route := range base.TradeRoutes {

			build.RegComps(
				builder.NewComponent(
					utils_types.FilePath(front.RouteUrl(combo_route.Transport.Route)),
					front.TradeRouteInfo(combo_route.Transport.Route, combo_route.Frigate.Route, combo_route.Freighter.Route, shared),
				),
			)
		}

	}
}
