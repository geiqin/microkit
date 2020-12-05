package session

import "github.com/geiqin/microkit/cache"

type SessConfig struct {
	Driver      string     `json:"driver"`
	CookieName  string     `json:"cookie_name"`
	MaxLifeTime int64      `json:"max_life_time"`
	Provider    *cache.RedisConfig `json:"provider"`
}