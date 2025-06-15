package config

import "flame_clouds/config/types"

type Config struct {
	Monitor Monitor `yaml:"monitor"`
	Bot     Bot     `yaml:"bot"`
}

type Monitor struct {
	City    string       `yaml:"city"`
	Evening MonitorEvent `yaml:"evening"` // 晚霞
	Morning MonitorEvent `yaml:"morning"` // 朝霞
	Map     Map          `yaml:"map"`     // 地图
}

type Map struct {
	Enable bool   `yaml:"enable"`
	Region string `yaml:"region"`
}

type MonitorEvent struct {
	EventType types.EventType `yaml:"eventType"`
	Enable    bool            `yaml:"enable"`
	Quality   float64         `yaml:"quality"`
	Time      string          `yaml:"time"`
}

type Bot struct {
	Enable  bool                `yaml:"enable"`
	Target  types.BotTargetType `yaml:"target"`
	SendKey string              `yaml:"sendKey"`
}
