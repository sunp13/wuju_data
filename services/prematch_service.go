package services

import (
	"fmt"
	"wuju_data/entity"
	"wuju_data/global"
	"wuju_data/models"
	"wuju_data/utils"

	"github.com/patrickmn/go-cache"
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

// 从数据库中重新加载 最后的比分信息
func (s *prematchService) Reload(conf *entity.Conf) {
	for commID := range global.C_COMM.Items() {
		//
		Adata, err := models.AsianHandicapModel.GetListByCommID(commID)
		if err != nil {
			log.Error().Str("err", err.Error()).Send()
			continue
		}
		for _, dataA := range Adata {
			asianValue := fmt.Sprintf("%s|%s|%s|%s",
				dataA["home_odds"],
				dataA["home_handicap"],
				dataA["away_odds"],
				dataA["away_handicap"],
			)
			asianTime := fmt.Sprintf("%v", dataA["update_at"])
			global.C_ASIA.Add(commID, fmt.Sprintf("%s|%s", asianTime, asianValue), cache.NoExpiration)
		}

		// G
		Gdata, err := models.GoalLineModel.GetListByCommID(commID)
		if err != nil {
			log.Error().Str("err", err.Error()).Send()
			continue
		}
		for _, dataG := range Gdata {
			gValue := fmt.Sprintf("%s|%s|%s|%s",
				dataG["home_odds"],
				dataG["home_handicap"],
				dataG["away_odds"],
				dataG["away_handicap"],
			)
			gTime := fmt.Sprintf("%v", dataG["update_at"])
			global.C_GOALLINE.Add(commID, fmt.Sprintf("%s|%s", gTime, gValue), cache.NoExpiration)
		}

		//F
		Fdata, err := models.FullTimeModel.GetListByCommID(commID)
		if err != nil {
			log.Error().Str("err", err.Error()).Send()
			continue
		}
		for _, dataF := range Fdata {
			fValue := fmt.Sprintf("%s|%s|%s",
				dataF["home_odds"],
				dataF["draw_odds"],
				dataF["away_odds"],
			)
			fTime := fmt.Sprintf("%v", dataF["update_at"])
			global.C_FULLTIME.Add(commID, fmt.Sprintf("%s|%s", fTime, fValue), cache.NoExpiration)
		}
	}
	log.Info().Msgf("reload prematch success")
}

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
	asianLinesTime := data.Results[0].AsianLines.UpdatedAt
	asianHandicapOdds := data.Results[0].AsianLines.SP.AsianHandicap.Odds
	asianValue := fmt.Sprintf("%s|%s|%s|%s",
		asianHandicapOdds[0].Odds,
		asianHandicapOdds[0].Handicap,
		asianHandicapOdds[1].Odds,
		asianHandicapOdds[1].Handicap,
	)
	// 读取缓存是否存在
	asianCacheData, ok := global.C_ASIA.Get(commID)
	if !ok {
		// 不存在
		// 写缓存
		global.C_ASIA.Add(commID, fmt.Sprintf("%s|%s", asianLinesTime, asianValue), cache.NoExpiration)
		// 写数据库
		_, err := models.AsianHandicapModel.AddList(
			commID,
			asianHandicapOdds[0].Odds,
			asianHandicapOdds[0].Handicap,
			asianHandicapOdds[1].Odds,
			asianHandicapOdds[1].Handicap,
			asianLinesTime,
		)
		if err != nil {
			log.Error().Str("err", err.Error()).Send()
		}
	} else {
		// 存在判断是否一致
		if asianCacheData.(string)[11:] != asianValue {
			// 不一致
			// 写缓存
			global.C_ASIA.Set(commID, fmt.Sprintf("%s|%s", asianLinesTime, asianValue), cache.NoExpiration)
			// 写数据库
			_, err := models.AsianHandicapModel.AddList(
				commID,
				asianHandicapOdds[0].Odds,
				asianHandicapOdds[0].Handicap,
				asianHandicapOdds[1].Odds,
				asianHandicapOdds[1].Handicap,
				asianLinesTime,
			)
			if err != nil {
				log.Error().Str("err", err.Error()).Send()
			}
		}
	}

	return nil
}

