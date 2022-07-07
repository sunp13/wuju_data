package models

import "github.com/sunp13/dbtool"

type goalLineModel struct{}

func (m *goalLineModel) GetListByCommID(commID string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.goal_line
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

func (m *goalLineModel) AddList(snowID, commID, overOdds, overName, underOdds, underName, addTime, dataHash string) (res int64, err error) {
	sql := `
	insert into b365api.goal_line(
		goal_id,
		comm_id,
		over_odds,
		over_name,
		under_odds,
		under_name,
		add_time,
		data_hash
	) values (?,?,?,?,?,?,?,?) on duplicate key update update_time = ?
	`
	params := []interface{}{
		snowID,
		commID,
		overOdds,
		overName,
		underOdds,
		underName,
		addTime,
		dataHash,
		addTime,
	}
	res, err = dbtool.D.UpdateSQL(sql, params)
	return
}
