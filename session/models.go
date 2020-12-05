package session

type SessConfig struct {
	Driver      string     `json:"driver"`
	CookieName  string     `json:"cookie_name"`
	MaxLifeTime int64      `json:"max_life_time"`
	Provider    *RedisProviderConfig `json:"provider"`
}

type RedisProviderConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database int    `json:"database"`
}