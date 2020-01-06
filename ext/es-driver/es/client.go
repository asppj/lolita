package es

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/olivere/elastic"
	"github.com/olivere/elastic/trace/opentracing"
)

var client *elastic.Client
var clientLock sync.Mutex

// 链路追踪
func newTraceClient() *http.Client {
	tr := opentracing.NewTransport()
	return &http.Client{Transport: tr}
}

// Init 初始化es客户端连接
func Init(config *Config) error {
	if client != nil {
		return nil
	}

	clientLock.Lock()
	if client == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var err error
		client, err = elastic.NewClient(
			elastic.SetURL(config.Addr),
			elastic.SetBasicAuth(config.Username, config.Password),
			elastic.SetHttpClient(newTraceClient()),
			elastic.SetSniff(false),
		)
		if err != nil {
			return err
		}
		info, code, err := client.Ping(config.Addr).Do(ctx)
		if err != nil {
			return err
		}
		if code == 200 {
			fmt.Printf("connected to es: %s ,version: %s", info.ClusterName, info.Version.Number)
		}
	}
	clientLock.Unlock()

	return nil
}

// Client 获取es连接
func Client() *elastic.Client {
	return client
}
