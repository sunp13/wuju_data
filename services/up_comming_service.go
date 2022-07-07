package services

import (
	"fmt"
	"strconv"
	"time"
	"wuju_data/entity"
	"wuju_data/global"
	"wuju_data/models"
	"wuju_data/utils"

	"github.com/rs/zerolog/log"
)

type upCommingService struct {
	token string
	url   string
}

// 初始化
func (s *upCommingService) Setup(conf *entity.Conf) {
	s.token = conf.Token
	s.url = conf.UpComing
}

func (s *upCommingService) Reload(conf *entity.Conf) {
	for _, v := range conf.Leagues {
		data, err := models.UpCommingModel.GetListByLeagueID(v.ID)
		if err != nil {
			log.Error().Str("err", err.Error()).Send()
			continue
		}
		for _, dataV := range data {
			commID := fmt.Sprintf("%v", dataV["comm_id"])
			commTime := fmt.Sprintf("%v", dataV["comm_time"])
			commTimeInt64, _ := strconv.ParseInt(commTime, 10, 64)
			if commTimeInt64 > time.Now().Unix() {
				expireSec := commTimeInt64 - time.Now().Unix() + 5*60
				if err := global.C_COMM.Add(commID, commTime, time.Duration(expireSec)*time.Second); err != nil {
					log.Error().Str("err", err.Error()).Send()
				}
			}
		}
	}
	log.Info().Msgf("reload upcomming success")
}

// 开始写库, 写缓存
func (s *upCommingService) GetCommingData(leagueID string, page int) error {
	url := fmt.Sprintf(s.url, s.token, leagueID, page)

	respByte, err := utils.HTTPGet(url)
	if err != nil {
		return err
	}

	var respComm entity.RespComming
	err = global.JSON.Unmarshal(respByte, &respComm)
	if err != nil {
		return err
	}
	if respComm.Success != 1 {
		return fmt.Errorf("resp_comm.success != 1 (%d)", respComm.Success)
	}

	i := 0
	for _, v := range respComm.Results {
		_, err := models.UpCommingModel.AddList(
			SnowFlakeService.NextID(),
			v.ID,
			v.Time,
			v.League.ID,
			v.League.Name,
			v.Home.ID,
			v.Home.Name,
			v.Away.ID,
			v.Away.Name,
			v.SS,
			v.OurEventID,
			v.RID,
			time.Now().Format("2006-01-02 15:04:05"),
		)
		if err != nil {
			continue
		}

		commTime, _ := strconv.ParseInt(v.Time, 10, 64)
		if commTime > time.Now().Unix() {
			// 数据过期时间为 开赛过后300秒(5分钟)
			expireSec := commTime - time.Now().Unix() + 5*60
			if err := global.C_COMM.Add(v.ID, commTime, time.Duration(expireSec)*time.Second); err != nil {
				log.Error().Str("err", err.Error()).Send()
			}
		}
		i++
	}
	log.Info().Int("count", len(respComm.Results)).Int("new", i).Msgf("up_comm update result")
	return nil
}
