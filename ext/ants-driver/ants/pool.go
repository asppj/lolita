package ants

import (
	"log"
	"sync"

	ant "github.com/panjf2000/ants"
)

// init Init
func init() {
	if _, err := DefaultPool(); err != nil {
		panic(err)
	}
}

// Pool pool
type Pool = ant.Pool

// PoolWithFunc PoolWithFunc
type PoolWithFunc = ant.PoolWithFunc

// Option opt
type Option = ant.Option

var _pool *Pool
var _poolOnce sync.Once

// DefaultPool 协程池
func DefaultPool() (p *Pool, err error) {
	if _pool != nil {
		return _pool, nil
	}
	_poolOnce.Do(func() {
		p, err = ant.NewPool(PoolSize)
	})
	if err != nil {
		return
	}
	_pool = p
	return _pool, nil
}

// NewPool 新建
func NewPool(size int, opts ...Option) (*Pool, error) {
	return ant.NewPool(size, opts...)
}

// Go submits a task to pool.
func Go(task func()) {
	if err := _pool.Submit(task); err != nil {
		log.Fatalln(err)
	}
}

// Submit submits a task to pool.
func Submit(task func()) error {
	return _pool.Submit(task)
}

// Running returns the number of the currently running goroutines.
func Running() int {
	return _pool.Running()
}

// Cap returns the capacity of this default pool.
func Cap() int {
	return _pool.Cap()
}

// Free returns the available goroutines to work.
func Free() int {
	return _pool.Free()
}

// Release Closes the default pool.
func Release() {
	_pool.Release()
}
