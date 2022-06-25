package entity

// 获取upcoming数据
type RespComming struct {
	Success int `json:"success"`
	Paper   struct {
		Page    int `json:"page"`
		PerPage int `json:"per_page"`
		Total   int `json:"total"`
	} `json:"paper"`

	Results []CommingData `json:"results"`
}
type CommingData struct {
	ID      string `json:"id"`
	SportID string `json:"sport_id"`
	Time    string `json:"time"`
	League  struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"league"`
	Home struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"home"`
	Away struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"away"`
	SS         string `json:"ss"`
	OurEventID string `json:"our_event_id"`
	RID        string `json:"r_id"`
	UpdatedAt  string `json:"updated_at"`
}
