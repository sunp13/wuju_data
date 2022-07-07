package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
	"wuju_data/entity"
	"wuju_data/global"
	"wuju_data/models"
	"wuju_data/utils"
)

type eventService struct {
	token string
	url   string
}

func (s *eventService) Setup(c *entity.Conf) {
	s.token = c.Token
	s.url = c.Event

}

// 根据commid 获取数据
func (s *eventService) GetEventData(commID string) error {
	// 调用
	url := fmt.Sprintf(s.url, s.token, commID)

	respByte, err := utils.HTTPGet(url)
	if err != nil {
		return err
	}
	var res entity.RespEvent

	err = global.JSON.Unmarshal(respByte, &res)
	if err != nil {
		return err
	}

	if res.Success != 1 {
		return fmt.Errorf("resp.success != 1 (%d) ,%s", res.Success, url)
	}

	if len(res.Result) == 0 {
		return fmt.Errorf("resp.result.length == 0")
	}

	// 解析数据
	// 拿到开始下标
	beginID := 0
	for i, v := range res.Result[0] {
		t := fmt.Sprintf("%v", v["type"])
		id := fmt.Sprintf("%v", v["ID"])
		if t == "MG" && id == "12" {
			beginID = i
		}
	}

	dataList := make(map[string][]map[string]string)
	dataList["1"] = make([]map[string]string, 0)
	dataList["2"] = make([]map[string]string, 0)
	id := ""
	for i := beginID + 1; i < beginID+50; i++ {
		t := fmt.Sprintf("%v", res.Result[0][i]["type"])
		if t != "MA" && t != "PA" {
			break
		}
		if t == "MA" {
			id = fmt.Sprintf("%v", res.Result[0][i]["ID"])
			continue
		}
		if t == "PA" {
			od := res.Result[0][i]["OD"]
			odSpl := strings.Split(od, "/")

			a, _ := strconv.ParseFloat(odSpl[0], 64)
			b, _ := strconv.ParseFloat(odSpl[1], 64)
			s := a/b + 1

			sStr := fmt.Sprintf("%.3f", s)
			res.Result[0][i]["OD"] = sStr
			dataList[id] = append(dataList[id], res.Result[0][i])
		}
	}

	homeDataByte, _ := global.JSON.Marshal(dataList["1"])
	awayDataByte, _ := global.JSON.Marshal(dataList["2"])

	// 计算MD5
	h := md5.New()
	h.Write([]byte(commID))
	h.Write(homeDataByte)
	h.Write(awayDataByte)
	hRes := hex.EncodeToString(h.Sum(nil))

	_, err = models.InPlayEventModel.AddList(
		SnowFlakeService.NextID(),
		commID,
		string(homeDataByte),
		string(awayDataByte),
		hRes,
		time.Now().Format("2006-01-02 15:04:05"),
	)
	return err
}
