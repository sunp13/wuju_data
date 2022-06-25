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

func (m *goalLineModel) AddList(commID, overOdds, overName, underOdds, underName, updateAt string) (res int64, err error) {
	sql := `
	insert into b365api.goal_line(
		comm_id,
		over_odds,
		over_name,
		under_odds,
		under_name,
		update_at
	) values (?,?,?,?,?,?)
	`
	params := []interface{}{
		commID,
		overOdds,
		overName,
		underOdds,
		underName,
		updateAt,
	}
	res, err = dbtool.D.UpdateSQL(sql, params)
	return
}
