package redis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

// Cao 缓存访问结构
type Cao struct {
	client *redis.Client
}

// NewCao 创建新的缓存访问对象
func NewCao(client *redis.Client) *Cao {
	return &Cao{client}
}

// Get 获取redis缓存
func (c *Cao) Get(key string) (string, error) {
	return c.client.Get(key).Result()
}

// SetByTTL 设置 带有有效时间
func (c *Cao) SetByTTL(key string, value string, extime int64) error {
	return c.client.Set(key, value, time.Duration(extime)*time.Duration(time.Second)).Err()
}

// Set 设置
func (c *Cao) Set(key string, value string) error {
	return c.client.Set(key, value, time.Duration(-1)*time.Second).Err()
}

// Keys Keys
func (c *Cao) Keys(pattern string) ([]string, error) {
	return c.client.Keys(pattern).Result()
}

// Scan Scan
func (c *Cao) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	return c.client.Scan(cursor, match, count).Result()
}

// Expire 过期
func (c *Cao) Expire(key string, extime int64) error {
	return c.client.Expire(key, time.Duration(extime)*time.Second).Err()
}

// ExpireAt 设置过期时间
func (c *Cao) ExpireAt(key string, ex time.Time) error {
	return c.client.ExpireAt(key, ex).Err()
}

// HSet Hset
func (c *Cao) HSet(key string, field string, value interface{}) error {
	return c.client.HSet(key, field, value).Err()
}

// HSetJSON HSet for json
func (c *Cao) HSetJSON(key, field string, value interface{}) error {
	btr, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.HSet(key, field, string(btr)).Err()
}

// HMSet HMSet
func (c *Cao) HMSet(key string, fields map[string]interface{}) error {
	return c.client.HMSet(key, fields).Err()
}

// SetTTL SetTTL
func (c *Cao) SetTTL(key string, extime int64) error {
	return c.client.Expire(key, time.Duration(extime)*time.Duration(time.Second)).Err()
}

// Del Del
func (c *Cao) Del(key string) error {
	return c.client.Del(key).Err()
}

// DelPipelined 批量删除
func (c *Cao) DelPipelined(keys []string) error {
	_, err := c.client.Pipelined(func(pipe redis.Pipeliner) error {
		for _, key := range keys {
			err := pipe.Del(key).Err()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// DelPipeline 批量删除
func (c *Cao) DelPipeline(keys []string) error {
	pipe := c.client.Pipeline()
	for _, key := range keys {
		err := pipe.Del(key).Err()
		if err != nil {
			return err
		}
	}
	_, err := pipe.Exec()
	return err
}

// HDel Hdel
func (c *Cao) HDel(key string, fields string) error {
	return c.client.HDel(key, fields).Err()
}

// HGet Hget
func (c *Cao) HGet(key string, field string) (string, error) {
	return c.client.HGet(key, field).Result()
}

// HGetJSON HGetJSON
func (c *Cao) HGetJSON(key, field string, T interface{}) error {
	str, err := c.client.HGet(key, field).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(str), T)
	if err != nil {
		return err
	}
	return nil
}

// HGetAll 获取全部
func (c *Cao) HGetAll(key string) map[string]interface{} {
	return ConvertStringToMap(c.client.HGetAll(key).Val())
}

// ConvertStringToMap 转换从redis获取的数据
func ConvertStringToMap(base map[string]string) map[string]interface{} {
	resultMap := make(map[string]interface{})
	for k, v := range base {
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(v), &dat); err == nil {
			resultMap[k] = dat
		} else {
			resultMap[k] = v
		}
	}
	return resultMap
}
