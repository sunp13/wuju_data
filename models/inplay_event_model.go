package models

import "github.com/sunp13/dbtool"

type inPlayEventModel struct{}

func (m *inPlayEventModel) AddList(snowID, commID, homeData, awayData, dataHash, addTime string) (res int64, err error) {
	sql := `
	insert into b365api.inplay_event(
		event_id,
		comm_id,
		home_data,
		away_data,
		data_hash,
		add_time
	) values (?,?,?,?,?,?) on duplicate key update update_time = ?
	`
	params := []interface{}{
		snowID,
		commID,
		homeData,
		awayData,
		dataHash,
		addTime,
		addTime,
	}
	res, err = dbtool.D.UpdateSQL(sql, params)
	return
}
