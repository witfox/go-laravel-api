package app

import (
	"gohub/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsLTesting() bool {
	return config.Get("app.env") == "testing"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

//获取当前时间，支持时区
func TimeNowInTimezone() time.Time {

	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}
