package models

import (
	"github.com/sunp13/dbtool"
)

type asianHandicapModel struct{}

//
func (m *asianHandicapModel) GetListByCommID(commID string) (res []map[string]interface{}, err error) {
	sql := `
	select * from b365api.asian_handicap
	where comm_id = ? 
	order by update_at desc
	limit 1
	`
	params := []interface{}{
		commID,
	}
	res, err = dbtool.D.QuerySQL(sql, params)
	return
}

// 添加入库
func (m *asianHandicapModel) AddList(snowID, commID, homeOdds, homeHandIcap, awayOdds, awayHandIcap, addTime, dataHash string) (res int64, err error) {
	sql := `
	insert into b365api.asian_handicap(
		hand_id,
		comm_id,
		home_odds,
		home_handicap,
		away_odds,
		away_handicap,
		add_time,
		data_hash
	) values (?,?,?,?,?,?,?,?) on duplicate key update update_time = ?
	`
	params := []interface{}{
		snowID,
		commID,
		homeOdds,
		homeHandIcap,
		awayOdds,
		awayHandIcap,
		addTime,
		dataHash,
		addTime,
	}
	res, err = dbtool.D.UpdateSQL(sql, params)
	return
}
