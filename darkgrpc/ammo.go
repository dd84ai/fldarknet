package darkgrpc

import (
	"context"

	pb "github.com/darklab8/fl-darkstat/darkgrpc/statproto"
)

func (s *Server) GetAmmos(_ context.Context, in *pb.Empty) (*pb.GetAmmoReply, error) {
	if s.app_data != nil {
		s.app_data.Lock()
		defer s.app_data.Unlock()
	}

	var items []*pb.Ammo
	for _, item := range s.app_data.Configs.Ammos {
		item := &pb.Ammo{
			Name:                item.Name, //
			Price:               int64(item.Price),
			HitPts:              int64(item.HitPts),
			Volume:              item.Volume,
			MunitionLifetime:    item.MunitionLifetime,
			Nickname:            item.Nickname,
			NameID:              int64(item.NameID),
			InfoID:              int64(item.InfoID),
			SeekerType:          item.SeekerType,
			SeekerRange:         int64(item.SeekerRange),
			SeekerFovDeg:        int64(item.SeekerFovDeg),
			Bases:               NewBases(item.Bases),
			DiscoveryTechCompat: NewTechCompat(item.DiscoveryTechCompat),
			AmountInCatridge:    NewInt64(item.AmmoLimit.AmountInCatridge),
			MaxCatridges:        NewInt64(item.AmmoLimit.MaxCatridges),
			Mass:                item.Mass,
		}

		items = append(items, item)
	}
	return &pb.GetAmmoReply{Items: items}, nil
}

func (s *Server) GetAmmosMarketGoods(_ context.Context, in *pb.GetMarketGoodsInput) (*pb.GetMarketGoodsReply, error) {
	return GetMarketGoods(s.app_data.Configs.Ammos, in)
}

func (s *Server) GetAmmosTechCompat(_ context.Context, in *pb.GetTechCompatInput) (*pb.GetTechCompatReply, error) {
	return GetTechCompat(s.app_data.Configs.Ammos, in)
}
