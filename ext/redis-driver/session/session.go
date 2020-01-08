package session

import (
	"encoding/json"
	"sync"

	"github.com/asppj/t-mk-opentrace/ext/redis-driver/redis"

	"github.com/go-errors/errors"
)

// RoleIds 角色信息
type RoleIds struct {
	PcMenuLimits []string          `json:"pcMenuLimits"` // Pc菜单权限Id
	PcFuncLimits []string          `json:"pcFuncLimits"` // Pc功能权限Id
	PcDataLimits map[string]string `json:"pcDataLimits"` // Pc数据权限Id
}

// Session session结构
type Session struct {
	ID           string  `json:"id"`           // sessionId
	UID          string  `json:"uid"`          // 座席id
	Name         string  `json:"name"`         // 座席姓名
	Phone        string  `json:"phone"`        // 坐席手机号
	Account      string  `json:"account"`      // 账号id
	AccountDB    string  `json:"accountDb"`    // 分库db
	DepartmentID string  `json:"departmentId"` // 部门Id
	Identity     string  `json:"identity"`     // 用户身份
	RoleLimits   RoleIds `json:"roleLimits"`   // 角色Id
	CreateTime   string  `json:"createTime"`   // 创建时间
	LastTime     string  `json:"lastTime"`     // 最后更新时间
}

var cao *redis.Cao
var caoLock sync.Mutex

// InitSession 初始化session
func InitSession() (*redis.Cao, error) {
	if cao != nil {
		return cao, nil
	}
	caoLock.Lock()
	if redis.Client(redis.MKSession) == nil {
		return nil, errors.New("init redis session client first")
	}
	if cao == nil {
		cao = redis.NewCao(redis.Client(redis.MKSession))
	}
	caoLock.Unlock()
	return cao, nil
}

// Get 获取session
func Get(sid string) (*Session, error) {
	if sid == "" {
		return nil, errors.New("sid为空")
	}
	ss, err := cao.Get(redis.SessionPrefix + sid)
	if err != nil {
		return nil, err
	}
	s := &Session{}
	if ss != "" {
		err := json.Unmarshal([]byte(ss), s)
		if err != nil {
			return nil, err
		}
		err = cao.Expire(redis.SessionPrefix+sid, redis.DAY)
		if err != nil {
			return nil, err
		}
	}
	if s.ID == "" {
		return nil, errors.New("session不合法")
	}
	return s, nil
}

// GetForMobile 获取session移动端
func GetForMobile(sid string) (*MobileSession, error) {
	if sid == "" {
		return nil, errors.New("sid为空")
	}

	ss, err := cao.Get(redis.SessionForMobilePrefix + sid)
	if err != nil {
		return nil, err
	}

	s := &MobileSession{}
	if ss != "" {
		err := json.Unmarshal([]byte(ss), s)
		if err != nil {
			return nil, err
		}
		err = cao.Expire(redis.SessionForMobilePrefix+sid, redis.DAY)
		if err != nil {
			return nil, err
		}
	}

	if s.ID == "" {
		return nil, errors.New("session不合法")
	}

	return s, nil
}

// Set 设置session
// TODO: 保留方法，营销板块的业务中不涉及到session的设置功能
func (s *Session) Set(user *Session) (string, error) {
	return "", nil
}

// MobileSession session结构
type MobileSession struct {
	ID           string  `json:"id"`           // sessionId
	UID          string  `json:"uid"`          // 座席id
	Name         string  `json:"name"`         // 座席姓名
	Phone        string  `json:"phone"`        // 坐席手机号
	Account      string  `json:"account"`      // 账号id
	AccountDB    string  `json:"accountDb"`    // 分库db
	DepartmentID string  `json:"departmentId"` // 部门Id
	Identity     string  `json:"identity"`     // 用户身份
	RoleLimits   RoleIds `json:"roleLimits"`   // 角色Id
	CreateTime   string  `json:"createTime"`   // 创建时间
	LastTime     int64   `json:"lastTime"`     // 最后更新时间
}
