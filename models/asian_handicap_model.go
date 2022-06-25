package models

import "github.com/sunp13/dbtool"

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
func (m *asianHandicapModel) AddList(commID, homeOdds, homeHandIcap, awayOdds, awayHandIcap, updateAt string) (res int64, err error) {
	sql := `
	insert into b365api.asian_handicap(
		comm_id,
		home_odds,
		home_handicap,
		away_odds,
		away_handicap,
		update_at
	) values (?,?,?,?,?,?)
	`
	params := []interface{}{
		commID,
		homeOdds,
		homeHandIcap,
		awayOdds,
		awayHandIcap,
		updateAt,
	}
	res, err = dbtool.D.UpdateSQL(sql, params)
	return
}
