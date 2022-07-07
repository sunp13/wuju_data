package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
	"wuju_data/entity"
	"wuju_data/global"
	"wuju_data/initlize"
	"wuju_data/services"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

var (
	conf entity.Conf
)

func init() {

	if err := initlize.InitAll(); err != nil {
		panic(err)
	}

	confByte, err := ioutil.ReadFile("./conf/conf.yml")
	if err != nil {
		log.Fatal().Str("err", err.Error()).Send()
	}

	if err = yaml.Unmarshal(confByte, &conf); err != nil {
		log.Fatal().Str("err", err.Error()).Send()
	}

	// 初始化并发chan,每秒50最大请求
	go func() {
		t := time.NewTicker(1 * time.Second)
		for range t.C {
			for i := 0; i < 100; i++ {
				select {
				case global.CH_BULK <- struct{}{}:
				default:
				}
			}
		}
	}()

}

func main() {
	// 初始化缓存数据
	go func() {
		t := time.NewTicker(1 * time.Second)
		for range t.C {
			fmt.Println(len(global.CH_BULK))
		}
	}()

	// 初始化UpComing
	services.UpCommingService.Setup(&conf)
	services.UpCommingService.Reload(&conf)
	// 初始化prematch
	services.PrematchService.Setup(&conf)
	// 初始化event
	services.EventService.Setup(&conf)
	// 初始化snow
	services.SnowFlakeService.Setup(&conf)

	// 1. 跑upComing数据
	go func() {
		t := time.NewTicker(time.Duration(conf.UpComingInterval) * time.Second)
		for range t.C {
			for _, v := range conf.Leagues {
				for i := 0; i < v.MaxPage; i++ {
					// 同样算在并发里面
					<-global.CH_BULK
					page := i + 1
					if err := services.UpCommingService.GetCommingData(v.ID, page); err != nil {
						log.Error().Str("err", err.Error()).Send()
					}
				}
			}
		}
	}()

	// 2. 跑permatch数据
	go func() {
		t := time.NewTicker(1 * time.Second)
		for range t.C {
			for commID := range global.C_COMM.Items() {
				<-global.CH_BULK
				if err := services.PrematchService.GetPrematchData(commID); err != nil {
					log.Error().Str("err", err.Error()).Send()
				}
			}
		}
	}()

	// 3. 跑event数据
	go func() {
		t := time.NewTicker(1 * time.Second)
		for range t.C {
			for commID, commValue := range global.C_COMM.Items() {
				// 判断开始时间
				commTime := fmt.Sprintf("%v", commValue.Object)
				commTimeUnix, _ := strconv.ParseInt(commTime, 10, 64)
				nowUnix := time.Now().Unix()
				// 判断是否距离开赛15分钟之内
				if commTimeUnix-nowUnix < 15*60 {
					// 获取并发权限
					<-global.CH_BULK
					err := services.EventService.GetEventData(commID)
					if err != nil {
						log.Error().Str("err", err.Error()).Send()
					}
				}
			}
		}
	}()
	select {}
}
