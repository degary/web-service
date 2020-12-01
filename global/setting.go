package global

import (
	"github.com/degary/web-service/pkg/logger"
	"github.com/degary/web-service/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSeting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger *logger.Logger
)
