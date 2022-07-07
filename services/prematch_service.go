package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
	"wuju_data/entity"
	"wuju_data/global"
	"wuju_data/models"
	"wuju_data/utils"

	"github.com/rs/zerolog/log"
)

type prematchService struct {
	token string
	url   string
}

func (s *prematchService) Setup(conf *entity.Conf) {
	s.token = conf.Token
	s.url = conf.PreMatch
}

// // 从数据库中重新加载 最后的比分信息
// func (s *prematchService) Reload(conf *entity.Conf) {
// 	for commID := range global.C_COMM.Items() {
// 		//
// 		Adata, err := models.AsianHandicapModel.GetListByCommID(commID)
// 		if err != nil {
// 			log.Error().Str("err", err.Error()).Send()
// 			continue
// 		}
// 		for _, dataA := range Adata {
// 			asianValue := fmt.Sprintf("%s|%s|%s|%s",
// 				dataA["home_odds"],
// 				dataA["home_handicap"],
// 				dataA["away_odds"],
// 				dataA["away_handicap"],
// 			)
// 			asianTime := fmt.Sprintf("%v", dataA["update_at"])
// 			global.C_ASIA.Add(commID, fmt.Sprintf("%s|%s", asianTime, asianValue), cache.NoExpiration)
// 		}

// 		// G
// 		Gdata, err := models.GoalLineModel.GetListByCommID(commID)
// 		if err != nil {
// 			log.Error().Str("err", err.Error()).Send()
// 			continue
// 		}
// 		for _, dataG := range Gdata {
// 			gValue := fmt.Sprintf("%s|%s|%s|%s",
// 				dataG["home_odds"],
// 				dataG["home_handicap"],
// 				dataG["away_odds"],
// 				dataG["away_handicap"],
// 			)
// 			gTime := fmt.Sprintf("%v", dataG["update_at"])
// 			global.C_GOALLINE.Add(commID, fmt.Sprintf("%s|%s", gTime, gValue), cache.NoExpiration)
// 		}

// 		//F
// 		Fdata, err := models.FullTimeModel.GetListByCommID(commID)
// 		if err != nil {
// 			log.Error().Str("err", err.Error()).Send()
// 			continue
// 		}
// 		for _, dataF := range Fdata {
// 			fValue := fmt.Sprintf("%s|%s|%s",
// 				dataF["home_odds"],
// 				dataF["draw_odds"],
// 				dataF["away_odds"],
// 			)
// 			fTime := fmt.Sprintf("%v", dataF["update_at"])
// 			global.C_FULLTIME.Add(commID, fmt.Sprintf("%s|%s", fTime, fValue), cache.NoExpiration)
// 		}
// 	}
// 	log.Info().Msgf("reload prematch success")
// }

func (s *prematchService) GetPrematchData(commID string) error {
	url := fmt.Sprintf(s.url, s.token, commID)

	respByte, err := utils.HTTPGet(url)
	if err != nil {
		return err
	}

	var respPrematch entity.RespPrematch

	err = global.JSON.Unmarshal(respByte, &respPrematch)
	if err != nil {
		return err
	}

	if respPrematch.Success != 1 {
		return fmt.Errorf("resp_prematch.success != 1 (%d)", respPrematch.Success)
	}

	// 解析亚洲
	if err = s.ParseAsiaHandicap(commID, &respPrematch); err != nil {
		log.Error().Str("err", err.Error()).Str("comm_id", commID).Send()
	}
	// 解析大小
	if err = s.ParseGoalLine(commID, &respPrematch); err != nil {
		log.Error().Str("err", err.Error()).Str("comm_id", commID).Send()
	}
	// 解析欧
	if err = s.ParseFullTime(commID, &respPrematch); err != nil {
		log.Error().Str("err", err.Error()).Str("comm_id", commID).Send()
	}

	return nil
}

func (s *prematchService) ParseAsiaHandicap(commID string, data *entity.RespPrematch) error {

	if len(data.Results) == 0 {
		return fmt.Errorf("result.length == 0")
	}

	if len(data.Results[0].AsianLines.SP.AsianHandicap.Odds) != 2 {
		return fmt.Errorf("result.asianLines.odds != 2")
	}

	// 亚盘结果保存
	// 1. 结果和缓存对比,如果是新的-入库,如果改变-入库,如果没变-丢弃
	asianHandicapOdds := data.Results[0].AsianLines.SP.AsianHandicap.Odds

	h := md5.New()
	h.Write([]byte(commID))
	h.Write([]byte(asianHandicapOdds[0].Odds))
	h.Write([]byte(asianHandicapOdds[0].Handicap))
	h.Write([]byte(asianHandicapOdds[1].Odds))
	h.Write([]byte(asianHandicapOdds[1].Handicap))
	hRes := hex.EncodeToString(h.Sum(nil))
	_, err := models.AsianHandicapModel.AddList(
		SnowFlakeService.NextID(),
		commID,
		asianHandicapOdds[0].Odds,
		asianHandicapOdds[0].Handicap,
		asianHandicapOdds[1].Odds,
		asianHandicapOdds[1].Handicap,
		time.Now().Format("2006-01-02 15:04:05"),
		hRes,
	)
	return err
}

func (s *prematchService) ParseGoalLine(commID string, data *entity.RespPrematch) error {

	if len(data.Results) == 0 {
		return fmt.Errorf("result.length == 0")
	}

	if len(data.Results[0].AsianLines.SP.GoalLine.Odds) != 2 {
		return fmt.Errorf("goalLine.odds != 2")
	}

	goalOdds := data.Results[0].AsianLines.SP.GoalLine.Odds
	h := md5.New()
	h.Write([]byte(commID))
	h.Write([]byte(goalOdds[0].Odds))
	h.Write([]byte(goalOdds[0].Name))
	h.Write([]byte(goalOdds[1].Odds))
	h.Write([]byte(goalOdds[1].Name))
	hRes := hex.EncodeToString(h.Sum(nil))

	_, err := models.GoalLineModel.AddList(
		SnowFlakeService.NextID(),
		commID,
		goalOdds[0].Odds,
		goalOdds[0].Name,
		goalOdds[1].Odds,
		goalOdds[1].Name,
		time.Now().Format("2006-01-02 15:04:05"),
		hRes,
	)

	return err
}

func (s *prematchService) ParseFullTime(commID string, data *entity.RespPrematch) error {
	if len(data.Results) == 0 {
		return fmt.Errorf("result.length == 0")
	}

	if len(data.Results[0].Main.SP.FullTimeResult.Odds) != 3 {
		return fmt.Errorf("fulltime.odds != 3")
	}

	fOdds := data.Results[0].Main.SP.FullTimeResult.Odds

	h := md5.New()
	h.Write([]byte(commID))
	h.Write([]byte(fOdds[0].Odds))
	h.Write([]byte(fOdds[1].Odds))
	h.Write([]byte(fOdds[2].Odds))
	hRes := hex.EncodeToString(h.Sum(nil))

	_, err := models.FullTimeModel.AddList(
		SnowFlakeService.NextID(),
		commID,
		fOdds[0].Odds,
		fOdds[1].Odds,
		fOdds[2].Odds,
		time.Now().Format("2006-01-02 15:04:05"),
		hRes,
	)
	return err
}
