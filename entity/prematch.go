package entity

// 赛前数据
type RespPrematch struct {
	Success int              `json:"success"`
	Results []PrematchResult `json:"results"`
}

type PrematchResult struct {
	FI         string `json:"FI"`
	EventID    string `json:"event_id"`
	AsianLines struct {
		UpdatedAt string `json:"updated_at"`
		SP        struct {
			// 亚盘
			AsianHandicap struct {
				ID   string              `json:"id"`
				Name string              `json:"name"`
				Odds []AsianHandicapOdds `json:"odds"`
			} `json:"asian_handicap"`
			// 大小
			GoalLine struct {
				ID   string         `json:"id"`
				Name string         `json:"name"`
				Odds []GoalLineOdds `json:"odds"`
			} `json:"goal_line"`
		} `json:"sp"`
	} `json:"asian_lines"`
	Main struct {
		UpdateAt string `json:"updated_at"`
		SP       struct {
			// 欧赔
			FullTimeResult struct {
				ID   string               `json:"id"`
				Name string               `json:"name"`
				Odds []FullTimeResultOdds `json:"odds"`
			} `json:"full_time_result"`
		} `json:"sp"`
	} `json:"main"`
}
type AsianHandicapOdds struct {
	ID       string `json:"id"`
	Odds     string `json:"odds"`
	Header   string `json:"header"`
	Handicap string `json:"handicap"`
}

type GoalLineOdds struct {
	ID     string `json:"id"`
	Odds   string `json:"odds"`
	Header string `json:"header"`
	Name   string `json:"name"`
}

type FullTimeResultOdds struct {
	ID   string `json:"id"`
	Odds string `json:"odds"`
	Name string `json:"name"`
}
