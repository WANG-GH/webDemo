package global

import (
	"webDemo/pkg/setting"

	"github.com/jinzhu/gorm"
)

var (
    DatabaseSetting *setting.DatabaseSettingS
	ServerSetting *setting.ServerSettingS
	JWTSetting *setting.JWTSettingS
	DBEngine *gorm.DB 
	EmailSetting *setting.EmailSettingS
)
