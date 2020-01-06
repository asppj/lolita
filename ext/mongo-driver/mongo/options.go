package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Cursor 游标
type Cursor = mongo.Cursor

// SingleResult 单个结果
type SingleResult = mongo.SingleResult

// InsertOneResult 插入的单个结果
type InsertOneResult = mongo.InsertOneResult

// InsertManyResult 多个结果
type InsertManyResult = mongo.InsertManyResult

// DeleteResult 删除结果
type DeleteResult = mongo.DeleteResult

// FindOptions 查找option
type FindOptions = options.FindOptions

// FindOneOptions 查找option
type FindOneOptions = options.FindOneOptions

// DeleteOptions 删除option
type DeleteOptions = options.DeleteOptions

// UpdateResult 更新结果
type UpdateResult = mongo.UpdateResult

// UpdateOptions 更新option
type UpdateOptions = options.UpdateOptions

// FindOneAndUpdateOptions 更新一个
type FindOneAndUpdateOptions = options.FindOneAndUpdateOptions

// FindOneAndDeleteOptions 删除一个
type FindOneAndDeleteOptions = options.FindOneAndDeleteOptions

// AggregateOptions 聚合
type AggregateOptions = options.AggregateOptions

// CountOptions count
type CountOptions = options.CountOptions

// SessionContext 事务session
type SessionContext = mongo.SessionContext

// WriteException 写入错误
type WriteException = mongo.WriteException

// InsertOneOptions insert
type InsertOneOptions = options.InsertOneOptions

// InsertManyOptions insertMany
type InsertManyOptions = options.InsertManyOptions

// ErrNoDocuments is returned by Decode when an operation that returns a
// SingleResult doesn't return any documents.
var ErrNoDocuments = mongo.ErrNoDocuments
