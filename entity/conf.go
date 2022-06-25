package entity

// 配置文件
type Conf struct {
	Token            string        `yaml:"token"`
	UpComing         string        `yaml:"upcoming"`
	UpComingInterval int           `yaml:"upcoming_interval"`
	PreMatch         string        `yaml:"prematch"`
	Leagues          []ConfLeagues `yaml:"leagues"`
}
type ConfLeagues struct {
	ID      string `yaml:"id"`
	Name    string `yaml:"name"`
	MaxPage int    `yaml:"max_page"`
}
