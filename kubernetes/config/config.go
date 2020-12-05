// Package config implements go-config with env and k8s configmap
package config

import (
	"github.com/geiqin/microkit/plugins/config/source/configmap"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/env"
)

// NewConfig returns config with env and k8s configmap setup
func NewConfig(opts ...config.Option) config.Config {
	cfg, _ := config.NewConfig()
	cfg.Load(
		env.NewSource(),
		configmap.NewSource(),
	)
	return cfg
}
