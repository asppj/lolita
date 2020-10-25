package es

import "time"

const (
	// IndexPrefix 索引前缀
	IndexPrefix = "mk_"

	// IndexAliasPostfix 索引别名后缀
	IndexAliasPostfix = "_current"

	// DefaultQueryTimeOut 默认查询超时时间
	DefaultQueryTimeOut time.Duration = 5
)
