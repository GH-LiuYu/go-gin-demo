package global

import (
	"demo/config"
	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
	"github.com/go-redis/redis"
)

var (
	DEMO_DB     *gorm.DB
	DEMO_CONFIG config.Server
	DEMO_LOG    *oplogging.Logger
	DEMO_VP     *viper.Viper
	DEMO_REDIS  *redis.Client
)

