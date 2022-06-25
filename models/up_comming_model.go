package models

import (
	"time"

	"github.com/sunp13/dbtool"
)

type upCommingModel struct{}

// GetList
func (m *upCommingModel) GetList() (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.up_coming
	where comm_time > ?
	`
	params := []interface{}{
		time.Now().Add(-1 * time.Hour).Unix(),
	}
	res, err = dbtool.D.QuerySQL(sql, params)
	return
}

// 根据联赛ID获取已存在数据
func (m *upCommingModel) GetListByLeagueID(leagueID string) (res []map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.up_coming
	where league_id = ?
	and comm_time > ?
	`
	params := []interface{}{
		leagueID,
		time.Now().Add(-1 * time.Hour).Unix(),
	}
	res, err = dbtool.D.QuerySQL(sql, params)
	return
}

// GetListByID 获取单场比赛
func (m *upCommingModel) GetListByID(id string) (res map[string]interface{}, err error) {
	sql := `
	SELECT * FROM b365api.up_coming
	where comm_id = ?
	`
	params := []interface{}{
		id,
	}
	var result []map[string]interface{}
	result, err = dbtool.D.QuerySQL(sql, params)

	if len(result) > 0 {
		res = result[0]
	}
	return
}

func (m *upCommingModel) AddList(commID, commTime, leagueID, leagueName, homeID, homeName, awayID, awayName, ss, ourEventID, rID, updateAt string) (res int64, err error) {

	sql := `
	insert into b365api.up_coming(
		comm_id,
		comm_time,
		league_id,
		league_name,
		home_id,
		home_name,
		away_id,
		away_name,
		ss,
		our_event_id,
		r_id,
		update_at
	) value (?,?,?,?,?,?,?,?,?,?,?,?)
	`
	params := []interface{}{
		commID,
		commTime,
		leagueID,
		leagueName,
		homeID,
		homeName,
		awayID,
		awayName,
		ss,
		ourEventID,
		rID,
		updateAt,
	}
	res, err = dbtool.D.UpdateSQL(sql, params)
	return
}

func (m *upCommingModel) DeleteList(id string) (res int64, err error) {
	sql := `
	update b365api.up_coming set
	is_deleted = 1
	where comm_id = ?
	`
	params := []interface{}{
		id,
	}
	res, err = dbtool.D.UpdateSQL(sql, params)
	return
}
