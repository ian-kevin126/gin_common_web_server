package initialize

import (
	"context"
	"ewa_admin_server/global"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.EWA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.EWA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		fmt.Println("====4-redis====: redis init success")
		global.EWA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.EWA_REDIS = client
	}
}
