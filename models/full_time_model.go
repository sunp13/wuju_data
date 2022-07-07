package models

import "github.com/sunp13/dbtool"

type fullTimeModel struct {
}

// GetListByCommID
func (m *fullTimeModel) GetListByCommID(commID string) (res []map[string]interface{}, err error) {
	sql := `
	select * from b365api.full_time
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

// AddList
func (m *fullTimeModel) AddList(snowID, commID, homeOdds, drawOdds, awayOdds, addTime, dataHash string) (res int64, err error) {
	sql := `
	insert into b365api.full_time(
		full_id,
		comm_id,
		home_odds,
		draw_odds,
		away_odds,
		add_time,
		data_hash
	) values (?,?,?,?,?,?,?) on duplicate key update update_time = ?
	`
	params := []interface{}{
		snowID,
		commID,
		homeOdds,
		drawOdds,
		awayOdds,
		addTime,
		dataHash,
		addTime,
	}
	res, err = dbtool.D.UpdateSQL(sql, params)
	return

}
