package rpc

import (
	"context"
	"fmt"
	"game/internal/utils/common"
	"game/proto/gamepb"
	"time"

	"gitlab.lucky9studio.com/sdk-service/project-ranking-service/proto/rank"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var RankService rank.RankServiceClient

func InitRankService(addr string) {
	cc, err := grpc.NewClient(addr)
	if err != nil {
		panic(err)
	}
	RankService = rank.NewRankServiceClient(cc)
}

func UpdateRank(appId int64, gameId int64, name string, timezero int, ctx context.Context, data []*rank.RankData) {
	var date *timestamppb.Timestamp
	var expiredTime *timestamppb.Timestamp
	now := time.Now()
	loc := time.FixedZone(fmt.Sprintf("UTC%+d", timezero), timezero*60*60)
	switch name {
	case gamepb.LeaderboardType_LT_Today.String():
		date = timestamppb.New(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc))
		expiredTime = timestamppb.New(time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc))
	case gamepb.LeaderboardType_LT_Week.String():
		offset := (int(now.Weekday()) - int(time.Monday) + 7) % 7
		monday := now.AddDate(0, 0, -offset)
		nextMonday := monday.AddDate(0, 0, 7)
	
		date = timestamppb.New(time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, loc))
		expiredTime = timestamppb.New(time.Date(nextMonday.Year(), nextMonday.Month(), nextMonday.Day(), 0, 0, 0, 0, loc).Add(-time.Second))
	case gamepb.LeaderboardType_LT_Total.String():
		date = timestamppb.New(time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, loc))
		expiredTime = timestamppb.New(time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, loc).Add(-time.Second))
	}

	req := &rank.RankDataUpdateReq{
		RequestId: common.GetOnlyId(),
		AppId:       appId,
		GameId:      gameId,
		Name: 	  name,
		Date:       date,
		ExpiredTime: expiredTime,
		Data: 	 data,
	}

	RankService.RankDataUpdate(ctx, req)
}