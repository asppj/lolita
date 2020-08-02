package cache

import (
	"encoding/json"
	"errors"

	"github.com/asppj/lolita/ext/redis-driver/redis"
)

// FIFO fifo缓存
type FIFO struct {
	cao            *redis.Cao
	prefix         string
	Duration       int64
	ResourcePrefix string
}

// NewFIFO 新建新的fifo缓存
func NewFIFO(prefix string, duration int64) (*FIFO, error) {
	client := redis.Client(redis.MKCache)
	if client == nil {
		return nil, errors.New("init redis cache client first")
	}
	return &FIFO{
		cao:            redis.NewCao(client),
		prefix:         redis.FIFOCachePrefix,
		Duration:       duration,
		ResourcePrefix: prefix,
	}, nil
}

func (fifo *FIFO) id(rid string) string {
	if fifo == nil {
		return ""
	}
	return fifo.prefix + fifo.ResourcePrefix + rid
}

// FindJSON 获取fifo缓存
func (fifo *FIFO) FindJSON(rid string, v interface{}) error {
	if rid == "" {
		return errors.New("FIFO.Find id is required")
	}
	if v == nil {
		return errors.New("FIFO.Find v is nil")
	}
	id := fifo.id(rid)
	cs, err := fifo.cao.Get(id)
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil
		}
		return err
	}
	err = json.Unmarshal([]byte(cs), v)
	return err
}

// FindString 获取fifo缓存
func (fifo *FIFO) FindString(rid string) (string, error) {
	if rid == "" {
		return "", errors.New("FIFO.Find id is required")
	}
	id := fifo.id(rid)
	cs, err := fifo.cao.Get(id)
	if err != nil {
		if err.Error() == "redis: nil" {
			return "", nil
		}
		return "", err
	}
	return cs, nil
}

// CreateJSON 创建fifo缓存
func (fifo *FIFO) CreateJSON(rid string, v interface{}) error {
	id := fifo.id(rid)
	cs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return fifo.cao.SetByTTL(id, string(cs), fifo.Duration)
}

// CreateString 创建fifo缓存
func (fifo *FIFO) CreateString(rid string, v string) error {
	id := fifo.id(rid)
	return fifo.cao.SetByTTL(id, v, fifo.Duration)
}

// Drop 删除缓存
func (fifo *FIFO) Drop(rid string) error {
	id := fifo.id(rid)
	return fifo.cao.Del(id)
}

// DropMany 删除多个缓存
func (fifo *FIFO) DropMany(rids []string) error {
	return fifo.cao.DelPipelined(rids)
}