func (s *prematchService) ParseGoalLine(commID string, data *entity.RespPrematch) error {

	if len(data.Results) == 0 {
		return fmt.Errorf("result.length == 0")
	}

	if len(data.Results[0].AsianLines.SP.GoalLine.Odds) != 2 {
		return fmt.Errorf("goalLine.odds != 2")
	}

	goalLineTime := data.Results[0].AsianLines.UpdatedAt
	goalOdds := data.Results[0].AsianLines.SP.GoalLine.Odds
	goalValue := fmt.Sprintf("%s|%s|%s|%s",
		goalOdds[0].Odds,
		goalOdds[0].Name,
		goalOdds[1].Odds,
		goalOdds[1].Name,
	)

	// 读取缓存是否存在
	cacheData, ok := global.C_GOALLINE.Get(commID)
	if !ok {
		// 不存在
		// 写缓存
		global.C_GOALLINE.Add(commID, fmt.Sprintf("%s|%s", goalLineTime, goalValue), cache.NoExpiration)
		// 写数据库
		_, err := models.GoalLineModel.AddList(
			commID,
			goalOdds[0].Odds,
			goalOdds[0].Name,
			goalOdds[1].Odds,
			goalOdds[1].Name,
			goalLineTime,
		)
		if err != nil {
			log.Error().Str("err", err.Error()).Send()
		}
	} else {
		// 存在判断是否一致
		if cacheData.(string)[11:] != goalValue {
			// 不一致
			// 写缓存
			global.C_GOALLINE.Set(commID, fmt.Sprintf("%s|%s", goalLineTime, goalValue), cache.NoExpiration)
			// 写数据库
			_, err := models.GoalLineModel.AddList(
				commID,
				goalOdds[0].Odds,
				goalOdds[0].Name,
				goalOdds[1].Odds,
				goalOdds[1].Name,
				goalLineTime,
			)
			if err != nil {
				log.Error().Str("err", err.Error()).Send()
			}
		}
	}
	return nil
}

func (s *prematchService) ParseFullTime(commID string, data *entity.RespPrematch) error {
	if len(data.Results) == 0 {
		return fmt.Errorf("result.length == 0")
	}

	if len(data.Results[0].Main.SP.FullTimeResult.Odds) != 3 {
		return fmt.Errorf("fulltime.odds != 3")
	}

	fTime := data.Results[0].Main.UpdateAt
	fOdds := data.Results[0].Main.SP.FullTimeResult.Odds
	fValue := fmt.Sprintf("%s|%s|%s",
		fOdds[0].Odds,
		fOdds[1].Odds,
		fOdds[2].Odds,
	)

	// 读取缓存是否存在
	cacheData, ok := global.C_FULLTIME.Get(commID)
	if !ok {
		// 不存在
		// 写缓存
		global.C_FULLTIME.Add(commID, fmt.Sprintf("%s|%s", fTime, fValue), cache.NoExpiration)
		// 写数据库
		_, err := models.FullTimeModel.AddList(
			commID,
			fOdds[0].Odds,
			fOdds[1].Odds,
			fOdds[2].Odds,
			fTime,
		)
		if err != nil {
			log.Error().Str("err", err.Error()).Send()
		}
	} else {
		// 存在判断是否一致
		if cacheData.(string)[11:] != fValue {
			// 不一致
			// 写缓存
			global.C_GOALLINE.Set(commID, fmt.Sprintf("%s|%s", fTime, fValue), cache.NoExpiration)
			// 写数据库
			_, err := models.FullTimeModel.AddList(
				commID,
				fOdds[0].Odds,
				fOdds[1].Odds,
				fOdds[2].Odds,
				fTime,
			)
			if err != nil {
				log.Error().Str("err", err.Error()).Send()
			}
		}
	}
	return nil
}
