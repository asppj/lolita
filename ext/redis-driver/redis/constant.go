package redis

// ClientName redis连接名称
type ClientName string

// ClientName redis连接名称
const (
	MKSession ClientName = "session"
	MKCache   ClientName = "cache"
)

// Prefix 缓存前缀配置（一级命名空间管理，放在底层package中，方式各自维护各自的导致命名空间冲突）
const (
	SessionPrefix          = "wpmanage." // SessionPrefix 工作手机团队session前缀（所有团队要统一）
	LRUCachePrefix         = "lru."      // LRUCachePrefix lru缓存前缀
	FIFOCachePrefix        = "fifo."     // FIFOCachePrefix fifo缓存前缀
	SessionForMobilePrefix = "wpapp."    // SessionForMobilePrefix 工作手机团队session前缀（所有团队要统一）
)

// time 缓存时间配置
const (
	MINUTE   = 60          // 分钟
	HOUR     = MINUTE * 60 // 小时
	HALFHOUR = MINUTE * 30 // 半小时
	DAY      = HOUR * 24   // 天

	LRUDuration      = MINUTE * 10 // lru持续时间
	LRUDurationByDay = DAY * 10    // lru持续时间 day

	FIFODuration = MINUTE * 10 // fifo缓存持续时间
)
