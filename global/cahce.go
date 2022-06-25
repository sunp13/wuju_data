package global

import "github.com/patrickmn/go-cache"

var (
	// 缓存未开始比赛信息
	// key = comm_id(比赛ID)
	// value = 比赛时间(unix) // 只有未开始的比赛才会被缓存 (时间unix > 当前时间unix)
	C_COMM = cache.New(cache.NoExpiration, cache.NoExpiration)

	// 亚盘结果缓存
	// key = comm_id(比赛id)
	// value = 时间UNIX|主队odds|主队handicap|客队odds|客队handicap
	C_ASIA = cache.New(cache.NoExpiration, cache.NoExpiration)

	// 大小结果缓存
	// key = comm_id(比赛ID)
	// value = 时间UNIX|OverODDS|OverName|UnderOdds|UnderName
	C_GOALLINE = cache.New(cache.NoExpiration, cache.NoExpiration)

	// 欧赔结果缓存
	// key = comm_id(比赛ID)
	// value = 时间UNIX|主队Odds|平局ODDS|客队ODDS
	C_FULLTIME = cache.New(cache.NoExpiration, cache.NoExpiration)
)
