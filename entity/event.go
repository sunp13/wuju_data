package entity

type RespEvent struct {
	Success int                   `json:"success"`
	Result  [][]map[string]string `json:"results"`
	Stats   struct {
		UpdateAt string `json:"update_at"`
		UpdateDt string `json:"update_dt"`
	} `json:"stats"`
}
