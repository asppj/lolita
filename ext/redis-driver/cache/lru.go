package cache

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/asppj/t-mk-opentrace/ext/redis-driver/redis"
)

// LRU lru缓存
type LRU struct {
	cao            *redis.Cao
	prefix         string
	Duration       int64
	ResourcePrefix string
}

// NewLRU 获取LRU缓存
func NewLRU(prefix string, duration int64) (*LRU, error) {
	client := redis.Client(redis.MKCache)
	if client == nil {
		return nil, errors.New("init redis cache client first")
	}
	return &LRU{
		cao:            redis.NewCao(client),
		prefix:         redis.LRUCachePrefix,
		Duration:       duration,
		ResourcePrefix: prefix,
	}, nil
}

func (lru *LRU) id(rid string) string {
	if lru == nil {
		return ""
	}
	return lru.prefix + lru.ResourcePrefix + rid
}

// FindJSON 获取lru缓存
func (lru *LRU) FindJSON(rid string, v interface{}) error {
	if rid == "" {
		return errors.New("LRU.Find id is required")
	}
	if v == nil {
		return errors.New("LRU.Find v is nil")
	}
	id := lru.id(rid)
	cs, err := lru.cao.Get(id)
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil
		}
		return err
	}
	if cs != "" {
		err = lru.cao.Expire(id, lru.Duration)
		if err != nil {
			return fmt.Errorf("LRU.Find touch expire failed %v", err)
		}
	}
	err = json.Unmarshal([]byte(cs), v)
	return err
}

// CreateJSON 创建lru缓存
func (lru *LRU) CreateJSON(rid string, v interface{}) error {
	id := lru.id(rid)
	cs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return lru.cao.SetByTTL(id, string(cs), lru.Duration)
}

// Drop 删除缓存
func (lru *LRU) Drop(rid string) error {
	id := lru.id(rid)
	return lru.cao.Del(id)
}

// DropMany 删除多个缓存
func (lru *LRU) DropMany(rids []string) error {
	ids := make([]string, len(rids))
	for index, rid := range rids {
		if rid != "" {
			ids[index] = lru.id(rid)
		}
	}
	return lru.cao.DelPipeline(ids)
	// return lru.cao.DelPipelined(rids)
}
