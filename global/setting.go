package global

import (
	"github.com/yunshu2009/blog-api/pkg/logger"
	"github.com/yunshu2009/blog-api/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	EmailSetting    *setting.EmailSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
