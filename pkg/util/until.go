package util

import (
	"gravitation/pkg/setting"
)

func Setup() {
	JwtSecret := []byte(setting.AppSetting.JwtSecret)
}
