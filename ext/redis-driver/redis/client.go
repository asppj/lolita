package redis

import (
	"sync"

	"github.com/go-redis/redis"
)

var clientMapLock sync.Mutex
var clientMap = make(map[ClientName]*redis.Client)

// Client 根据name获取客户端连接
func Client(name ClientName) *redis.Client {
	return clientMap[name]
}

// InitClient 初始化连接
func InitClient(config *Config) (*redis.Client, error) {
	if clientMap[config.ClientName] != nil {
		return clientMap[config.ClientName], nil
	}
	var err error
	var client *redis.Client
	clientMapLock.Lock()
	if clientMap[config.ClientName] == nil {
		client, err = connect(config)
		if err != nil {
			return nil, err
		}
		if client != nil {
			clientMap[config.ClientName] = client
		}
	}
	clientMapLock.Unlock()
	return client, err
}

// connect 建立redis连接
func connect(config *Config) (*redis.Client, error) {
	opt := &redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	}
	client := redis.NewClient(opt)

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
