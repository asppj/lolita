package mongo

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clientMapLock sync.Mutex
var clientMap = make(map[ClientName]*mongo.Client)

// Client 根据name获取client
// 注意：本方法直接返回client结果，是因为传入的name一定是写好的constant，而且一定已经成功初始化完成，否则程序不可能执行到这里，所以不用做任何判断
func Client(name ClientName) *mongo.Client {
	return clientMap[name]
}

// InitClient 根据配置初始化Client
func InitClient(c *Config) (*mongo.Client, error) {
	if clientMap[c.ClientName] != nil {
		return clientMap[c.ClientName], nil
	}
	var err error
	var client *mongo.Client
	clientMapLock.Lock()
	if clientMap[c.ClientName] == nil {
		client, err = connect(c)
		if client != nil {
			clientMap[c.ClientName] = client
		}
	}
	clientMapLock.Unlock()
	return client, err
}

// connect 建立数据库连接
func connect(c *Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultConnectTimeout*time.Second)
	defer cancel()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(c.Addr),
		options.Client().SetConnectTimeout(DefaultConnectTimeout*time.Second),
		options.Client().SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-1",
			AuthSource:    c.Auth.AuthSource,
			Username:      c.Auth.Username,
			Password:      c.Auth.Password,
			PasswordSet:   c.Auth.PasswordSet,
		}),
	)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
