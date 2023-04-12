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

	// 使用client.Ping()测试与Redis服务器的连接状态，并响应该连接状况。
	pong, err := client.Ping(context.Background()).Result()

	/**
	在调用 client.Ping() 方法时，传入了 context.Background() 作为参数。 context 是 Go 语言提供的标准库，它主要用于在程序中
	跟踪请求的上下文信息，并且可以将这些上下文信息传递给其他程序（例如：流、RPC等）。

	context.Background() 是创建一个空的上下文信息，也可以使用 context.WithCancel(), context.WithTimeout()
	和 context.WithDeadline() 等函数来创建具有不同属性的上下文信息。

	在这个方法中，使用 context.Background() 是因为不需要跟踪请求的具体信息，只是连接Redis并测试连接状态而已。
	实际上，我们无法从上下文中获取任何有用的信息。
	*/

	if err != nil {
		global.EWA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		fmt.Println("====4-redis====: redis init success")
		// 成功则打印出：redis connect ping response:	{"pong": "PONG"}
		global.EWA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		// 如果成功连接，则将全局变量global.EWA_REDIS设置为client实例。
		global.EWA_REDIS = client
	}
}

/**
`client.Ping(context.Background()).Result()` 这一行代码的作用是使用 Redis 客户端对象 `client` 发送一个 ping 命令来测试
与 Redis 服务器的连接是否正常，该命令实际上不会改变 Redis 数据库中的任何数据。

在这个过程中，函数将向 Redis 服务器发送一个 Ping 请求，如果 Redis 服务器正在运行且已准备好处理请求，则它将回复 Pong。
此时，客户端将接收到服务器返回的 Pong 响应。

`Ping()` 方法有一个额外参数，即 `context` 对象。在这里，我们使用 `context.Background()` 函数来创建一个空的
上下文对象（也可称为空白上下文），表示没有任何特殊属性的上下文。

对于这个方法的调用结果，我们使用 Go 语言的多重赋值语法来解包返回的两个值 `pong` 和 `err`。其中，`pong` 是字符串类型的值，
表示从 Redis 服务器返回的 Pong 响应；`err` 表示在执行函数期间可能发生的任何错误，它可以为 `nil`，如果没有出现错误。

因此，整行代码的含义是使用 Redis 客户端对象 `client` 对 Redis 服务器发送 Ping 命令，通过上下文对象 `context.Background()` 的
配置和控制等待和取消执行延迟任务，然后将响应结果保存在变量 `pong` 中并检查是否存在错误，最后返回结果或者错误信息。
*/
