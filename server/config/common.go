package config

import "github.com/cngamesdk/go-core/config"

type CommonConfig struct {
	config.CommonConfig `yaml:",inline" mapstructure:",squash"`
	GamePackaging       GamePackaging `mapstructure:"game_packaging" json:"game_packaging" yaml:"game_packaging"` // 游戏打包配置
	Endpoint            string        `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`                   // 接入点
}

type GamePackaging struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                            // 打包存储路径
	ProcessCount int    `mapstructure:"process_count" json:"process_count" yaml:"process_count"` // 进程数
}
