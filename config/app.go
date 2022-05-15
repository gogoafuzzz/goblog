package config

import "goblog/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"name":     config.Env("APP_NAME", "Goblog"),
			"env":      config.Env("APP_ENV", "production"),
			"debug":    config.Env("APP_DEBUG", false),
			"port":     config.Env("APP_PORT", "8888"),
			"key":      config.Env("APP_KEY", "YPDwNCEEMlqyWcSRzD23u67PBR64ay6C"),
			"url":      config.Env("APP_URL", "http://localhost:8888"),
			"timezone": config.Env("TIMEZONE", "Asia/Taipei"),
		}
	})
}
